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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// DcimFrontPortsUpdateReader is a Reader for the DcimFrontPortsUpdate structure.
type DcimFrontPortsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimFrontPortsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimFrontPortsUpdateOK creates a DcimFrontPortsUpdateOK with default headers values
func NewDcimFrontPortsUpdateOK() *DcimFrontPortsUpdateOK {
	return &DcimFrontPortsUpdateOK{}
}

/* DcimFrontPortsUpdateOK describes a response with status code 200, with default header values.

DcimFrontPortsUpdateOK dcim front ports update o k
*/
type DcimFrontPortsUpdateOK struct {
	Payload *models.FrontPort
}

func (o *DcimFrontPortsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/front-ports/{id}/][%d] dcimFrontPortsUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimFrontPortsUpdateOK) GetPayload() *models.FrontPort {
	return o.Payload
}

func (o *DcimFrontPortsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimFrontPortsUpdateDefault creates a DcimFrontPortsUpdateDefault with default headers values
func NewDcimFrontPortsUpdateDefault(code int) *DcimFrontPortsUpdateDefault {
	return &DcimFrontPortsUpdateDefault{
		_statusCode: code,
	}
}

/* DcimFrontPortsUpdateDefault describes a response with status code -1, with default header values.

DcimFrontPortsUpdateDefault dcim front ports update default
*/
type DcimFrontPortsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim front ports update default response
func (o *DcimFrontPortsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimFrontPortsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/front-ports/{id}/][%d] dcim_front-ports_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimFrontPortsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimFrontPortsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
