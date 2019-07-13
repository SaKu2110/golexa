# Golaxa
A Go package for interacting with Amazon Alexa.  
Designed to work with Gin.

An example application implementing the golexa can be played with [here](https://github.com/SaKu2110/golexa_example).
## Installation
To install Gin package, you need to install Go and set your Go workspace first.  
1. The first need Go installed (version 1.10+ is required), then you can use the below Go command to install Gin and Golexa.  

```sh
$ go get github.com/gin-gonic/gin
$ go get github.com/SaKu2110/golexa
```

2. Import it in your code

```go
import(
  "github.com/gin-gonic/gin"
  "github.com/SaKu2110/golexa"
)
```

## Usage
### Getting Started
```go
package main

import(
	"github.com/gin-gonic/gin"
	"github.com/SaKu2110/golexa"
)

func main(){
	gin := gin.Default()
  
	golexa := golexa.Default()
	golexa.SetIntent(intent)

	gin.POST("/", golexa.Handler)
	gin.Run()
}

func intent (c *golexa.Context) {
  // Your request processing
}
```
### About 'Response'
##### `Ask` and `Tell`

There are two kinds of responses you can send to Alexa: asks and tells. An ask should ask the user a question, and expect them to reply. A tell should end the conversation.
```go
func intent (c *golexa.Context) {
  switch c.Intent{
  case "AskMoreIntent":
    c.Ask("What next?")
  case "FinishTalkIntent":
    c.Tell("That's all.")
  default:
    c.Tell("There is no such intent.")
  }
}
```

##### `slots`

Alexa sometimes has a programming language argument called slot.
```go
func intent (c *golexa.Context) {
  switch c.Intent{
  case "EchoColorIntent":
    // c.Slots("Slot Name")
    c.Ask(c.Slots("color"))
  default:
    c.Tell("There is no such intent.")
  }
}
```
## License
The package is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).
