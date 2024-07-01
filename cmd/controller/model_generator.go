package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/turbovladimir/csv-parser/pkg/db"
	"net/http"
)

type GenerateModelsRequest struct {
}

func GenerateModels(c *gin.Context) {
	var r GenerateModelsRequest

	err := json.NewDecoder(c.Request.Body).Decode(&r)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	}

	g := db.ModelGenerator{}
	g.GenerateAll()
}

type GenerateModelRequest struct {
	ModelName string `json:"model_name"`
}

func GenerateModel(c *gin.Context) {
	var r GenerateModelRequest

	err := json.NewDecoder(c.Request.Body).Decode(&r)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	}

	g := db.ModelGenerator{}
	g.GenerateModel(r.ModelName)
}
