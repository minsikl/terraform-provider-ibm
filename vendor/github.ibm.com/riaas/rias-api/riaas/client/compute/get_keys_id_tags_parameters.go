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

// NewGetKeysIDTagsParams creates a new GetKeysIDTagsParams object
// with the default values initialized.
func NewGetKeysIDTagsParams() *GetKeysIDTagsParams {
	var ()
	return &GetKeysIDTagsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKeysIDTagsParamsWithTimeout creates a new GetKeysIDTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKeysIDTagsParamsWithTimeout(timeout time.Duration) *GetKeysIDTagsParams {
	var ()
	return &GetKeysIDTagsParams{

		timeout: timeout,
	}
}

// NewGetKeysIDTagsParamsWithContext creates a new GetKeysIDTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKeysIDTagsParamsWithContext(ctx context.Context) *GetKeysIDTagsParams {
	var ()
	return &GetKeysIDTagsParams{

		Context: ctx,
	}
}

// NewGetKeysIDTagsParamsWithHTTPClient creates a new GetKeysIDTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKeysIDTagsParamsWithHTTPClient(client *http.Client) *GetKeysIDTagsParams {
	var ()
	return &GetKeysIDTagsParams{
		HTTPClient: client,
	}
}

/*GetKeysIDTagsParams contains all the parameters to send to the API endpoint
for the get keys ID tags operation typically these are written to a http.Request
*/
type GetKeysIDTagsParams struct {

	/*ID
	  The key ID

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

// WithTimeout adds the timeout to the get keys ID tags params
func (o *GetKeysIDTagsParams) WithTimeout(timeout time.Duration) *GetKeysIDTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get keys ID tags params
func (o *GetKeysIDTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get keys ID tags params
func (o *GetKeysIDTagsParams) WithContext(ctx context.Context) *GetKeysIDTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get keys ID tags params
func (o *GetKeysIDTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get keys ID tags params
func (o *GetKeysIDTagsParams) WithHTTPClient(client *http.Client) *GetKeysIDTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get keys ID tags params
func (o *GetKeysIDTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get keys ID tags params
func (o *GetKeysIDTagsParams) WithID(id string) *GetKeysIDTagsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get keys ID tags params
func (o *GetKeysIDTagsParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the get keys ID tags params
func (o *GetKeysIDTagsParams) WithVersion(version string) *GetKeysIDTagsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get keys ID tags params
func (o *GetKeysIDTagsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetKeysIDTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
