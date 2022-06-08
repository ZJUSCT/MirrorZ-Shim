package models

import "time"

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
	HelpUrl       string    `json:"helpUrl"`
	Upstream      string    `json:"upstream"`
	Status        string    `json:"status"`
	LastUpdated   time.Time `json:"lastUpdated"`
	NextScheduled time.Time `json:"nextScheduled"`
	LastSuccess   time.Time `json:"lastSuccess"`
	Files         []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
}
