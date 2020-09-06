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

package authorizations

import (
	"fmt"
	"github.com/jasonrig/paypayopa-sdk-golang/api"
	"github.com/jasonrig/paypayopa-sdk-golang/request"
	"net/url"
)

// API Endpoint: Delete /v2/user/authorizations/{userAuthorizationId}
// Unlink user
// - Unlink a user from the client
// Documentation references:
// * https://www.paypay.ne.jp/opa/doc/v1.0/direct_debit#operation/unlinkUser
// * https://www.paypay.ne.jp/opa/doc/v1.0/preauth_capture#operation/unlinkUser
// * https://www.paypay.ne.jp/opa/doc/v1.0/continuous_payments#operation/unlinkUser

// URL Parameters
type DeleteParams struct {
	UserAuthorizationId string
}

type Delete struct {
	api.Environment
	Params DeleteParams
}

func (req *Delete) MakeRequest() *request.Request {
	return req.Environment.ToRequest(fmt.Sprintf("/v2/user/authorizations/%s", url.QueryEscape(req.Params.UserAuthorizationId)), "DELETE", nil, nil)
}
