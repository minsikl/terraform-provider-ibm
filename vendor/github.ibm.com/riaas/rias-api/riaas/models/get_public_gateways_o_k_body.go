// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetPublicGatewaysOKBody PublicGatewayCollection
// swagger:model getPublicGatewaysOKBody
type GetPublicGatewaysOKBody struct {

	// first
	First *GetPublicGatewaysOKBodyFirst `json:"first,omitempty"`

	// The maximum number of resources can be returned by the request
	// Maximum: 100
	// Minimum: 1
	Limit int64 `json:"limit,omitempty"`

	// next
	Next *Next `json:"next,omitempty"`

	// Collection of public gateways
	PublicGateways []*PublicGateway `json:"public_gateways,omitempty"`

	// The total number of resources across all pages
	// Minimum: 0
	TotalCount *int64 `json:"total_count,omitempty"`
}

// Validate validates this get public gateways o k body
func (m *GetPublicGatewaysOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFirst(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicGateways(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetPublicGatewaysOKBody) validateFirst(formats strfmt.Registry) error {

	if swag.IsZero(m.First) { // not required
		return nil
	}

	if m.First != nil {
		if err := m.First.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("first")
			}
			return err
		}
	}

	return nil
}

func (m *GetPublicGatewaysOKBody) validateLimit(formats strfmt.Registry) error {

	if swag.IsZero(m.Limit) { // not required
		return nil
	}

	if err := validate.MinimumInt("limit", "body", int64(m.Limit), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("limit", "body", int64(m.Limit), 100, false); err != nil {
		return err
	}

	return nil
}

func (m *GetPublicGatewaysOKBody) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next")
			}
			return err
		}
	}

	return nil
}

func (m *GetPublicGatewaysOKBody) validatePublicGateways(formats strfmt.Registry) error {

	if swag.IsZero(m.PublicGateways) { // not required
		return nil
	}

	for i := 0; i < len(m.PublicGateways); i++ {
		if swag.IsZero(m.PublicGateways[i]) { // not required
			continue
		}

		if m.PublicGateways[i] != nil {
			if err := m.PublicGateways[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("public_gateways" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetPublicGatewaysOKBody) validateTotalCount(formats strfmt.Registry) error {

	if swag.IsZero(m.TotalCount) { // not required
		return nil
	}

	if err := validate.MinimumInt("total_count", "body", int64(*m.TotalCount), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetPublicGatewaysOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetPublicGatewaysOKBody) UnmarshalBinary(b []byte) error {
	var res GetPublicGatewaysOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
