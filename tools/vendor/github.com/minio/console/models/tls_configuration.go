// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2021 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TLSConfiguration tls configuration
//
// swagger:model tlsConfiguration
type TLSConfiguration struct {

	// ca certificates
	CaCertificates []string `json:"ca_certificates"`

	// console
	Console *KeyPairConfiguration `json:"console,omitempty"`

	// console ca certificates
	ConsoleCaCertificates []string `json:"console_ca_certificates"`

	// minio
	Minio []*KeyPairConfiguration `json:"minio"`
}

// Validate validates this tls configuration
func (m *TLSConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConsole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinio(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TLSConfiguration) validateConsole(formats strfmt.Registry) error {
	if swag.IsZero(m.Console) { // not required
		return nil
	}

	if m.Console != nil {
		if err := m.Console.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("console")
			}
			return err
		}
	}

	return nil
}

func (m *TLSConfiguration) validateMinio(formats strfmt.Registry) error {
	if swag.IsZero(m.Minio) { // not required
		return nil
	}

	for i := 0; i < len(m.Minio); i++ {
		if swag.IsZero(m.Minio[i]) { // not required
			continue
		}

		if m.Minio[i] != nil {
			if err := m.Minio[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("minio" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this tls configuration based on the context it is used
func (m *TLSConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConsole(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMinio(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TLSConfiguration) contextValidateConsole(ctx context.Context, formats strfmt.Registry) error {

	if m.Console != nil {
		if err := m.Console.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("console")
			}
			return err
		}
	}

	return nil
}

func (m *TLSConfiguration) contextValidateMinio(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Minio); i++ {

		if m.Minio[i] != nil {
			if err := m.Minio[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("minio" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TLSConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TLSConfiguration) UnmarshalBinary(b []byte) error {
	var res TLSConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}