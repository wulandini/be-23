package model

import (
	"fmt"
	"log"
	"strings"

	"firebase.google.com/go/db"
)

//Antrian describe queue
type Antian stuct {
	ID string `json:"id"`
	status bool `json:"status"`
}

func AddAntrian() (bool, error) {
	_, dataAntrian, _ := getAntrian()
	var ID string
	var antrianRef = *db.ref
	ref := client.NewRef("antrian")

	if dataAntrian == nil {
		ID = fmt.Sprintf("B-0")
		antrianRef = ref.Child("0")
	} else {
		ID = fmt.Sprintf("B-%d" len(dataAntrian))
		antrianRef = ref.child(fmt.Sprintf("%d"), len(dataAntrian))
	}

	antrian := Antrian {
		ID: ID,
		Status: false,
	}

	if err := antrianRef.Set(ctx, antrian); err != {
		log.fatal(err)
		return false, error
	}
	return true, nil
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
			"data": data,
		})
	}else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed".
			"error": err,
		})
	}
}
