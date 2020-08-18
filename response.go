package golexa

type Response struct {
	Version				string					`json:"version"`
	SessionAttributes	map[string]interface{}	`json:"sessionAttributes,omitempty"`
	Response			*response				`json:"response"`
}

type response struct {
	OutputSpeech		*outputSpeech `json:"outputSpeech,omitempty"`
	Card				*card         `json:"card,omitempty"`
	Reprompt			*reprompt     `json:"reprompt,omitempty"`
	ShouldEndSession	bool          `json:"shouldEndSession,omitempty"`
}

type outputSpeech struct {
	Type			string	`json:"type"`
	Text			string	`json:"text,omitempty"`
	SSML			string	`json:"ssml,omitempty"`			
	PlayBehavior	string	`json:"playBehavior,omitempty"`
}

type card struct {
	Type	string	`json:"type"`
	Title	string	`json:"title,omitempty"`
	Content	string	`json:"content,omitempty"`
	Text	string	`json:"text,omitempty"`
	Image	*image	`json:"image,omitempty"`
}

type image struct {
	SmallImageUrl	string	`json:"smallImageUrl,omitempty"`
	LargeImageUrl	string	`json:"largeImageUrl,omitempty"`
}

type reprompt struct {
	OutputSpeech	*outputSpeech	`json:"shouldEndSession,omitempty"`
}