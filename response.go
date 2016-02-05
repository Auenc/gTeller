package gTeller

import (
	"bytes"
	"encoding/json"
)

//Response is the object returned from an API call
type Response struct {
	URL          string
	Method       string
	RequestData  interface{}
	ResponseCode int
	ResponseData []byte
}

//Get attempts to save the data contained within the Response object to the specified interface
//Accepts the object you wish to save the data to as obj
//Accepts the interface the response data should be as d
func (res *Response) Get(obj interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(res.ResponseData))
	err := decoder.Decode(obj)
	return err
}
