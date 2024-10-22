package context_test_spies

import (
	"errors"
	"net/http"
)

func SpyResponseWriter() *spyResponseWriter {
	return &spyResponseWriter{}
}

type spyResponseWriter struct {
	Written bool
}

func (s *spyResponseWriter) Header() http.Header {
	s.Written = true
	return nil
}

func (s *spyResponseWriter) Write([]byte) (int, error) {
	s.Written = true
	return 0, errors.New("not implemented")
}

func (s *spyResponseWriter) WriteHeader(statusCode int) {
	s.Written = true
}
