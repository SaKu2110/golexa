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
	slots		interface{}
	Intent		string
}

func (c *Context) Slots(s string)(string){
	return c.slots.(map[string]interface{})[s].(map[string]interface{})["value"].(string)
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
	var request interface{}

	g.BindJSON(&request)
	context.slots = request.(map[string]interface{})["request"].(interface{})
	context.Intent = request.(map[string]interface{})["request"].(map[string]interface{})["type"].(string)
	if context.Intent == "IntentRequest" {
		context.Intent = request.(map[string]interface{})["request"].(map[string]interface{})["intent"].(map[string]interface{})["name"].(string)
	}

	context.gin = g
	golexa.Intent(&context)
}
