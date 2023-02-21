package models

type RepositoryError struct {
	Code string
}

func (er *RepositoryError) Error() string {
	return er.Code
}

const RepositoryErrorKey_ = ""
