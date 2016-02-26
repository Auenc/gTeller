package gTeller

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type initResponse struct {
	Token     string
	PublicKey string
}

//Request is an object that processes the api request
type Request struct {
	Host        string
	URL         string
	method      string
	RequestData interface{}
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

func (req *Request) Authorise(username, password, key string) (AuthorisedRequest, error) {

	auth := AuthorisedRequest{Username: username, Password: password,
		Request: req, Token: key}

	err := auth.Init()
	if err != nil {
		return auth, err
	}

	return auth, nil
}

func (req *Request) Send() (Response, error) {
	var response Response
	var data []byte
	var err error
	//If there is data to be sent
	if req.RequestData != nil {

		reqD := struct {
			RequestData interface{}
		}{
			req.RequestData,
		}

		data, err = json.Marshal(reqD)
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
	r, err := http.NewRequest(req.method, req.Host+req.URL, bytes.NewReader(data))
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

	if resp.StatusCode != http.StatusOK {
		return response, errors.New(string(resp.Status))
	}

	response = Response{URL: req.URL, Method: req.method,
		RequestData: req.RequestData, ResponseCode: resp.StatusCode, ResponseData: body}

	return response, nil
}
