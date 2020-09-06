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

package authorizations

import (
	"net/url"
	"paypayopa-sdk/api"
	"paypayopa-sdk/request"
)

// API Endpoint: Get /v2/user/authorizations?userAuthorizationId={userAuthorizationId}
// Get user authorization status
// - Get the authorization status of a user
// Documentation references:
// * https://www.paypay.ne.jp/opa/doc/v1.0/direct_debit#operation/getUserAuthorizationStatus
// * https://www.paypay.ne.jp/opa/doc/v1.0/preauth_capture#operation/getUserAuthorizationStatus
// * https://www.paypay.ne.jp/opa/doc/v1.0/continuous_payments#operation/getUserAuthorizationStatus

// URL Parameters
type GetParams struct {
	UserAuthorizationId string
}

// Server response
type GetResponse struct {
	UserAuthorizationId *string   `json:"userAuthorizationId"`
	ReferenceIds        *[]string `json:"referenceIds"`
	Status              *string   `json:"status"`
	Scopes              *[]string `json:"scopes"`
	ExpireAt            *int      `json:"expireAt"`
	IssuedAt            *int      `json:"issuedAt"`
}

type Get struct {
	api.Environment
	Params GetParams
}

func (req *Get) MakeRequest() *request.Request {
	queryParams := &url.Values{}
	queryParams.Add("userAuthorizationId", req.Params.UserAuthorizationId)
	queryString := queryParams.Encode()
	return req.Environment.ToRequest("/v2/user/authorizations", "GET", &queryString, nil)
}
