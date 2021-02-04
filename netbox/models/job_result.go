// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// JobResult job result
//
// swagger:model JobResult
type JobResult struct {

	// Completed
	// Format: date-time
	Completed *strfmt.DateTime `json:"completed,omitempty"`

	// Created
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// Data
	Data *string `json:"data,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Job id
	// Required: true
	// Format: uuid
	JobID *strfmt.UUID `json:"job_id"`

	// Name
	// Required: true
	// Max Length: 255
	// Min Length: 1
	Name *string `json:"name"`

	// Obj type
	// Read Only: true
	ObjType string `json:"obj_type,omitempty"`

	// status
	Status *JobResultStatus `json:"status,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// user
	User *NestedUser `json:"user,omitempty"`
}

// Validate validates this job result
func (m *JobResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobResult) validateCompleted(formats strfmt.Registry) error {
	if swag.IsZero(m.Completed) { // not required
		return nil
	}

	if err := validate.FormatOf("completed", "body", "date-time", m.Completed.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JobResult) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JobResult) validateJobID(formats strfmt.Registry) error {

	if err := validate.Required("job_id", "body", m.JobID); err != nil {
		return err
	}

	if err := validate.FormatOf("job_id", "body", "uuid", m.JobID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JobResult) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 255); err != nil {
		return err
	}

	return nil
}

func (m *JobResult) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *JobResult) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JobResult) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this job result based on the context it is used
func (m *JobResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateObjType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobResult) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *JobResult) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *JobResult) contextValidateObjType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "obj_type", "body", string(m.ObjType)); err != nil {
		return err
	}

	return nil
}

func (m *JobResult) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *JobResult) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

func (m *JobResult) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {
		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobResult) UnmarshalBinary(b []byte) error {
	var res JobResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// JobResultStatus Status
//
// swagger:model JobResultStatus
type JobResultStatus struct {

	// label
	// Required: true
	// Enum: [Pending Running Completed Errored Failed]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [pending running completed errored failed]
	Value *string `json:"value"`
}

// Validate validates this job result status
func (m *JobResultStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var jobResultStatusTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Pending","Running","Completed","Errored","Failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		jobResultStatusTypeLabelPropEnum = append(jobResultStatusTypeLabelPropEnum, v)
	}
}

const (

	// JobResultStatusLabelPending captures enum value "Pending"
	JobResultStatusLabelPending string = "Pending"

	// JobResultStatusLabelRunning captures enum value "Running"
	JobResultStatusLabelRunning string = "Running"

	// JobResultStatusLabelCompleted captures enum value "Completed"
	JobResultStatusLabelCompleted string = "Completed"

	// JobResultStatusLabelErrored captures enum value "Errored"
	JobResultStatusLabelErrored string = "Errored"

	// JobResultStatusLabelFailed captures enum value "Failed"
	JobResultStatusLabelFailed string = "Failed"
)

// prop value enum
func (m *JobResultStatus) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, jobResultStatusTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *JobResultStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("status"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var jobResultStatusTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pending","running","completed","errored","failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		jobResultStatusTypeValuePropEnum = append(jobResultStatusTypeValuePropEnum, v)
	}
}

const (

	// JobResultStatusValuePending captures enum value "pending"
	JobResultStatusValuePending string = "pending"

	// JobResultStatusValueRunning captures enum value "running"
	JobResultStatusValueRunning string = "running"

	// JobResultStatusValueCompleted captures enum value "completed"
	JobResultStatusValueCompleted string = "completed"

	// JobResultStatusValueErrored captures enum value "errored"
	JobResultStatusValueErrored string = "errored"

	// JobResultStatusValueFailed captures enum value "failed"
	JobResultStatusValueFailed string = "failed"
)

// prop value enum
func (m *JobResultStatus) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, jobResultStatusTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *JobResultStatus) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("status"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this job result status based on the context it is used
func (m *JobResultStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *JobResultStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobResultStatus) UnmarshalBinary(b []byte) error {
	var res JobResultStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
