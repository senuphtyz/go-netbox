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

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// CircuitsCircuitTerminationsCreateReader is a Reader for the CircuitsCircuitTerminationsCreate structure.
type CircuitsCircuitTerminationsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTerminationsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCircuitsCircuitTerminationsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsCircuitTerminationsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsCircuitTerminationsCreateCreated creates a CircuitsCircuitTerminationsCreateCreated with default headers values
func NewCircuitsCircuitTerminationsCreateCreated() *CircuitsCircuitTerminationsCreateCreated {
	return &CircuitsCircuitTerminationsCreateCreated{}
}

/* CircuitsCircuitTerminationsCreateCreated describes a response with status code 201, with default header values.

CircuitsCircuitTerminationsCreateCreated circuits circuit terminations create created
*/
type CircuitsCircuitTerminationsCreateCreated struct {
	Payload *models.CircuitTermination
}

func (o *CircuitsCircuitTerminationsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /circuits/circuit-terminations/][%d] circuitsCircuitTerminationsCreateCreated  %+v", 201, o.Payload)
}
func (o *CircuitsCircuitTerminationsCreateCreated) GetPayload() *models.CircuitTermination {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsCircuitTerminationsCreateDefault creates a CircuitsCircuitTerminationsCreateDefault with default headers values
func NewCircuitsCircuitTerminationsCreateDefault(code int) *CircuitsCircuitTerminationsCreateDefault {
	return &CircuitsCircuitTerminationsCreateDefault{
		_statusCode: code,
	}
}

/* CircuitsCircuitTerminationsCreateDefault describes a response with status code -1, with default header values.

CircuitsCircuitTerminationsCreateDefault circuits circuit terminations create default
*/
type CircuitsCircuitTerminationsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits circuit terminations create default response
func (o *CircuitsCircuitTerminationsCreateDefault) Code() int {
	return o._statusCode
}

func (o *CircuitsCircuitTerminationsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /circuits/circuit-terminations/][%d] circuits_circuit-terminations_create default  %+v", o._statusCode, o.Payload)
}
func (o *CircuitsCircuitTerminationsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
