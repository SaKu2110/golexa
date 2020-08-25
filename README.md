# Golexa
A Go package for interacting with Amazon Alexa.  

An example application implementing the golexa can be played with [here](https://github.com/SaKu2110/golexa_example).
## Installation
To install Gin package, you need to install Go and set your Go workspace first.  
1. The first need Go installed (version 1.10+ is required), then you can use the below Go command to install Gin and Golexa.  

```sh
$ go get github.com/SaKu2110/golexa
```

2. Import it in your code

```go
import(
  "github.com/SaKu2110/golexa"
)
```

## Usage
### Getting Started
```go
package main

import(
	"github.com/SaKu2110/golexa"
)

func main(){
  app := golexa.New()
  app.LaunchSession(LaunchRequest)
  app.CustomIntent("SayHelloIntent", SayHelloIntent)
  app.ClosedSession(SessionEndedRequest)
  gin.Run()
}

func LaunchRequest (c *golexa.Context) {
  // Your request processing
}

func SayHelloIntent (c *golexa.Context) {
  // Your request processing
}

func SessionEndedRequest (c *golexa.Context) {
  // Your request processing
}
```
### About 'Response'
#### `Ask` and `Tell`

There are two kinds of responses you can send to Alexa: asks and tells. An ask should ask the user a question, and expect them to reply. A tell should end the conversation.
```go
func RequestHandler (c *golexa.Context) {
  /* If you want to continue the conversation with this speech,
   * you can use Ask().
   * Json parameter: shouldEndSession = false
   */
  c.Ask().WithText("Hello World.")
  /*  If you want to end the conversation with this utterance,
   *  you can use Tell().
   *  Json parameter: shouldEndSession = true
   *  c.Tell()
   */ 
}
```

#### `slots`

Alexa sometimes has a programming language argument called slot.
```go
func RequestHandler (c *golexa.Context) {
  // c.LoadSlot("Slot Name")
  fmt.Printf("%s\n", c.LoadSlot("color"))
}
```
#### `attribute`
Reading and setting session attributes.  

You can persist data to an Alexa session:  
```go
func RequestHandler (c *golexa.Context) {
  // Transcribe attributes to shouldEndSession.
  c.CopyAttributes()
}
```
You can read it:
```go
func RequestHandler (c *golexa.Context) {
  // c.LoadAttribute("Attribute Key")
  _ = c.LoadAttribute("episode")
}
```
You cat regist new attribute:
```go
func RequestHandler (c *golexa.Context) {
  // c.SetAttribute("Key", Value)
  c.SetAttribute("episode", "one")
}
```

## License
The package is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).
