package router

import (
	"KeyValTest/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	_"io/ioutil"
	"io/ioutil"
)

//do we want this package global?
//TODO: temporarily global public to access for testing
var KeyValDAL model.KeyValDALInterface

//Helper to create gin Engine
func CreateServiceHandlers() *gin.Engine  {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.Data(200, "text/plain", []byte("OK"))
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/key/:key", getKeyVal)
	r.GET("/keys", getAllKeyVals)

	r.POST("/key", createKeyVal)
	r.PUT("/key/:key", updateKeyVal)
	r.DELETE("/key/:key", deleteKeyVal)

	return r
}

func StartService(hostingIP string, hostingPort int, dal *model.KeyValDAL) {
	KeyValDAL = dal

	r := CreateServiceHandlers()
	ipPort := fmt.Sprintf("%s:%d", hostingIP, hostingPort)
	r.Run(ipPort)

}

func getKeyVal(c *gin.Context) {
	key := c.Param("key")

	keyValData, err := KeyValDAL.GetKeyVal(key)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else if keyValData != nil {
		c.JSON(200, keyValData)
	} else {
		c.JSON(404, gin.H{"error": "Not Found"})
	}

}

func getAllKeyVals(c *gin.Context) {
	keyValArray, err := KeyValDAL.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else if keyValArray != nil {
		c.JSON(200, keyValArray)
	} else {
		c.JSON(404, gin.H{"error": "Not Found"})
	}
}

func createKeyVal(c *gin.Context) {
	var newKeyVal model.KeyValData

	if c.BindJSON(&newKeyVal) == nil {
		keyValData, err := KeyValDAL.CreateKeyVal(newKeyVal.Key, newKeyVal.Value)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		} else if keyValData != nil {
			c.JSON(200, keyValData)
		} else {
			c.JSON(500, gin.H{"error": "Unable to create"})
		}
	} else{
		if b, err := ioutil.ReadAll(c.Request.Body); err == nil {
			log.Printf("Received: %s",string(b))
		}
		log.Print("Was unable to bind to model")

		c.JSON(500, gin.H{"error": "Invalid JSON"})
	}
}

func updateKeyVal(c *gin.Context) {
	key := c.Param("key")
	var updateKeyVal model.KeyValData
	if c.BindJSON(&updateKeyVal) == nil {
		if key != updateKeyVal.Key {
			c.JSON(500, gin.H{"error": "Key in URI does not match post parameter"})
			return
		}

		keyValData, err := KeyValDAL.UpdateKeyVal(updateKeyVal.Key, updateKeyVal.Value)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		} else if keyValData != nil {
			c.JSON(200, keyValData)
		} else {
			c.JSON(500, gin.H{"error": "Unable to update"})
		}
	}
}

func deleteKeyVal(c *gin.Context) {
	key := c.Param("key")
	deleted, err := KeyValDAL.DeleteKeyVal(key)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else if deleted == false {
		c.JSON(500, gin.H{"error": "Unable to delete"})
	} else if deleted == true {
		c.JSON(200, deleted)
	} else {
		c.JSON(500, gin.H{"error": "Unable to delete"})
	}
}
