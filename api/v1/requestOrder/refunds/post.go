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

package refunds

import (
	"github.com/jasonrig/paypayopa-sdk-golang/api"
	"github.com/jasonrig/paypayopa-sdk-golang/request"
)

// API Endpoint: Post /v1/requestOrder/refunds
// Refund a payment
// - Refund a payment.
// Documentation references:
// * https://www.paypay.ne.jp/opa/doc/v1.0/pending_payments#operation/refundPayment

// Request payload
type PostPayload struct {
	MerchantRefundId string     `json:"merchantRefundId"`
	PaymentId        string     `json:"paymentId"`
	Amount           api.Amount `json:"amount"`
	RequestedAt      int        `json:"requestedAt"`
	Reason           *string    `json:"reason"`
}

// Server response
type PostResponse struct {
	Status           *string     `json:"status"`
	AcceptedAt       *int        `json:"acceptedAt"`
	MerchantRefundId *string     `json:"merchantRefundId"`
	PaymentId        *string     `json:"paymentId"`
	Amount           *api.Amount `json:"amount"`
	RequestedAt      *int        `json:"requestedAt"`
	Reason           *string     `json:"reason"`
}

type Post struct {
	api.Environment
	Payload PostPayload
}

func (req *Post) MakeRequest() *request.Request {
	var payload interface{} = req.Payload
	return req.Environment.ToRequest("/v1/requestOrder/refunds", "POST", nil, &payload)
}
