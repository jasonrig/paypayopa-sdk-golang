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

package api

import (
	"github.com/jasonrig/paypayopa-sdk-golang/request"
	"net/url"
)

type Requester interface {
	MakeRequest() *request.Request
}

type Environment struct {
	Host string
}

func (environment *Environment) ToRequest(path string, method string, query *string, payload *interface{}) *request.Request {
	queryString := ""
	if query != nil {
		queryString = *query
	}
	return &request.Request{
		Url: url.URL{
			Scheme:   "https",
			Host:     environment.Host,
			Path:     path,
			RawQuery: queryString,
		},
		Method:      method,
		ContentType: "application/json",
		Body:        payload,
	}
}

var ProductionEnvironment = Environment{
	Host: "api.paypay.ne.jp",
}

var StagingEnvironment = Environment{
	Host: "stg-api.paypay.ne.jp",
}

var SandboxEnvironment = Environment{
	Host: "stg-api.sandbox.paypay.ne.jp",
}
