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

// NewGetInstancesIDTagsTagNameParams creates a new GetInstancesIDTagsTagNameParams object
// with the default values initialized.
func NewGetInstancesIDTagsTagNameParams() *GetInstancesIDTagsTagNameParams {
	var ()
	return &GetInstancesIDTagsTagNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstancesIDTagsTagNameParamsWithTimeout creates a new GetInstancesIDTagsTagNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInstancesIDTagsTagNameParamsWithTimeout(timeout time.Duration) *GetInstancesIDTagsTagNameParams {
	var ()
	return &GetInstancesIDTagsTagNameParams{

		timeout: timeout,
	}
}

// NewGetInstancesIDTagsTagNameParamsWithContext creates a new GetInstancesIDTagsTagNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInstancesIDTagsTagNameParamsWithContext(ctx context.Context) *GetInstancesIDTagsTagNameParams {
	var ()
	return &GetInstancesIDTagsTagNameParams{

		Context: ctx,
	}
}

// NewGetInstancesIDTagsTagNameParamsWithHTTPClient creates a new GetInstancesIDTagsTagNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInstancesIDTagsTagNameParamsWithHTTPClient(client *http.Client) *GetInstancesIDTagsTagNameParams {
	var ()
	return &GetInstancesIDTagsTagNameParams{
		HTTPClient: client,
	}
}

/*GetInstancesIDTagsTagNameParams contains all the parameters to send to the API endpoint
for the get instances ID tags tag name operation typically these are written to a http.Request
*/
type GetInstancesIDTagsTagNameParams struct {

	/*ID
	  The resource identifier

	*/
	ID string
	/*TagName
	  The name of the tag

	*/
	TagName string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get instances ID tags tag name params
func (o *GetInstancesIDTagsTagNameParams) WithTimeout(timeout time.Duration) *GetInstancesIDTagsTagNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get instances ID tags tag name params
func (o *GetInstancesIDTagsTagNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get instances ID tags tag name params
func (o *GetInstancesIDTagsTagNameParams) WithContext(ctx context.Context) *GetInstancesIDTagsTagNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get instances ID tags tag name params
func (o *GetInstancesIDTagsTagNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get instances ID tags tag name params
func (o *GetInstancesIDTagsTagNameParams) WithHTTPClient(client *http.Client) *GetInstancesIDTagsTagNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get instances ID tags tag name params
func (o *GetInstancesIDTagsTagNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get instances ID tags tag name params
func (o *GetInstancesIDTagsTagNameParams) WithID(id string) *GetInstancesIDTagsTagNameParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get instances ID tags tag name params
func (o *GetInstancesIDTagsTagNameParams) SetID(id string) {
	o.ID = id
}

// WithTagName adds the tagName to the get instances ID tags tag name params
func (o *GetInstancesIDTagsTagNameParams) WithTagName(tagName string) *GetInstancesIDTagsTagNameParams {
	o.SetTagName(tagName)
	return o
}

// SetTagName adds the tagName to the get instances ID tags tag name params
func (o *GetInstancesIDTagsTagNameParams) SetTagName(tagName string) {
	o.TagName = tagName
}

// WithVersion adds the version to the get instances ID tags tag name params
func (o *GetInstancesIDTagsTagNameParams) WithVersion(version string) *GetInstancesIDTagsTagNameParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get instances ID tags tag name params
func (o *GetInstancesIDTagsTagNameParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstancesIDTagsTagNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param tag_name
	if err := r.SetPathParam("tag_name", o.TagName); err != nil {
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
