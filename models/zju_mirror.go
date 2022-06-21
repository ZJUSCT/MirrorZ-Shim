package models

type ZjuMirror struct {
	Id   string `json:"id"`
	Url  string `json:"url"`
	Name struct {
		Zh string `json:"zh"`
		En string `json:"en"`
	} `json:"name"`
	Desc struct {
		Zh string `json:"zh"`
		En string `json:"en"`
	} `json:"desc"`
	HelpUrl       string `json:"helpUrl"`
	Upstream      string `json:"upstream"`
	Size          string `json:"size"`
	IndexFileType string `json:"type"`
	Status        string `json:"status"`
	LastUpdated   string `json:"lastUpdated"`
	NextScheduled string `json:"nextScheduled"`
	LastSuccess   string `json:"lastSuccess"`
	Files         []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
}
