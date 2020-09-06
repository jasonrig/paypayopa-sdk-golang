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