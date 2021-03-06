// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: image.proto

package pb

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
var _image_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Image with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Image) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for OwnerId

	// no validation rules for OwnerType

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ImageValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ImageValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ImageValidationError is the validation error returned by Image.Validate if
// the designated constraints aren't met.
type ImageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImageValidationError) ErrorName() string { return "ImageValidationError" }

// Error satisfies the builtin error interface
func (e ImageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImageValidationError{}

// Validate checks the field values on ImageInfo with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ImageInfo) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetOwnerId() < 1 {
		return ImageInfoValidationError{
			field:  "OwnerId",
			reason: "value must be greater than or equal to 1",
		}
	}

	if _, ok := OwnerType_name[int32(m.GetOwnerType())]; !ok {
		return ImageInfoValidationError{
			field:  "OwnerType",
			reason: "value must be one of the defined enum values",
		}
	}

	return nil
}

// ImageInfoValidationError is the validation error returned by
// ImageInfo.Validate if the designated constraints aren't met.
type ImageInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImageInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImageInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImageInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImageInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImageInfoValidationError) ErrorName() string { return "ImageInfoValidationError" }

// Error satisfies the builtin error interface
func (e ImageInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImageInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImageInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImageInfoValidationError{}

// Validate checks the field values on ListImagesByOwnerIDReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListImagesByOwnerIDReq) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetOwnerId() < 1 {
		return ListImagesByOwnerIDReqValidationError{
			field:  "OwnerId",
			reason: "value must be greater than or equal to 1",
		}
	}

	if _, ok := OwnerType_name[int32(m.GetOwnerType())]; !ok {
		return ListImagesByOwnerIDReqValidationError{
			field:  "OwnerType",
			reason: "value must be one of the defined enum values",
		}
	}

	return nil
}

// ListImagesByOwnerIDReqValidationError is the validation error returned by
// ListImagesByOwnerIDReq.Validate if the designated constraints aren't met.
type ListImagesByOwnerIDReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListImagesByOwnerIDReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListImagesByOwnerIDReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListImagesByOwnerIDReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListImagesByOwnerIDReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListImagesByOwnerIDReqValidationError) ErrorName() string {
	return "ListImagesByOwnerIDReqValidationError"
}

// Error satisfies the builtin error interface
func (e ListImagesByOwnerIDReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListImagesByOwnerIDReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListImagesByOwnerIDReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListImagesByOwnerIDReqValidationError{}

// Validate checks the field values on ListImagesByOwnerIDRes with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListImagesByOwnerIDRes) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetImages() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListImagesByOwnerIDResValidationError{
					field:  fmt.Sprintf("Images[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListImagesByOwnerIDResValidationError is the validation error returned by
// ListImagesByOwnerIDRes.Validate if the designated constraints aren't met.
type ListImagesByOwnerIDResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListImagesByOwnerIDResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListImagesByOwnerIDResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListImagesByOwnerIDResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListImagesByOwnerIDResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListImagesByOwnerIDResValidationError) ErrorName() string {
	return "ListImagesByOwnerIDResValidationError"
}

// Error satisfies the builtin error interface
func (e ListImagesByOwnerIDResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListImagesByOwnerIDRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListImagesByOwnerIDResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListImagesByOwnerIDResValidationError{}

// Validate checks the field values on BatchCreateImagesRes with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *BatchCreateImagesRes) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetImages() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BatchCreateImagesResValidationError{
					field:  fmt.Sprintf("Images[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// BatchCreateImagesResValidationError is the validation error returned by
// BatchCreateImagesRes.Validate if the designated constraints aren't met.
type BatchCreateImagesResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BatchCreateImagesResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BatchCreateImagesResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BatchCreateImagesResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BatchCreateImagesResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BatchCreateImagesResValidationError) ErrorName() string {
	return "BatchCreateImagesResValidationError"
}

// Error satisfies the builtin error interface
func (e BatchCreateImagesResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBatchCreateImagesRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BatchCreateImagesResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BatchCreateImagesResValidationError{}

// Validate checks the field values on BatchCreateImagesReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *BatchCreateImagesReq) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Data.(type) {

	case *BatchCreateImagesReq_Info:

		if v, ok := interface{}(m.GetInfo()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BatchCreateImagesReqValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *BatchCreateImagesReq_Chunk:

		if len(m.GetChunk()) > 65536 {
			return BatchCreateImagesReqValidationError{
				field:  "Chunk",
				reason: "value length must be at most 65536 bytes",
			}
		}

	default:
		return BatchCreateImagesReqValidationError{
			field:  "Data",
			reason: "value is required",
		}

	}

	return nil
}

// BatchCreateImagesReqValidationError is the validation error returned by
// BatchCreateImagesReq.Validate if the designated constraints aren't met.
type BatchCreateImagesReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BatchCreateImagesReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BatchCreateImagesReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BatchCreateImagesReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BatchCreateImagesReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BatchCreateImagesReqValidationError) ErrorName() string {
	return "BatchCreateImagesReqValidationError"
}

// Error satisfies the builtin error interface
func (e BatchCreateImagesReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBatchCreateImagesReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BatchCreateImagesReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BatchCreateImagesReqValidationError{}

// Validate checks the field values on BatchDeleteImagesReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *BatchDeleteImagesReq) Validate() error {
	if m == nil {
		return nil
	}

	_BatchDeleteImagesReq_Ids_Unique := make(map[int64]struct{}, len(m.GetIds()))

	for idx, item := range m.GetIds() {
		_, _ = idx, item

		if _, exists := _BatchDeleteImagesReq_Ids_Unique[item]; exists {
			return BatchDeleteImagesReqValidationError{
				field:  fmt.Sprintf("Ids[%v]", idx),
				reason: "repeated value must contain unique items",
			}
		} else {
			_BatchDeleteImagesReq_Ids_Unique[item] = struct{}{}
		}

		// no validation rules for Ids[idx]
	}

	return nil
}

// BatchDeleteImagesReqValidationError is the validation error returned by
// BatchDeleteImagesReq.Validate if the designated constraints aren't met.
type BatchDeleteImagesReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BatchDeleteImagesReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BatchDeleteImagesReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BatchDeleteImagesReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BatchDeleteImagesReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BatchDeleteImagesReqValidationError) ErrorName() string {
	return "BatchDeleteImagesReqValidationError"
}

// Error satisfies the builtin error interface
func (e BatchDeleteImagesReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBatchDeleteImagesReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BatchDeleteImagesReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BatchDeleteImagesReqValidationError{}

// Validate checks the field values on DeleteImagesByOwnerIDReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteImagesByOwnerIDReq) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetOwnerId() < 1 {
		return DeleteImagesByOwnerIDReqValidationError{
			field:  "OwnerId",
			reason: "value must be greater than or equal to 1",
		}
	}

	if _, ok := OwnerType_name[int32(m.GetOwnerType())]; !ok {
		return DeleteImagesByOwnerIDReqValidationError{
			field:  "OwnerType",
			reason: "value must be one of the defined enum values",
		}
	}

	return nil
}

// DeleteImagesByOwnerIDReqValidationError is the validation error returned by
// DeleteImagesByOwnerIDReq.Validate if the designated constraints aren't met.
type DeleteImagesByOwnerIDReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteImagesByOwnerIDReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteImagesByOwnerIDReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteImagesByOwnerIDReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteImagesByOwnerIDReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteImagesByOwnerIDReqValidationError) ErrorName() string {
	return "DeleteImagesByOwnerIDReqValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteImagesByOwnerIDReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteImagesByOwnerIDReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteImagesByOwnerIDReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteImagesByOwnerIDReqValidationError{}
