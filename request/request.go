/*
 *    Copyright 2020 Jason Rigby
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Request struct {
	Url         url.URL
	Method      string
	ContentType string
	Body        *interface{}
}

// Sets the Authorization header for the HTTP request
func (request *Request) setHeaders(req *http.Request, auth *Authentication) error {
	authHeader, err := auth.CreateAuthHeader(request)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", *authHeader)
	req.Header.Add("Content-Type", "application/json")
	return nil
}

// Executes the HTTP request
// The result is stored in responsePayload
func (request *Request) Call(auth *Authentication, client *http.Client, responsePayload interface{}) error {
	body, err := json.Marshal(request.Body)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(body)
	req, err := http.NewRequest(request.Method, request.Url.String(), reader)
	if err != nil {
		return err
	}
	err = request.setHeaders(req, auth)
	if err != nil {
		return err
	}
	httpResponse, err := client.Do(req)
	defer closeResponse(httpResponse)
	if err != nil {
		return err
	}
	rawUnmarshalledResponse := &response{}
	err = rawUnmarshalledResponse.unmarshal(httpResponse, responsePayload)
	if err != nil {
		return err
	}
	return nil
}

func closeResponse(response *http.Response) {
	err := response.Body.Close()
	if err != nil {
		panic(err)
	}
}

type response struct {
	ResultInfo struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		CodeId  string `json:"codeId"`
	} `json:"resultInfo"`
	Data json.RawMessage `json:"data"`
}

// Unmarshals the server response
func (response *response) unmarshal(httpResponse *http.Response, resp interface{}) error {
	bodyBytes, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bodyBytes, response)
	if err != nil {
		return err
	}
	if response.ResultInfo.Code != "SUCCESS" {
		return errors.New(fmt.Sprintf("Code: %s\nMessage: %s\nError Id: %s", response.ResultInfo.Code, response.ResultInfo.Message, response.ResultInfo.CodeId))
	}
	err = json.Unmarshal(response.Data, resp)
	if err != nil {
		return err
	}
	return nil
}
