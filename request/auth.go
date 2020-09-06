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
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"os"
	"strconv"
	"strings"
	"time"
)

// Authentication parameters
type Authentication struct {
	clientId     string
	clientSecret string
}

// Prepares authentication parameters for API requests
// If clientId or clientSecret are nil, the parameters are loaded from
// environment variables PAYPAY_CLIENT_ID and PAYPAY_CLIENT_SECRET, respectively.
func NewAuth(clientId *string, clientSecret *string) (*Authentication, error) {
	var clientId0 string
	var clientSecret0 string
	if clientId == nil {
		var present bool
		clientId0, present = os.LookupEnv("PAYPAY_CLIENT_ID")
		if !present {
			return nil, errors.New("missing client id - hint: try setting the PAYPAY_CLIENT_ID environment variable")
		}
	} else {
		clientId0 = *clientId
	}
	if clientSecret == nil {
		var present bool
		clientSecret0, present = os.LookupEnv("PAYPAY_CLIENT_SECRET")
		if !present {
			return nil, errors.New("missing client secret - hint: try setting the PAYPAY_CLIENT_SECRET environment variable")
		}
	} else {
		clientSecret0 = *clientSecret
	}
	return &Authentication{
		clientId:     clientId0,
		clientSecret: clientSecret0,
	}, nil
}

// Generates the Authorization header for a given request
func (auth *Authentication) CreateAuthHeader(request *Request) (*string, error) {
	epoch := time.Now().Unix()
	nonce, _ := uuid.NewRandom()
	return auth.createAuthHeader(request, nonce.String(), epoch)
}

func (auth *Authentication) createAuthHeader(request *Request, nonce string, epoch int64) (*string, error) {
	body, err := json.Marshal(request.Body)
	if err != nil {
		return nil, err
	}
	md5HashFn := md5.New()
	md5HashFn.Write([]byte(request.ContentType))
	md5HashFn.Write(body)
	hashBytes := md5HashFn.Sum(nil)
	payloadHash := base64.StdEncoding.EncodeToString(hashBytes)

	epochStr := strconv.FormatInt(epoch, 10)
	signatureList := strings.Join([]string{
		request.Url.Path,
		request.Method,
		nonce,
		epochStr,
		request.ContentType,
		payloadHash,
	}, "\n")

	hmacFn := hmac.New(sha256.New, []byte(auth.clientSecret))
	hmacFn.Write([]byte(signatureList))
	signatureHash := base64.StdEncoding.EncodeToString(hmacFn.Sum(nil))
	header := strings.Join([]string{
		auth.clientId,
		signatureHash,
		nonce,
		epochStr,
		payloadHash,
	}, ":")
	header = fmt.Sprintf("hmac OPA-Auth:%s", header)
	return &header, nil
}
