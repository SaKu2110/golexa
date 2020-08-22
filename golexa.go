package golexa

import(
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
)

const(
	DEAFULT_PORT = ":9020"
	CONTENT_TYPE_JSON = "application/json; application/json"
)

var listenPort string

type IntentFunc func(*Context)

type Engine struct {
	IntentGroup	map[string]IntentFunc
}

// Create new golexa engine.
func New() (*Engine) {
	engine := &Engine{
		IntentGroup: make(map[string]IntentFunc),
	}
	return engine
}

/****************************/
/* Intent handler registers */
/****************************/

func (engine *Engine) LaunchSession(intent IntentFunc) {
	engine.IntentGroup["LaunchRequest"] = intent
}

func (engine *Engine) ClosedSession(intent IntentFunc) {
	engine.IntentGroup["SessionEndedRequest"] = intent
}

func (engine *Engine) CustomIntent(name string, intent IntentFunc) {
	engine.IntentGroup[name] = intent
}

/***************************/
/* Response Json generator */
/***************************/

func (engine *Engine) Response(body []byte) (response []byte,err error) {
	status := http.StatusOK
	cx := GenContext(body)
	engine.IntentGroup[cx.params.name](&cx)
	response, err = cx.JSON()
	if err != nil {
		status = http.StatusInternalServerError
		debugPrint("[ERROR] Faild Marshal Json.")
	}
	if !cx.integrity {
		status = http.StatusInternalServerError
		err = fmt.Errorf("[ERROR] Inconsistent JSON.")
	}
	handledebugPrint(status, cx.params.requestType, cx.params.name)
	return
}

/************************/
/* Launch golexa server */
/************************/

func (engine *Engine) Run(addr ...string) (err error){
	switch len(addr) {
	case 0:
		listenPort = DEAFULT_PORT
		if port := os.Getenv("PORT"); port != "" {
			debugPrint("Environment variable PORT=\"%s\"", port)
			listenPort = port
		}
	case 1:
		listenPort = addr[0]
	default:
		panic("[ERROR] too many parameters.")
	}
	http.HandleFunc("/" , func(res http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic("[ERROR] Faild read Request Json body.")
			res.WriteHeader(http.StatusInternalServerError)
		}
		body, err = engine.Response(body)
		if err != nil {
			panic("[ERROR] Faild create Response Json.")
			res.WriteHeader(http.StatusInternalServerError)
		}
		setResponseHeader(res, http.StatusOK)
		res.Write(body)
	})
	debugPrint("Listening and serving HTTP on %s\n", listenPort)
	err = http.ListenAndServe(listenPort, nil)
	return
}

func setResponseHeader(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", CONTENT_TYPE_JSON)
	w.WriteHeader(code)
}