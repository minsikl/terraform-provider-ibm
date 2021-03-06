// Code generated by go-swagger; DO NOT EDIT.

package network

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

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// NewPatchPublicGatewaysIDParams creates a new PatchPublicGatewaysIDParams object
// with the default values initialized.
func NewPatchPublicGatewaysIDParams() *PatchPublicGatewaysIDParams {
	var ()
	return &PatchPublicGatewaysIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchPublicGatewaysIDParamsWithTimeout creates a new PatchPublicGatewaysIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchPublicGatewaysIDParamsWithTimeout(timeout time.Duration) *PatchPublicGatewaysIDParams {
	var ()
	return &PatchPublicGatewaysIDParams{

		timeout: timeout,
	}
}

// NewPatchPublicGatewaysIDParamsWithContext creates a new PatchPublicGatewaysIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchPublicGatewaysIDParamsWithContext(ctx context.Context) *PatchPublicGatewaysIDParams {
	var ()
	return &PatchPublicGatewaysIDParams{

		Context: ctx,
	}
}

// NewPatchPublicGatewaysIDParamsWithHTTPClient creates a new PatchPublicGatewaysIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchPublicGatewaysIDParamsWithHTTPClient(client *http.Client) *PatchPublicGatewaysIDParams {
	var ()
	return &PatchPublicGatewaysIDParams{
		HTTPClient: client,
	}
}

/*PatchPublicGatewaysIDParams contains all the parameters to send to the API endpoint
for the patch public gateways ID operation typically these are written to a http.Request
*/
type PatchPublicGatewaysIDParams struct {

	/*Body*/
	Body *models.PatchPublicGatewaysIDParamsBody
	/*ID
	  The public gateway identifier

	*/
	ID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch public gateways ID params
func (o *PatchPublicGatewaysIDParams) WithTimeout(timeout time.Duration) *PatchPublicGatewaysIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch public gateways ID params
func (o *PatchPublicGatewaysIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch public gateways ID params
func (o *PatchPublicGatewaysIDParams) WithContext(ctx context.Context) *PatchPublicGatewaysIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch public gateways ID params
func (o *PatchPublicGatewaysIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch public gateways ID params
func (o *PatchPublicGatewaysIDParams) WithHTTPClient(client *http.Client) *PatchPublicGatewaysIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch public gateways ID params
func (o *PatchPublicGatewaysIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch public gateways ID params
func (o *PatchPublicGatewaysIDParams) WithBody(body *models.PatchPublicGatewaysIDParamsBody) *PatchPublicGatewaysIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch public gateways ID params
func (o *PatchPublicGatewaysIDParams) SetBody(body *models.PatchPublicGatewaysIDParamsBody) {
	o.Body = body
}

// WithID adds the id to the patch public gateways ID params
func (o *PatchPublicGatewaysIDParams) WithID(id string) *PatchPublicGatewaysIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch public gateways ID params
func (o *PatchPublicGatewaysIDParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the patch public gateways ID params
func (o *PatchPublicGatewaysIDParams) WithVersion(version string) *PatchPublicGatewaysIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the patch public gateways ID params
func (o *PatchPublicGatewaysIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PatchPublicGatewaysIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
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
