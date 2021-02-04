// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasScriptsReadReader is a Reader for the ExtrasScriptsRead structure.
type ExtrasScriptsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasScriptsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasScriptsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasScriptsReadOK creates a ExtrasScriptsReadOK with default headers values
func NewExtrasScriptsReadOK() *ExtrasScriptsReadOK {
	return &ExtrasScriptsReadOK{}
}

/* ExtrasScriptsReadOK describes a response with status code 200, with default header values.

ExtrasScriptsReadOK extras scripts read o k
*/
type ExtrasScriptsReadOK struct {
}

func (o *ExtrasScriptsReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/scripts/{id}/][%d] extrasScriptsReadOK ", 200)
}

func (o *ExtrasScriptsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
