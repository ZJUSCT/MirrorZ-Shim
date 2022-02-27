package models

type MirrorZExtension struct {
	Extension string `json:"extension"`
	Endpoints []struct {
		Label   string        `json:"label"`
		Public  bool          `json:"public"`
		Resolve string        `json:"resolve"`
		Filter  []string      `json:"filter"`
		Range   []interface{} `json:"range"`
	} `json:"endpoints"`
}
