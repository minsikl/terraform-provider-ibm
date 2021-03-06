// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GetPublicGatewaysIDTagsOKBody get public gateways Id tags o k body
// swagger:model getPublicGatewaysIdTagsOKBody
type GetPublicGatewaysIDTagsOKBody struct {

	// A collection of tags for this resource
	Tags []string `json:"tags,omitempty"`
}

// Validate validates this get public gateways Id tags o k body
func (m *GetPublicGatewaysIDTagsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetPublicGatewaysIDTagsOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetPublicGatewaysIDTagsOKBody) UnmarshalBinary(b []byte) error {
	var res GetPublicGatewaysIDTagsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
