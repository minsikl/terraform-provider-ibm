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

// GetNetworkAclsIDTagsTagNameReader is a Reader for the GetNetworkAclsIDTagsTagName structure.
type GetNetworkAclsIDTagsTagNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkAclsIDTagsTagNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewGetNetworkAclsIDTagsTagNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetNetworkAclsIDTagsTagNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetNetworkAclsIDTagsTagNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewGetNetworkAclsIDTagsTagNameNoContent creates a GetNetworkAclsIDTagsTagNameNoContent with default headers values
func NewGetNetworkAclsIDTagsTagNameNoContent() *GetNetworkAclsIDTagsTagNameNoContent {
	return &GetNetworkAclsIDTagsTagNameNoContent{}
}

/*GetNetworkAclsIDTagsTagNameNoContent handles this case with default header values.

error
*/
type GetNetworkAclsIDTagsTagNameNoContent struct {
	Payload *models.Riaaserror
}

func (o *GetNetworkAclsIDTagsTagNameNoContent) Error() string {
	return fmt.Sprintf("[GET /network_acls/{id}/tags/{tag_name}][%d] getNetworkAclsIdTagsTagNameNoContent  %+v", 204, o.Payload)
}

func (o *GetNetworkAclsIDTagsTagNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkAclsIDTagsTagNameNotFound creates a GetNetworkAclsIDTagsTagNameNotFound with default headers values
func NewGetNetworkAclsIDTagsTagNameNotFound() *GetNetworkAclsIDTagsTagNameNotFound {
	return &GetNetworkAclsIDTagsTagNameNotFound{}
}

/*GetNetworkAclsIDTagsTagNameNotFound handles this case with default header values.

error
*/
type GetNetworkAclsIDTagsTagNameNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetNetworkAclsIDTagsTagNameNotFound) Error() string {
	return fmt.Sprintf("[GET /network_acls/{id}/tags/{tag_name}][%d] getNetworkAclsIdTagsTagNameNotFound  %+v", 404, o.Payload)
}

func (o *GetNetworkAclsIDTagsTagNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkAclsIDTagsTagNameDefault creates a GetNetworkAclsIDTagsTagNameDefault with default headers values
func NewGetNetworkAclsIDTagsTagNameDefault(code int) *GetNetworkAclsIDTagsTagNameDefault {
	return &GetNetworkAclsIDTagsTagNameDefault{
		_statusCode: code,
	}
}

/*GetNetworkAclsIDTagsTagNameDefault handles this case with default header values.

unexpectederror
*/
type GetNetworkAclsIDTagsTagNameDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the get network acls ID tags tag name default response
func (o *GetNetworkAclsIDTagsTagNameDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworkAclsIDTagsTagNameDefault) Error() string {
	return fmt.Sprintf("[GET /network_acls/{id}/tags/{tag_name}][%d] GetNetworkAclsIDTagsTagName default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetworkAclsIDTagsTagNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
