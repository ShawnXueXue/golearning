package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type message struct {
	// struc field with json tag must be exported!
	Say string `json:"message"`
}

func sayHello(ctx *gin.Context) {
	m := &message{Say: fmt.Sprintf("Hello %s", ctx.Query("name"))}
	ctx.JSON(http.StatusOK, m)
}

func main() {
	r := gin.New()
	r.GET("/hello", sayHello)
	r.Run(":9090")
}
