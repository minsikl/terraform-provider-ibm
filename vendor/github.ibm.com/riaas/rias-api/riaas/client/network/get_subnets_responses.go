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

// GetSubnetsReader is a Reader for the GetSubnets structure.
type GetSubnetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubnetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSubnetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetSubnetsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewGetSubnetsOK creates a GetSubnetsOK with default headers values
func NewGetSubnetsOK() *GetSubnetsOK {
	return &GetSubnetsOK{}
}

/*GetSubnetsOK handles this case with default header values.

dummy
*/
type GetSubnetsOK struct {
	Payload *models.GetSubnetsOKBody
}

func (o *GetSubnetsOK) Error() string {
	return fmt.Sprintf("[GET /subnets][%d] getSubnetsOK  %+v", 200, o.Payload)
}

func (o *GetSubnetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetSubnetsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubnetsDefault creates a GetSubnetsDefault with default headers values
func NewGetSubnetsDefault(code int) *GetSubnetsDefault {
	return &GetSubnetsDefault{
		_statusCode: code,
	}
}

/*GetSubnetsDefault handles this case with default header values.

unexpectederror
*/
type GetSubnetsDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the get subnets default response
func (o *GetSubnetsDefault) Code() int {
	return o._statusCode
}

func (o *GetSubnetsDefault) Error() string {
	return fmt.Sprintf("[GET /subnets][%d] GetSubnets default  %+v", o._statusCode, o.Payload)
}

func (o *GetSubnetsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
