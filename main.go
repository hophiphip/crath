package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

const defaultHost = "0.0.0.0"
const defaultPort = "8080"

type OrdSubmit struct {
	Element string `json:"element"`
	Module  string `json:"module"`
}

func main() {
	host, isHostSet := os.LookupEnv("CRATH_HOST")
	if !isHostSet {
		host = defaultHost
	}

	port, isPortSet := os.LookupEnv("CRATH_PORT")
	if !isPortSet {
		port = defaultPort
	}

	router := gin.Default()
	router.Use(CORSMiddleware())

	api := router.Group("/api")
	{
		api.GET("ord", getOrd)
	}

	if err := router.Run(fmt.Sprintf("%s:%s", host, port)); err != nil {
		log.Fatal(err)
	}
}
