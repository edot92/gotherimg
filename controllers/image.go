package controllers

import (
	"encoding/json"
	"fmt"
	"image"
	"net/http"
	"os"
	"strconv"

	"github.com/danbondd/tempconv"
	"github.com/gin-gonic/gin"
	colorful "github.com/lucasb-eyer/go-colorful"
	"github.com/maruel/temperature"
)

type ResponseJSON struct {
	status  string
	data    interface{}
	message string
}

//GetImageValue
func GetImageValue(c *gin.Context) {
	pointx, _ := strconv.Atoi(c.Param("pointx"))
	pointy, _ := strconv.Atoi(c.Param("pointy"))
	message := "pointx:" + strconv.Itoa(pointx) + " dan pointy :" + strconv.Itoa(pointy)
	imgfile, err := os.Open("static/tes.jpg")

	if err != nil {
		fmt.Println("img.jpg file not found!")
		os.Exit(1)
	}
	defer imgfile.Close()
	imgCfg, _, err := image.DecodeConfig(imgfile)

	if err != nil {
		fmt.Println("image.DecodeConfig", err)
		os.Exit(1)
	}

	width := imgCfg.Width
	height := imgCfg.Height

	fmt.Println("Width : ", width)
	fmt.Println("Height : ", height)

	// we need to reset the io.Reader again for image.Decode() function below to work
	// otherwise we will  - panic: runtime error: invalid memory address or nil pointer dereference
	// there is no build in rewind for io.Reader, use Seek(0,0)
	imgfile.Seek(0, 0)

	// get the image
	img, _, err := image.Decode(imgfile)
	var w colorful.Color
	r, g, b, _ := img.At(pointx, pointx).RGBA()

	w.R = float64(r)
	w.G = float64(g)
	w.B = float64(b)

	warna := colorful.Color(w)
	hex := warna.Hex()
	type temperatureS struct {
		Celcius    tempconv.Celsius
		Fahrenheit tempconv.Fahrenheit
	}
	var tempconvV temperatureS
	message = message + " hex code :" + hex

	var res ResponseJSON
	res.status = "success"
	type bit8ya struct {
		r uint8
		g uint8
		b uint8
	}
	var ww bit8ya
	ww.r = hilo(int(w.R))
	ww.g = hilo(int(w.G))
	ww.b = hilo(int(w.B))
	d := map[string]interface{}{"Hex": hex, "RGB": ww, "Suhu:": tempconvV}
	_ = d
	www, _ := json.Marshal(ww)
	res.data = &www
	_ = www
	fmt.Println("res:", res)

	var kelvin tempconv.Kelvin = tempconv.Kelvin(temperature.ToKelvin((ww.r), (ww.g), (ww.b)))
	tempconvV.Celcius = tempconv.KelvinToCelcius(kelvin)
	tempconvV.Fahrenheit = tempconv.Fahrenheit(kelvin)
	fmt.Println(tempconvV)
	fmt.Println("kelvin:", kelvin)
	c.JSON(http.StatusOK, gin.H{"Hex": hex, "R": ww.r, "G": ww.g, "B": ww.b, "Suhu:": tempconvV})

}

func hilo(intValue int) uint8 {
	ucHigh := (uint8)((intValue >> 0x08) & 0xff)
	ucLow := (uint8)(intValue & 0xff)
	_ = ucHigh
	// fmt.Println(ucHigh)
	// fmt.Println(ucLow)
	// return strconv.Itoa(ucLow)
	return ucLow

}
