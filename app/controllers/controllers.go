package controllers

import (
	"fmt"

	"github.com/gpaulo00/velocity-test/config"
)

// Declare Variables in package controllers
var err error

// Recover = handle panic error
func catch() {
	// recover() = return error panic
	if err := recover(); err != nil {
		fmt.Println("Terjadi kesalahan :", err)
	}
}

func getKey() string {
	config := config.GetConfig()

	return config.GetString("server.api_key")
}
