// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/admin/admin.proto

package adminv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UserMultiError, or nil if none found.
func (m *User) ValidateAll() error {
	return m.validate(true)
}

func (m *User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Email

	// no validation rules for CreatedAt

	if len(errors) > 0 {
		return UserMultiError(errors)
	}

	return nil
}

// UserMultiError is an error wrapping multiple validation errors returned by
// User.ValidateAll() if the designated constraints aren't met.
type UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMultiError) AllErrors() []error { return m }

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

// Validate checks the field values on GetUserDetailsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetUserDetailsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserDetailsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserDetailsRequestMultiError, or nil if none found.
func (m *GetUserDetailsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserDetailsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	if len(errors) > 0 {
		return GetUserDetailsRequestMultiError(errors)
	}

	return nil
}

// GetUserDetailsRequestMultiError is an error wrapping multiple validation
// errors returned by GetUserDetailsRequest.ValidateAll() if the designated
// constraints aren't met.
type GetUserDetailsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserDetailsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserDetailsRequestMultiError) AllErrors() []error { return m }

// GetUserDetailsRequestValidationError is the validation error returned by
// GetUserDetailsRequest.Validate if the designated constraints aren't met.
type GetUserDetailsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserDetailsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserDetailsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserDetailsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserDetailsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserDetailsRequestValidationError) ErrorName() string {
	return "GetUserDetailsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserDetailsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserDetailsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserDetailsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserDetailsRequestValidationError{}

// Validate checks the field values on GetUserDetailsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetUserDetailsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserDetailsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserDetailsResponseMultiError, or nil if none found.
func (m *GetUserDetailsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserDetailsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetUserDetailsResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetUserDetailsResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetUserDetailsResponseValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetUserDetailsResponseMultiError(errors)
	}

	return nil
}

// GetUserDetailsResponseMultiError is an error wrapping multiple validation
// errors returned by GetUserDetailsResponse.ValidateAll() if the designated
// constraints aren't met.
type GetUserDetailsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserDetailsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserDetailsResponseMultiError) AllErrors() []error { return m }

// GetUserDetailsResponseValidationError is the validation error returned by
// GetUserDetailsResponse.Validate if the designated constraints aren't met.
type GetUserDetailsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserDetailsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserDetailsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserDetailsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserDetailsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserDetailsResponseValidationError) ErrorName() string {
	return "GetUserDetailsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserDetailsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserDetailsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserDetailsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserDetailsResponseValidationError{}
