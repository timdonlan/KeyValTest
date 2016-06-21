package router

import (
	"github.com/gin-gonic/gin"
	"KeyValTest/model"
)

func StartService(){
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/key/:key",getKeyVal)
	r.GET("/keys",getAllKeyVals)
	r.POST("/key", createKeyVal)


	r.Run() // listen and server on 0.0.0.0:8080
}

func getKeyVal(c *gin.Context) {
	key := c.Param("key")

	//c.JSON(200, key)
	//return

	keyValData, err := model.GetKeyVal(key)
	if (err != nil) {
		c.JSON(500, gin.H{"error":err})
	}else if keyValData != nil {
		c.JSON(200, keyValData)
	} else {
		c.JSON(404, gin.H{"error": "Not Found"})
	}

}

func getAllKeyVals(c *gin.Context){
	keyValArray, err := model.GetAllKeyVal()
	if (err != nil) {
		c.JSON(500, gin.H{"error":err})
	} else if keyValArray != nil {
		c.JSON(200, keyValArray)
	} else {
		c.JSON(404, gin.H{"error": "Not Found"})
	}

}

func createKeyVal(c *gin.Context){
	var newKeyVal model.KeyValData

	if c.BindJSON(&newKeyVal) == nil {
		keyValData, err := model.CreateKeyVal(newKeyVal.Key, newKeyVal.Value)
		if (err != nil) {
			c.JSON(500, gin.H{"error":err})
		} else if keyValData != nil {
			c.JSON(200, keyValData)
		} else {
			c.JSON(500, gin.H{"error": "Unable to create"})
		}
	}
}

func getAllKeyVal(c *gin.Context){

}