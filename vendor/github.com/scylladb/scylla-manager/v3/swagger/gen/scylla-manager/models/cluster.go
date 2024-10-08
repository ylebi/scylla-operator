// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Cluster cluster
//
// swagger:model Cluster
type Cluster struct {

	// auth token
	AuthToken string `json:"auth_token,omitempty"`

	// force non ssl session port
	ForceNonSslSessionPort bool `json:"force_non_ssl_session_port,omitempty"`

	// force tls disabled
	ForceTLSDisabled bool `json:"force_tls_disabled,omitempty"`

	// host
	Host string `json:"host,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// port
	Port int64 `json:"port,omitempty"`

	// ssl user cert file
	// Format: byte
	SslUserCertFile strfmt.Base64 `json:"ssl_user_cert_file,omitempty"`

	// ssl user key file
	// Format: byte
	SslUserKeyFile strfmt.Base64 `json:"ssl_user_key_file,omitempty"`

	// username
	Username string `json:"username,omitempty"`

	// without repair
	WithoutRepair bool `json:"without_repair,omitempty"`
}

// Validate validates this cluster
func (m *Cluster) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Cluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cluster) UnmarshalBinary(b []byte) error {
	var res Cluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
