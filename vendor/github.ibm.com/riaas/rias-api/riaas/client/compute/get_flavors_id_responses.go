// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// GetFlavorsIDReader is a Reader for the GetFlavorsID structure.
type GetFlavorsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlavorsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFlavorsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetFlavorsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetFlavorsIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewGetFlavorsIDOK creates a GetFlavorsIDOK with default headers values
func NewGetFlavorsIDOK() *GetFlavorsIDOK {
	return &GetFlavorsIDOK{}
}

/*GetFlavorsIDOK handles this case with default header values.

dummy
*/
type GetFlavorsIDOK struct {
	Payload *models.Flavor
}

func (o *GetFlavorsIDOK) Error() string {
	return fmt.Sprintf("[GET /flavors/{id}][%d] getFlavorsIdOK  %+v", 200, o.Payload)
}

func (o *GetFlavorsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Flavor)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlavorsIDNotFound creates a GetFlavorsIDNotFound with default headers values
func NewGetFlavorsIDNotFound() *GetFlavorsIDNotFound {
	return &GetFlavorsIDNotFound{}
}

/*GetFlavorsIDNotFound handles this case with default header values.

error
*/
type GetFlavorsIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetFlavorsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /flavors/{id}][%d] getFlavorsIdNotFound  %+v", 404, o.Payload)
}

func (o *GetFlavorsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlavorsIDDefault creates a GetFlavorsIDDefault with default headers values
func NewGetFlavorsIDDefault(code int) *GetFlavorsIDDefault {
	return &GetFlavorsIDDefault{
		_statusCode: code,
	}
}

/*GetFlavorsIDDefault handles this case with default header values.

unexpectederror
*/
type GetFlavorsIDDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the get flavors ID default response
func (o *GetFlavorsIDDefault) Code() int {
	return o._statusCode
}

func (o *GetFlavorsIDDefault) Error() string {
	return fmt.Sprintf("[GET /flavors/{id}][%d] GetFlavorsID default  %+v", o._statusCode, o.Payload)
}

func (o *GetFlavorsIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
