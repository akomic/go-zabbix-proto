package agent

import (
	"encoding/json"
	"log"
)

// Response
type Response struct {
	Data     []uint8 `json:"data"`
	Response string  `json:"response"`
	Info     string  `json:"info"`
	Json     string  `json:"json"`
}

// Response class constructor.
func NewResponse(data []uint8) (r *Response, err error) {
	jsonData := data[13:]

	r = &Response{Data: data, Json: string(jsonData)}
	err = json.Unmarshal(jsonData, r)
	if err != nil {
		log.Printf("Error decoding response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("Syntax error at byte offset %d", e.Offset)
		}
		return
	}
	return
}
