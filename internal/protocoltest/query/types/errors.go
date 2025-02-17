// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// This error is thrown when a request is invalid.
type ComplexError struct {
	Message *string

	ErrorCodeOverride *string

	TopLevel *string
	Nested   *ComplexNestedErrorData

	noSmithyDocumentSerde
}

func (e *ComplexError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ComplexError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ComplexError) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ComplexError"
	}
	return *e.ErrorCodeOverride
}
func (e *ComplexError) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type CustomCodeError struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *CustomCodeError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CustomCodeError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CustomCodeError) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "Customized"
	}
	return *e.ErrorCodeOverride
}
func (e *CustomCodeError) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This error is thrown when an invalid greeting value is provided.
type InvalidGreeting struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidGreeting) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidGreeting) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidGreeting) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InvalidGreeting"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidGreeting) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
