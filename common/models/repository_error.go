package models

type RepositoryError struct {
	Code RepositoryErrorCode
}

const RepositoryErrorKey_ = ""

type RepositoryErrorCode int64

const (
	RepositoryErrorCode_OK RepositoryErrorCode = iota
	RepositoryErrorCode_Fail
	RepositoryErrorCode_NotFound
	RepositoryErrorCode_UpdateFail
)

func (code RepositoryErrorCode) Error() string {
	switch code {
	case RepositoryErrorCode_Fail:
		return "CreateFail"
	case RepositoryErrorCode_OK:
		return "OK"
	case RepositoryErrorCode_NotFound:
		return "Item not found"
	case RepositoryErrorCode_UpdateFail:
		return "Update fail"

	}

	return "unknown"
}
