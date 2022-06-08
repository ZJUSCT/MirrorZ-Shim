package models

type MirrorZ struct {
	Version   float64         `json:"version,omitempty"`
	Site      MirrorzSite     `json:"site"`
	Info      []MirrorzInfo   `json:"info"`
	Mirrors   []MirrorzMirror `json:"mirrors"`
	Extension string          `json:"extension"`
	Endpoints []struct {
		Label   string        `json:"label"`
		Public  bool          `json:"public"`
		Resolve string        `json:"resolve"`
		Filter  []string      `json:"filter"`
		Range   []interface{} `json:"range"`
	} `json:"endpoints"`
}

type MirrorzSite struct {
	URL          string `json:"url"`
	Logo         string `json:"logo,omitempty"`
	LogoDarkmode string `json:"logo_darkmode,omitempty"`
	Abbr         string `json:"abbr"`
	Name         string `json:"name,omitempty"`
	Homepage     string `json:"homepage,omitempty"`
	Issue        string `json:"issue,omitempty"`
	Request      string `json:"request,omitempty"`
	Email        string `json:"email,omitempty"`
	Group        string `json:"group,omitempty"`
	Disk         string `json:"disk,omitempty"`
	Note         string `json:"note,omitempty"`
	Big          string `json:"big,omitempty"`
}

type MirrorzMirror struct {
	Cname    string `json:"cname"`
	Desc     string `json:"desc"`
	URL      string `json:"url"`
	Status   string `json:"status"`
	Help     string `json:"help"`
	Upstream string `json:"upstream"`
	Size     string `json:"size,omitempty"`
}

type MirrorzInfo struct {
	Distro   string `json:"distro"`
	Category string `json:"category"`
	Urls     []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"urls"`
}
