// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetInstancesInstanceIDNetworkInterfacesIDTagsParams creates a new GetInstancesInstanceIDNetworkInterfacesIDTagsParams object
// with the default values initialized.
func NewGetInstancesInstanceIDNetworkInterfacesIDTagsParams() *GetInstancesInstanceIDNetworkInterfacesIDTagsParams {
	var ()
	return &GetInstancesInstanceIDNetworkInterfacesIDTagsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstancesInstanceIDNetworkInterfacesIDTagsParamsWithTimeout creates a new GetInstancesInstanceIDNetworkInterfacesIDTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInstancesInstanceIDNetworkInterfacesIDTagsParamsWithTimeout(timeout time.Duration) *GetInstancesInstanceIDNetworkInterfacesIDTagsParams {
	var ()
	return &GetInstancesInstanceIDNetworkInterfacesIDTagsParams{

		timeout: timeout,
	}
}

// NewGetInstancesInstanceIDNetworkInterfacesIDTagsParamsWithContext creates a new GetInstancesInstanceIDNetworkInterfacesIDTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInstancesInstanceIDNetworkInterfacesIDTagsParamsWithContext(ctx context.Context) *GetInstancesInstanceIDNetworkInterfacesIDTagsParams {
	var ()
	return &GetInstancesInstanceIDNetworkInterfacesIDTagsParams{

		Context: ctx,
	}
}

// NewGetInstancesInstanceIDNetworkInterfacesIDTagsParamsWithHTTPClient creates a new GetInstancesInstanceIDNetworkInterfacesIDTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInstancesInstanceIDNetworkInterfacesIDTagsParamsWithHTTPClient(client *http.Client) *GetInstancesInstanceIDNetworkInterfacesIDTagsParams {
	var ()
	return &GetInstancesInstanceIDNetworkInterfacesIDTagsParams{
		HTTPClient: client,
	}
}

/*GetInstancesInstanceIDNetworkInterfacesIDTagsParams contains all the parameters to send to the API endpoint
for the get instances instance ID network interfaces ID tags operation typically these are written to a http.Request
*/
type GetInstancesInstanceIDNetworkInterfacesIDTagsParams struct {

	/*ID
	  The resource identifier

	*/
	ID string
	/*InstanceID
	  The instance identifier

	*/
	InstanceID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get instances instance ID network interfaces ID tags params
func (o *GetInstancesInstanceIDNetworkInterfacesIDTagsParams) WithTimeout(timeout time.Duration) *GetInstancesInstanceIDNetworkInterfacesIDTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get instances instance ID network interfaces ID tags params
func (o *GetInstancesInstanceIDNetworkInterfacesIDTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get instances instance ID network interfaces ID tags params
func (o *GetInstancesInstanceIDNetworkInterfacesIDTagsParams) WithContext(ctx context.Context) *GetInstancesInstanceIDNetworkInterfacesIDTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get instances instance ID network interfaces ID tags params
func (o *GetInstancesInstanceIDNetworkInterfacesIDTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get instances instance ID network interfaces ID tags params
func (o *GetInstancesInstanceIDNetworkInterfacesIDTagsParams) WithHTTPClient(client *http.Client) *GetInstancesInstanceIDNetworkInterfacesIDTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get instances instance ID network interfaces ID tags params
func (o *GetInstancesInstanceIDNetworkInterfacesIDTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get instances instance ID network interfaces ID tags params
func (o *GetInstancesInstanceIDNetworkInterfacesIDTagsParams) WithID(id string) *GetInstancesInstanceIDNetworkInterfacesIDTagsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get instances instance ID network interfaces ID tags params
func (o *GetInstancesInstanceIDNetworkInterfacesIDTagsParams) SetID(id string) {
	o.ID = id
}

// WithInstanceID adds the instanceID to the get instances instance ID network interfaces ID tags params
func (o *GetInstancesInstanceIDNetworkInterfacesIDTagsParams) WithInstanceID(instanceID string) *GetInstancesInstanceIDNetworkInterfacesIDTagsParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the get instances instance ID network interfaces ID tags params
func (o *GetInstancesInstanceIDNetworkInterfacesIDTagsParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WithVersion adds the version to the get instances instance ID network interfaces ID tags params
func (o *GetInstancesInstanceIDNetworkInterfacesIDTagsParams) WithVersion(version string) *GetInstancesInstanceIDNetworkInterfacesIDTagsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get instances instance ID network interfaces ID tags params
func (o *GetInstancesInstanceIDNetworkInterfacesIDTagsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstancesInstanceIDNetworkInterfacesIDTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param instance_id
	if err := r.SetPathParam("instance_id", o.InstanceID); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
