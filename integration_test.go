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

package paypayopa_sdk_golang

import (
	"github.com/google/uuid"
	"net/http"
	"paypayopa-sdk/api"
	"paypayopa-sdk/api/v2/codes"
	"paypayopa-sdk/request"
	"testing"
)

func TestGenerateQR(t *testing.T) {
	auth, err := request.NewAuth(nil, nil)
	if err != nil {
		t.Skipf("No authentication, skipping: %s", err)
	}

	t.Run("API generates a payment QR code for a valid request", func(t *testing.T) {
		paymentId, _ := uuid.NewRandom()
		apiRequest := &codes.Post{
			Environment: api.SandboxEnvironment,
			Payload: codes.PostPayload{
				MerchantPaymentId: paymentId.String(),
				Amount: api.Amount{
					Amount:   100,
					Currency: "JPY",
				},
				CodeType: "ORDER_QR",
				OrderItems: &[]api.OrderItems{
					{
						Name: "Fun thing",
						UnitPrice: &api.Amount{
							Amount:   100,
							Currency: "JPY",
						},
						Quantity: 1,
					},
				},
			},
		}

		response := &codes.PostResponse{}
		err = apiRequest.MakeRequest().Call(auth, &http.Client{}, response)
		if err != nil {
			t.Errorf("API call returned an error, %s", err)
		} else {
			t.Logf("Got response %s", *response.Url)
		}
	})

	t.Run("API returns an error when an invalid request is made (negative order price)", func(t *testing.T) {
		paymentId, _ := uuid.NewRandom()
		apiRequest := &codes.Post{
			Environment: api.SandboxEnvironment,
			Payload: codes.PostPayload{
				MerchantPaymentId: paymentId.String(),
				Amount: api.Amount{
					Amount:   -1,
					Currency: "JPY",
				},
				CodeType: "ORDER_QR",
			},
		}

		response := &codes.PostResponse{}
		err = apiRequest.MakeRequest().Call(auth, &http.Client{}, response)
		if err == nil {
			t.Error("API did not return an error when it should have")
		} else {
			t.Logf("API correctly returned an error:\n%s", err)
		}
	})
}
