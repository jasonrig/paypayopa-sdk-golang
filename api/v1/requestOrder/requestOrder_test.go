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
	"net/url"
	"paypayopa-sdk/api"
	"paypayopa-sdk/request"
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
					MerchantPaymentId: "fake payment id",
				},
			},
			want: &request.Request{
				Url: url.URL{
					Scheme: "https",
					Host:   api.SandboxEnvironment.Host,
					Path:   "/v1/requestOrder/fake+payment+id",
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
			if got := req.MakeRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeRequest() = %v, want %v", got, tt.want)
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
					MerchantPaymentId: "fake payment id",
				},
			},
			want: &request.Request{
				Url: url.URL{
					Scheme: "https",
					Host:   api.SandboxEnvironment.Host,
					Path:   "/v1/requestOrder/fake+payment+id",
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

func TestPost_MakeRequest(t *testing.T) {
	var expectedPayload interface{} = PostPayload{
		MerchantPaymentId:   "fake payment id",
		UserAuthorizationId: "fake user auth",
		Amount: api.Amount{
			Amount:   100,
			Currency: "JPY",
		},
		RequestedAt: 1598797689,
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
					MerchantPaymentId:   "fake payment id",
					UserAuthorizationId: "fake user auth",
					Amount: api.Amount{
						Amount:   100,
						Currency: "JPY",
					},
					RequestedAt: 1598797689,
				},
			},
			want: &request.Request{
				Url: url.URL{
					Scheme: "https",
					Host:   api.SandboxEnvironment.Host,
					Path:   "/v1/requestOrder",
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
			if got := req.MakeRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
