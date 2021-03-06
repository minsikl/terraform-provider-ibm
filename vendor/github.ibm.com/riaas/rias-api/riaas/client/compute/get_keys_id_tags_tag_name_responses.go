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

// GetKeysIDTagsTagNameReader is a Reader for the GetKeysIDTagsTagName structure.
type GetKeysIDTagsTagNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKeysIDTagsTagNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewGetKeysIDTagsTagNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetKeysIDTagsTagNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetKeysIDTagsTagNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewGetKeysIDTagsTagNameNoContent creates a GetKeysIDTagsTagNameNoContent with default headers values
func NewGetKeysIDTagsTagNameNoContent() *GetKeysIDTagsTagNameNoContent {
	return &GetKeysIDTagsTagNameNoContent{}
}

/*GetKeysIDTagsTagNameNoContent handles this case with default header values.

error
*/
type GetKeysIDTagsTagNameNoContent struct {
	Payload *models.Riaaserror
}

func (o *GetKeysIDTagsTagNameNoContent) Error() string {
	return fmt.Sprintf("[GET /keys/{id}/tags/{tag_name}][%d] getKeysIdTagsTagNameNoContent  %+v", 204, o.Payload)
}

func (o *GetKeysIDTagsTagNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKeysIDTagsTagNameNotFound creates a GetKeysIDTagsTagNameNotFound with default headers values
func NewGetKeysIDTagsTagNameNotFound() *GetKeysIDTagsTagNameNotFound {
	return &GetKeysIDTagsTagNameNotFound{}
}

/*GetKeysIDTagsTagNameNotFound handles this case with default header values.

error
*/
type GetKeysIDTagsTagNameNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetKeysIDTagsTagNameNotFound) Error() string {
	return fmt.Sprintf("[GET /keys/{id}/tags/{tag_name}][%d] getKeysIdTagsTagNameNotFound  %+v", 404, o.Payload)
}

func (o *GetKeysIDTagsTagNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKeysIDTagsTagNameDefault creates a GetKeysIDTagsTagNameDefault with default headers values
func NewGetKeysIDTagsTagNameDefault(code int) *GetKeysIDTagsTagNameDefault {
	return &GetKeysIDTagsTagNameDefault{
		_statusCode: code,
	}
}

/*GetKeysIDTagsTagNameDefault handles this case with default header values.

unexpectederror
*/
type GetKeysIDTagsTagNameDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the get keys ID tags tag name default response
func (o *GetKeysIDTagsTagNameDefault) Code() int {
	return o._statusCode
}

func (o *GetKeysIDTagsTagNameDefault) Error() string {
	return fmt.Sprintf("[GET /keys/{id}/tags/{tag_name}][%d] GetKeysIDTagsTagName default  %+v", o._statusCode, o.Payload)
}

func (o *GetKeysIDTagsTagNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
