package errors

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"runtime"
)

type Error interface {
	Error() string
	StatusCode() int
	GetCause() Error
	GetStackTrace() *string
}

type BaseError struct {
	Message    string
	Cause      Error
	StackTrace *string
}

type ErrorToJson struct {
	Type    string `json:"errorType"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type ErrorsToJson struct {
	Errors []*ErrorToJson `json:"errors"`
}

func NewBaseError(message string, cause Error, captureStackTrace bool) BaseError {
	var stackString *string = nil

	if captureStackTrace {
		buffer := make([]byte, 4096) // adjust buffer size to be larger than expected stack
		stackLength := runtime.Stack(buffer, false)
		s := string(buffer[:stackLength])
		stackString = &s
	}

	return BaseError{
		Message:    message,
		Cause:      cause,
		StackTrace: stackString,
	}
}

func (err *BaseError) Error() string {
	return err.Message
}

func (err *BaseError) StatusCode() int {
	return http.StatusInternalServerError
}

func (err *BaseError) GetCause() Error {
	return err.Cause
}

func (err *BaseError) GetStackTrace() *string {
	return err.StackTrace
}

func IsError(err any) (bool, Error) {
	thisErrType, ok := err.(Error)
	if ok {
		return true, thisErrType
	}
	return false, nil
}

func ToError(err error) Error {
	return &BaseError{
		Message: err.Error(),
	}
}

func ToJson(err Error) []byte {
	var thisError Error
	var errorsArray ErrorsToJson

	thisError = err
	for thisError != nil {
		errorsArray.Errors = append(errorsArray.Errors, &ErrorToJson{
			Type:    fmt.Sprintf("%s", reflect.TypeOf(thisError)),
			Message: thisError.Error(),
			Status:  thisError.StatusCode(),
		})

		thisError = thisError.GetCause()
	}

	jsonResponse, jsonError := json.Marshal(errorsArray)

	if jsonError != nil {
		log.Println("Unable to encode JSON")
		panic(jsonError)
	}

	return jsonResponse
}

func ToJsonString(err Error) string {
	bytes := ToJson(err)
	return string(bytes)
}
