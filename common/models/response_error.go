package models

type ResponseError struct {
	Code ResponseErrorCode
}

type Message struct {
	Title   string
	Content string
}

type ResponseErrorCode int64

const (
	ReponseErr_Badaccss ResponseErrorCode = iota
	// "ReponseErr_Badaccss"
)

func (code ResponseErrorCode) String() string {
	switch code {
	case ReponseErr_Badaccss:
		return "ReponseErr_Badaccss"
	}
	return "unknown"
}
