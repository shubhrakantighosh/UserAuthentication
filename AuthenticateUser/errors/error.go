package errors

import (
	"encoding/json"
	"regexp"
	"time"
)

type CustomError struct {
	Time         time.Time         `json:"time"`
	UserMessage  string            `json:"-"`
	ErrorExist   bool              `json:"error_exist"`
	Error        error             `json:"-"`
	StructErrors map[string]string `json:"struct_errors"`
}

func (c CustomError) Exist() bool {
	return c.ErrorExist
}

type gormErrorMessage struct {
	Severity       string `json:"Severity"`
	Code           string `json:"Code"`
	Message        string `json:"Message"`
	Detail         string `json:"Detail"`
	Where          string `json:"Where"`
	SchemaName     string `json:"SchemaName"`
	TableName      string `json:"TableName"`
	ConstraintName string `json:"ConstraintName"`
}

func DbError(e error) CustomError {
	errMessage := map[string]string{}
	errMessage["Error"] = e.Error()

	errorStruct, err := json.Marshal(e)
	if err != nil {
		return CustomError{Time: time.Now(), StructErrors: errMessage, Error: err, ErrorExist: true}
	}
	var cusErr gormErrorMessage
	err = json.Unmarshal(errorStruct, &cusErr)
	if err != nil {
		return CustomError{Time: time.Now(), StructErrors: errMessage, Error: err, ErrorExist: true}
	}
	message := userErrorMessage(cusErr.Detail, e)
	return CustomError{Time: time.Now(), StructErrors: message, Error: err, ErrorExist: true}
}

func userErrorMessage(str string, e error) map[string]string {
	var key, value string
	message := map[string]string{}
	regex := regexp.MustCompile(`\((.*?)\)`)
	match := regex.FindStringSubmatch(str)
	if len(match) > 1 {
		key = match[1]
	} else {
		message["Error"] = e.Error()
	}
	regex = regexp.MustCompile(`\)\s+(.*)`)
	match = regex.FindStringSubmatch(str)
	if len(match) > 1 {
		value = match[1]
	} else {
		message["Error"] = e.Error()
	}
	message[key] = value
	return message
}
