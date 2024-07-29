// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/review/v1/review.proto

package v1

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

// Validate checks the field values on CreateReviewRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateReviewRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateReviewRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateReviewRequestMultiError, or nil if none found.
func (m *CreateReviewRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateReviewRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserID() <= 0 {
		err := CreateReviewRequestValidationError{
			field:  "UserID",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetOrderID() <= 0 {
		err := CreateReviewRequestValidationError{
			field:  "OrderID",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetScore(); val <= 0 || val > 5 {
		err := CreateReviewRequestValidationError{
			field:  "Score",
			reason: "value must be inside range (0, 5]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetServiceScore(); val <= 0 || val > 5 {
		err := CreateReviewRequestValidationError{
			field:  "ServiceScore",
			reason: "value must be inside range (0, 5]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetExpressScore(); val <= 0 || val > 5 {
		err := CreateReviewRequestValidationError{
			field:  "ExpressScore",
			reason: "value must be inside range (0, 5]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetContent()); l < 8 || l > 255 {
		err := CreateReviewRequestValidationError{
			field:  "Content",
			reason: "value length must be between 8 and 255 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for PicInfo

	// no validation rules for VideoInfo

	// no validation rules for Anonymous

	if len(errors) > 0 {
		return CreateReviewRequestMultiError(errors)
	}

	return nil
}

// CreateReviewRequestMultiError is an error wrapping multiple validation
// errors returned by CreateReviewRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateReviewRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateReviewRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateReviewRequestMultiError) AllErrors() []error { return m }

// CreateReviewRequestValidationError is the validation error returned by
// CreateReviewRequest.Validate if the designated constraints aren't met.
type CreateReviewRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateReviewRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateReviewRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateReviewRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateReviewRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateReviewRequestValidationError) ErrorName() string {
	return "CreateReviewRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateReviewRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateReviewRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateReviewRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateReviewRequestValidationError{}

// Validate checks the field values on CreateReviewReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateReviewReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateReviewReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateReviewReplyMultiError, or nil if none found.
func (m *CreateReviewReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateReviewReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ReviewID

	if len(errors) > 0 {
		return CreateReviewReplyMultiError(errors)
	}

	return nil
}

// CreateReviewReplyMultiError is an error wrapping multiple validation errors
// returned by CreateReviewReply.ValidateAll() if the designated constraints
// aren't met.
type CreateReviewReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateReviewReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateReviewReplyMultiError) AllErrors() []error { return m }

// CreateReviewReplyValidationError is the validation error returned by
// CreateReviewReply.Validate if the designated constraints aren't met.
type CreateReviewReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateReviewReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateReviewReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateReviewReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateReviewReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateReviewReplyValidationError) ErrorName() string {
	return "CreateReviewReplyValidationError"
}

// Error satisfies the builtin error interface
func (e CreateReviewReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateReviewReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateReviewReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateReviewReplyValidationError{}

// Validate checks the field values on UpdateReviewRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateReviewRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateReviewRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateReviewRequestMultiError, or nil if none found.
func (m *UpdateReviewRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateReviewRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateReviewRequestMultiError(errors)
	}

	return nil
}

// UpdateReviewRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateReviewRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateReviewRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateReviewRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateReviewRequestMultiError) AllErrors() []error { return m }

// UpdateReviewRequestValidationError is the validation error returned by
// UpdateReviewRequest.Validate if the designated constraints aren't met.
type UpdateReviewRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateReviewRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateReviewRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateReviewRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateReviewRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateReviewRequestValidationError) ErrorName() string {
	return "UpdateReviewRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateReviewRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateReviewRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateReviewRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateReviewRequestValidationError{}

// Validate checks the field values on UpdateReviewReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateReviewReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateReviewReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateReviewReplyMultiError, or nil if none found.
func (m *UpdateReviewReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateReviewReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateReviewReplyMultiError(errors)
	}

	return nil
}

// UpdateReviewReplyMultiError is an error wrapping multiple validation errors
// returned by UpdateReviewReply.ValidateAll() if the designated constraints
// aren't met.
type UpdateReviewReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateReviewReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateReviewReplyMultiError) AllErrors() []error { return m }

// UpdateReviewReplyValidationError is the validation error returned by
// UpdateReviewReply.Validate if the designated constraints aren't met.
type UpdateReviewReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateReviewReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateReviewReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateReviewReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateReviewReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateReviewReplyValidationError) ErrorName() string {
	return "UpdateReviewReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateReviewReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateReviewReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateReviewReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateReviewReplyValidationError{}

// Validate checks the field values on DeleteReviewRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteReviewRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteReviewRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteReviewRequestMultiError, or nil if none found.
func (m *DeleteReviewRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteReviewRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteReviewRequestMultiError(errors)
	}

	return nil
}

// DeleteReviewRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteReviewRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteReviewRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteReviewRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteReviewRequestMultiError) AllErrors() []error { return m }

// DeleteReviewRequestValidationError is the validation error returned by
// DeleteReviewRequest.Validate if the designated constraints aren't met.
type DeleteReviewRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteReviewRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteReviewRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteReviewRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteReviewRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteReviewRequestValidationError) ErrorName() string {
	return "DeleteReviewRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteReviewRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteReviewRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteReviewRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteReviewRequestValidationError{}

// Validate checks the field values on DeleteReviewReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteReviewReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteReviewReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteReviewReplyMultiError, or nil if none found.
func (m *DeleteReviewReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteReviewReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteReviewReplyMultiError(errors)
	}

	return nil
}

// DeleteReviewReplyMultiError is an error wrapping multiple validation errors
// returned by DeleteReviewReply.ValidateAll() if the designated constraints
// aren't met.
type DeleteReviewReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteReviewReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteReviewReplyMultiError) AllErrors() []error { return m }

// DeleteReviewReplyValidationError is the validation error returned by
// DeleteReviewReply.Validate if the designated constraints aren't met.
type DeleteReviewReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteReviewReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteReviewReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteReviewReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteReviewReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteReviewReplyValidationError) ErrorName() string {
	return "DeleteReviewReplyValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteReviewReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteReviewReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteReviewReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteReviewReplyValidationError{}

// Validate checks the field values on GetReviewRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetReviewRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetReviewRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetReviewRequestMultiError, or nil if none found.
func (m *GetReviewRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetReviewRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetReviewRequestMultiError(errors)
	}

	return nil
}

// GetReviewRequestMultiError is an error wrapping multiple validation errors
// returned by GetReviewRequest.ValidateAll() if the designated constraints
// aren't met.
type GetReviewRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetReviewRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetReviewRequestMultiError) AllErrors() []error { return m }

