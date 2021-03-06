// Code generated by go-swagger; DO NOT EDIT.

package v_p_naa_s

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

// NewPostVpnGatewaysVpnGatewayIDConnectionsParams creates a new PostVpnGatewaysVpnGatewayIDConnectionsParams object
// with the default values initialized.
func NewPostVpnGatewaysVpnGatewayIDConnectionsParams() *PostVpnGatewaysVpnGatewayIDConnectionsParams {
	var ()
	return &PostVpnGatewaysVpnGatewayIDConnectionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostVpnGatewaysVpnGatewayIDConnectionsParamsWithTimeout creates a new PostVpnGatewaysVpnGatewayIDConnectionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostVpnGatewaysVpnGatewayIDConnectionsParamsWithTimeout(timeout time.Duration) *PostVpnGatewaysVpnGatewayIDConnectionsParams {
	var ()
	return &PostVpnGatewaysVpnGatewayIDConnectionsParams{

		timeout: timeout,
	}
}

// NewPostVpnGatewaysVpnGatewayIDConnectionsParamsWithContext creates a new PostVpnGatewaysVpnGatewayIDConnectionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostVpnGatewaysVpnGatewayIDConnectionsParamsWithContext(ctx context.Context) *PostVpnGatewaysVpnGatewayIDConnectionsParams {
	var ()
	return &PostVpnGatewaysVpnGatewayIDConnectionsParams{

		Context: ctx,
	}
}

// NewPostVpnGatewaysVpnGatewayIDConnectionsParamsWithHTTPClient creates a new PostVpnGatewaysVpnGatewayIDConnectionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostVpnGatewaysVpnGatewayIDConnectionsParamsWithHTTPClient(client *http.Client) *PostVpnGatewaysVpnGatewayIDConnectionsParams {
	var ()
	return &PostVpnGatewaysVpnGatewayIDConnectionsParams{
		HTTPClient: client,
	}
}

/*PostVpnGatewaysVpnGatewayIDConnectionsParams contains all the parameters to send to the API endpoint
for the post vpn gateways vpn gateway ID connections operation typically these are written to a http.Request
*/
type PostVpnGatewaysVpnGatewayIDConnectionsParams struct {

	/*Body
	  The VPN connection template.

	*/
	Body *models.VPNGatewayConnectionTemplate
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string
	/*VpnGatewayID
	  The VPN gateway identifier

	*/
	VpnGatewayID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post vpn gateways vpn gateway ID connections params
func (o *PostVpnGatewaysVpnGatewayIDConnectionsParams) WithTimeout(timeout time.Duration) *PostVpnGatewaysVpnGatewayIDConnectionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post vpn gateways vpn gateway ID connections params
func (o *PostVpnGatewaysVpnGatewayIDConnectionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post vpn gateways vpn gateway ID connections params
func (o *PostVpnGatewaysVpnGatewayIDConnectionsParams) WithContext(ctx context.Context) *PostVpnGatewaysVpnGatewayIDConnectionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post vpn gateways vpn gateway ID connections params
func (o *PostVpnGatewaysVpnGatewayIDConnectionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post vpn gateways vpn gateway ID connections params
func (o *PostVpnGatewaysVpnGatewayIDConnectionsParams) WithHTTPClient(client *http.Client) *PostVpnGatewaysVpnGatewayIDConnectionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post vpn gateways vpn gateway ID connections params
func (o *PostVpnGatewaysVpnGatewayIDConnectionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post vpn gateways vpn gateway ID connections params
func (o *PostVpnGatewaysVpnGatewayIDConnectionsParams) WithBody(body *models.VPNGatewayConnectionTemplate) *PostVpnGatewaysVpnGatewayIDConnectionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post vpn gateways vpn gateway ID connections params
func (o *PostVpnGatewaysVpnGatewayIDConnectionsParams) SetBody(body *models.VPNGatewayConnectionTemplate) {
	o.Body = body
}

// WithVersion adds the version to the post vpn gateways vpn gateway ID connections params
func (o *PostVpnGatewaysVpnGatewayIDConnectionsParams) WithVersion(version string) *PostVpnGatewaysVpnGatewayIDConnectionsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the post vpn gateways vpn gateway ID connections params
func (o *PostVpnGatewaysVpnGatewayIDConnectionsParams) SetVersion(version string) {
	o.Version = version
}

// WithVpnGatewayID adds the vpnGatewayID to the post vpn gateways vpn gateway ID connections params
func (o *PostVpnGatewaysVpnGatewayIDConnectionsParams) WithVpnGatewayID(vpnGatewayID string) *PostVpnGatewaysVpnGatewayIDConnectionsParams {
	o.SetVpnGatewayID(vpnGatewayID)
	return o
}

// SetVpnGatewayID adds the vpnGatewayId to the post vpn gateways vpn gateway ID connections params
func (o *PostVpnGatewaysVpnGatewayIDConnectionsParams) SetVpnGatewayID(vpnGatewayID string) {
	o.VpnGatewayID = vpnGatewayID
}

// WriteToRequest writes these params to a swagger request
func (o *PostVpnGatewaysVpnGatewayIDConnectionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param vpn_gateway_id
	if err := r.SetPathParam("vpn_gateway_id", o.VpnGatewayID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
