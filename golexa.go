package golexa

import(
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type golexa struct{
	// Handle intents
	Intent	func(*Context)
}

type Context struct{
	body	response
	slots	interface{}
	Intent	string
}

func initializeContext() (Context) {
	context := Context{}
	context.body.Version = "1.0"
	context.body.Response.OutputSpeech.Type = "PlainText"
	context.body.Response.OutputSpeech.PlayBehavior = "REPLACE_ENQUEUED"
	context.body.Response.Type = "_DEFAULT_RESPONSE"
	return context
}

func (c *Context) Slots(s string)(string){
	return c.slots.(map[string]interface{})[s].(map[string]interface{})["value"].(string)
}

func (c *Context) Ask(message string){
	c.body.Response.OutputSpeech.Text = message
    c.body.Response.ShouldEndSession = false
}

func (c *Context) Tell(message string){
	c.body.Response.OutputSpeech.Text = message
    c.body.Response.ShouldEndSession = true
}

func Default() *golexa{
	return &golexa{}
}

func (golexa *golexa) SetIntent(intent func(*Context)){
	golexa.Intent = intent
}

func (golexa *golexa)Handler(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	request := make(map[string]interface{})
    if err := json.Unmarshal(body, &request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}
	context := initializeContext()
	context.slots = request["request"]
	if context.Intent = request["request"].(map[string]interface{})["type"].(string);
		context.Intent == "IntentRequest" {
			context.Intent = request["request"].
			(map[string]interface{})["intent"].(map[string]interface{})["name"].(string)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	golexa.Intent(&context)
	body, err = json.Marshal(context.body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(body)
}
