package models

type ServiceErrorCode int64

type ServiceError struct {
	Code ServiceErrorCode
}

const (
	ServiceErrorCode_OK ServiceErrorCode = iota
	ServiceErrorCode_Fail
)

func (code ServiceErrorCode) String() string {
	switch code {
	case ServiceErrorCode_Fail:
		return "ServiceErrorCode"
	}
	return "unknown"
}

func (se ServiceError) Error() string {
	switch se.Code {
	case ServiceErrorCode_Fail:
		return "ServiceErrorCode"
	}
	return "unknown"
}
