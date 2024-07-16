package data

type ApiConvertible interface {
	MarshalApiResponse() ([]byte, error)
}
