package main

import (
	"encoding/json"
	"github.com/ZJUSCT/MirrorZ-Shim/models"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		e.Logger.Info("Incoming ping request")
		return c.String(http.StatusOK, "[INFO] ZJU-Mirror MirrorZ Shim Working.")
	})
	e.GET("/mirrorz.json", func(c echo.Context) error {
		e.Logger.Info("Incoming mirrorz.json request")

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

		data := new(models.MirrorZ)
		data.Site = mirrorData.Site
		data.Version = mirrorData.Version
		data.Info = mirrorData.Info
		data.Mirrors = mirrorData.Mirrors

		if err := c.Bind(data); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, data)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
