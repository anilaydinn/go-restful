package main

import (
	"database/sql"
	"log"
	"net/http"

	"railApiGin/dbutils"

	"github.com/gin-gonic/gin"
	"gopkg.in/gin-gonic/gin.v1"
)

// DB driver visible to whole program
var DB *sql.DB

// StationResource holds information about locations
type StationResource struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`
}

//GetStation returns the station detail
func GetStation(c *gin.Context) {
	var station StationResource

	id := c.Param("station_id")

	err := DB.QueryRow("SELECT ID, NAME, CAST(OPENING_TIME as CHAR), CAST(CLOSING_TIME as CHAR) FROM station WHERE id=?", id).Scan(&station.ID, &station.Name, &station.OpeningTime, &station.ClosingTime)

	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"result": station,
		})
	}
}

// CreateStation handles POST
func CreateStation(c *gin.Context) {
	var station StationResource

	// Parse the body into our resource
	if err := c.BindJSON(&station); err == nil {
		//Format Time to Go time format
		statement, _ := DB.Prepare("INSERT INTO station (NAME, OPENING_TIME, CLOSING_TIME) VALUES(?,?,?)")

		result, err := statement.Exec(station.Name, station.OpeningTime, station.ClosingTime)

		if err == nil {
			newID, _ := result.LastInsertId()
			station.ID = int(newID)
			c.JSON(http.StatusOK, gin.H{
				"result": station,
			})
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

// RemoveStation handles the removing of resource
func RemoveStation(c *gin.Context) {
	id := c.Param("station_id")

	statement, _ := DB.Prepare("DELETE FROM station WHERE id=?")

	_, err := statement.Exec(id)

	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.String(http.StatusOK, "")
	}
}

func main() {
	var err error

	DB, err = sql.Open("sqlite3", "./railapi.db")

	if err != nil {
		log.Println("Driver creation failed!")
	}
	dbutils.Initialize(DB)

	r := gin.Default()

	//Add routes to REST verbs
	r.GET("/v1/stations/:station_id", GetStation)
	r.POST("/v1/stations", CreateStation)
	r.DELETE("/v1/stations/:station_id", RemoveStation)

	r.Run(":8000") //Default listen and serve on localhost:8000
}
