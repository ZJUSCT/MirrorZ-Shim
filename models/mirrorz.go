package models

type MirrorZ struct {
	Version float64 `json:"version,omitempty"`
	Site    struct {
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
	} `json:"site"`
	Info []struct {
		Distro   string `json:"distro"`
		Category string `json:"category"`
		Urls     []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"urls"`
	} `json:"info"`
	Mirrors []struct {
		Cname    string `json:"cname"`
		Desc     string `json:"desc"`
		URL      string `json:"url"`
		Status   string `json:"status"`
		Help     string `json:"help"`
		Upstream string `json:"upstream"`
		Size     string `json:"size,omitempty"`
	} `json:"mirrors"`
}
