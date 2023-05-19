package api

import "time"

type Respond struct {
	Code      int
	Message   string
	Timestamp string
	Ttl       int
	Data      interface{}
}

func ReturnResponse(code int, message string) *Respond {
	return &Respond{Code: code, Message: message, Timestamp: time.Now().Format("2006-01-02 15:04:05"), Ttl: 1, Data: nil}
}

func ReturnResponseWithData(code int, message string, data interface{}) *Respond {
	return &Respond{Code: code, Message: message, Timestamp: time.Now().Format("2006-01-02 15:04:05"), Ttl: 1, Data: data}
}
