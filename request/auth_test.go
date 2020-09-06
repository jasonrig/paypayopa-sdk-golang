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

package request

import (
	"net/url"
	"os"
	"testing"
)

type testBody struct {
	A string `json:"a"`
}

func TestAuthentication_createAuthHeader(t *testing.T) {
	var body interface{} = &testBody{
		"test",
	}
	t.Run("Auth header computed correctly", func(t *testing.T) {
		epoch := int64(1598797689)
		nonce := "a0dce033-4f5e-428b-b4d2-a8c167c1192e"
		clientId := "abc"
		clientSecret := "def"
		request := &Request{
			Method: "get",
			Url: url.URL{
				Path: "/",
			},
			ContentType: "application/json",
			Body:        &body,
		}
		auth, _ := NewAuth(&clientId, &clientSecret)
		expectedHeader := "hmac OPA-Auth:abc:4ecQ6PEHbkIW4XrIv1YNZpZqxZg1SmYa+B/OAWnF1lI=:a0dce033-4f5e-428b-b4d2-a8c167c1192e:1598797689:nDPsVlcbd45BVy/vgnhHJg=="
		if got, _ := auth.createAuthHeader(request, nonce, epoch); *got != expectedHeader {
			t.Errorf("createAuthHeader() = \n %v, \n want \n %v", *got, expectedHeader)
		}
	})

	t.Run("Auth header computed using client id and secret from environment if not explicitly provided", func(t *testing.T) {

		err := os.Setenv("PAYPAY_CLIENT_ID", "abc")
		if err != nil {
			panic(err)
		}
		err = os.Setenv("PAYPAY_CLIENT_SECRET", "def")
		if err != nil {
			panic(err)
		}

		epoch := int64(1598797689)
		nonce := "a0dce033-4f5e-428b-b4d2-a8c167c1192e"
		request := &Request{
			Method: "get",
			Url: url.URL{
				Path: "/",
			},
			ContentType: "application/json",
			Body:        &body,
		}
		auth, _ := NewAuth(nil, nil)
		expectedHeader := "hmac OPA-Auth:abc:4ecQ6PEHbkIW4XrIv1YNZpZqxZg1SmYa+B/OAWnF1lI=:a0dce033-4f5e-428b-b4d2-a8c167c1192e:1598797689:nDPsVlcbd45BVy/vgnhHJg=="
		if got, _ := auth.createAuthHeader(request, nonce, epoch); *got != expectedHeader {
			t.Errorf("createAuthHeader() = \n %v, \n want \n %v", *got, expectedHeader)
		}
	})

	t.Run("Auth header cannot be computed when credentials are missing", func(t *testing.T) {
		err := os.Unsetenv("PAYPAY_CLIENT_ID")
		if err != nil {
			panic(err)
		}
		err = os.Unsetenv("PAYPAY_CLIENT_SECRET")
		if err != nil {
			panic(err)
		}
		auth, err := NewAuth(nil, nil)
		if err == nil {
			t.Errorf("NewAuth() = \n %v, \n want \n %v", auth, nil)
		}
	})
}
