
# Velocity Test
Prueba técnica para la empresa Velocity

Hecha en Golang, con las siguientes librerías:
- Gin
- GORM

Se utilizó `go mod` para el manejo de librerías

Nota: utilicé la API de openstreetmap para la geolocalización inversa, no utilicé mapbox

## Requisitos
- Go >1.16
- MySQL (se utilizó MariaDB 5.5)

## Configuración
Los archivos de configuración estan en el modulo config, en formato YAML.

## Ejemplo de Uso
Luego de ejecutar el proyecto se puede hacer este comando para probar la API
```sh
curl -X POST localhost:8000/v1/order/process/1
```

## Autor
Gustavo Paulo <gustavo.paulo.segura@gmail.com>