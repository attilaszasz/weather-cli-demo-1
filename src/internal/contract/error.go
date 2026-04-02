package contract

import (
	"encoding/json"
	"errors"
	"io"
	"strings"
)

const (
	ExitCodeUsage      = 2
	ExitCodeDownstream = 3
	ExitCodeInternal   = 4
)

type ErrorCategory string

const (
	ErrorCategoryValidation ErrorCategory = "validation"
	ErrorCategoryTransport  ErrorCategory = "downstream_transport"
	ErrorCategoryProvider   ErrorCategory = "downstream_provider"
	ErrorCategoryInternal   ErrorCategory = "internal"
)

type ErrorResponse struct {
	Category ErrorCategory `json:"category"`
	Code     string        `json:"code"`
	Message  string        `json:"message"`
	ExitCode int           `json:"exitCode"`
}

type CLIError struct {
	response ErrorResponse
	cause    error
}

func (e CLIError) Error() string {
	return e.response.Message
}

func (e CLIError) Unwrap() error {
	return e.cause
}

func (e CLIError) ExitCode() int {
	return e.response.ExitCode
}

func (e CLIError) Response() ErrorResponse {
	return e.response
}

func NewCLIError(category ErrorCategory, code string, message string, exitCode int, cause error) CLIError {
	return CLIError{
		response: ErrorResponse{
			Category: category,
			Code:     code,
			Message:  message,
			ExitCode: exitCode,
		},
		cause: cause,
	}
}

func ClassifyError(err error) CLIError {
	var cliErr CLIError
	if errors.As(err, &cliErr) {
		return cliErr
	}

	message := err.Error()
	switch {
	case strings.Contains(message, "required"), strings.Contains(message, "must be between"), strings.Contains(message, "unexpected positional"):
		return NewCLIError(ErrorCategoryValidation, "invalid_input", message, ExitCodeUsage, err)
	case strings.Contains(message, "execute open-meteo request"), strings.Contains(message, "timeout"), strings.Contains(message, "deadline exceeded"), strings.Contains(message, "no such host"):
		return NewCLIError(ErrorCategoryTransport, "downstream_transport_failure", "weather provider transport failure", ExitCodeDownstream, err)
	case strings.Contains(message, "open-meteo returned status"), strings.Contains(message, "decode open-meteo response"), strings.Contains(message, "missing current time"):
		return NewCLIError(ErrorCategoryProvider, "downstream_provider_failure", "weather provider returned an invalid response", ExitCodeDownstream, err)
	default:
		return NewCLIError(ErrorCategoryInternal, "internal_failure", "internal command failure", ExitCodeInternal, err)
	}
}

func WriteErrorResponse(writer io.Writer, err error) error {
	cliErr := ClassifyError(err)
	encoder := json.NewEncoder(writer)
	if encodeErr := encoder.Encode(cliErr.Response()); encodeErr != nil {
		fallback := NewCLIError(ErrorCategoryInternal, "internal_failure", "internal command failure", ExitCodeInternal, encodeErr)
		return fallback
	}
	return cliErr
}

func ExitCode(err error) int {
	return ClassifyError(err).ExitCode()
}