// GetReviewRequestValidationError is the validation error returned by
// GetReviewRequest.Validate if the designated constraints aren't met.
type GetReviewRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetReviewRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetReviewRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetReviewRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetReviewRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetReviewRequestValidationError) ErrorName() string { return "GetReviewRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetReviewRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetReviewRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetReviewRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetReviewRequestValidationError{}

// Validate checks the field values on GetReviewReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetReviewReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetReviewReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetReviewReplyMultiError,
// or nil if none found.
func (m *GetReviewReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetReviewReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetReviewReplyMultiError(errors)
	}

	return nil
}

// GetReviewReplyMultiError is an error wrapping multiple validation errors
// returned by GetReviewReply.ValidateAll() if the designated constraints
// aren't met.
type GetReviewReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetReviewReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetReviewReplyMultiError) AllErrors() []error { return m }

// GetReviewReplyValidationError is the validation error returned by
// GetReviewReply.Validate if the designated constraints aren't met.
type GetReviewReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetReviewReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetReviewReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetReviewReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetReviewReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetReviewReplyValidationError) ErrorName() string { return "GetReviewReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetReviewReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetReviewReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetReviewReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetReviewReplyValidationError{}

// Validate checks the field values on ListReviewRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListReviewRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListReviewRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListReviewRequestMultiError, or nil if none found.
func (m *ListReviewRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListReviewRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListReviewRequestMultiError(errors)
	}

	return nil
}

// ListReviewRequestMultiError is an error wrapping multiple validation errors
// returned by ListReviewRequest.ValidateAll() if the designated constraints
// aren't met.
type ListReviewRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListReviewRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListReviewRequestMultiError) AllErrors() []error { return m }

// ListReviewRequestValidationError is the validation error returned by
// ListReviewRequest.Validate if the designated constraints aren't met.
type ListReviewRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListReviewRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListReviewRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListReviewRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListReviewRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListReviewRequestValidationError) ErrorName() string {
	return "ListReviewRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListReviewRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListReviewRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListReviewRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListReviewRequestValidationError{}

// Validate checks the field values on ListReviewReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListReviewReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListReviewReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListReviewReplyMultiError, or nil if none found.
func (m *ListReviewReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListReviewReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListReviewReplyMultiError(errors)
	}

	return nil
}

// ListReviewReplyMultiError is an error wrapping multiple validation errors
// returned by ListReviewReply.ValidateAll() if the designated constraints
// aren't met.
type ListReviewReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListReviewReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListReviewReplyMultiError) AllErrors() []error { return m }

// ListReviewReplyValidationError is the validation error returned by
// ListReviewReply.Validate if the designated constraints aren't met.
type ListReviewReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListReviewReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListReviewReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListReviewReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListReviewReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListReviewReplyValidationError) ErrorName() string { return "ListReviewReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListReviewReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListReviewReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListReviewReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListReviewReplyValidationError{}
