package golexa

type Request struct {
	Version	string	`json:"version"`
	Session	session	`json:"session"`
	Context	context `json:"context"`
	Request	request	`json:"request"`
}

type session struct {
	New			bool					`json:"new"`
	SessionId	string					`json:"sessionId"`
	Attributes	map[string]interface{}	`json:"attributes"`
	Application	application				`json:"application"`
	User		user					`json:"user"`
}

type application struct {
	ApplicationId	string	`json:"applicationId"`
}

type user struct {
	UserId		string	`json:"userId"`
	AccessToken	string	`json:"accessToken"`
}

type context struct {
	System		system		`json:"System"`
	AudioPlayer	audioPlayer	`json:"AudioPlayer"`
}

type system struct {
	ApiAccessToken	string		`json:"apiAccessToken"`
	ApiEndpoint		string		`json:"apiEndpoint"`
	Application		application	`json:"application"`
	Device			device		`json:"device"`
	Person			person		`json:"person"`
	User			user		`json:"user"`
}

type device struct {
	DeviceId	string	`json:"deviceId"`
}

type person struct {
	PersonId	string	`json:"deviceId"`
	AccessToken	string	`json:"accessToken"`
}

type audioPlayer struct {
	Token					string	`json:"token"`
	OffsetInMilliseconds	string	`json:"offsetInMilliseconds"`
	PlayerActivity			string	`json:"playerActivity"`
}

type request struct {
	Type	string	`json:"type"`
	Intent	intent	`json:"intent"`
}

type intent struct {
	Name	string			`json:"name"`
	Slots	map[string]slot	`json:"slots"`
}

type slot struct {
	Name	string	`json:"name"`
	Version	string	`json:"value"`
}