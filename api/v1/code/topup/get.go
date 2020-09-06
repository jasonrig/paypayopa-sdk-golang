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

package topup

import (
	"fmt"
	"net/url"
	"paypayopa-sdk/api"
	"paypayopa-sdk/request"
)

// API Endpoint: Get /v1/code/topup/{merchantTopUpId}
// Get topup details
// - Get the details of topup done using QR Code
// Documentation references:
// * https://www.paypay.ne.jp/opa/doc/v1.0/direct_debit#operation/getTopUpQRDetails
// * https://www.paypay.ne.jp/opa/doc/v1.0/preauth_capture#operation/getTopUpQRDetails

// URL Parameters
type GetParams struct {
	MerchantTopUpId string
}

// Server response
type GetResponse struct {
	TopUpId             *string      `json:"topUpId"`
	MerchantTopUpId     *string      `json:"merchantTopUpId"`
	UserAuthorizationId *string      `json:"userAuthorizationId"`
	RequestedAt         *int         `json:"requestedAt"`
	AcceptedAt          *int         `json:"acceptedAt"`
	ExpiryDate          *int         `json:"expiryDate"`
	Status              *string      `json:"status"`
	Metadata            *interface{} `json:"metadata"`
}

type Get struct {
	api.Environment
	Params GetParams
}

func (req *Get) MakeRequest() *request.Request {
	return req.Environment.ToRequest(fmt.Sprintf("/v1/code/topup/%s", url.QueryEscape(req.Params.MerchantTopUpId)), "GET", nil, nil)
}
