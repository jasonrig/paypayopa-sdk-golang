# Unofficial PayPay Golang SDK

This SDK provides a Golang API for accessing all API endpoints described in the following documentation:
* [Web Payment (1.1)](https://www.paypay.ne.jp/opa/doc/v1.0/webcashier)
* [Native Payment (1.1)](https://www.paypay.ne.jp/opa/doc/v1.0/direct_debit)
* [PreAuth & Capture (1.0)](https://www.paypay.ne.jp/opa/doc/v1.0/preauth_capture)
* [Dynamic QR (1.1)](https://www.paypay.ne.jp/opa/doc/v1.0/dynamicqrcode)
* [App Invoke (1.1)](https://www.paypay.ne.jp/opa/doc/v1.0/appinvoke)
* [Request Money (1.0)](https://www.paypay.ne.jp/opa/doc/v1.0/pending_payments)
* [Continuous Payment (1.0)](https://www.paypay.ne.jp/opa/doc/v1.0/continuous_payments)

## Implementation example
Please see GitHub repositiory [jasonrig/paypayopa-sdk-golang-example](https://github.com/jasonrig/paypayopa-sdk-golang-example)
for an implementation example showing how to generate a payment QR code.

## Project structure
The structure of this project follows the same structure as the documented API endpoints. For example,
the POST method for the [`/v2/codes`](https://www.paypay.ne.jp/opa/doc/v1.0/webcashier#operation/createQRCode) endpoint
is implemented in [api/v2/codes/post.go](api/v2/codes/post.go).

Each API endpoint is defined by a combination of structs:
* `MethodParams` (e.g. `GetParams`) - The URL parameters
* `MethodPayload` (e.g. `PostPayload`) - The request body
* `MethodResponse` (e.g. `PostResponse`) - The server response
* `Method` (e.g. `Post`) - Provides methods to execute the request

Please see GitHub repositiory [jasonrig/paypayopa-sdk-golang-example](https://github.com/jasonrig/paypayopa-sdk-golang-example)
for an implementation example, or check the tests within this repository.

### Official API documentation
Within each source file under the [api/](api/) directory contains references to the official PayPay API documentation.
Note that a REST endpoint may be referenced in more than one set of official PayPay documentation.

## Tests
Run the unit tests using the following command from within the root of this repository:
```shell script
PAYPAY_CLIENT_ID="xxxxxxx" PAYPAY_CLIENT_SECRET="xxxxxxx" go test -v ./...
```
Example output:
```
=== RUN   TestGenerateQR
=== RUN   TestGenerateQR/API_generates_a_payment_QR_code_for_a_valid_request
    integration_test.go:63: Got response https://qr-stg.sandbox.paypay.ne.jp/xxxxxxxxxx
=== RUN   TestGenerateQR/API_returns_an_error_when_an_invalid_request_is_made_(negative_order_price)
    integration_test.go:86: API correctly returned an error:
        Code: INVALID_REQUEST_PARAMS
        Message: Amount should be greater than 0 | [OPA0602A5F9AC8F41F78D686DC054D6522F]
        Error Id: 08100006
--- PASS: TestGenerateQR (2.73s)
    --- PASS: TestGenerateQR/API_generates_a_payment_QR_code_for_a_valid_request (2.28s)
    --- PASS: TestGenerateQR/API_returns_an_error_when_an_invalid_request_is_made_(negative_order_price) (0.44s)
PASS
ok      github.com/jasonrig/paypayopa-sdk-golang        4.001s
?       github.com/jasonrig/paypayopa-sdk-golang/api    [no test files]
=== RUN   TestDelete_MakeRequest
=== RUN   TestDelete_MakeRequest/A_simple_delete_request_is_constructed
--- PASS: TestDelete_MakeRequest (0.00s)
    --- PASS: TestDelete_MakeRequest/A_simple_delete_request_is_constructed (0.00s)
=== RUN   TestGet_MakeRequest
=== RUN   TestGet_MakeRequest/A_simple_get_request_is_constructed
--- PASS: TestGet_MakeRequest (0.00s)
    --- PASS: TestGet_MakeRequest/A_simple_get_request_is_constructed (0.00s)
=== RUN   TestPost_MakeRequest
=== RUN   TestPost_MakeRequest/A_simple_post_request_is_constructed
--- PASS: TestPost_MakeRequest (0.00s)
    --- PASS: TestPost_MakeRequest/A_simple_post_request_is_constructed (0.00s)
PASS
ok      github.com/jasonrig/paypayopa-sdk-golang/api/v1/code/topup      (cached)
=== RUN   TestDelete_MakeRequest
=== RUN   TestDelete_MakeRequest/A_simple_delete_request_is_constructed
--- PASS: TestDelete_MakeRequest (0.00s)
    --- PASS: TestDelete_MakeRequest/A_simple_delete_request_is_constructed (0.00s)
=== RUN   TestGet_MakeRequest
=== RUN   TestGet_MakeRequest/A_simple_get_request_is_constructed
--- PASS: TestGet_MakeRequest (0.00s)
    --- PASS: TestGet_MakeRequest/A_simple_get_request_is_constructed (0.00s)
=== RUN   TestPost_MakeRequest
=== RUN   TestPost_MakeRequest/A_simple_post_request_is_constructed
--- PASS: TestPost_MakeRequest (0.00s)
    --- PASS: TestPost_MakeRequest/A_simple_post_request_is_constructed (0.00s)
PASS
ok      github.com/jasonrig/paypayopa-sdk-golang/api/v1/requestOrder    (cached)
=== RUN   TestPost_MakeRequest
=== RUN   TestPost_MakeRequest/A_simple_post_request_is_constructed
--- PASS: TestPost_MakeRequest (0.00s)
    --- PASS: TestPost_MakeRequest/A_simple_post_request_is_constructed (0.00s)
PASS
ok      github.com/jasonrig/paypayopa-sdk-golang/api/v1/requestOrder/refunds    (cached)
=== RUN   TestPost_MakeRequest
=== RUN   TestPost_MakeRequest/A_simple_post_request_is_constructed
--- PASS: TestPost_MakeRequest (0.00s)
    --- PASS: TestPost_MakeRequest/A_simple_post_request_is_constructed (0.00s)
PASS
ok      github.com/jasonrig/paypayopa-sdk-golang/api/v1/subscription/payments   (cached)
=== RUN   TestDelete_MakeRequest
=== RUN   TestDelete_MakeRequest/A_simple_delete_request_is_constructed
--- PASS: TestDelete_MakeRequest (0.00s)
    --- PASS: TestDelete_MakeRequest/A_simple_delete_request_is_constructed (0.00s)
=== RUN   TestPost_MakeRequest
=== RUN   TestPost_MakeRequest/A_simple_post_request_is_constructed
--- PASS: TestPost_MakeRequest (0.00s)
    --- PASS: TestPost_MakeRequest/A_simple_post_request_is_constructed (0.00s)
PASS
ok      github.com/jasonrig/paypayopa-sdk-golang/api/v2/codes   (cached)
=== RUN   TestGet_MakeRequest
=== RUN   TestGet_MakeRequest/A_simple_get_request_is_constructed
--- PASS: TestGet_MakeRequest (0.00s)
    --- PASS: TestGet_MakeRequest/A_simple_get_request_is_constructed (0.00s)
PASS
ok      github.com/jasonrig/paypayopa-sdk-golang/api/v2/codes/payments  (cached)
=== RUN   TestDelete_MakeRequest
=== RUN   TestDelete_MakeRequest/A_simple_delete_request_is_constructed
--- PASS: TestDelete_MakeRequest (0.00s)
    --- PASS: TestDelete_MakeRequest/A_simple_delete_request_is_constructed (0.00s)
=== RUN   TestGet_MakeRequest
=== RUN   TestGet_MakeRequest/A_simple_get_request_is_constructed
--- PASS: TestGet_MakeRequest (0.00s)
    --- PASS: TestGet_MakeRequest/A_simple_get_request_is_constructed (0.00s)
=== RUN   TestPost_MakeRequest
=== RUN   TestPost_MakeRequest/A_simple_post_request_is_constructed
=== RUN   TestPost_MakeRequest/A_post_request_is_constructed_with_query_string
--- PASS: TestPost_MakeRequest (0.00s)
    --- PASS: TestPost_MakeRequest/A_simple_post_request_is_constructed (0.00s)
    --- PASS: TestPost_MakeRequest/A_post_request_is_constructed_with_query_string (0.00s)
PASS
ok      github.com/jasonrig/paypayopa-sdk-golang/api/v2/payments        (cached)
=== RUN   TestPost_MakeRequest
=== RUN   TestPost_MakeRequest/A_simple_post_request_is_constructed
--- PASS: TestPost_MakeRequest (0.00s)
    --- PASS: TestPost_MakeRequest/A_simple_post_request_is_constructed (0.00s)
PASS
ok      github.com/jasonrig/paypayopa-sdk-golang/api/v2/payments/capture        (cached)
=== RUN   TestPost_MakeRequest
=== RUN   TestPost_MakeRequest/A_simple_post_request_is_constructed
--- PASS: TestPost_MakeRequest (0.00s)
    --- PASS: TestPost_MakeRequest/A_simple_post_request_is_constructed (0.00s)
PASS
ok      github.com/jasonrig/paypayopa-sdk-golang/api/v2/payments/preauthorize   (cached)
=== RUN   TestPost_MakeRequest
=== RUN   TestPost_MakeRequest/A_simple_post_request_is_constructed
--- PASS: TestPost_MakeRequest (0.00s)
    --- PASS: TestPost_MakeRequest/A_simple_post_request_is_constructed (0.00s)
PASS
ok      github.com/jasonrig/paypayopa-sdk-golang/api/v2/payments/preauthorize/revert    (cached)
=== RUN   TestGet_MakeRequest
=== RUN   TestGet_MakeRequest/A_simple_get_request_is_constructed
--- PASS: TestGet_MakeRequest (0.00s)
    --- PASS: TestGet_MakeRequest/A_simple_get_request_is_constructed (0.00s)
=== RUN   TestPost_MakeRequest
=== RUN   TestPost_MakeRequest/A_simple_post_request_is_constructed
--- PASS: TestPost_MakeRequest (0.00s)
    --- PASS: TestPost_MakeRequest/A_simple_post_request_is_constructed (0.00s)
PASS
ok      github.com/jasonrig/paypayopa-sdk-golang/api/v2/refunds (cached)
=== RUN   TestDelete_MakeRequest
=== RUN   TestDelete_MakeRequest/A_simple_delete_request_is_constructed
--- PASS: TestDelete_MakeRequest (0.00s)
    --- PASS: TestDelete_MakeRequest/A_simple_delete_request_is_constructed (0.00s)
PASS
ok      github.com/jasonrig/paypayopa-sdk-golang/api/v2/user/authorizations     (cached)
=== RUN   TestGet_MakeRequest
=== RUN   TestGet_MakeRequest/A_simple_get_request_is_constructed
--- PASS: TestGet_MakeRequest (0.00s)
    --- PASS: TestGet_MakeRequest/A_simple_get_request_is_constructed (0.00s)
PASS
ok      github.com/jasonrig/paypayopa-sdk-golang/api/v2/user/profile/secure     (cached)
=== RUN   TestGet_MakeRequest
=== RUN   TestGet_MakeRequest/A_simple_get_request_is_constructed
=== RUN   TestGet_MakeRequest/A_get_request_is_constructed_with_product_type
--- PASS: TestGet_MakeRequest (0.00s)
    --- PASS: TestGet_MakeRequest/A_simple_get_request_is_constructed (0.00s)
    --- PASS: TestGet_MakeRequest/A_get_request_is_constructed_with_product_type (0.00s)
PASS
ok      github.com/jasonrig/paypayopa-sdk-golang/api/v2/wallet/check_balance    (cached)
=== RUN   TestGet_MakeRequest
=== RUN   TestGet_MakeRequest/A_simple_get_request_is_constructed
=== RUN   TestGet_MakeRequest/A_get_request_is_constructed_with_product_type
--- PASS: TestGet_MakeRequest (0.00s)
    --- PASS: TestGet_MakeRequest/A_simple_get_request_is_constructed (0.00s)
    --- PASS: TestGet_MakeRequest/A_get_request_is_constructed_with_product_type (0.00s)
PASS
ok      github.com/jasonrig/paypayopa-sdk-golang/api/v6/wallet/balance  (cached)
=== RUN   TestAuthentication_createAuthHeader
=== RUN   TestAuthentication_createAuthHeader/Auth_header_computed_correctly
=== RUN   TestAuthentication_createAuthHeader/Auth_header_computed_using_client_id_and_secret_from_environment_if_not_explicitly_provided
=== RUN   TestAuthentication_createAuthHeader/Auth_header_cannot_be_computed_when_credentials_are_missing
--- PASS: TestAuthentication_createAuthHeader (0.00s)
    --- PASS: TestAuthentication_createAuthHeader/Auth_header_computed_correctly (0.00s)
    --- PASS: TestAuthentication_createAuthHeader/Auth_header_computed_using_client_id_and_secret_from_environment_if_not_explicitly_provided (0.00s)
    --- PASS: TestAuthentication_createAuthHeader/Auth_header_cannot_be_computed_when_credentials_are_missing (0.00s)
PASS
ok      github.com/jasonrig/paypayopa-sdk-golang/request        0.927s
```

Omitting the environment variables will skip the integration test with the Sandbox API.