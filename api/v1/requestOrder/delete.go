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

package requestOrder

import (
	"fmt"
	"github.com/jasonrig/paypayopa-sdk-golang/api"
	"github.com/jasonrig/paypayopa-sdk-golang/request"
	"net/url"
)

// API Endpoint: Delete /v1/requestOrder/{merchantPaymentId}
// Cancel a pending order
// - This api is used delete the pending order
// Documentation references:
// * https://www.paypay.ne.jp/opa/doc/v1.0/pending_payments#operation/cancelPendingOrder

// URL Parameters
type DeleteParams struct {
	MerchantPaymentId string
}

type Delete struct {
	api.Environment
	Params DeleteParams
}

func (req *Delete) MakeRequest() *request.Request {
	return req.Environment.ToRequest(fmt.Sprintf("/v1/requestOrder/%s", url.QueryEscape(req.Params.MerchantPaymentId)), "DELETE", nil, nil)
}
