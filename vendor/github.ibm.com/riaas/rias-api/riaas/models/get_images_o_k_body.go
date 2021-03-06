// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetImagesOKBody ImageCollection
// swagger:model getImagesOKBody
type GetImagesOKBody struct {
	Pagination

	// Collection of images
	Images []*Image `json:"images"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GetImagesOKBody) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Pagination
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Pagination = aO0

	// now for regular properties
	var propsGetImagesOKBody struct {
		Images []*Image `json:"images"`
	}
	if err := swag.ReadJSON(raw, &propsGetImagesOKBody); err != nil {
		return err
	}
	m.Images = propsGetImagesOKBody.Images

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GetImagesOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.Pagination)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsGetImagesOKBody struct {
		Images []*Image `json:"images"`
	}
	propsGetImagesOKBody.Images = m.Images

	jsonDataPropsGetImagesOKBody, errGetImagesOKBody := swag.WriteJSON(propsGetImagesOKBody)
	if errGetImagesOKBody != nil {
		return nil, errGetImagesOKBody
	}
	_parts = append(_parts, jsonDataPropsGetImagesOKBody)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get images o k body
func (m *GetImagesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Pagination
	if err := m.Pagination.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetImagesOKBody) validateImages(formats strfmt.Registry) error {

	if swag.IsZero(m.Images) { // not required
		return nil
	}

	for i := 0; i < len(m.Images); i++ {
		if swag.IsZero(m.Images[i]) { // not required
			continue
		}

		if m.Images[i] != nil {
			if err := m.Images[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetImagesOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetImagesOKBody) UnmarshalBinary(b []byte) error {
	var res GetImagesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
