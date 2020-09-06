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

// API Endpoint: Delete /v2/payments/{merchantPaymentId}
// Cancel a payment
// - This api is used in case, while creating a payment, the client can not determine the status of the payment.
//   For example, client get timeout or the response cannot contain the information to indicate the exact payment status.
//   By calling this api, if accepted, the OPA will guarantee the money eventually goes back to user's account.
// - Note: The Cancel API can be used until 00:14:59 AM the day after the Payment has happened.
//   For 00:15 AM or later, please call the refund API to refund the payment.
// Documentation references:
// * https://www.paypay.ne.jp/opa/doc/v1.0/webcashier#operation/cancelPayment
// * https://www.paypay.ne.jp/opa/doc/v1.0/direct_debit#operation/cancelPayment
// * https://www.paypay.ne.jp/opa/doc/v1.0/preauth_capture#operation/cancelPayment
// * https://www.paypay.ne.jp/opa/doc/v1.0/dynamicqrcode#operation/cancelPayment
// * https://www.paypay.ne.jp/opa/doc/v1.0/appinvoke#operation/cancelPayment
// * https://www.paypay.ne.jp/opa/doc/v1.0/continuous_payments#operation/cancelPayment

// URL Parameters
type DeleteParams struct {
	MerchantPaymentId string
}

type Delete struct {
	api.Environment
	Params DeleteParams
}

func (req *Delete) MakeRequest() *request.Request {
	return req.Environment.ToRequest(fmt.Sprintf("/v2/payments/%s", url.QueryEscape(req.Params.MerchantPaymentId)), "DELETE", nil, nil)
}
