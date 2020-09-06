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

package capture

import (
	"github.com/jasonrig/paypayopa-sdk-golang/api"
	"github.com/jasonrig/paypayopa-sdk-golang/request"
)

// API Endpoint: Post /v2/payments/capture
// Capture a payment authorization
// - This api is used to capture the payment authorization for a payment.
// - If you want to increase the amount, we will send a notification to the user asking for consent.
// Documentation references:
// * https://www.paypay.ne.jp/opa/doc/v1.0/webcashier#operation/capturePaymentAuth
// * https://www.paypay.ne.jp/opa/doc/v1.0/preauth_capture#operation/capturePaymentAuth
// * https://www.paypay.ne.jp/opa/doc/v1.0/dynamicqrcode#operation/capturePaymentAuth
// * https://www.paypay.ne.jp/opa/doc/v1.0/appinvoke#operation/capturePaymentAuth

// Request payload
type PostPayload struct {
	MerchantPaymentId string     `json:"merchantPaymentId"`
	Amount            api.Amount `json:"amount"`
	MerchantCaptureId string     `json:"merchantCaptureId"`
	RequestedAt       int        `json:"requestedAt"`
	OrderDescription  string     `json:"orderDescription"`
}

// Server response
type PostResponse struct {
	PaymentId  *string `json:"paymentId"`
	Status     *string `json:"status"`
	AcceptedAt *int    `json:"acceptedAt"`
	Refunds    *struct {
		Data *[]struct {
			Status           *string     `json:"status"`
			AcceptedAt       *int        `json:"acceptedAt"`
			MerchantRefundId *string     `json:"merchantRefundId"`
			PaymentId        *string     `json:"paymentId"`
			Amount           *api.Amount `json:"amount"`
			RequestedAt      *int        `json:"requestedAt"`
			Reason           *string     `json:"reason"`
		} `json:"data"`
	} `json:"refunds"`
	Captures *struct {
		Data *[]struct {
			AcceptedAt        *int        `json:"acceptedAt"`
			MerchantCaptureId *string     `json:"merchantCaptureId"`
			Amount            *api.Amount `json:"amount"`
			OrderDescription  *string     `json:"orderDescription"`
			RequestedAt       *int        `json:"requestedAt"`
			Status            *string     `json:"status"`
		} `json:"data"`
	} `json:"captures"`
	MerchantPaymentId   *string     `json:"merchantPaymentId"`
	UserAuthorizationId *string     `json:"userAuthorizationId"`
	Amount              *api.Amount `json:"amount"`
	RequestedAt         *int        `json:"requestedAt"`
	ExpiresAt           *int        `json:"expiresAt"`
	StoreId             *string     `json:"storeId"`
	TerminalId          *string     `json:"terminalId"`
	OrderReceiptNumber  *string     `json:"orderReceiptNumber"`
	OrderDescription    *string     `json:"orderDescription"`
	OrderItems          *[]struct {
		Name      *string     `json:"name"`
		Category  *string     `json:"category"`
		Quantity  *int        `json:"quantity"`
		ProductId *string     `json:"productId"`
		UnitPrice *api.Amount `json:"unitPrice"`
	} `json:"orderItems"`
	Metadata       *interface{} `json:"metadata"`
	AssumeMerchant *string      `json:"assumeMerchant"`
}

type Post struct {
	api.Environment
	Payload PostPayload
}

func (req *Post) MakeRequest() *request.Request {
	var payload interface{} = req.Payload
	return req.Environment.ToRequest("/v2/payments/capture", "POST", nil, &payload)
}
