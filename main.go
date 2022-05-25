package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type control struct {
	Mode       string `json:"mode"`
	Brightness int    `json:"brightness"`
	Color      string `json:"color"`
}

type sensor struct {
	Brightness int    `json:"brightness"`
	Color      string `json:"color"`
}

var data_sensor = sensor{
	Brightness: 80,
	Color:      "#ebb552",
}

var data_control = control{
	Mode:       "auto",
	Brightness: 80,
	Color:      "#ebb552",
}

func getLEDstatus(context *gin.Context) {
	if data_control.Mode == "off" {
		var tmp = sensor{
			Brightness: 0,
			Color:      "#000000",
		}
		context.IndentedJSON(http.StatusOK, tmp)
		return 
	}
	if data_control.Mode == "normal" {
		var tmp = sensor{
			Brightness: data_control.Brightness,
			Color:      data_control.Color,
		}
		context.IndentedJSON(http.StatusOK, tmp)
		return
	}
	if data_control.Mode == "auto" {
		var tmp = sensor{
			Brightness: data_sensor.Brightness,
			Color:      data_sensor.Color,
		}
		context.IndentedJSON(http.StatusOK, tmp)
		return
	}
}

func postSensor(context *gin.Context) {
	var temp sensor
	err := context.BindJSON(&temp)
	if err != nil {
		context.String(http.StatusBadRequest, "Decode error! please check your JSON formating.")
		return
	}
	if temp.Brightness > 100 || temp.Brightness < 0 {
		context.String(http.StatusBadRequest, "Wrong data! brightness in json must be int between 0 to 100")
		return
	}
	if len(temp.Color) != 7 || temp.Color[0] != '#' {
		context.String(http.StatusBadRequest, "Wrong data! color in json must be start by # follow by 6 charater")
		return
	}
	data_sensor = temp
	context.IndentedJSON(http.StatusCreated, data_sensor)
}

func postData(context *gin.Context) {
	var temp control
	err := context.BindJSON(&temp)
	if err != nil {
		context.String(http.StatusBadRequest, "Decode error! please check your JSON formating.")
		return
	}
	if temp.Mode != "auto" && temp.Mode != "off" && temp.Mode != "normal" {
		context.String(http.StatusBadRequest, "Wrong data! mode in json must be auto or normal or off in all lowercase only")
		return
	}
	if temp.Brightness > 100 || temp.Brightness < 0 {
		context.String(http.StatusBadRequest, "Wrong data! brightness in json must be int between 0 to 100")
		return
	}
	if len(temp.Color) != 7 || temp.Color[0] != '#' {
		context.String(http.StatusBadRequest, "Wrong data! color in json must be start by # follow by 6 charater")
		return
	}

	data_control = temp
	context.IndentedJSON(http.StatusCreated, data_control)
}

func main() {
	router := gin.Default()
	router.GET("/led", getLEDstatus)
	router.POST("/control", postData)
	router.POST("/sensor", postSensor)
	router.Run("localhost:8080")
}
