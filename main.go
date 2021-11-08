package main

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
)

func IndexHandler(ctx *gin.Context) {
	ctx.XML(200, Person{
		FirstName: "YUAN",
		LastName:  "HAIYUE",
	})
}

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"firstName,attr"`
	LastName  string   `xml:"lastName,attr"`
}

func main() {
	router := gin.Default()
	router.GET("/", IndexHandler)
	router.Run(":8081")
}
