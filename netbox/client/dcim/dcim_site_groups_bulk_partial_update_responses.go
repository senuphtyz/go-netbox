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

// DcimSiteGroupsBulkPartialUpdateReader is a Reader for the DcimSiteGroupsBulkPartialUpdate structure.
type DcimSiteGroupsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSiteGroupsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimSiteGroupsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimSiteGroupsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimSiteGroupsBulkPartialUpdateOK creates a DcimSiteGroupsBulkPartialUpdateOK with default headers values
func NewDcimSiteGroupsBulkPartialUpdateOK() *DcimSiteGroupsBulkPartialUpdateOK {
	return &DcimSiteGroupsBulkPartialUpdateOK{}
}

/* DcimSiteGroupsBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimSiteGroupsBulkPartialUpdateOK dcim site groups bulk partial update o k
*/
type DcimSiteGroupsBulkPartialUpdateOK struct {
	Payload *models.SiteGroup
}

func (o *DcimSiteGroupsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/site-groups/][%d] dcimSiteGroupsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimSiteGroupsBulkPartialUpdateOK) GetPayload() *models.SiteGroup {
	return o.Payload
}

func (o *DcimSiteGroupsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SiteGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimSiteGroupsBulkPartialUpdateDefault creates a DcimSiteGroupsBulkPartialUpdateDefault with default headers values
func NewDcimSiteGroupsBulkPartialUpdateDefault(code int) *DcimSiteGroupsBulkPartialUpdateDefault {
	return &DcimSiteGroupsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* DcimSiteGroupsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimSiteGroupsBulkPartialUpdateDefault dcim site groups bulk partial update default
*/
type DcimSiteGroupsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim site groups bulk partial update default response
func (o *DcimSiteGroupsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimSiteGroupsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/site-groups/][%d] dcim_site-groups_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimSiteGroupsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimSiteGroupsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
