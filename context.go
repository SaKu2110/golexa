package golexa

import(
	"encoding/json"
)

const(
	// RESPONSE QUEUE BEHAVIOR
	RESPONSE_QUEUE_ADDED = "ENQUEUE"
	RESPONSE_QUEUE_INTERRUPT = "REPLACE_ALL"
	RESPONSE_QUEUE_REPLACE = "REPLACE_ENQUEUED"

	// CARD TYPE
	CARD_TYPE_SIMPLE = "Simple"
	CARD_TYPE_STANDARD = "Standard"
	CARD_TYPE_LINK_ACCOUNT = "LinkAccount"

	// MESSAGE TYPE
	MESSAGE_TYPE_TEXT = "PlainText"
	MESSAGE_TYPE_SSML = "SSML"
)

type Context struct {
	// request params
	params		*Params

	attributes	map[string]interface{}
	response	*Response
	integrity	bool
}

type Params struct {
	requestType	string
	// session Data
	attributes	map[string]interface{}

	// Intent  Data
	name		string
	slots		map[string]Slot
}

func GenContext(body []byte) (cx Context) {
	request := RequestBody{}
	json.Unmarshal(body, &request)
	
	cx = Context{
		params:		parseRequest(request),
		attributes:	make(map[string]interface{}),
		response:	makeNewResponse(),
		integrity:	true,
	}
	return
}

func parseRequest(req RequestBody) (params *Params) {
	intentName := req.Request.Type
	if intentName == "IntentRequest" {
		intentName = req.Request.Intent.Name
	}
	attributeMap := make(map[string]interface{})
	if req.Session.Attributes != nil {
		attributeMap = req.Session.Attributes
	}
	params = &Params{
		requestType:	req.Request.Type,
		attributes:		attributeMap,
		name:			intentName,
		slots:			req.Request.Intent.Slots,
	}
	return
}

func makeNewResponse() (res *Response) {
	res = &Response{}
	return
}

func (cx *Context) packData() (body *ResponseBody) {
	body = &ResponseBody{
		Version:	"1.0",
		Attributes:	cx.attributes,
		Response:	cx.response,
	}
	return
}

func (cx *Context) JSON() (body []byte, err error) {
	body, err = json.Marshal(cx.packData())
	if err != nil {
		cx.integrity = false
		return
	}
	return
}

/******************************/
/* OutputSpeech object setter */
/******************************/

func (cx *Context) Ask()  *Context {
	cx.response.ShouldEndSession = false
	return cx
}

func (cx *Context) Tell() *Context {
	cx.response.ShouldEndSession = true
	return cx
}

func (cx *Context) WithText(text string, params ...string) *Context {
	behavior := ""
	if len(params) > 0 {
		behavior = checkBehavior(params[0])
	}
	cx.response.OutputSpeech = setOutputSpeech("PlainText", text, behavior)
	return cx
}

func (cx *Context) WithSSML(ssml string, params ...string) *Context {
	behavior := ""
	if len(params) > 0 {
		behavior = checkBehavior(params[0])
	}
	cx.response.OutputSpeech = setOutputSpeech("SSML", ssml, behavior)
	return cx
}

func (cx *Context) RepromptWithText(text string, params ...string) *Context {
	behavior := ""
	if len(params) > 0 {
		behavior = checkBehavior(params[0])
	}
	cx.response.Reprompt = &Reprompt{
		OutputSpeech:	setOutputSpeech("PlainText", text, behavior),
	}
	return cx
}

func (cx *Context) RepromptWithSSML(ssml string, params ...string) *Context {
	behavior := ""
	if len(params) > 0 {
		behavior = checkBehavior(params[0])
	}
	cx.response.Reprompt = &Reprompt{
		OutputSpeech:	setOutputSpeech("SSMl", ssml, behavior),
	}
	return cx
}

func setOutputSpeech(msgType, msg, behavior string) (out *OutputSpeech) {
	out = &OutputSpeech{
		Type:			msgType,
		PlayBehavior:	behavior,
	}
	switch out.Type {
	case MESSAGE_TYPE_TEXT:
		out.Text = msg
	case MESSAGE_TYPE_SSML:
		out.SSML = msg
	}
	return
}

func checkBehavior(behavior string) string {
	switch behavior {
	case RESPONSE_QUEUE_ADDED:
		return RESPONSE_QUEUE_REPLACE
	case RESPONSE_QUEUE_INTERRUPT:
		return RESPONSE_QUEUE_INTERRUPT
	case RESPONSE_QUEUE_REPLACE:
		return RESPONSE_QUEUE_REPLACE
	}
	return ""
}

/**********************/
/* Card object setter */
/**********************/

func (cx *Context) SimpleCard(title, content string) *Context {
	cx.response.Card = &Card{
		Type:		CARD_TYPE_SIMPLE,
		Title:		title,
		Content:	content,
	}
	return cx
}

func (cx *Context) StandardCard(title, text, smallURL, largeURL string) *Context {
	cx.response.Card = &Card{
		Type:	CARD_TYPE_STANDARD,
		Title:	title,
		Text:	text,
		Image:	&Image{
			SmallImageUrl:	smallURL,
			LargeImageUrl:	largeURL,
		},
	}
	return cx
}

func (cx *Context) LinkAccountCard() *Context {
	cx.response.Card = &Card{
		Type:	CARD_TYPE_LINK_ACCOUNT,
	}
	return cx
}

/*********************/
/* Attribute Manager */
/*********************/

func (cx *Context) SetAttribute(key string, obj interface{}) *Context {
	cx.attributes[key] = obj
	return cx
}

func (cx *Context) CopyAttributes() *Context {
	cx.attributes = cx.params.attributes
	return cx
}

func (cx *Context) LoadAttribute(key string) interface{} {
	return cx.params.attributes[key]
}

/*********************/
/* Slot value getter */
/*********************/

func (cx *Context) LoadSlot(key string) string {
	return cx.params.slots[key].Value
}
