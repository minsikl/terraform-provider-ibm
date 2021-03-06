// Code generated by go-swagger; DO NOT EDIT.

package l_baas

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

// NewPostLoadBalancersParams creates a new PostLoadBalancersParams object
// with the default values initialized.
func NewPostLoadBalancersParams() *PostLoadBalancersParams {
	var ()
	return &PostLoadBalancersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLoadBalancersParamsWithTimeout creates a new PostLoadBalancersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLoadBalancersParamsWithTimeout(timeout time.Duration) *PostLoadBalancersParams {
	var ()
	return &PostLoadBalancersParams{

		timeout: timeout,
	}
}

// NewPostLoadBalancersParamsWithContext creates a new PostLoadBalancersParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLoadBalancersParamsWithContext(ctx context.Context) *PostLoadBalancersParams {
	var ()
	return &PostLoadBalancersParams{

		Context: ctx,
	}
}

// NewPostLoadBalancersParamsWithHTTPClient creates a new PostLoadBalancersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLoadBalancersParamsWithHTTPClient(client *http.Client) *PostLoadBalancersParams {
	var ()
	return &PostLoadBalancersParams{
		HTTPClient: client,
	}
}

/*PostLoadBalancersParams contains all the parameters to send to the API endpoint
for the post load balancers operation typically these are written to a http.Request
*/
type PostLoadBalancersParams struct {

	/*Body
	  The load balancer template

	*/
	Body *models.LoadBalancerTemplate
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post load balancers params
func (o *PostLoadBalancersParams) WithTimeout(timeout time.Duration) *PostLoadBalancersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post load balancers params
func (o *PostLoadBalancersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post load balancers params
func (o *PostLoadBalancersParams) WithContext(ctx context.Context) *PostLoadBalancersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post load balancers params
func (o *PostLoadBalancersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post load balancers params
func (o *PostLoadBalancersParams) WithHTTPClient(client *http.Client) *PostLoadBalancersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post load balancers params
func (o *PostLoadBalancersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post load balancers params
func (o *PostLoadBalancersParams) WithBody(body *models.LoadBalancerTemplate) *PostLoadBalancersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post load balancers params
func (o *PostLoadBalancersParams) SetBody(body *models.LoadBalancerTemplate) {
	o.Body = body
}

// WithVersion adds the version to the post load balancers params
func (o *PostLoadBalancersParams) WithVersion(version string) *PostLoadBalancersParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the post load balancers params
func (o *PostLoadBalancersParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PostLoadBalancersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
