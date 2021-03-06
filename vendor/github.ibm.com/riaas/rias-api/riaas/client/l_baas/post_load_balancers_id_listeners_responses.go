// Code generated by go-swagger; DO NOT EDIT.

package l_baas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// PostLoadBalancersIDListenersReader is a Reader for the PostLoadBalancersIDListeners structure.
type PostLoadBalancersIDListenersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLoadBalancersIDListenersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostLoadBalancersIDListenersCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostLoadBalancersIDListenersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLoadBalancersIDListenersCreated creates a PostLoadBalancersIDListenersCreated with default headers values
func NewPostLoadBalancersIDListenersCreated() *PostLoadBalancersIDListenersCreated {
	return &PostLoadBalancersIDListenersCreated{}
}

/*PostLoadBalancersIDListenersCreated handles this case with default header values.

The listener was created successfully.
*/
type PostLoadBalancersIDListenersCreated struct {
	Payload *models.Listener
}

func (o *PostLoadBalancersIDListenersCreated) Error() string {
	return fmt.Sprintf("[POST /load_balancers/{id}/listeners][%d] postLoadBalancersIdListenersCreated  %+v", 201, o.Payload)
}

func (o *PostLoadBalancersIDListenersCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Listener)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLoadBalancersIDListenersBadRequest creates a PostLoadBalancersIDListenersBadRequest with default headers values
func NewPostLoadBalancersIDListenersBadRequest() *PostLoadBalancersIDListenersBadRequest {
	return &PostLoadBalancersIDListenersBadRequest{}
}

/*PostLoadBalancersIDListenersBadRequest handles this case with default header values.

An invalid listener template was provided.
*/
type PostLoadBalancersIDListenersBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PostLoadBalancersIDListenersBadRequest) Error() string {
	return fmt.Sprintf("[POST /load_balancers/{id}/listeners][%d] postLoadBalancersIdListenersBadRequest  %+v", 400, o.Payload)
}

func (o *PostLoadBalancersIDListenersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
