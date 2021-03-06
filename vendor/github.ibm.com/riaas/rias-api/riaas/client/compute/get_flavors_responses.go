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

// GetFlavorsReader is a Reader for the GetFlavors structure.
type GetFlavorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlavorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFlavorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetFlavorsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewGetFlavorsOK creates a GetFlavorsOK with default headers values
func NewGetFlavorsOK() *GetFlavorsOK {
	return &GetFlavorsOK{}
}

/*GetFlavorsOK handles this case with default header values.

dummy
*/
type GetFlavorsOK struct {
	Payload *models.GetFlavorsOKBody
}

func (o *GetFlavorsOK) Error() string {
	return fmt.Sprintf("[GET /flavors][%d] getFlavorsOK  %+v", 200, o.Payload)
}

func (o *GetFlavorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetFlavorsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlavorsDefault creates a GetFlavorsDefault with default headers values
func NewGetFlavorsDefault(code int) *GetFlavorsDefault {
	return &GetFlavorsDefault{
		_statusCode: code,
	}
}

/*GetFlavorsDefault handles this case with default header values.

unexpectederror
*/
type GetFlavorsDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the get flavors default response
func (o *GetFlavorsDefault) Code() int {
	return o._statusCode
}

func (o *GetFlavorsDefault) Error() string {
	return fmt.Sprintf("[GET /flavors][%d] GetFlavors default  %+v", o._statusCode, o.Payload)
}

func (o *GetFlavorsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
