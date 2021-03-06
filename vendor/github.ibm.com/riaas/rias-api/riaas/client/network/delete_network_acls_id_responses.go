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

// DeleteNetworkAclsIDReader is a Reader for the DeleteNetworkAclsID structure.
type DeleteNetworkAclsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkAclsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteNetworkAclsIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteNetworkAclsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteNetworkAclsIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewDeleteNetworkAclsIDNoContent creates a DeleteNetworkAclsIDNoContent with default headers values
func NewDeleteNetworkAclsIDNoContent() *DeleteNetworkAclsIDNoContent {
	return &DeleteNetworkAclsIDNoContent{}
}

/*DeleteNetworkAclsIDNoContent handles this case with default header values.

error
*/
type DeleteNetworkAclsIDNoContent struct {
	Payload *models.Riaaserror
}

func (o *DeleteNetworkAclsIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /network_acls/{id}][%d] deleteNetworkAclsIdNoContent  %+v", 204, o.Payload)
}

func (o *DeleteNetworkAclsIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkAclsIDNotFound creates a DeleteNetworkAclsIDNotFound with default headers values
func NewDeleteNetworkAclsIDNotFound() *DeleteNetworkAclsIDNotFound {
	return &DeleteNetworkAclsIDNotFound{}
}

/*DeleteNetworkAclsIDNotFound handles this case with default header values.

error
*/
type DeleteNetworkAclsIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *DeleteNetworkAclsIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /network_acls/{id}][%d] deleteNetworkAclsIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteNetworkAclsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkAclsIDDefault creates a DeleteNetworkAclsIDDefault with default headers values
func NewDeleteNetworkAclsIDDefault(code int) *DeleteNetworkAclsIDDefault {
	return &DeleteNetworkAclsIDDefault{
		_statusCode: code,
	}
}

/*DeleteNetworkAclsIDDefault handles this case with default header values.

unexpectederror
*/
type DeleteNetworkAclsIDDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the delete network acls ID default response
func (o *DeleteNetworkAclsIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNetworkAclsIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /network_acls/{id}][%d] DeleteNetworkAclsID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteNetworkAclsIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
