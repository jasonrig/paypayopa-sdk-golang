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

package revert

import (
	"net/url"
	"paypayopa-sdk/api"
	"paypayopa-sdk/request"
	"reflect"
	"testing"
)

func TestPost_MakeRequest(t *testing.T) {
	var expectedPayload interface{} = PostPayload{
		MerchantRevertId: "fake revert id",
		PaymentId:        "fake payment id",
		RequestedAt:      1598797689,
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
					MerchantRevertId: "fake revert id",
					PaymentId:        "fake payment id",
					RequestedAt:      1598797689,
				},
			},
			want: &request.Request{
				Url: url.URL{
					Scheme: "https",
					Host:   api.SandboxEnvironment.Host,
					Path:   "/v2/payments/preauthorize/revert",
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
