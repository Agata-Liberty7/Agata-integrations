package clients

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	db = make(map[string]ClientBundles)
	// dbTimestemps = make(map[string]string)
)

type StatsType struct {
	Count uint32 `json:"count"`
	Time  string `json:"time"`
}

func InitRoutes(r *gin.Engine) {
	f := r.Group("/clients")

	db["0"] = nil
	db["1"] = nil
	db["2"] = nil

	f.POST("/:clientID", func(c *gin.Context) {
		clientID := c.Params.ByName("clientID")

		var bundle ClientBundle
		err := c.BindJSON(&bundle)
		if err != nil {
			log.Panic(err)
		}

		bundle.Timestamp = time.Now().Format(time.RFC3339)
		db[clientID] = append(db[clientID], bundle)

		c.JSON(http.StatusOK, nil)
	})

	f.GET("/:clientID", func(c *gin.Context) {
		clientID := c.Params.ByName("clientID")
		resp, ok := getBundle(clientID)

		if ok {
			c.JSON(http.StatusOK, resp)
		} else {
			c.JSON(http.StatusOK, nil)
		}
	})

	f.GET("/", func(c *gin.Context) {
		result := []ClientResponse{}

		for clientID := range db {
			resp, ok := getBundle(clientID)
			if ok {
				result = append(result, resp)
			}
		}

		if len(result) > 0 {
			c.JSON(http.StatusOK, result)
		} else {
			c.JSON(http.StatusOK, nil)
		}
	})
}

func getBundle(clientID string) (resp ClientResponse, ok bool) {
	value, ok := db[clientID]

	resp = ClientResponse{
		ClientID: clientID,
		Items:    []ClientResponseMeta{},
	}

	resp.Data = make(map[string]string)

	for _, bundle := range value {
		resp.Items = append(resp.Items, ClientResponseMeta{
			Timestamp: bundle.Timestamp,
			Status:    "Данные получены",
		})
		for _, field := range bundle.Fields {
			resp.Data[field.Map] = field.Value
		}
	}

	return
}
