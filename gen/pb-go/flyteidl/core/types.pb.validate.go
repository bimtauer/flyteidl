// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: flyteidl/core/types.proto

package core

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _types_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on SchemaType with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *SchemaType) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetColumns() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SchemaTypeValidationError{
					field:  fmt.Sprintf("Columns[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// SchemaTypeValidationError is the validation error returned by
// SchemaType.Validate if the designated constraints aren't met.
type SchemaTypeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SchemaTypeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SchemaTypeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SchemaTypeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SchemaTypeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SchemaTypeValidationError) ErrorName() string { return "SchemaTypeValidationError" }

// Error satisfies the builtin error interface
func (e SchemaTypeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSchemaType.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SchemaTypeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SchemaTypeValidationError{}

// Validate checks the field values on BlobType with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *BlobType) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Format

	// no validation rules for Dimensionality

	return nil
}

// BlobTypeValidationError is the validation error returned by
// BlobType.Validate if the designated constraints aren't met.
type BlobTypeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BlobTypeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BlobTypeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BlobTypeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BlobTypeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BlobTypeValidationError) ErrorName() string { return "BlobTypeValidationError" }

// Error satisfies the builtin error interface
func (e BlobTypeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBlobType.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BlobTypeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BlobTypeValidationError{}

// Validate checks the field values on LiteralType with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *LiteralType) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LiteralTypeValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.Type.(type) {

	case *LiteralType_Simple:
		// no validation rules for Simple

	case *LiteralType_Schema:

		if v, ok := interface{}(m.GetSchema()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LiteralTypeValidationError{
					field:  "Schema",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LiteralType_CollectionType:

		if v, ok := interface{}(m.GetCollectionType()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LiteralTypeValidationError{
					field:  "CollectionType",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LiteralType_MapValueType:

		if v, ok := interface{}(m.GetMapValueType()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LiteralTypeValidationError{
					field:  "MapValueType",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LiteralType_Blob:

		if v, ok := interface{}(m.GetBlob()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LiteralTypeValidationError{
					field:  "Blob",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// LiteralTypeValidationError is the validation error returned by
// LiteralType.Validate if the designated constraints aren't met.
type LiteralTypeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LiteralTypeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LiteralTypeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LiteralTypeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LiteralTypeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LiteralTypeValidationError) ErrorName() string { return "LiteralTypeValidationError" }

// Error satisfies the builtin error interface
func (e LiteralTypeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLiteralType.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LiteralTypeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LiteralTypeValidationError{}

// Validate checks the field values on OutputReference with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *OutputReference) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for NodeId

	// no validation rules for Var

	return nil
}

// OutputReferenceValidationError is the validation error returned by
// OutputReference.Validate if the designated constraints aren't met.
type OutputReferenceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OutputReferenceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OutputReferenceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OutputReferenceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OutputReferenceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OutputReferenceValidationError) ErrorName() string { return "OutputReferenceValidationError" }

// Error satisfies the builtin error interface
func (e OutputReferenceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOutputReference.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OutputReferenceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OutputReferenceValidationError{}

// Validate checks the field values on Error with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Error) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for FailedNodeId

	// no validation rules for Message

	return nil
}

// ErrorValidationError is the validation error returned by Error.Validate if
// the designated constraints aren't met.
type ErrorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ErrorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ErrorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ErrorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ErrorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ErrorValidationError) ErrorName() string { return "ErrorValidationError" }

// Error satisfies the builtin error interface
func (e ErrorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sError.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ErrorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ErrorValidationError{}

// Validate checks the field values on Secret with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Secret) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Type

	return nil
}

// SecretValidationError is the validation error returned by Secret.Validate if
// the designated constraints aren't met.
type SecretValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SecretValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SecretValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SecretValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SecretValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SecretValidationError) ErrorName() string { return "SecretValidationError" }

// Error satisfies the builtin error interface
func (e SecretValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSecret.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SecretValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SecretValidationError{}

// Validate checks the field values on OAuth2Client with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *OAuth2Client) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ClientId

	if v, ok := interface{}(m.GetClientSecret()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OAuth2ClientValidationError{
				field:  "ClientSecret",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// OAuth2ClientValidationError is the validation error returned by
// OAuth2Client.Validate if the designated constraints aren't met.
type OAuth2ClientValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OAuth2ClientValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OAuth2ClientValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OAuth2ClientValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OAuth2ClientValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OAuth2ClientValidationError) ErrorName() string { return "OAuth2ClientValidationError" }

// Error satisfies the builtin error interface
func (e OAuth2ClientValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOAuth2Client.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OAuth2ClientValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OAuth2ClientValidationError{}

// Validate checks the field values on Identity with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Identity) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for IamRole

	// no validation rules for K8SServiceAccount

	if v, ok := interface{}(m.GetOauth2Client()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IdentityValidationError{
				field:  "Oauth2Client",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// IdentityValidationError is the validation error returned by
// Identity.Validate if the designated constraints aren't met.
type IdentityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdentityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdentityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdentityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdentityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdentityValidationError) ErrorName() string { return "IdentityValidationError" }

// Error satisfies the builtin error interface
func (e IdentityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIdentity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdentityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdentityValidationError{}

// Validate checks the field values on OAuth2TokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *OAuth2TokenRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Type

	if v, ok := interface{}(m.GetClient()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OAuth2TokenRequestValidationError{
				field:  "Client",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for IdpDiscoveryEndpoint

	// no validation rules for TokenEndpoint

	return nil
}

// OAuth2TokenRequestValidationError is the validation error returned by
// OAuth2TokenRequest.Validate if the designated constraints aren't met.
type OAuth2TokenRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OAuth2TokenRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OAuth2TokenRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OAuth2TokenRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OAuth2TokenRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OAuth2TokenRequestValidationError) ErrorName() string {
	return "OAuth2TokenRequestValidationError"
}

// Error satisfies the builtin error interface
func (e OAuth2TokenRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOAuth2TokenRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OAuth2TokenRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OAuth2TokenRequestValidationError{}

// Validate checks the field values on SecurityContext with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *SecurityContext) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRunAs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SecurityContextValidationError{
				field:  "RunAs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetSecrets() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SecurityContextValidationError{
					field:  fmt.Sprintf("Secrets[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetTokens() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SecurityContextValidationError{
					field:  fmt.Sprintf("Tokens[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// SecurityContextValidationError is the validation error returned by
// SecurityContext.Validate if the designated constraints aren't met.
type SecurityContextValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SecurityContextValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SecurityContextValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SecurityContextValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SecurityContextValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SecurityContextValidationError) ErrorName() string { return "SecurityContextValidationError" }

// Error satisfies the builtin error interface
func (e SecurityContextValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSecurityContext.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SecurityContextValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SecurityContextValidationError{}

// Validate checks the field values on SchemaType_SchemaColumn with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SchemaType_SchemaColumn) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Type

	return nil
}

// SchemaType_SchemaColumnValidationError is the validation error returned by
// SchemaType_SchemaColumn.Validate if the designated constraints aren't met.
type SchemaType_SchemaColumnValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SchemaType_SchemaColumnValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SchemaType_SchemaColumnValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SchemaType_SchemaColumnValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SchemaType_SchemaColumnValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SchemaType_SchemaColumnValidationError) ErrorName() string {
	return "SchemaType_SchemaColumnValidationError"
}

// Error satisfies the builtin error interface
func (e SchemaType_SchemaColumnValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSchemaType_SchemaColumn.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SchemaType_SchemaColumnValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SchemaType_SchemaColumnValidationError{}
