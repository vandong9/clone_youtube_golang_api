package models

type RepositoryError struct {
	Code RepositoryErrorCode
}

const RepositoryErrorKey_ = ""

type RepositoryErrorCode int64

const (
	RepositoryErrorCode_OK RepositoryErrorCode = iota
	RepositoryErrorCode_Fail
)

func (code RepositoryErrorCode) String() string {
	switch code {
	case RepositoryErrorCode_Fail:
		return "CreateFail"
	case RepositoryErrorCode_OK:
		return "OK"
	}

	return "unknown"
}
