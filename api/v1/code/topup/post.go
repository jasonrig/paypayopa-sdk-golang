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
	"paypayopa-sdk/api"
	"paypayopa-sdk/request"
)

// API Endpoint: Post /v1/code/topup
// Create topup QR code
// - This service will allow a merchant to create qr code for the user.
// - The user can scan the QR code to topUp the his/her account.
// Documentation references:
// * https://www.paypay.ne.jp/opa/doc/v1.0/direct_debit#operation/createTopUpQRCode
// * https://www.paypay.ne.jp/opa/doc/v1.0/preauth_capture#operation/createTopUpQRCode

// Request payload
type PostPayload struct {
	MerchantTopUpId     string       `json:"merchantTopUpId"`
	UserAuthorizationId string       `json:"userAuthorizationId"`
	MinimumTopUpAmount  api.Amount   `json:"minimumTopUpAmount"`
	Metadata            *interface{} `json:"metadata"`
	CodeType            string       `json:"codeType"`
	RequestedAt         *int         `json:"requestedAt"`
	RedirectType        *string      `json:"redirectType"`
	RedirectUrl         *string      `json:"redirectUrl"`
	UserAgent           *string      `json:"userAgent"`
}

// Server response
type PostResponse struct {
	CodeId              *string      `json:"codeId"`
	Url                 *string      `json:"url"`
	Status              *string      `json:"status"`
	MerchantTopUpId     *string      `json:"merchantTopUpId"`
	UserAuthorizationId *string      `json:"userAuthorizationId"`
	MinimumTopUpAmount  *api.Amount  `json:"minimumTopUpAmount"`
	Metadata            *interface{} `json:"metadata"`
	ExpiryDate          *int         `json:"expiryDate"`
	CodeType            *string      `json:"codeType"`
	RequestedAt         *int         `json:"requestedAt"`
	RedirectType        *string      `json:"redirectType"`
	RedirectUrl         *string      `json:"redirectUrl"`
	UserAgent           *string      `json:"userAgent"`
}

type Post struct {
	api.Environment
	Payload PostPayload
}

func (req *Post) MakeRequest() *request.Request {
	var payload interface{} = req.Payload
	return req.Environment.ToRequest("/v1/code/topup", "POST", nil, &payload)
}
