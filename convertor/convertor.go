package convertor

import (
	"github.com/ZJUSCT/MirrorZ-Shim/models"
	"github.com/sirupsen/logrus"
)

func Convert(mirrorData []models.ZjuMirror) *models.MirrorZ {
	var cnameMapper map[string]string
	if err := readJson(cnameMapper, "./configs/mirrorz-cname.json"); err != nil {
		logrus.Error(err)
		return nil
	}

	var mirrorzExt models.MirrorZExtension
	if err := readJson(mirrorzExt, "./configs/extension.json"); err != nil {
		logrus.Error(err)
		return nil
	}

	data := new(models.MirrorZ)
	data.Version = 1.5 // FIXME: do not hard core
	// data.Site = mirrorData.Site FIXME: read config file here
	data.Info = convertToMirrorzInfo(mirrorData)
	data.Mirrors = convertToMirrorzMirrors(mirrorData)
	data.Extension = mirrorzExt.Extension
	data.Endpoints = mirrorzExt.Endpoints

	// Do convert
	for i := range data.Info {
		cname := cnameMapper[data.Info[i].Distro]
		if cname != "" {
			data.Info[i].Distro = cname
		} else {
			logrus.
				WithField("zju_distro_name", data.Info[i].Distro).
				Warn("key not found in cname map")
		}
	}
	for i := range data.Mirrors {
		cname := cnameMapper[data.Mirrors[i].Cname]
		if cname != "" {
			data.Mirrors[i].Cname = cname
		} else {
			logrus.
				WithField("zju_mirror_cname", data.Mirrors[i].Cname).
				Warn("key not found in cname map")
		}
	}
	return data
}

func convertToMirrorzInfo(mirrorData []models.ZjuMirror) []models.MirrorzInfo {
	var mirrorzInfo []models.MirrorzInfo
	for _, v := range mirrorData {
		mirrorzInfo = append(
			mirrorzInfo,
			models.MirrorzInfo{Distro: v.Name.Zh, Category: "os", Urls: v.Files}, // FIXME: do not hardcode os
		)
	}
	return mirrorzInfo
}

func convertToMirrorzMirrors(mirrorData []models.ZjuMirror) []models.MirrorzMirror {
	var mirrozMirror []models.MirrorzMirror
	for _, v := range mirrorData {
		mirrozMirror = append(
			mirrozMirror,
			models.MirrorzMirror{Cname: v.Name.Zh, Desc: v.Desc.Zh, URL: v.Url, Help: v.HelpUrl, Upstream: v.Upstream},
		)
	}
	return mirrozMirror
}
