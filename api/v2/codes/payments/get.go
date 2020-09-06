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

package payments

import (
	"fmt"
	"github.com/jasonrig/paypayopa-sdk-golang/api"
	"github.com/jasonrig/paypayopa-sdk-golang/request"
	"net/url"
)

// API Endpoint: Get /v2/codes/payments/{merchantPaymentId}
// Get payment details
// - Get payment details.
// Documentation references:
// * https://www.paypay.ne.jp/opa/doc/v1.0/webcashier#operation/getPaymentDetails
// * https://www.paypay.ne.jp/opa/doc/v1.0/dynamicqrcode#operation/getPaymentDetails
// * https://www.paypay.ne.jp/opa/doc/v1.0/appinvoke#operation/getPaymentDetails

// URL Parameters
type GetParams struct {
	MerchantPaymentId string
}

// Server response
type GetResponse struct {
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
	Captures struct {
		Data []struct {
			AcceptedAt        *int        `json:"acceptedAt"`
			MerchantCaptureId *string     `json:"merchantCaptureId"`
			Amount            *api.Amount `json:"amount"`
			OrderDescription  *string     `json:"orderDescription"`
			RequestedAt       *int        `json:"requestedAt"`
			ExpiresAt         *int        `json:"expiresAt"`
			Status            *string     `json:"status"`
		} `json:"data"`
	} `json:"captures"`
	Revert struct {
		AcceptedAt       *int    `json:"acceptedAt"`
		MerchantRevertId *string `json:"merchantRevertId"`
		RequestedAt      *int    `json:"requestedAt"`
		Reason           *string `json:"reason"`
	} `json:"revert"`
	MerchantPaymentId   *string     `json:"merchantPaymentId"`
	UserAuthorizationId *string     `json:"userAuthorizationId"`
	Amount              *api.Amount `json:"amount"`
	RequestedAt         *int        `json:"requestedAt"`
	ExpiresAt           *int        `json:"expiresAt"`
	CanceledAt          *int        `json:"canceledAt"`
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
	Metadata *interface{} `json:"metadata"`
}

type Get struct {
	api.Environment
	Params GetParams
}

func (req *Get) MakeRequest() *request.Request {
	return req.Environment.ToRequest(fmt.Sprintf("/v2/codes/payments/%s", url.QueryEscape(req.Params.MerchantPaymentId)), "GET", nil, nil)
}
