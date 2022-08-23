package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gpaulo00/velocity-test/app/models"
	"github.com/jftuga/geodist"
)

type Result struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

func (r Result) toCoor() geodist.Coord {
	lat, _ := strconv.ParseFloat(r.Lat, 64)
	lon, _ := strconv.ParseFloat(r.Lon, 64)
	return geodist.Coord{Lat: lat, Lon: lon}
}

type OrderController struct{}

var orderModel = new(models.Order)

func (h OrderController) Process(c *gin.Context) {
	defer catch()

	// obtener orden
	orderID := c.Param("id")
	order, err := orderModel.GetOrder(orderID)
	if err != nil {
		c.JSON(400, gin.H{"error": "No se pudo encontrar la orden", "status": 200})
		return
	}

	// get shipping information coordinates
	myurl, _ := url.Parse("https://nominatim.openstreetmap.org/search")
	//65 East 8th Street, Huntington Station, Huntington, NY 11746, United States of America
	q := myurl.Query()
	q.Set("q", order.ShippingInformation.Address)
	q.Set("format", "json")
	q.Set("polygon", "1")
	q.Set("addressdetails", "1")
	myurl.RawQuery = q.Encode()

	// request
	res, err := http.Get(myurl.String())
	if err != nil {
		c.JSON(400, gin.H{"error": "Dirección de envio invalida"})
		return
	}
	var data []Result
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil || len(data) < 1 {
		c.JSON(400, gin.H{"error": "json error", "message": err.Error()})
		return
	}

	// calculate distance
	pointA := geodist.Coord{Lat: order.Warehouse.Lat, Lon: order.Warehouse.Lng}
	pointB := data[0].toCoor()
	_, km, err := geodist.VincentyDistance(pointA, pointB)
	if err != nil {
		c.JSON(400, gin.H{"error": "No se pudo calcular distancia"})
		return
	}

	// check distance
	order.TotalKms = km
	if orderModel.UpdateOrder(order) != nil {
		c.JSON(400, gin.H{"error": "No se pudo actualizar"})
		return
	}

	c.JSON(200, gin.H{"message": "Success"})
}
