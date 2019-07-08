package golexa

type request struct {
	Request struct {
		Type                       string    `json:"type"`
		Intent struct {
			Name	string `json: name`
			Slots struct {
				Slot struct {
					Value	string	`json: value`
				}	`json: slot`
			}	`json: slots`
		} `json: Intent`
	} `json:"request"`
}
