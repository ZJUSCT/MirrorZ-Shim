package main

import (
	"encoding/json"
	"fmt"
	"github.com/ZJUSCT/MirrorZ-Shim/convertor"
	"github.com/ZJUSCT/MirrorZ-Shim/models"
	"github.com/dgraph-io/ristretto"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"time"
)

func main() {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 100,
		MaxCost:     1000,
		BufferItems: 64,
	})
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		e.Logger.Info("Incoming ping request")
		return c.String(http.StatusOK, "[INFO] ZJU-Mirror MirrorZ Shim Working.")
	})
	e.GET("/mirrorz.json", func(c echo.Context) error {
		e.Logger.Info("Incoming mirrorz.json request")

		var data *models.MirrorZ
		val, found := cache.Get("mirrorz")
		if !found {
			fmt.Println("no cache")
			resp, err := http.Get("https://mirrors.zju.edu.cn/api/mirrors")
			if err != nil {
				return err
			}
			defer func(Body io.ReadCloser) {
				_ = Body.Close()
			}(resp.Body)

			body, err := io.ReadAll(resp.Body)
			mirrorData := models.Mirror{}
			err = json.Unmarshal(body, &mirrorData)
			if err != nil {
				return err
			}

			data = convertor.Convert(mirrorData)

			cache.SetWithTTL("mirrorz", data, 1, time.Minute*5)
		} else {
			fmt.Println("has cache")
			data = val.(*models.MirrorZ)
		}

		if err := c.Bind(data); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, data)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
