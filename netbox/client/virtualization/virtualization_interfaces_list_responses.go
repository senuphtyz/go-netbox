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

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// VirtualizationInterfacesListReader is a Reader for the VirtualizationInterfacesList structure.
type VirtualizationInterfacesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationInterfacesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationInterfacesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationInterfacesListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationInterfacesListOK creates a VirtualizationInterfacesListOK with default headers values
func NewVirtualizationInterfacesListOK() *VirtualizationInterfacesListOK {
	return &VirtualizationInterfacesListOK{}
}

/* VirtualizationInterfacesListOK describes a response with status code 200, with default header values.

VirtualizationInterfacesListOK virtualization interfaces list o k
*/
type VirtualizationInterfacesListOK struct {
	Payload *VirtualizationInterfacesListOKBody
}

func (o *VirtualizationInterfacesListOK) Error() string {
	return fmt.Sprintf("[GET /virtualization/interfaces/][%d] virtualizationInterfacesListOK  %+v", 200, o.Payload)
}
func (o *VirtualizationInterfacesListOK) GetPayload() *VirtualizationInterfacesListOKBody {
	return o.Payload
}

func (o *VirtualizationInterfacesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(VirtualizationInterfacesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationInterfacesListDefault creates a VirtualizationInterfacesListDefault with default headers values
func NewVirtualizationInterfacesListDefault(code int) *VirtualizationInterfacesListDefault {
	return &VirtualizationInterfacesListDefault{
		_statusCode: code,
	}
}

/* VirtualizationInterfacesListDefault describes a response with status code -1, with default header values.

VirtualizationInterfacesListDefault virtualization interfaces list default
*/
type VirtualizationInterfacesListDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization interfaces list default response
func (o *VirtualizationInterfacesListDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationInterfacesListDefault) Error() string {
	return fmt.Sprintf("[GET /virtualization/interfaces/][%d] virtualization_interfaces_list default  %+v", o._statusCode, o.Payload)
}
func (o *VirtualizationInterfacesListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationInterfacesListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*VirtualizationInterfacesListOKBody virtualization interfaces list o k body
swagger:model VirtualizationInterfacesListOKBody
*/
type VirtualizationInterfacesListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*models.VMInterface `json:"results"`
}

// Validate validates this virtualization interfaces list o k body
func (o *VirtualizationInterfacesListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VirtualizationInterfacesListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("virtualizationInterfacesListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *VirtualizationInterfacesListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("virtualizationInterfacesListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *VirtualizationInterfacesListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("virtualizationInterfacesListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *VirtualizationInterfacesListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("virtualizationInterfacesListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("virtualizationInterfacesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("virtualizationInterfacesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this virtualization interfaces list o k body based on the context it is used
func (o *VirtualizationInterfacesListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VirtualizationInterfacesListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("virtualizationInterfacesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("virtualizationInterfacesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *VirtualizationInterfacesListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VirtualizationInterfacesListOKBody) UnmarshalBinary(b []byte) error {
	var res VirtualizationInterfacesListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
