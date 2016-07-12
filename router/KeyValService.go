package router

import (
	"KeyValTest/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

//do we want this package global?
var keyValDAL model.KeyValDALInterface

func StartService(hostingIP string, hostingPort int, dal *model.KeyValDAL) {

	keyValDAL = dal

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

	ipPort := fmt.Sprintf("%s:%d", hostingIP, hostingPort)
	r.Run(ipPort)

}

func getKeyVal(c *gin.Context) {
	key := c.Param("key")

	keyValData, err := keyValDAL.GetKeyVal(key)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else if keyValData != nil {
		c.JSON(200, keyValData)
	} else {
		c.JSON(404, gin.H{"error": "Not Found"})
	}

}

func getAllKeyVals(c *gin.Context) {
	keyValArray, err := keyValDAL.GetAll()
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
		keyValData, err := keyValDAL.CreateKeyVal(newKeyVal.Key, newKeyVal.Value)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		} else if keyValData != nil {
			c.JSON(200, keyValData)
		} else {
			c.JSON(500, gin.H{"error": "Unable to create"})
		}
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

		keyValData, err := keyValDAL.UpdateKeyVal(updateKeyVal.Key, updateKeyVal.Value)
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
	deleted, err := keyValDAL.DeleteKeyVal(key)
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
