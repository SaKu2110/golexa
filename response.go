package golexa

type ResponseBody struct {
	Version		string						`json:"version"`
	Attributes	map[string]interface{}		`json:"sessionAttributes,omitempty"`
	Response	*Response					`json:"response"`
}

type Response struct {
	OutputSpeech		*OutputSpeech	`json:"outputSpeech,omitempty"`
	Card				*Card			`json:"card,omitempty"`
	Reprompt			*Reprompt		`json:"reprompt,omitempty"`
	ShouldEndSession	bool			`json:"shouldEndSession"`
}

type OutputSpeech struct {
	Type			string	`json:"type"`
	Text			string	`json:"text,omitempty"`
	SSML			string	`json:"ssml,omitempty"`			
	PlayBehavior	string	`json:"playBehavior,omitempty"`
}

type Card struct {
	Type	string	`json:"type"`
	Title	string	`json:"title,omitempty"`
	Content	string	`json:"content,omitempty"`
	Text	string	`json:"text,omitempty"`
	Image	*Image	`json:"image,omitempty"`
}

type Image struct {
	SmallImageUrl	string	`json:"smallImageUrl,omitempty"`
	LargeImageUrl	string	`json:"largeImageUrl,omitempty"`
}

type Reprompt struct {
	OutputSpeech	*OutputSpeech	`json:"outputSpeech,omitempty"`
}