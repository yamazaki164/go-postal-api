package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/yamazaki164/go-postal/postal"
)

var (
	config *Config
)

type Item map[string]interface{}

func endpointAction(c echo.Context) error {
	postalCode := c.QueryParam("code")
	if len(postalCode) < 3 {
		return c.JSON(http.StatusBadRequest, Item{
			"status": http.StatusBadRequest,
		})
	}

	data, e := ioutil.ReadFile(config.JsonFile(postalCode[0:3]))
	if e != nil {
		return c.JSON(http.StatusBadRequest, Item{
			"status": http.StatusBadRequest,
		})
	}

	var p postal.AreaPostal
	if e := json.Unmarshal(data, &p); e != nil {
		return c.JSON(http.StatusBadRequest, Item{
			"status": http.StatusBadRequest,
		})
	}

	result, err := p[postalCode]
	if !err {
		return c.JSON(http.StatusBadRequest, Item{
			"status": http.StatusBadRequest,
		})
	}

	return c.JSON(http.StatusOK, Item{
		"status": http.StatusOK,
		"result": result,
	})
}

func main() {
	configFileOpt := flag.String("c", "./server.conf", "/path/to/config/file")
	flag.Parse()

	var err error
	config, err = LoadToml(*configFileOpt)
	if err != nil {
		fmt.Println(config)
		panic(err)
	}

	if !config.IsValidConfig() {
		panic(errors.New("invalid config file"))
	}

	e := echo.New()
	e.Use(middleware.Gzip())
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())
	e.GET(config.Endpoint, endpointAction)
	e.Logger.Fatal(e.Start(config.BindAddress()))
}
