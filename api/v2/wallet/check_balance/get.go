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

package check_balance

import (
	"github.com/jasonrig/paypayopa-sdk-golang/api"
	"github.com/jasonrig/paypayopa-sdk-golang/request"
	"net/url"
	"strconv"
)

// API Endpoint: Get /v2/wallet/check_balance?userAuthorizationId={userAuthorizationId}&amount={amount}&currency={currency}
// Check user wallet balance
// - Check if user has enough balance to make a payment
// Documentation references:
// * https://www.paypay.ne.jp/opa/doc/v1.0/direct_debit#operation/checkWalletBalance
// * https://www.paypay.ne.jp/opa/doc/v1.0/preauth_capture#operation/checkWalletBalance

// URL Parameters
type GetParams struct {
	UserAuthorizationId string
	Amount              int
	Currency            string
	ProductType         *string
}

// Server response
type GetResponse struct {
	HasEnoughBalance *bool `json:"hasEnoughBalance"`
}

type Get struct {
	api.Environment
	Params GetParams
}

func (req *Get) MakeRequest() *request.Request {
	queryParams := url.Values{}
	queryParams.Add("userAuthorizationId", req.Params.UserAuthorizationId)
	queryParams.Add("amount", strconv.FormatInt(int64(req.Params.Amount), 10))
	queryParams.Add("currency", req.Params.Currency)
	if req.Params.ProductType != nil {
		queryParams.Add("productType", *req.Params.ProductType)
	}
	queryString := queryParams.Encode()
	return req.Environment.ToRequest("/v2/wallet/check_balance", "GET", &queryString, nil)
}
