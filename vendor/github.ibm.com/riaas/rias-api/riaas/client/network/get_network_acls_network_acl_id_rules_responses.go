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

// GetNetworkAclsNetworkACLIDRulesReader is a Reader for the GetNetworkAclsNetworkACLIDRules structure.
type GetNetworkAclsNetworkACLIDRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkAclsNetworkACLIDRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNetworkAclsNetworkACLIDRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetNetworkAclsNetworkACLIDRulesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewGetNetworkAclsNetworkACLIDRulesOK creates a GetNetworkAclsNetworkACLIDRulesOK with default headers values
func NewGetNetworkAclsNetworkACLIDRulesOK() *GetNetworkAclsNetworkACLIDRulesOK {
	return &GetNetworkAclsNetworkACLIDRulesOK{}
}

/*GetNetworkAclsNetworkACLIDRulesOK handles this case with default header values.

dummy
*/
type GetNetworkAclsNetworkACLIDRulesOK struct {
	Payload *models.GetNetworkAclsNetworkACLIDRulesOKBody
}

func (o *GetNetworkAclsNetworkACLIDRulesOK) Error() string {
	return fmt.Sprintf("[GET /network_acls/{network_acl_id}/rules][%d] getNetworkAclsNetworkAclIdRulesOK  %+v", 200, o.Payload)
}

func (o *GetNetworkAclsNetworkACLIDRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetNetworkAclsNetworkACLIDRulesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkAclsNetworkACLIDRulesDefault creates a GetNetworkAclsNetworkACLIDRulesDefault with default headers values
func NewGetNetworkAclsNetworkACLIDRulesDefault(code int) *GetNetworkAclsNetworkACLIDRulesDefault {
	return &GetNetworkAclsNetworkACLIDRulesDefault{
		_statusCode: code,
	}
}

/*GetNetworkAclsNetworkACLIDRulesDefault handles this case with default header values.

unexpectederror
*/
type GetNetworkAclsNetworkACLIDRulesDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the get network acls network ACL ID rules default response
func (o *GetNetworkAclsNetworkACLIDRulesDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworkAclsNetworkACLIDRulesDefault) Error() string {
	return fmt.Sprintf("[GET /network_acls/{network_acl_id}/rules][%d] GetNetworkAclsNetworkACLIDRules default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetworkAclsNetworkACLIDRulesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
