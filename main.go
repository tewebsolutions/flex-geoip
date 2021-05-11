package main

import (
	"log"
	"net"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
)

func main() {
	db, err := geoip2.Open("GeoLite2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()
	r.GET("/country/:country_code", func(c *gin.Context) {
		country_code := c.Param("country_code")
		ip := c.ClientIP()
		ipObj := net.ParseIP(ip)

		record, err := db.Country(ipObj)
		if err != nil {
			c.JSON(404, gin.H{
				"message": "Not found",
			})
		}
		if strings.EqualFold(country_code, record.Country.IsoCode) {
			c.JSON(200, gin.H{
				"message": "Matches country code",
			})
		} else {
			c.JSON(503, gin.H{
				"message": "Does not match country code",
			})
		}
	})
	r.Run()
}
