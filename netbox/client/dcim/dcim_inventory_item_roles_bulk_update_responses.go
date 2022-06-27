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

// DcimInventoryItemRolesBulkUpdateReader is a Reader for the DcimInventoryItemRolesBulkUpdate structure.
type DcimInventoryItemRolesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemRolesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInventoryItemRolesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInventoryItemRolesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInventoryItemRolesBulkUpdateOK creates a DcimInventoryItemRolesBulkUpdateOK with default headers values
func NewDcimInventoryItemRolesBulkUpdateOK() *DcimInventoryItemRolesBulkUpdateOK {
	return &DcimInventoryItemRolesBulkUpdateOK{}
}

/* DcimInventoryItemRolesBulkUpdateOK describes a response with status code 200, with default header values.

DcimInventoryItemRolesBulkUpdateOK dcim inventory item roles bulk update o k
*/
type DcimInventoryItemRolesBulkUpdateOK struct {
	Payload *models.InventoryItemRole
}

func (o *DcimInventoryItemRolesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/inventory-item-roles/][%d] dcimInventoryItemRolesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimInventoryItemRolesBulkUpdateOK) GetPayload() *models.InventoryItemRole {
	return o.Payload
}

func (o *DcimInventoryItemRolesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItemRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInventoryItemRolesBulkUpdateDefault creates a DcimInventoryItemRolesBulkUpdateDefault with default headers values
func NewDcimInventoryItemRolesBulkUpdateDefault(code int) *DcimInventoryItemRolesBulkUpdateDefault {
	return &DcimInventoryItemRolesBulkUpdateDefault{
		_statusCode: code,
	}
}

/* DcimInventoryItemRolesBulkUpdateDefault describes a response with status code -1, with default header values.

DcimInventoryItemRolesBulkUpdateDefault dcim inventory item roles bulk update default
*/
type DcimInventoryItemRolesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim inventory item roles bulk update default response
func (o *DcimInventoryItemRolesBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimInventoryItemRolesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/inventory-item-roles/][%d] dcim_inventory-item-roles_bulk_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimInventoryItemRolesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemRolesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
