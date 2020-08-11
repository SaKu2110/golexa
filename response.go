package golexa

type response struct {
        Version  string `json:"version"`
        Response struct {
                OutputSpeech struct {
                        Type string `json:"type"`
                        Text string `json:"text"`
                        PlayBehavior string `json:"playBehavior"`
                } `json:"outputSpeech"`
                Card struct {
                        Type    string `json:"type"`
                        Title   string `json:"title"`
                        Image struct {
                                SampleImageUrl  string `json:"smallImageUrl"`
                                LargeImageUrl   string `json:"largeImageUrl"`
                        } `json:"image"`
                        Reprompt struct {
                                OutputSpeech struct {
                                        Type string `json:"type"`
                                        Text string `json:"text"`
                                        PlayBehavior string `json:"playBehavior"`
                                }  `json:"outputSpeech"`
                        } `json:"reprompt"`
                } `json:"card"`
                ShouldEndSession bool   `json:"shouldEndSession"`
                Type             string `json:"type"`
        } `json:"response"`
}