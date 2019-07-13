package golexa

type request struct {
	Request struct {
		Type	string `json:"type"`
		Intent struct {
			Name	string `json: name`
			Slots interface{} `json: slots`
		} `json: Intent`
	} `json:"request"`
}
