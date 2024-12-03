package middleware

import (
	"context"
	"encoding/json"
	"github.com/Technology-99/third_party/response"
	"github.com/Technology-99/third_party/sony"
	"net/http"
)

func CommonErrResponse(w http.ResponseWriter, r *http.Request, Code int32, v ...any) {

	msg := ""
	if len(v) > 0 {
		msg = v[0].(string)
	} else {
		msg = response.StatusText(Code)
	}

	requestID := ""
	if r.Context().Value("RequestID") == nil {
		requestID = sony.NextId()
	} else {
		requestID = r.Context().Value("RequestID").(string)
	}
	ctx := context.WithValue(r.Context(), "RequestID", requestID)
	r = r.WithContext(ctx)

	resp := response.CommonResponse{
		Code:      Code,
		Msg:       msg,
		RequestID: requestID,
		Path:      r.RequestURI,
	}
	body, _ := json.Marshal(&resp)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
