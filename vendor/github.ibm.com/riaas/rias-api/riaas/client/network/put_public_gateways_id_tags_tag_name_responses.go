// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// PutPublicGatewaysIDTagsTagNameReader is a Reader for the PutPublicGatewaysIDTagsTagName structure.
type PutPublicGatewaysIDTagsTagNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutPublicGatewaysIDTagsTagNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPutPublicGatewaysIDTagsTagNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutPublicGatewaysIDTagsTagNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutPublicGatewaysIDTagsTagNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutPublicGatewaysIDTagsTagNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewPutPublicGatewaysIDTagsTagNameNoContent creates a PutPublicGatewaysIDTagsTagNameNoContent with default headers values
func NewPutPublicGatewaysIDTagsTagNameNoContent() *PutPublicGatewaysIDTagsTagNameNoContent {
	return &PutPublicGatewaysIDTagsTagNameNoContent{}
}

/*PutPublicGatewaysIDTagsTagNameNoContent handles this case with default header values.

error
*/
type PutPublicGatewaysIDTagsTagNameNoContent struct {
	Payload *models.Riaaserror
}

func (o *PutPublicGatewaysIDTagsTagNameNoContent) Error() string {
	return fmt.Sprintf("[PUT /public_gateways/{id}/tags/{tag_name}][%d] putPublicGatewaysIdTagsTagNameNoContent  %+v", 204, o.Payload)
}

func (o *PutPublicGatewaysIDTagsTagNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutPublicGatewaysIDTagsTagNameBadRequest creates a PutPublicGatewaysIDTagsTagNameBadRequest with default headers values
func NewPutPublicGatewaysIDTagsTagNameBadRequest() *PutPublicGatewaysIDTagsTagNameBadRequest {
	return &PutPublicGatewaysIDTagsTagNameBadRequest{}
}

/*PutPublicGatewaysIDTagsTagNameBadRequest handles this case with default header values.

error
*/
type PutPublicGatewaysIDTagsTagNameBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PutPublicGatewaysIDTagsTagNameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /public_gateways/{id}/tags/{tag_name}][%d] putPublicGatewaysIdTagsTagNameBadRequest  %+v", 400, o.Payload)
}

func (o *PutPublicGatewaysIDTagsTagNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutPublicGatewaysIDTagsTagNameNotFound creates a PutPublicGatewaysIDTagsTagNameNotFound with default headers values
func NewPutPublicGatewaysIDTagsTagNameNotFound() *PutPublicGatewaysIDTagsTagNameNotFound {
	return &PutPublicGatewaysIDTagsTagNameNotFound{}
}

/*PutPublicGatewaysIDTagsTagNameNotFound handles this case with default header values.

error
*/
type PutPublicGatewaysIDTagsTagNameNotFound struct {
	Payload *models.Riaaserror
}

func (o *PutPublicGatewaysIDTagsTagNameNotFound) Error() string {
	return fmt.Sprintf("[PUT /public_gateways/{id}/tags/{tag_name}][%d] putPublicGatewaysIdTagsTagNameNotFound  %+v", 404, o.Payload)
}

func (o *PutPublicGatewaysIDTagsTagNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutPublicGatewaysIDTagsTagNameDefault creates a PutPublicGatewaysIDTagsTagNameDefault with default headers values
func NewPutPublicGatewaysIDTagsTagNameDefault(code int) *PutPublicGatewaysIDTagsTagNameDefault {
	return &PutPublicGatewaysIDTagsTagNameDefault{
		_statusCode: code,
	}
}

/*PutPublicGatewaysIDTagsTagNameDefault handles this case with default header values.

unexpectederror
*/
type PutPublicGatewaysIDTagsTagNameDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the put public gateways ID tags tag name default response
func (o *PutPublicGatewaysIDTagsTagNameDefault) Code() int {
	return o._statusCode
}

func (o *PutPublicGatewaysIDTagsTagNameDefault) Error() string {
	return fmt.Sprintf("[PUT /public_gateways/{id}/tags/{tag_name}][%d] PutPublicGatewaysIDTagsTagName default  %+v", o._statusCode, o.Payload)
}

func (o *PutPublicGatewaysIDTagsTagNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
