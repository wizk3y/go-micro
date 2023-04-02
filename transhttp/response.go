package transhttp

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// RespondJSON -- makes the response with payload as json format
func RespondJSON(w http.ResponseWriter, httpStatusCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(httpStatusCode)
	w.Write(data)
}

// RespondError -- makes the error response with payload as json format
func RespondError(w http.ResponseWriter, httpStatusCode int, message string) {
	RespondJSON(w, httpStatusCode, map[string]string{"error": message})
}

// RespondJSONError --
func RespondJSONError(w http.ResponseWriter, payload interface{}) {
	var httpStatusCode = http.StatusInternalServerError
	if err, ok := payload.(error); ok {
		httpStatusCode = GetStatusCode(err)
	}
	RespondJSON(w, httpStatusCode, payload)
}

// RespondMessage -- makes the message response with payload as json format
func RespondMessage(w http.ResponseWriter, httpStatusCode int, message string) {
	RespondJSON(w, httpStatusCode, map[string]string{"message": message})
}

// RespondMessageWithContentType -- makes the message response with payload as json format
func RespondMessageWithContentType(w http.ResponseWriter, httpStatusCode int, message string, contentType string) {
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", strconv.Itoa(len(message)))
	w.WriteHeader(httpStatusCode)
	w.Write([]byte(message))
}

// Redirect -- redirect
func Redirect(w http.ResponseWriter, r *http.Request, url string) {
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}

// RespondFile --
func RespondFile(w http.ResponseWriter, r *http.Request, fileName string) {
	http.ServeFile(w, r, fileName)
}
