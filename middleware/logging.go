package middleware

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

// LogRequest logs the details of the incoming HTTP request
func LogRequest(r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		return
	}

	log.Printf("Incoming Request: %s %s %s\nHeaders: %v\nBody: %s",
		r.Method, r.RequestURI, r.Proto, r.Header, string(requestBody))

	// Reset the request body for further processing
	r.Body = io.NopCloser(bytes.NewBuffer(requestBody))
}

// LogResponse logs the details of the outgoing HTTP response
func LogResponse(w http.ResponseWriter, r *http.Request, next http.Handler) {
	responseRecorder := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
	next.ServeHTTP(responseRecorder, r)

	log.Printf("Outgoing Response: Status: %d\nHeaders: %v\nBody: %s",
		responseRecorder.statusCode, responseRecorder.Header(), responseRecorder.body.String())
}

// responseWriter is a custom http.ResponseWriter that captures the response details
type responseWriter struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body.Write(b)
	return rw.ResponseWriter.Write(b)
}

// LoggingMiddleware logs the requests and responses
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		LogRequest(r)
		LogResponse(w, r, next)
	})
}
