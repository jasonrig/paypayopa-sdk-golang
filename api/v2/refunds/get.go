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
	"fmt"
	"github.com/jasonrig/paypayopa-sdk-golang/api"
	"github.com/jasonrig/paypayopa-sdk-golang/request"
	"net/url"
)

// API Endpoint: Get /v2/refunds/{merchantRefundId}
// Get refund details
// - Get refund details.
// Documentation references:
// * https://www.paypay.ne.jp/opa/doc/v1.0/webcashier#operation/getRefundDetails
// * https://www.paypay.ne.jp/opa/doc/v1.0/direct_debit#operation/getRefundDetails
// * https://www.paypay.ne.jp/opa/doc/v1.0/preauth_capture#operation/getRefundDetails
// * https://www.paypay.ne.jp/opa/doc/v1.0/dynamicqrcode#operation/getRefundDetails
// * https://www.paypay.ne.jp/opa/doc/v1.0/appinvoke#operation/getRefundDetails
// * https://www.paypay.ne.jp/opa/doc/v1.0/pending_payments#operation/getRefundDetails
// * https://www.paypay.ne.jp/opa/doc/v1.0/continuous_payments#operation/getRefundDetails

// URL Parameters
type GetParams struct {
	MerchantRefundId string
}

// Server response
type GetResponse struct {
	Status           *string     `json:"status"`
	AcceptedAt       *int        `json:"acceptedAt"`
	MerchantRefundId *string     `json:"merchantRefundId"`
	PaymentId        *string     `json:"paymentId"`
	Amount           *api.Amount `json:"amount"`
	RequestedAt      *int        `json:"requestedAt"`
	Reason           *string     `json:"reason"`
	AssumeMerchant   *string     `json:"assumeMerchant"`
}

type Get struct {
	api.Environment
	Params GetParams
}

func (req *Get) MakeRequest() *request.Request {
	return req.Environment.ToRequest(fmt.Sprintf("/v2/refunds/%s", url.QueryEscape(req.Params.MerchantRefundId)), "GET", nil, nil)
}
