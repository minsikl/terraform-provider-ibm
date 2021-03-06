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

// GetNetworkAclsIDTagsReader is a Reader for the GetNetworkAclsIDTags structure.
type GetNetworkAclsIDTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkAclsIDTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNetworkAclsIDTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetNetworkAclsIDTagsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetNetworkAclsIDTagsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewGetNetworkAclsIDTagsOK creates a GetNetworkAclsIDTagsOK with default headers values
func NewGetNetworkAclsIDTagsOK() *GetNetworkAclsIDTagsOK {
	return &GetNetworkAclsIDTagsOK{}
}

/*GetNetworkAclsIDTagsOK handles this case with default header values.

dummy
*/
type GetNetworkAclsIDTagsOK struct {
	Payload *models.GetNetworkAclsIDTagsOKBody
}

func (o *GetNetworkAclsIDTagsOK) Error() string {
	return fmt.Sprintf("[GET /network_acls/{id}/tags][%d] getNetworkAclsIdTagsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkAclsIDTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetNetworkAclsIDTagsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkAclsIDTagsNotFound creates a GetNetworkAclsIDTagsNotFound with default headers values
func NewGetNetworkAclsIDTagsNotFound() *GetNetworkAclsIDTagsNotFound {
	return &GetNetworkAclsIDTagsNotFound{}
}

/*GetNetworkAclsIDTagsNotFound handles this case with default header values.

error
*/
type GetNetworkAclsIDTagsNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetNetworkAclsIDTagsNotFound) Error() string {
	return fmt.Sprintf("[GET /network_acls/{id}/tags][%d] getNetworkAclsIdTagsNotFound  %+v", 404, o.Payload)
}

func (o *GetNetworkAclsIDTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkAclsIDTagsDefault creates a GetNetworkAclsIDTagsDefault with default headers values
func NewGetNetworkAclsIDTagsDefault(code int) *GetNetworkAclsIDTagsDefault {
	return &GetNetworkAclsIDTagsDefault{
		_statusCode: code,
	}
}

/*GetNetworkAclsIDTagsDefault handles this case with default header values.

unexpectederror
*/
type GetNetworkAclsIDTagsDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the get network acls ID tags default response
func (o *GetNetworkAclsIDTagsDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworkAclsIDTagsDefault) Error() string {
	return fmt.Sprintf("[GET /network_acls/{id}/tags][%d] GetNetworkAclsIDTags default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetworkAclsIDTagsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
