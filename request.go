package golexa

import(
	"time"
)

type request struct {
	Request struct {
		Type                       string    `json:"type"`
		RequestID                  string    `json:"requestId"`
		Timestamp                  time.Time `json:"timestamp"`
		Locale                     string    `json:"locale"`
		ShouldLinkResultBeReturned bool      `json:"shouldLinkResultBeReturned"`
		Intent struct {
			Name	string `json: name`
			ConfirmationStatus	string `json: confirmationStatus`
		} `json: Intent`
	} `json:"request"`
}
