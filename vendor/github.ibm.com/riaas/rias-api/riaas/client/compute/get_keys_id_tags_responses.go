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

// GetKeysIDTagsReader is a Reader for the GetKeysIDTags structure.
type GetKeysIDTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKeysIDTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetKeysIDTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetKeysIDTagsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetKeysIDTagsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewGetKeysIDTagsOK creates a GetKeysIDTagsOK with default headers values
func NewGetKeysIDTagsOK() *GetKeysIDTagsOK {
	return &GetKeysIDTagsOK{}
}

/*GetKeysIDTagsOK handles this case with default header values.

dummy
*/
type GetKeysIDTagsOK struct {
	Payload *models.GetKeysIDTagsOKBody
}

func (o *GetKeysIDTagsOK) Error() string {
	return fmt.Sprintf("[GET /keys/{id}/tags][%d] getKeysIdTagsOK  %+v", 200, o.Payload)
}

func (o *GetKeysIDTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetKeysIDTagsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKeysIDTagsNotFound creates a GetKeysIDTagsNotFound with default headers values
func NewGetKeysIDTagsNotFound() *GetKeysIDTagsNotFound {
	return &GetKeysIDTagsNotFound{}
}

/*GetKeysIDTagsNotFound handles this case with default header values.

error
*/
type GetKeysIDTagsNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetKeysIDTagsNotFound) Error() string {
	return fmt.Sprintf("[GET /keys/{id}/tags][%d] getKeysIdTagsNotFound  %+v", 404, o.Payload)
}

func (o *GetKeysIDTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKeysIDTagsDefault creates a GetKeysIDTagsDefault with default headers values
func NewGetKeysIDTagsDefault(code int) *GetKeysIDTagsDefault {
	return &GetKeysIDTagsDefault{
		_statusCode: code,
	}
}

/*GetKeysIDTagsDefault handles this case with default header values.

unexpectederror
*/
type GetKeysIDTagsDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the get keys ID tags default response
func (o *GetKeysIDTagsDefault) Code() int {
	return o._statusCode
}

func (o *GetKeysIDTagsDefault) Error() string {
	return fmt.Sprintf("[GET /keys/{id}/tags][%d] GetKeysIDTags default  %+v", o._statusCode, o.Payload)
}

func (o *GetKeysIDTagsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
