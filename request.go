package gTeller

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/auenc/gTeller/filter"
	"github.com/auenc/gTeller/items"
)

//Request is an object that processes the api request
type Request struct {
	Host        string
	URL         string
	method      string
	RequestData interface{}
}

const (
	ItemURI = "item/"
)

//Items configures the Request object to point towards the item API endpoints
func (req *Request) Items(filter filter.ItemsFilter) ([]items.Item, error) {
	var list []items.Item

	//configuring request
	req.URL = req.URL + ItemURI
	req.method = "POST"
	req.RequestData = ListItemsRequest{Filter: filter}

	resp, err := req.List()
	if err != nil {
		return list, err
	}
	err = resp.Get(&list)
	if err != nil {
		return list, err
	}

	return list, nil
}

//List calls the list action of the specified end point
//If an end point is not specified, an error will be returned
func (req *Request) List() (Response, error) {
	var response Response

	//Checking to see if end point is configured
	if req.URL == "" {
		return response, errors.New("Endpoint not defined")
	}

	req.URL = req.URL + "list"

	return req.Send()
}
func (req *Request) Send() (Response, error) {
	var response Response
	var data []byte
	var err error
	//If there is data to be sent
	if req.RequestData != nil {
		data, err = json.Marshal(req.RequestData)
		if err != nil {
			return response, err
		}
	}
	//Checking URL
	if req.URL == "" {
		return response, errors.New("Send::URL not specified.")
	}
	//Checking host
	if req.Host == "" {
		return response, errors.New("Send::Host not specified.")
	}
	//Checking method
	if req.method == "" {
		return response, errors.New("Send::Method not specified.")
	}

	//Creating request
	r, err := http.NewRequest(req.method, req.Host+req.URL, bytes.NewBuffer(data))
	if err != nil {
		return response, err
	}
	//Setting request values
	r.Header.Set("Content-Type", "application/json")

	//Getting client
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return response, err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	response = Response{URL: req.URL, Method: req.method,
		RequestData: req.RequestData, ResponseCode: resp.StatusCode, ResponseData: body}

	return response, nil
}
