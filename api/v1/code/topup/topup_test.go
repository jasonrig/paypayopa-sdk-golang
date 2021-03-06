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
	"github.com/jasonrig/paypayopa-sdk-golang/api"
	"github.com/jasonrig/paypayopa-sdk-golang/request"
	"net/url"
	"reflect"
	"testing"
)

func TestDelete_MakeRequest(t *testing.T) {
	type fields struct {
		Environment api.Environment
		Params      DeleteParams
	}
	tests := []struct {
		name   string
		fields fields
		want   *request.Request
	}{
		{
			name: "A simple delete request is constructed",
			fields: fields{
				Environment: api.SandboxEnvironment,
				Params: DeleteParams{
					CodeId: "fake code id",
				},
			},
			want: &request.Request{
				Url: url.URL{
					Scheme: "https",
					Host:   api.SandboxEnvironment.Host,
					Path:   "/v1/code/topup/fake+code+id",
				},
				Method:      "DELETE",
				ContentType: "application/json",
				Body:        nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &Delete{
				Environment: tt.fields.Environment,
				Params:      tt.fields.Params,
			}
			got := req.MakeRequest()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGet_MakeRequest(t *testing.T) {
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
					MerchantTopUpId: "fake topup id",
				},
			},
			want: &request.Request{
				Url: url.URL{
					Scheme: "https",
					Host:   api.SandboxEnvironment.Host,
					Path:   "/v1/code/topup/fake+topup+id",
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
			got := req.MakeRequest()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPost_MakeRequest(t *testing.T) {
	var expectedPayload interface{} = PostPayload{
		MerchantTopUpId:     "fake topup id",
		UserAuthorizationId: "fake user authorization",
		MinimumTopUpAmount: api.Amount{
			Amount:   100,
			Currency: "JPY",
		},
		CodeType: "TOPUP_QR",
	}

	type fields struct {
		Environment api.Environment
		Payload     PostPayload
	}
	tests := []struct {
		name   string
		fields fields
		want   *request.Request
	}{
		{
			name: "A simple post request is constructed",
			fields: fields{
				Environment: api.SandboxEnvironment,
				Payload: PostPayload{
					MerchantTopUpId:     "fake topup id",
					UserAuthorizationId: "fake user authorization",
					MinimumTopUpAmount: api.Amount{
						Amount:   100,
						Currency: "JPY",
					},
					CodeType: "TOPUP_QR",
				},
			},
			want: &request.Request{
				Url: url.URL{
					Scheme: "https",
					Host:   api.SandboxEnvironment.Host,
					Path:   "/v1/code/topup",
				},
				Method:      "POST",
				ContentType: "application/json",
				Body:        &expectedPayload,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &Post{
				Environment: tt.fields.Environment,
				Payload:     tt.fields.Payload,
			}
			got := req.MakeRequest()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}
