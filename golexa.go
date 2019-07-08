package golexa

import(
	"github.com/gin-gonic/gin"
	"encoding/json"
)

type alexa struct{
	FUNC	func(string, string)string
}

func Default(f func(string, string)string) *alexa{
	init := &alexa{FUNC: f}
	return init
}

func (alexa *alexa)CallHandler(g *gin.Context){
	var request_type string
	req := request{}
	g.BindJSON(&req)
	request_type = req.Request.Type
	if request_type == "IntentRequest" {
		request_type = req.Request.Intent.Name
	}
	json := alexa.FUNC(request_type, req.Request.Intent.Slots.Slot.Value)
	g.JSON(200, json)
}

func Ask(msg string) string{
	res := response{}

	// データ入力
        res.Version = "1.0"
        res.Response.OutputSpeech.Type = "PlainText"
        res.Response.OutputSpeech.Text = msg
        res.Response.ShouldEndSession = false
        res.Response.Type = "_DEFAULT_RESPONSE"

        byte_code, _ := json.Marshal(res)
        json := string(byte_code)

	return json
}

func Tell(msg string) string{
	res := response{}

	// データ入力
        res.Version = "1.0"
        res.Response.OutputSpeech.Type = "PlainText"
        res.Response.OutputSpeech.Text = msg
        res.Response.ShouldEndSession = true
        res.Response.Type = "_DEFAULT_RESPONSE"

        byte_code, _ := json.Marshal(res)
        json := string(byte_code)


	return json
}
