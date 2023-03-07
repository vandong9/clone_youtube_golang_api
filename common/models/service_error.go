package models

type ServiceErrorCode int64

type ServiceError struct {
	Code ServiceErrorCode
}

const (
	ServiceErrorCode_OK ServiceErrorCode = iota
	ServiceErrorCode_Fail
	ServiceErrorCode_NotFound
)

func (code ServiceErrorCode) String() string {
	switch code {
	case ServiceErrorCode_Fail:
		return "ServiceErrorCode"
	case ServiceErrorCode_NotFound:
		return "Item NotFound"
	}
	return "unknown"
}

func (se ServiceError) Error() string {
	switch se.Code {
	case ServiceErrorCode_Fail:
		return "ServiceErrorCode"
	case ServiceErrorCode_NotFound:
		return "Item NotFound"

	}
	return "unknown"
}
