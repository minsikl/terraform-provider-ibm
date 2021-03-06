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

// NewGetNetworkAclsIDTagsTagNameParams creates a new GetNetworkAclsIDTagsTagNameParams object
// with the default values initialized.
func NewGetNetworkAclsIDTagsTagNameParams() *GetNetworkAclsIDTagsTagNameParams {
	var ()
	return &GetNetworkAclsIDTagsTagNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkAclsIDTagsTagNameParamsWithTimeout creates a new GetNetworkAclsIDTagsTagNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkAclsIDTagsTagNameParamsWithTimeout(timeout time.Duration) *GetNetworkAclsIDTagsTagNameParams {
	var ()
	return &GetNetworkAclsIDTagsTagNameParams{

		timeout: timeout,
	}
}

// NewGetNetworkAclsIDTagsTagNameParamsWithContext creates a new GetNetworkAclsIDTagsTagNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkAclsIDTagsTagNameParamsWithContext(ctx context.Context) *GetNetworkAclsIDTagsTagNameParams {
	var ()
	return &GetNetworkAclsIDTagsTagNameParams{

		Context: ctx,
	}
}

// NewGetNetworkAclsIDTagsTagNameParamsWithHTTPClient creates a new GetNetworkAclsIDTagsTagNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkAclsIDTagsTagNameParamsWithHTTPClient(client *http.Client) *GetNetworkAclsIDTagsTagNameParams {
	var ()
	return &GetNetworkAclsIDTagsTagNameParams{
		HTTPClient: client,
	}
}

/*GetNetworkAclsIDTagsTagNameParams contains all the parameters to send to the API endpoint
for the get network acls ID tags tag name operation typically these are written to a http.Request
*/
type GetNetworkAclsIDTagsTagNameParams struct {

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

// WithTimeout adds the timeout to the get network acls ID tags tag name params
func (o *GetNetworkAclsIDTagsTagNameParams) WithTimeout(timeout time.Duration) *GetNetworkAclsIDTagsTagNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network acls ID tags tag name params
func (o *GetNetworkAclsIDTagsTagNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network acls ID tags tag name params
func (o *GetNetworkAclsIDTagsTagNameParams) WithContext(ctx context.Context) *GetNetworkAclsIDTagsTagNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network acls ID tags tag name params
func (o *GetNetworkAclsIDTagsTagNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network acls ID tags tag name params
func (o *GetNetworkAclsIDTagsTagNameParams) WithHTTPClient(client *http.Client) *GetNetworkAclsIDTagsTagNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network acls ID tags tag name params
func (o *GetNetworkAclsIDTagsTagNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get network acls ID tags tag name params
func (o *GetNetworkAclsIDTagsTagNameParams) WithID(id string) *GetNetworkAclsIDTagsTagNameParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get network acls ID tags tag name params
func (o *GetNetworkAclsIDTagsTagNameParams) SetID(id string) {
	o.ID = id
}

// WithTagName adds the tagName to the get network acls ID tags tag name params
func (o *GetNetworkAclsIDTagsTagNameParams) WithTagName(tagName string) *GetNetworkAclsIDTagsTagNameParams {
	o.SetTagName(tagName)
	return o
}

// SetTagName adds the tagName to the get network acls ID tags tag name params
func (o *GetNetworkAclsIDTagsTagNameParams) SetTagName(tagName string) {
	o.TagName = tagName
}

// WithVersion adds the version to the get network acls ID tags tag name params
func (o *GetNetworkAclsIDTagsTagNameParams) WithVersion(version string) *GetNetworkAclsIDTagsTagNameParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get network acls ID tags tag name params
func (o *GetNetworkAclsIDTagsTagNameParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkAclsIDTagsTagNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
