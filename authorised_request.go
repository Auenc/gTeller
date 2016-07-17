package gTeller

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/auenc/gTeller/util"
)

//AuthorisedRequest is an object that represents a request that requires a User
//to be logged in
type AuthorisedRequest struct {
	Token     string
	PublicKey rsa.PublicKey
	Username  string
	Password  string
	*Request
}

type authorisedData struct {
	Token       string
	Username    string
	Password    string
	RequestData interface{}
}

//Checksum calls the checksum action of the specified end point
func (req *AuthorisedRequest) Checksum() (Response, error) {
	var response Response

	if req.URL == "" {
		return response, errors.New("Endpoint not defined")
	}
	req.URL = req.URL + "checksum"

	return req.Send()
}

//List calls the list action of the specified end point
//If an end point is not specified, an error will be returned
func (req *AuthorisedRequest) List() (Response, error) {
	var response Response

	//Checking to see if end point is configured
	if req.URL == "" {
		return response, errors.New("Endpoint not defined")
	}

	req.URL = req.URL + "list"

	return req.Send()
}

//Add calls the add action of the specified end point
//If an end point is not specified, an error will be returned
func (req *AuthorisedRequest) Add() (Response, error) {
	var response Response

	//Checking to see if end point is configured
	if req.URL == "" {
		return response, errors.New("Endpoint not defined")
	}

	req.URL = req.URL + "add"

	return req.Send()
}

//Update calls the Update action of the specified end point
//If an end point is not specified, an error will be returned
func (req *AuthorisedRequest) Update() (Response, error) {
	var response Response

	//Checking to see if end point is configured
	if req.URL == "" {
		return response, errors.New("Endpoint not defined")
	}

	req.URL = req.URL + "update"

	return req.Send()
}

//Remove calls the Remove action of the specified end point
//If an end point is not specified, an error will be returned
func (req *AuthorisedRequest) Remove() (Response, error) {
	var response Response

	//Checking to see if end point is configured
	if req.URL == "" {
		return response, errors.New("Endpoint not defined")
	}

	req.URL = req.URL + "remove"

	return req.Send()
}

func (req *AuthorisedRequest) Send() (Response, error) {
	var response Response
	var err error

	//encrypting username/password
	eUsernameR, err := util.RsaEncString(req.Username, &req.PublicKey)
	if err != nil {
		return response, err
	}
	ePasswordR, err := util.RsaEncString(req.Password, &req.PublicKey)
	if err != nil {
		return response, err
	}

	eUsername := util.EncString(eUsernameR)
	ePassword := util.EncString(ePasswordR)

	//Creating data object
	data := authorisedData{Token: req.Token, Username: eUsername,
		Password: ePassword, RequestData: req.RequestData}

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

	//Creating byte buffer for data
	var bData []byte
	bData, err = json.Marshal(data)
	if err != nil {
		return response, err
	}
	bbData := bytes.NewBuffer(bData)

	//Creating request
	r, err := http.NewRequest(req.method, req.Host+req.URL, bbData)
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

//Init returns a public key and a token. This should be removed before deploy
func (req *AuthorisedRequest) Init() error {
	var pubKey rsa.PublicKey

	pubKeyRaw := util.DencString(req.Token)

	if tmp, err := x509.ParsePKIXPublicKey([]byte(pubKeyRaw)); err == nil {
		pTmp := tmp.(*rsa.PublicKey)
		pubKey = *pTmp
	} else {
		return err
	}

	req.PublicKey = pubKey

	return nil
}
