package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.SetTrustedProxies([]string{"192.168.1.*"})

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World, time is "+fmt.Sprint(time.Now().Unix()))
	})

	r.Run(":3000")
}
