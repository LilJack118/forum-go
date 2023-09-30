package httpErrors

import (
	"encoding/json"
	"net/http"
)

func JSONError(w http.ResponseWriter, errMsg string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)

	// try to convert errMsg to json
	var detailMsg interface{}
	err := json.Unmarshal([]byte(errMsg), &detailMsg)
	if err != nil {
		// If errMsg is not a valid JSON, use it as a plain string.
		detailMsg = errMsg
	}

	err_map := map[string]interface{}{"type": "error", "detail": detailMsg}
	errString, err := json.Marshal(err_map)

	if err != nil {
		http.Error(w, errMsg, code)
	}

	w.Write(errString)
}
