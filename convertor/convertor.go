package convertor

import (
	"encoding/json"
	"fmt"
	"github.com/ZJUSCT/MirrorZ-Shim/models"
	"io/ioutil"
)

func Convert(mirrorData models.Mirror) *models.MirrorZ {
	var cnameMapper map[string]string
	file, err := ioutil.ReadFile("./configs/mirrorz-cname.json")
	if err != nil {
		fmt.Println("Read cname.json failed")
		return nil
	}
	err = json.Unmarshal(file, &cnameMapper)
	if err != nil {
		fmt.Println("Unmarshal cname.json failed")
		return nil
	}

	var mirrorzExt models.MirrorZExtension
	file, err = ioutil.ReadFile("./configs/extension.json")
	if err != nil {
		fmt.Println("Read extension.json failed")
		return nil
	}
	err = json.Unmarshal(file, &mirrorzExt)
	if err != nil {
		fmt.Println("Unmarshal extension.json failed")
		return nil
	}

	data := new(models.MirrorZ)
	data.Site = mirrorData.Site
	data.Version = mirrorData.Version
	data.Info = mirrorData.Info
	data.Mirrors = mirrorData.Mirrors
	data.Extension = mirrorzExt.Extension
	data.Endpoints = mirrorzExt.Endpoints

	// Do convert
	for i := range data.Info {
		cname := cnameMapper[data.Info[i].Distro]
		if cname != "" {
			data.Info[i].Distro = cname
		} else {
			fmt.Println(data.Info[i].Distro)
		}
	}
	for i := range data.Mirrors {
		cname := cnameMapper[data.Mirrors[i].Cname]
		if cname != "" {
			data.Mirrors[i].Cname = cname
		} else {
			fmt.Println(data.Mirrors[i].Cname)
		}
	}

	return data
}
