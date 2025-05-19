package responseModels

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Code int
	Err  error
}

func (e *ErrorResponse) Error() string {
	return e.Err.Error()
}

func (e *ErrorResponse) Write(w http.ResponseWriter) {
	w.WriteHeader(e.Code)
	w.Write([]byte(fmt.Sprintf(`{"error":"%s","code":%d}`, e.Err.Error(), e.Code)))
}

type SuccessResponse struct {
	Code int
	Data interface{}
}

func (s *SuccessResponse) Write(w http.ResponseWriter) {
	byt, err := json.Marshal(s.Data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(s.Code)
	w.Write(byt)
}
