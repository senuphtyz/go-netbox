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

// DcimConsoleServerPortsBulkPartialUpdateReader is a Reader for the DcimConsoleServerPortsBulkPartialUpdate structure.
type DcimConsoleServerPortsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsoleServerPortsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsoleServerPortsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsoleServerPortsBulkPartialUpdateOK creates a DcimConsoleServerPortsBulkPartialUpdateOK with default headers values
func NewDcimConsoleServerPortsBulkPartialUpdateOK() *DcimConsoleServerPortsBulkPartialUpdateOK {
	return &DcimConsoleServerPortsBulkPartialUpdateOK{}
}

/* DcimConsoleServerPortsBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimConsoleServerPortsBulkPartialUpdateOK dcim console server ports bulk partial update o k
*/
type DcimConsoleServerPortsBulkPartialUpdateOK struct {
	Payload *models.ConsoleServerPort
}

func (o *DcimConsoleServerPortsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/console-server-ports/][%d] dcimConsoleServerPortsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimConsoleServerPortsBulkPartialUpdateOK) GetPayload() *models.ConsoleServerPort {
	return o.Payload
}

func (o *DcimConsoleServerPortsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsoleServerPortsBulkPartialUpdateDefault creates a DcimConsoleServerPortsBulkPartialUpdateDefault with default headers values
func NewDcimConsoleServerPortsBulkPartialUpdateDefault(code int) *DcimConsoleServerPortsBulkPartialUpdateDefault {
	return &DcimConsoleServerPortsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* DcimConsoleServerPortsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimConsoleServerPortsBulkPartialUpdateDefault dcim console server ports bulk partial update default
*/
type DcimConsoleServerPortsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console server ports bulk partial update default response
func (o *DcimConsoleServerPortsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimConsoleServerPortsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/console-server-ports/][%d] dcim_console-server-ports_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimConsoleServerPortsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsoleServerPortsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
