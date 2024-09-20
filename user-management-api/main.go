package main

import (
	"user-management-api/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080") // Start the server on port 8080
}
