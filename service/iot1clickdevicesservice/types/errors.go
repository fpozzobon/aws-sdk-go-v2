// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

type ForbiddenException struct {
	Message *string

	ErrorCodeOverride *string

	Code *string

	noSmithyDocumentSerde
}

func (e *ForbiddenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ForbiddenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ForbiddenException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ForbiddenException"
	}
	return *e.ErrorCodeOverride
}
func (e *ForbiddenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type InternalFailureException struct {
	Message *string

	ErrorCodeOverride *string

	Code *string

	noSmithyDocumentSerde
}

func (e *InternalFailureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalFailureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalFailureException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InternalFailureException"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalFailureException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

type InvalidRequestException struct {
	Message *string

	ErrorCodeOverride *string

	Code *string

	noSmithyDocumentSerde
}

func (e *InvalidRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRequestException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InvalidRequestException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type PreconditionFailedException struct {
	Message *string

	ErrorCodeOverride *string

	Code *string

	noSmithyDocumentSerde
}

func (e *PreconditionFailedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PreconditionFailedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PreconditionFailedException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "PreconditionFailedException"
	}
	return *e.ErrorCodeOverride
}
func (e *PreconditionFailedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type RangeNotSatisfiableException struct {
	Message *string

	ErrorCodeOverride *string

	Code *string

	noSmithyDocumentSerde
}

func (e *RangeNotSatisfiableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RangeNotSatisfiableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RangeNotSatisfiableException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "RangeNotSatisfiableException"
	}
	return *e.ErrorCodeOverride
}
func (e *RangeNotSatisfiableException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type ResourceConflictException struct {
	Message *string

	ErrorCodeOverride *string

	Code *string

	noSmithyDocumentSerde
}

func (e *ResourceConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceConflictException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ResourceConflictException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type ResourceNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	Code *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ResourceNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
