package server

import "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

/*
httpBodyMarshaler is a wrapper around runtime.HTTPBodyMarshaler that allows
specifying a custom delimeter. This is because of grpc-gateway default of
appending a newline to each chunk, which sort of makes sense for JSON, but
breaks just about everything else.
*/
type httpBodyMarshaler struct {
	runtime.HTTPBodyMarshaler
	delimeter []byte
}

func (m *httpBodyMarshaler) Delimiter() []byte {
	return m.delimeter
}
