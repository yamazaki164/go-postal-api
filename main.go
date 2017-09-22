package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo"
	"github.com/yamazaki164/go-postal/postal"
)

var (
	config *Config
)

func endpointAction(c echo.Context) error {
	postalCode := c.QueryParam("code")
	if len(postalCode) < 3 {
		return c.JSON(http.StatusBadRequest, nil)
	}

	postalCodeShort := postalCode[0:3]

	file := filepath.Join(config.JsonDir, postalCodeShort+".json")
	b, e := ioutil.ReadFile(file)
	if e != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	var p postal.AreaPostal
	if e := json.Unmarshal(b, &p); e != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	rec, err := p[postalCode]
	if !err {
		return c.JSON(http.StatusBadRequest, nil)
	}

	return c.JSON(http.StatusOK, rec)
}

func main() {
	configFileOpt := flag.String("c", "./server.conf", "/path/to/config/file")
	flag.Parse()

	var err error
	config, err = loadToml(*configFileOpt)
	if err != nil {
		fmt.Println(config)
		panic(err)
	}

	if !config.IsValidConfig() {
		panic(errors.New("invalid config file"))
	}

	e := echo.New()
	e.GET(config.Endpoint, endpointAction)
	e.Logger.Fatal(e.Start(config.BindAddress()))
}
