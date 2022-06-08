package main

import (
	"encoding/json"
	"github.com/ZJUSCT/MirrorZ-Shim/convertor"
	"github.com/ZJUSCT/MirrorZ-Shim/models"
	"github.com/dgraph-io/ristretto"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"time"
)

func main() {
	// setup config system
	viper.SetEnvPrefix("MIRRORZ_SHIM")
	viper.AutomaticEnv()
	viper.SetDefault("URL", "https://mirrors.zju.edu.cn/api/mirrors")
	viper.SetDefault("CACHE_TTL", 5)

	// set up cache
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 100,
		MaxCost:     1000,
		BufferItems: 64,
	})
	if err != nil {
		panic(err)
	}

	// set up web server
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.GET("/", func(c echo.Context) error {
		e.Logger.Info("Incoming ping request")
		return c.String(http.StatusOK, "[INFO] ZJU-Mirror MirrorZ Shim Working.")
	})
	e.GET("/mirrorz.json", func(c echo.Context) error {
		e.Logger.Info("Incoming mirrorz.json request")

		var data *models.MirrorZ
		val, found := cache.Get("mirrorz")
		if !found {
			e.Logger.Info("no cache")
			url := viper.GetString("URL")
			resp, err := http.Get(url)
			if err != nil {
				return err
			}
			defer func(Body io.ReadCloser) {
				_ = Body.Close()
			}(resp.Body)

			body, err := io.ReadAll(resp.Body)
			var mirrorData []models.ZjuMirror
			err = json.Unmarshal(body, &mirrorData)
			if err != nil {
				return err
			}

			data = convertor.Convert(mirrorData)

			cacheTTL := viper.GetInt("CACHE_TTL")
			cache.SetWithTTL("mirrorz", data, 1, time.Duration(cacheTTL*int(time.Minute)))
		} else {
			e.Logger.Info("has cache")
			data = val.(*models.MirrorZ)
		}

		if err := c.Bind(data); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, data)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
