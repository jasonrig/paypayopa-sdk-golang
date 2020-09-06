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

package check_balance

import (
	"github.com/jasonrig/paypayopa-sdk-golang/api"
	"github.com/jasonrig/paypayopa-sdk-golang/request"
	"net/url"
	"reflect"
	"testing"
)

func TestGet_MakeRequest(t *testing.T) {
	testProductType := "fake product type"
	type fields struct {
		Environment api.Environment
		Params      GetParams
	}
	tests := []struct {
		name   string
		fields fields
		want   *request.Request
	}{
		{
			name: "A simple get request is constructed",
			fields: fields{
				Environment: api.SandboxEnvironment,
				Params: GetParams{
					UserAuthorizationId: "fake user auth",
					Amount:              100,
					Currency:            "JPY",
				},
			},
			want: &request.Request{
				Url: url.URL{
					Scheme:   "https",
					Host:     api.SandboxEnvironment.Host,
					Path:     "/v2/wallet/check_balance",
					RawQuery: "amount=100&currency=JPY&userAuthorizationId=fake+user+auth",
				},
				Method:      "GET",
				ContentType: "application/json",
				Body:        nil,
			},
		},
		{
			name: "A get request is constructed with product type",
			fields: fields{
				Environment: api.SandboxEnvironment,
				Params: GetParams{
					UserAuthorizationId: "fake user auth",
					Amount:              100,
					Currency:            "JPY",
					ProductType:         &testProductType,
				},
			},
			want: &request.Request{
				Url: url.URL{
					Scheme:   "https",
					Host:     api.SandboxEnvironment.Host,
					Path:     "/v2/wallet/check_balance",
					RawQuery: "amount=100&currency=JPY&productType=fake+product+type&userAuthorizationId=fake+user+auth",
				},
				Method:      "GET",
				ContentType: "application/json",
				Body:        nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &Get{
				Environment: tt.fields.Environment,
				Params:      tt.fields.Params,
			}
			if got := req.MakeRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
