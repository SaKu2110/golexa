package golexa

type RequestBody struct {
	Version	string		`json:"version"`
	Session	Session		`json:"session,omitempty"`
	Context	Status		`json:"context"`
	Request	struct {
		Type	string	`json:"type"`
		Intent	Intent	`json:"intent,omitempty"`
	}	`json:"request"`
}

type Session struct {
	New			bool					`json:"new"`
	SessionId	string					`json:"sessionId"`
	Attributes	map[string]interface{}	`json:"attributes"`
	Application	Application				`json:"application"`
	User		User					`json:"user"`
}

type Status struct {
	System		System		`json:"System"`
	AudioPlayer	AudioPlayer	`json:"AudioPlayer"`
}

type Application struct {
	Id	string	`json:"applicationId"`
}

type User struct {
	Id		string	`json:"userId"`
	Token	string	`json:"accessToken"`
}

type System struct {
	ApiAccessToken	string		`json:"apiAccessToken"`
	ApiEndpoint		string		`json:"apiEndpoint"`
	Application		Application	`json:"application"`
	Device			Device		`json:"device"`
	Person			Person		`json:"person"`
	User			User		`json:"user"`
}

type Device struct {
	Id	string	`json:"deviceId"`
}

type Person struct {
	Id	string			`json:"personId"`
	AccessToken	string	`json:"accessToken"`
}

type AudioPlayer struct {
	Token					string	`json:"token"`
	OffsetInMilliseconds	uint64	`json:"offsetInMilliseconds"`
	PlayerActivity			string	`json:"playerActivity"`
}

type Intent struct {
	Name	string			`json:"name,omitempty"`
	Slots	map[string]Slot	`json:"slots,omitempty"`
}

type Slot struct {
	Name	string	`json:"name"`
	Value	string	`json:"value"`
}
