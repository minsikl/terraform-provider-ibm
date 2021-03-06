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

// DeleteNetworkAclsIDTagsTagNameReader is a Reader for the DeleteNetworkAclsIDTagsTagName structure.
type DeleteNetworkAclsIDTagsTagNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkAclsIDTagsTagNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteNetworkAclsIDTagsTagNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteNetworkAclsIDTagsTagNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteNetworkAclsIDTagsTagNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewDeleteNetworkAclsIDTagsTagNameNoContent creates a DeleteNetworkAclsIDTagsTagNameNoContent with default headers values
func NewDeleteNetworkAclsIDTagsTagNameNoContent() *DeleteNetworkAclsIDTagsTagNameNoContent {
	return &DeleteNetworkAclsIDTagsTagNameNoContent{}
}

/*DeleteNetworkAclsIDTagsTagNameNoContent handles this case with default header values.

error
*/
type DeleteNetworkAclsIDTagsTagNameNoContent struct {
	Payload *models.Riaaserror
}

func (o *DeleteNetworkAclsIDTagsTagNameNoContent) Error() string {
	return fmt.Sprintf("[DELETE /network_acls/{id}/tags/{tag_name}][%d] deleteNetworkAclsIdTagsTagNameNoContent  %+v", 204, o.Payload)
}

func (o *DeleteNetworkAclsIDTagsTagNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkAclsIDTagsTagNameNotFound creates a DeleteNetworkAclsIDTagsTagNameNotFound with default headers values
func NewDeleteNetworkAclsIDTagsTagNameNotFound() *DeleteNetworkAclsIDTagsTagNameNotFound {
	return &DeleteNetworkAclsIDTagsTagNameNotFound{}
}

/*DeleteNetworkAclsIDTagsTagNameNotFound handles this case with default header values.

error
*/
type DeleteNetworkAclsIDTagsTagNameNotFound struct {
	Payload *models.Riaaserror
}

func (o *DeleteNetworkAclsIDTagsTagNameNotFound) Error() string {
	return fmt.Sprintf("[DELETE /network_acls/{id}/tags/{tag_name}][%d] deleteNetworkAclsIdTagsTagNameNotFound  %+v", 404, o.Payload)
}

func (o *DeleteNetworkAclsIDTagsTagNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkAclsIDTagsTagNameDefault creates a DeleteNetworkAclsIDTagsTagNameDefault with default headers values
func NewDeleteNetworkAclsIDTagsTagNameDefault(code int) *DeleteNetworkAclsIDTagsTagNameDefault {
	return &DeleteNetworkAclsIDTagsTagNameDefault{
		_statusCode: code,
	}
}

/*DeleteNetworkAclsIDTagsTagNameDefault handles this case with default header values.

unexpectederror
*/
type DeleteNetworkAclsIDTagsTagNameDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the delete network acls ID tags tag name default response
func (o *DeleteNetworkAclsIDTagsTagNameDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNetworkAclsIDTagsTagNameDefault) Error() string {
	return fmt.Sprintf("[DELETE /network_acls/{id}/tags/{tag_name}][%d] DeleteNetworkAclsIDTagsTagName default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteNetworkAclsIDTagsTagNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
