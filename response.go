package golexa

type response struct {
        Version  string `json:"version"`
        Response struct {
                OutputSpeech struct {
                        Type string `json:"type"`
                        Text string `json:"text"`
                } `json:"outputSpeech"`
                ShouldEndSession bool   `json:"shouldEndSession"`
                Type             string `json:"type"`
        } `json:"response"`
}
