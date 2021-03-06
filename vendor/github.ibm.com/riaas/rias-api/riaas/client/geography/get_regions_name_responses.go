// Code generated by go-swagger; DO NOT EDIT.

package geography

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// GetRegionsNameReader is a Reader for the GetRegionsName structure.
type GetRegionsNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRegionsNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRegionsNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetRegionsNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetRegionsNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewGetRegionsNameOK creates a GetRegionsNameOK with default headers values
func NewGetRegionsNameOK() *GetRegionsNameOK {
	return &GetRegionsNameOK{}
}

/*GetRegionsNameOK handles this case with default header values.

dummy
*/
type GetRegionsNameOK struct {
	Payload *models.Region
}

func (o *GetRegionsNameOK) Error() string {
	return fmt.Sprintf("[GET /regions/{name}][%d] getRegionsNameOK  %+v", 200, o.Payload)
}

func (o *GetRegionsNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Region)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegionsNameNotFound creates a GetRegionsNameNotFound with default headers values
func NewGetRegionsNameNotFound() *GetRegionsNameNotFound {
	return &GetRegionsNameNotFound{}
}

/*GetRegionsNameNotFound handles this case with default header values.

error
*/
type GetRegionsNameNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetRegionsNameNotFound) Error() string {
	return fmt.Sprintf("[GET /regions/{name}][%d] getRegionsNameNotFound  %+v", 404, o.Payload)
}

func (o *GetRegionsNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegionsNameDefault creates a GetRegionsNameDefault with default headers values
func NewGetRegionsNameDefault(code int) *GetRegionsNameDefault {
	return &GetRegionsNameDefault{
		_statusCode: code,
	}
}

/*GetRegionsNameDefault handles this case with default header values.

unexpectederror
*/
type GetRegionsNameDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the get regions name default response
func (o *GetRegionsNameDefault) Code() int {
	return o._statusCode
}

func (o *GetRegionsNameDefault) Error() string {
	return fmt.Sprintf("[GET /regions/{name}][%d] GetRegionsName default  %+v", o._statusCode, o.Payload)
}

func (o *GetRegionsNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
