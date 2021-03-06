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
	"fmt"
	"github.com/jasonrig/paypayopa-sdk-golang/api"
	"github.com/jasonrig/paypayopa-sdk-golang/request"
	"net/url"
)

// API Endpoint: Delete /v2/codes/{codeId}
// Delete a Code
// - Delete a created Code.
// Documentation references:
// * https://www.paypay.ne.jp/opa/doc/v1.0/webcashier#operation/deleteQRCode
// * https://www.paypay.ne.jp/opa/doc/v1.0/dynamicqrcode#operation/deleteQRCode
// * https://www.paypay.ne.jp/opa/doc/v1.0/appinvoke#operation/deleteQRCode

// URL Parameters
type DeleteParams struct {
	CodeId string
}

type Delete struct {
	api.Environment
	Params DeleteParams
}

func (req *Delete) MakeRequest() *request.Request {
	return req.Environment.ToRequest(fmt.Sprintf("/v2/codes/%s", url.QueryEscape(req.Params.CodeId)), "DELETE", nil, nil)
}
