package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Say string `json:"message"`
}

func SayHello(ctx *gin.Context) {
	m := &Message{Say: fmt.Sprintf("Hello %s", ctx.Query("name"))}
	ctx.JSON(http.StatusOK, m)
}

func main() {
	r := gin.New()
	r.GET("/hello", SayHello)
	r.Run(":9090")
}
