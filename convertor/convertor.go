package convertor

import (
	"github.com/ZJUSCT/MirrorZ-Shim/models"
	"github.com/sirupsen/logrus"
)

func Convert(mirrorData []models.ZjuMirror) *models.MirrorZ {
	var cnameMapper map[string]string
	if err := readJson(&cnameMapper, "./configs/mirrorz-cname.json"); err != nil {
		logrus.Error(err)
		return nil
	}

	var mirrorzExt models.MirrorZExtension
	if err := readJson(&mirrorzExt, "./configs/extension.json"); err != nil {
		logrus.Error(err)
		return nil
	}

	var mirrorzSite models.MirrorzSite
	if err := readJson(&mirrorzSite, "./configs/site-metadata.json"); err != nil {
		logrus.Error(err)
		return nil
	}

	data := new(models.MirrorZ)
	data.Version = 1.6 // FIXME: do not hard code
	data.Site = mirrorzSite
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
		if v.IndexFileType != "none" {
			mirrorzInfo = append(
				mirrorzInfo,
				models.MirrorzInfo{Distro: v.Id, Category: v.IndexFileType, Urls: v.Files},
			)
		}
	}
	return mirrorzInfo
}

func convertToMirrorzMirrors(mirrorData []models.ZjuMirror) []models.MirrorzMirror {
	var mirrorzMirror []models.MirrorzMirror
	for _, v := range mirrorData {
		statusMapper := map[string]string{
			"succeeded": "S",
			"syncing"  : "Y",
			"failed"   : "F",
			"pending"  : "D",
		}
		var status = "U"
		switch v.Status {
		case "succeeded", "syncing", "failed", "pending":
			status = statusMapper[v.Status] + v.LastUpdated + "X" + v.NextScheduled
			if v.Status != "succeeded" {
				// if a mirror status is succeeded, then no lastSuccess time will be reported
				status += "O" + v.LastSuccess
			}
		case "paused":
			status = "P"
		case "cached":
			status = "C"
		case "reverseProxied":
			status = "R"
			//case "unknown":
			//	status = "U"
			//default:
			//	status = "U"
		}
		mirrorzMirror = append(
			mirrorzMirror,
			models.MirrorzMirror{Cname: v.Id, Desc: v.Desc.Zh, URL: v.Url, Help: v.HelpUrl, Upstream: v.Upstream, Size: v.Size, Status: status},
		)
	}
	return mirrorzMirror
}
