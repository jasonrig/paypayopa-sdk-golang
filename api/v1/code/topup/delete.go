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

// API Endpoint: Delete /v1/code/topup/{codeId}
// Delete QR code
// - Delete the topup QR code
// Documentation references:
// * https://www.paypay.ne.jp/opa/doc/v1.0/direct_debit#operation/deleteTopUpQrCode
// * https://www.paypay.ne.jp/opa/doc/v1.0/preauth_capture#operation/deleteTopUpQrCode

// URL Parameters
type DeleteParams struct {
	CodeId string
}

// Server response
type DeleteResponse struct {
}

type Delete struct {
	api.Environment
	Params DeleteParams
}

func (req *Delete) MakeRequest() *request.Request {
	return req.Environment.ToRequest(fmt.Sprintf("/v1/code/topup/%s", url.QueryEscape(req.Params.CodeId)), "DELETE", nil, nil)
}
