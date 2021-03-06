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

// GetSubnetsSubnetIDPublicGatewayReader is a Reader for the GetSubnetsSubnetIDPublicGateway structure.
type GetSubnetsSubnetIDPublicGatewayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubnetsSubnetIDPublicGatewayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSubnetsSubnetIDPublicGatewayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetSubnetsSubnetIDPublicGatewayNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetSubnetsSubnetIDPublicGatewayDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewGetSubnetsSubnetIDPublicGatewayOK creates a GetSubnetsSubnetIDPublicGatewayOK with default headers values
func NewGetSubnetsSubnetIDPublicGatewayOK() *GetSubnetsSubnetIDPublicGatewayOK {
	return &GetSubnetsSubnetIDPublicGatewayOK{}
}

/*GetSubnetsSubnetIDPublicGatewayOK handles this case with default header values.

dummy
*/
type GetSubnetsSubnetIDPublicGatewayOK struct {
	Payload *models.PublicGateway
}

func (o *GetSubnetsSubnetIDPublicGatewayOK) Error() string {
	return fmt.Sprintf("[GET /subnets/{subnet_id}/public_gateway][%d] getSubnetsSubnetIdPublicGatewayOK  %+v", 200, o.Payload)
}

func (o *GetSubnetsSubnetIDPublicGatewayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicGateway)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubnetsSubnetIDPublicGatewayNotFound creates a GetSubnetsSubnetIDPublicGatewayNotFound with default headers values
func NewGetSubnetsSubnetIDPublicGatewayNotFound() *GetSubnetsSubnetIDPublicGatewayNotFound {
	return &GetSubnetsSubnetIDPublicGatewayNotFound{}
}

/*GetSubnetsSubnetIDPublicGatewayNotFound handles this case with default header values.

error
*/
type GetSubnetsSubnetIDPublicGatewayNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetSubnetsSubnetIDPublicGatewayNotFound) Error() string {
	return fmt.Sprintf("[GET /subnets/{subnet_id}/public_gateway][%d] getSubnetsSubnetIdPublicGatewayNotFound  %+v", 404, o.Payload)
}

func (o *GetSubnetsSubnetIDPublicGatewayNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubnetsSubnetIDPublicGatewayDefault creates a GetSubnetsSubnetIDPublicGatewayDefault with default headers values
func NewGetSubnetsSubnetIDPublicGatewayDefault(code int) *GetSubnetsSubnetIDPublicGatewayDefault {
	return &GetSubnetsSubnetIDPublicGatewayDefault{
		_statusCode: code,
	}
}

/*GetSubnetsSubnetIDPublicGatewayDefault handles this case with default header values.

unexpectederror
*/
type GetSubnetsSubnetIDPublicGatewayDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the get subnets subnet ID public gateway default response
func (o *GetSubnetsSubnetIDPublicGatewayDefault) Code() int {
	return o._statusCode
}

func (o *GetSubnetsSubnetIDPublicGatewayDefault) Error() string {
	return fmt.Sprintf("[GET /subnets/{subnet_id}/public_gateway][%d] GetSubnetsSubnetIDPublicGateway default  %+v", o._statusCode, o.Payload)
}

func (o *GetSubnetsSubnetIDPublicGatewayDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
