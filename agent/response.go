package agent

import (
	"encoding/json"
	"fmt"
)

// Response structure
type Response struct {
	Data     []uint8 `json:"data"`
	Response string  `json:"response"`
	Info     string  `json:"info"`
	JSON     string  `json:"json"`
}

// NewResponse is Response class constructor.
func NewResponse(data []uint8) (r *Response, err error) {
	jsonData := data[13:]

	r = &Response{Data: data, JSON: string(jsonData)}
	err = json.Unmarshal(jsonData, r)
	if err != nil {
		err = fmt.Errorf("Error decoding response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			err = fmt.Errorf("%s ; Syntax error at byte offset %d", err, e.Offset)
		}
		return
	}
	return
}
