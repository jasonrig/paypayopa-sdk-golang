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

package codes

import (
	"paypayopa-sdk/api"
	"paypayopa-sdk/request"
)

// API Endpoint: Post /v2/codes
// Create a Code
// - Create a Code to receive payments.
// - The expiration date of the created code is set to "expiryDate".
// Documentation references:
// * https://www.paypay.ne.jp/opa/doc/v1.0/webcashier#operation/createQRCode
// * https://www.paypay.ne.jp/opa/doc/v1.0/dynamicqrcode#operation/createQRCode
// * https://www.paypay.ne.jp/opa/doc/v1.0/appinvoke#operation/createQRCode

// Request payload
type PostPayload struct {
	MerchantPaymentId   string            `json:"merchantPaymentId"`
	Amount              api.Amount        `json:"amount"`
	OrderDescription    *string           `json:"orderDescription"`
	OrderItems          *[]api.OrderItems `json:"orderItems"`
	Metadata            *interface{}      `json:"metadata"`
	CodeType            string            `json:"codeType"`
	StoreInfo           *string           `json:"storeInfo"`
	StoreId             *string           `json:"storeId"`
	TerminalId          *string           `json:"terminalId"`
	RequestedAt         *int              `json:"requestedAt"`
	RedirectUrl         *string           `json:"redirectUrl"`
	RedirectType        *string           `json:"redirectType"`
	UserAgent           *string           `json:"userAgent"`
	IsAuthorization     *bool             `json:"isAuthorization"`
	AuthorizationExpiry *int              `json:"authorizationExpiry"`
}

// Server response
type PostResponse struct {
	CodeId            *string     `json:"codeId"`
	Url               *string     `json:"url"`
	Deeplink          *string     `json:"deeplink"`
	ExpiryDate        *int        `json:"expiryDate"`
	MerchantPaymentId *string     `json:"merchantPaymentId"`
	Amount            *api.Amount `json:"amount"`
	OrderDescription  *string     `json:"orderDescription"`
	OrderItems        *[]struct {
		Name      *string     `json:"name"`
		Category  *string     `json:"category"`
		Quantity  *int        `json:"quantity"`
		ProductId *string     `json:"productId"`
		UnitPrice *api.Amount `json:"unit_price"`
	} `json:"orderItems"`
	Metadata            *interface{} `json:"metadata"`
	CodeType            *string      `json:"codeType"`
	StoreInfo           *string      `json:"storeInfo"`
	StoreId             *string      `json:"storeId"`
	TerminalId          *string      `json:"terminalId"`
	RequestedAt         *int         `json:"requestedAt"`
	RedirectUrl         *string      `json:"redirectUrl"`
	RedirectType        *string      `json:"redirectType"`
	IsAuthorization     *bool        `json:"isAuthorization"`
	AuthorizationExpiry *int         `json:"authorizationExpiry"`
}

type Post struct {
	api.Environment
	Payload PostPayload
}

func (req *Post) MakeRequest() *request.Request {
	var payload interface{} = req.Payload
	return req.Environment.ToRequest("/v2/codes", "POST", nil, &payload)
}
