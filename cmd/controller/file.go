package controller

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/turbovladimir/csv-parser/pkg/db"
	"github.com/turbovladimir/csv-parser/pkg/db/model"
	"github.com/turbovladimir/csv-parser/pkg/file/csv"
	"net/http"
	"strconv"
	"time"
)

type FileParseRequest struct {
	FilePath string `json:"file_path"`
}

func FileParse(c *gin.Context) {
	// Declare a new Person struct.
	var r FileParseRequest

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(c.Request.Body).Decode(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	p := csv.Parser{}
	rows := p.ParseFile(r.FilePath)
	var contacts []*model.Contact
	added := time.Now()

	for _, row := range rows[:1] {
		h := md5.New()
		h.Write([]byte(row[1] + strconv.FormatInt(added.Unix(), 10)))
		id := hex.EncodeToString(h.Sum(nil))[0:15]
		contacts = append(contacts, &model.Contact{
			ContactID: id,
			Name:      row[0],
			AddedAt:   added,
			Phone:     row[1],
		})
	}

	res := db.NewConnection().Create(contacts)

	c.JSON(http.StatusOK, gin.H{
		"success":       true,
		"rows_affected": res.RowsAffected,
		"insert_errors": res.Error,
	})
}
