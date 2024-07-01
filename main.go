package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/turbovladimir/csv-parser/cmd/controller"
)

func main() {

	godotenv.Load(".env")

	router := gin.Default()
	router.POST("/file/parse", controller.FileParse)
	router.POST("/model/gen/all", controller.GenerateModels)
	router.POST("/model/gen/:model", controller.GenerateModels)
	router.Run()
}
