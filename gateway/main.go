package main

import (
	"gateway/common"
	"gateway/rest"
	"gateway/translate"
	"io"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func transform(context *gin.Context) {
	input, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.AsciiJSON(200, rest.EmptyResp(10000, err.Error()))
		return
	}

	input, err = common.Decompress(input)
	if err != nil {
		context.AsciiJSON(200, rest.EmptyResp(10000, err.Error()))
		return
	}

	xml, err := TranslateClient.CreateBuilding(string(input))
	if err != nil {
		context.AsciiJSON(200, rest.EmptyResp(10000, err.Error()))
		return
	}

	data := map[string]any{
		"xml": string(xml),
	}
	resp := rest.Resp(200, "ok", data)
	context.AsciiJSON(200, resp)
}

func status(context *gin.Context) {
	resp := rest.EmptyResp(200, "ok")
	context.AsciiJSON(200, resp)
}

var TranslateClient *translate.Client

func main() {
	TranslateClient = &translate.Client{Server: "127.0.0.1:8001"}
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.POST("/transform", transform)
	router.GET("/status", status)
	router.Run("0.0.0.0:8888")
}
