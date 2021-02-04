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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FrontPort front port
//
// swagger:model FrontPort
type FrontPort struct {

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Cable peer
	//
	//
	// Return the appropriate serializer for the cable termination model.
	//
	// Read Only: true
	CablePeer map[string]*string `json:"cable_peer,omitempty"`

	// Cable peer type
	// Read Only: true
	CablePeerType string `json:"cable_peer_type,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// device
	// Required: true
	Device *NestedDevice `json:"device"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Label
	//
	// Physical label
	// Max Length: 64
	Label string `json:"label,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// rear port
	// Required: true
	RearPort *FrontPortRearPort `json:"rear_port"`

	// Rear port position
	// Maximum: 1024
	// Minimum: 1
	RearPortPosition int64 `json:"rear_port_position,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// type
	// Required: true
	Type *FrontPortType `json:"type"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this front port
func (m *FrontPort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRearPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRearPortPosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FrontPort) validateCable(formats strfmt.Registry) error {
	if swag.IsZero(m.Cable) { // not required
		return nil
	}

	if m.Cable != nil {
		if err := m.Cable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *FrontPort) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *FrontPort) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *FrontPort) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *FrontPort) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 64); err != nil {
		return err
	}

	return nil
}

func (m *FrontPort) validateRearPort(formats strfmt.Registry) error {

	if err := validate.Required("rear_port", "body", m.RearPort); err != nil {
		return err
	}

	if m.RearPort != nil {
		if err := m.RearPort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rear_port")
			}
			return err
		}
	}

	return nil
}

func (m *FrontPort) validateRearPortPosition(formats strfmt.Registry) error {
	if swag.IsZero(m.RearPortPosition) { // not required
		return nil
	}

	if err := validate.MinimumInt("rear_port_position", "body", m.RearPortPosition, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("rear_port_position", "body", m.RearPortPosition, 1024, false); err != nil {
		return err
	}

	return nil
}

func (m *FrontPort) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FrontPort) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *FrontPort) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this front port based on the context it is used
func (m *FrontPort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCablePeer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCablePeerType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRearPort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FrontPort) contextValidateCable(ctx context.Context, formats strfmt.Registry) error {

	if m.Cable != nil {
		if err := m.Cable.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *FrontPort) contextValidateCablePeer(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *FrontPort) contextValidateCablePeerType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cable_peer_type", "body", string(m.CablePeerType)); err != nil {
		return err
	}

	return nil
}

func (m *FrontPort) contextValidateDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.Device != nil {
		if err := m.Device.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *FrontPort) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *FrontPort) contextValidateRearPort(ctx context.Context, formats strfmt.Registry) error {

	if m.RearPort != nil {
		if err := m.RearPort.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rear_port")
			}
			return err
		}
	}

	return nil
}

func (m *FrontPort) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FrontPort) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *FrontPort) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FrontPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FrontPort) UnmarshalBinary(b []byte) error {
	var res FrontPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FrontPortType Type
//
// swagger:model FrontPortType
type FrontPortType struct {

	// label
	// Required: true
	// Enum: [8P8C 8P6C 8P4C 8P2C GG45 TERA 4P TERA 2P TERA 1P 110 Punch BNC MRJ21 FC LC LC/APC LSH LSH/APC MPO MTRJ SC SC/APC ST CS SN Splice]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [8p8c 8p6c 8p4c 8p2c gg45 tera-4p tera-2p tera-1p 110-punch bnc mrj21 fc lc lc-apc lsh lsh-apc mpo mtrj sc sc-apc st cs sn splice]
	Value *string `json:"value"`
}

// Validate validates this front port type
func (m *FrontPortType) Validate(formats strfmt.Registry) error {
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

var frontPortTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["8P8C","8P6C","8P4C","8P2C","GG45","TERA 4P","TERA 2P","TERA 1P","110 Punch","BNC","MRJ21","FC","LC","LC/APC","LSH","LSH/APC","MPO","MTRJ","SC","SC/APC","ST","CS","SN","Splice"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontPortTypeTypeLabelPropEnum = append(frontPortTypeTypeLabelPropEnum, v)
	}
}

const (

	// FrontPortTypeLabelNr8P8C captures enum value "8P8C"
	FrontPortTypeLabelNr8P8C string = "8P8C"

	// FrontPortTypeLabelNr8P6C captures enum value "8P6C"
	FrontPortTypeLabelNr8P6C string = "8P6C"

	// FrontPortTypeLabelNr8P4C captures enum value "8P4C"
	FrontPortTypeLabelNr8P4C string = "8P4C"

	// FrontPortTypeLabelNr8P2C captures enum value "8P2C"
	FrontPortTypeLabelNr8P2C string = "8P2C"

	// FrontPortTypeLabelGG45 captures enum value "GG45"
	FrontPortTypeLabelGG45 string = "GG45"

	// FrontPortTypeLabelTERA4P captures enum value "TERA 4P"
	FrontPortTypeLabelTERA4P string = "TERA 4P"

	// FrontPortTypeLabelTERA2P captures enum value "TERA 2P"
	FrontPortTypeLabelTERA2P string = "TERA 2P"

	// FrontPortTypeLabelTERA1P captures enum value "TERA 1P"
	FrontPortTypeLabelTERA1P string = "TERA 1P"

	// FrontPortTypeLabelNr110Punch captures enum value "110 Punch"
	FrontPortTypeLabelNr110Punch string = "110 Punch"

	// FrontPortTypeLabelBNC captures enum value "BNC"
	FrontPortTypeLabelBNC string = "BNC"

	// FrontPortTypeLabelMRJ21 captures enum value "MRJ21"
	FrontPortTypeLabelMRJ21 string = "MRJ21"

	// FrontPortTypeLabelFC captures enum value "FC"
	FrontPortTypeLabelFC string = "FC"

	// FrontPortTypeLabelLC captures enum value "LC"
	FrontPortTypeLabelLC string = "LC"

	// FrontPortTypeLabelLCAPC captures enum value "LC/APC"
	FrontPortTypeLabelLCAPC string = "LC/APC"

	// FrontPortTypeLabelLSH captures enum value "LSH"
	FrontPortTypeLabelLSH string = "LSH"

	// FrontPortTypeLabelLSHAPC captures enum value "LSH/APC"
	FrontPortTypeLabelLSHAPC string = "LSH/APC"

	// FrontPortTypeLabelMPO captures enum value "MPO"
	FrontPortTypeLabelMPO string = "MPO"

	// FrontPortTypeLabelMTRJ captures enum value "MTRJ"
	FrontPortTypeLabelMTRJ string = "MTRJ"

	// FrontPortTypeLabelSC captures enum value "SC"
	FrontPortTypeLabelSC string = "SC"

	// FrontPortTypeLabelSCAPC captures enum value "SC/APC"
	FrontPortTypeLabelSCAPC string = "SC/APC"

	// FrontPortTypeLabelST captures enum value "ST"
	FrontPortTypeLabelST string = "ST"

	// FrontPortTypeLabelCS captures enum value "CS"
	FrontPortTypeLabelCS string = "CS"

	// FrontPortTypeLabelSN captures enum value "SN"
	FrontPortTypeLabelSN string = "SN"

	// FrontPortTypeLabelSplice captures enum value "Splice"
	FrontPortTypeLabelSplice string = "Splice"
)

// prop value enum
func (m *FrontPortType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, frontPortTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FrontPortType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var frontPortTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["8p8c","8p6c","8p4c","8p2c","gg45","tera-4p","tera-2p","tera-1p","110-punch","bnc","mrj21","fc","lc","lc-apc","lsh","lsh-apc","mpo","mtrj","sc","sc-apc","st","cs","sn","splice"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontPortTypeTypeValuePropEnum = append(frontPortTypeTypeValuePropEnum, v)
	}
}

const (

	// FrontPortTypeValueNr8p8c captures enum value "8p8c"
	FrontPortTypeValueNr8p8c string = "8p8c"

	// FrontPortTypeValueNr8p6c captures enum value "8p6c"
	FrontPortTypeValueNr8p6c string = "8p6c"

	// FrontPortTypeValueNr8p4c captures enum value "8p4c"
	FrontPortTypeValueNr8p4c string = "8p4c"

	// FrontPortTypeValueNr8p2c captures enum value "8p2c"
	FrontPortTypeValueNr8p2c string = "8p2c"

	// FrontPortTypeValueGg45 captures enum value "gg45"
	FrontPortTypeValueGg45 string = "gg45"

	// FrontPortTypeValueTeraDash4p captures enum value "tera-4p"
	FrontPortTypeValueTeraDash4p string = "tera-4p"

	// FrontPortTypeValueTeraDash2p captures enum value "tera-2p"
	FrontPortTypeValueTeraDash2p string = "tera-2p"

	// FrontPortTypeValueTeraDash1p captures enum value "tera-1p"
	FrontPortTypeValueTeraDash1p string = "tera-1p"

	// FrontPortTypeValueNr110DashPunch captures enum value "110-punch"
	FrontPortTypeValueNr110DashPunch string = "110-punch"

	// FrontPortTypeValueBnc captures enum value "bnc"
	FrontPortTypeValueBnc string = "bnc"

	// FrontPortTypeValueMrj21 captures enum value "mrj21"
	FrontPortTypeValueMrj21 string = "mrj21"

	// FrontPortTypeValueFc captures enum value "fc"
	FrontPortTypeValueFc string = "fc"

	// FrontPortTypeValueLc captures enum value "lc"
	FrontPortTypeValueLc string = "lc"

	// FrontPortTypeValueLcDashApc captures enum value "lc-apc"
	FrontPortTypeValueLcDashApc string = "lc-apc"

	// FrontPortTypeValueLsh captures enum value "lsh"
	FrontPortTypeValueLsh string = "lsh"

	// FrontPortTypeValueLshDashApc captures enum value "lsh-apc"
	FrontPortTypeValueLshDashApc string = "lsh-apc"

	// FrontPortTypeValueMpo captures enum value "mpo"
	FrontPortTypeValueMpo string = "mpo"

	// FrontPortTypeValueMtrj captures enum value "mtrj"
	FrontPortTypeValueMtrj string = "mtrj"

	// FrontPortTypeValueSc captures enum value "sc"
	FrontPortTypeValueSc string = "sc"

	// FrontPortTypeValueScDashApc captures enum value "sc-apc"
	FrontPortTypeValueScDashApc string = "sc-apc"

	// FrontPortTypeValueSt captures enum value "st"
	FrontPortTypeValueSt string = "st"

	// FrontPortTypeValueCs captures enum value "cs"
	FrontPortTypeValueCs string = "cs"

	// FrontPortTypeValueSn captures enum value "sn"
	FrontPortTypeValueSn string = "sn"

	// FrontPortTypeValueSplice captures enum value "splice"
	FrontPortTypeValueSplice string = "splice"
)

// prop value enum
func (m *FrontPortType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, frontPortTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FrontPortType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this front port type based on context it is used
func (m *FrontPortType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FrontPortType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FrontPortType) UnmarshalBinary(b []byte) error {
	var res FrontPortType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
