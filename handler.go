package controller

import {
	"net/http"

	"github.com/gin-gonic/gin"
}


/AddAntrianHandler is a function to add queue
func AddAntrianHandler(c *gin.Context) {
	flag, err := addAntrian()

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success"
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed".
			"error": err,
		})
	}
}

//GetAntrianHandler is a function to add queue
func GetAntrianHandler(c *gin.Context) {
	flag, data, err := getAntrian()

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
			"data": data,
		})
	}else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed".
			"error": err,
		})
	}
}

//UpdateAntrianHandler is a function to add queue
func UpdateAntrianHandler(c *gin.Context) {
	idAntrian := c.Param("idAntrian")
	flag, err := updateAntrian(idAntian)

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
			"data": data,
		})
	}else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed".
			"error": err,
		})
	}
}

//DeleteAntrianHandler is a function to add queue
func DeleteAntrianHandler(c *gin.Context) {
	idAntrian := c.Param("idAntrian")
	flag, err := deleteAntrian(idAntian)

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
		})
	}else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed".
			"error": err,
		})
	}
}

//PageAntrianHandler is a function to serve HTML
func PageAntrianHandler(c *gin.Context){
	flag, result, err := getAntrian()
	var currentAntrian map [string]interface{}

	for _, item := range result {
		if item != nil {
			currentAntrian = item
			break
		}
	}

	if flag && len(result) > 0 {
		c.HTML(http.statusOK, "index.html", gin.H{
			"antrian": currentAntrian["id"],
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error" : err,
		})
	}
}