package golexa

import(
	"github.com/gin-gonic/gin"
	"encoding/json"
)

type golexa struct{
	// Handle intents
	Intent	func(*Context)
}

type Context struct{
	gin		*gin.Context
	slotExist	bool
	slots		interface{}
	Intent		string
}

func (c *Context) Slots(s string)(string){
	value := c.slots.(map[string]interface{})[s].(map[string]interface{})["value"].(string)
	return value
}

func (c *Context) Ask(msg string){
	res := response{}

	// データ入力
        res.Version = "1.0"
        res.Response.OutputSpeech.Type = "PlainText"
        res.Response.OutputSpeech.Text = msg
        res.Response.ShouldEndSession = false
        res.Response.Type = "_DEFAULT_RESPONSE"

        byte_code, _ := json.Marshal(res)
	json := string(byte_code)

	c.gin.JSON(200, json)
}

func (c *Context) Tell(msg string){
	res := response{}

	// データ入力
        res.Version = "1.0"
        res.Response.OutputSpeech.Type = "PlainText"
        res.Response.OutputSpeech.Text = msg
        res.Response.ShouldEndSession = true
        res.Response.Type = "_DEFAULT_RESPONSE"

        byte_code, _ := json.Marshal(res)
	json := string(byte_code)

	c.gin.JSON(200, json)
}

func Default() *golexa{
	init := &golexa{}
	return init
}

func (golexa *golexa) SetIntent(f func(*Context)){
	golexa.Intent = f
}

func (golexa *golexa)Handler(g *gin.Context){
	context := Context{}
	request := request{}

	g.BindJSON(&request)
	context.slots = request.Request.Intent.Slots
	context.Intent = request.Request.Type
	if context.Intent == "IntentRequest" {
		context.Intent = request.Request.Intent.Name
	}

	context.gin = g
	golexa.Intent(&context)
}
