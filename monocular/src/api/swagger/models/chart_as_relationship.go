package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*ChartAsRelationship chart as relationship

swagger:model chartAsRelationship
*/
type ChartAsRelationship struct {

	/* data

	Required: true
	*/
	Data *Chart `json:"data"`

	/* links

	Required: true
	*/
	Links *ResourceLink `json:"links"`
}

// Validate validates this chart as relationship
func (m *ChartAsRelationship) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChartAsRelationship) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {

		if err := m.Data.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ChartAsRelationship) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("links", "body", m.Links); err != nil {
		return err
	}

	if m.Links != nil {

		if err := m.Links.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
