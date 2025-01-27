// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package secure

import (
	"github.com/google/go-safeweb/safehttp"
	"github.com/google/safehtml"
)

// TODO(clap|kele): comment on custom responses and error responses.

type ErrorResponse struct {
	code    safehttp.StatusCode
	message safehtml.HTML
}

func NewErrorResponse(code safehttp.StatusCode, message safehtml.HTML) ErrorResponse {
	return ErrorResponse{
		code:    code,
		message: message,
	}
}

func (e ErrorResponse) Code() safehttp.StatusCode {
	return e.code
}
