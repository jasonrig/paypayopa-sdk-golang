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

package balance

import (
	"github.com/jasonrig/paypayopa-sdk-golang/api"
	"github.com/jasonrig/paypayopa-sdk-golang/request"
	"net/url"
)

// API Endpoint: Get /v6/wallet/balance?userAuthorizationId={userAuthorizationId}&currency={currency}
// Get user wallet balance
// - Total Amount in User Wallet
// Documentation references:
// * https://www.paypay.ne.jp/opa/doc/v1.0/direct_debit#operation/getBalance
// * https://www.paypay.ne.jp/opa/doc/v1.0/preauth_capture#operation/getBalance

// URL Parameters
type GetParams struct {
	UserAuthorizationId string
	Currency            string
	ProductType         *string
}

// Server response
type GetResponse struct {
	UserAuthorizationId *string     `json:"userAuthorizationId"`
	TotalBalance        *api.Amount `json:"totalBalance"`
}

type Get struct {
	api.Environment
	Params GetParams
}

func (req *Get) MakeRequest() *request.Request {
	queryParams := url.Values{}
	queryParams.Add("userAuthorizationId", req.Params.UserAuthorizationId)
	queryParams.Add("currency", req.Params.Currency)
	if req.Params.ProductType != nil {
		queryParams.Add("productType", *req.Params.ProductType)
	}
	queryString := queryParams.Encode()
	return req.Environment.ToRequest("/v6/wallet/balance", "GET", &queryString, nil)
}
