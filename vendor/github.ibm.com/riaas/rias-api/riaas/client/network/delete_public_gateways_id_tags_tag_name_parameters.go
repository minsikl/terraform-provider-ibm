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
)

// NewDeletePublicGatewaysIDTagsTagNameParams creates a new DeletePublicGatewaysIDTagsTagNameParams object
// with the default values initialized.
func NewDeletePublicGatewaysIDTagsTagNameParams() *DeletePublicGatewaysIDTagsTagNameParams {
	var ()
	return &DeletePublicGatewaysIDTagsTagNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePublicGatewaysIDTagsTagNameParamsWithTimeout creates a new DeletePublicGatewaysIDTagsTagNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePublicGatewaysIDTagsTagNameParamsWithTimeout(timeout time.Duration) *DeletePublicGatewaysIDTagsTagNameParams {
	var ()
	return &DeletePublicGatewaysIDTagsTagNameParams{

		timeout: timeout,
	}
}

// NewDeletePublicGatewaysIDTagsTagNameParamsWithContext creates a new DeletePublicGatewaysIDTagsTagNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePublicGatewaysIDTagsTagNameParamsWithContext(ctx context.Context) *DeletePublicGatewaysIDTagsTagNameParams {
	var ()
	return &DeletePublicGatewaysIDTagsTagNameParams{

		Context: ctx,
	}
}

// NewDeletePublicGatewaysIDTagsTagNameParamsWithHTTPClient creates a new DeletePublicGatewaysIDTagsTagNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePublicGatewaysIDTagsTagNameParamsWithHTTPClient(client *http.Client) *DeletePublicGatewaysIDTagsTagNameParams {
	var ()
	return &DeletePublicGatewaysIDTagsTagNameParams{
		HTTPClient: client,
	}
}

/*DeletePublicGatewaysIDTagsTagNameParams contains all the parameters to send to the API endpoint
for the delete public gateways ID tags tag name operation typically these are written to a http.Request
*/
type DeletePublicGatewaysIDTagsTagNameParams struct {

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

// WithTimeout adds the timeout to the delete public gateways ID tags tag name params
func (o *DeletePublicGatewaysIDTagsTagNameParams) WithTimeout(timeout time.Duration) *DeletePublicGatewaysIDTagsTagNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete public gateways ID tags tag name params
func (o *DeletePublicGatewaysIDTagsTagNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete public gateways ID tags tag name params
func (o *DeletePublicGatewaysIDTagsTagNameParams) WithContext(ctx context.Context) *DeletePublicGatewaysIDTagsTagNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete public gateways ID tags tag name params
func (o *DeletePublicGatewaysIDTagsTagNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete public gateways ID tags tag name params
func (o *DeletePublicGatewaysIDTagsTagNameParams) WithHTTPClient(client *http.Client) *DeletePublicGatewaysIDTagsTagNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete public gateways ID tags tag name params
func (o *DeletePublicGatewaysIDTagsTagNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete public gateways ID tags tag name params
func (o *DeletePublicGatewaysIDTagsTagNameParams) WithID(id string) *DeletePublicGatewaysIDTagsTagNameParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete public gateways ID tags tag name params
func (o *DeletePublicGatewaysIDTagsTagNameParams) SetID(id string) {
	o.ID = id
}

// WithTagName adds the tagName to the delete public gateways ID tags tag name params
func (o *DeletePublicGatewaysIDTagsTagNameParams) WithTagName(tagName string) *DeletePublicGatewaysIDTagsTagNameParams {
	o.SetTagName(tagName)
	return o
}

// SetTagName adds the tagName to the delete public gateways ID tags tag name params
func (o *DeletePublicGatewaysIDTagsTagNameParams) SetTagName(tagName string) {
	o.TagName = tagName
}

// WithVersion adds the version to the delete public gateways ID tags tag name params
func (o *DeletePublicGatewaysIDTagsTagNameParams) WithVersion(version string) *DeletePublicGatewaysIDTagsTagNameParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete public gateways ID tags tag name params
func (o *DeletePublicGatewaysIDTagsTagNameParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePublicGatewaysIDTagsTagNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
