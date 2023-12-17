package middlewares

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log/slog"
	"time"
	"unsafe"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

type logMessage struct {
	StartTime string       `json:"start_time"`
	Latency   int          `json:"latency"`
	Request   *logRequest  `json:"request"`
	Response  *logResponse `json:"response"`
	Ip        string       `json:"ip"`
}

type logRequest struct {
	Method   string `json:"method"`
	Body     string `json:"body"`
	ClientIP string `json:"client_ip"`
	RemoteIP string `json:"remote_ip"`
	URL      string `json:"url"`
}

type logResponse struct {
	Body string `json:"body"`
}

func Logger(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 请求前
		var (
			writer    = &bodyLogWriter{body: bytes.NewBuffer(nil), ResponseWriter: c.Writer}
			bodyBytes []byte
		)

		c.Writer = writer

		if c.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}
		c.Next()

		latency := time.Since(t)
		respBody := writer.body.Bytes()
		s := logMessage{
			StartTime: t.Format(time.DateTime),
			Ip:        c.ClientIP(),
			Latency:   int(latency.Milliseconds()),
			Request: &logRequest{
				Body:   unsafe.String(unsafe.SliceData(bodyBytes), len(bodyBytes)),
				Method: c.Request.Method,
			},
			Response: &logResponse{
				Body: unsafe.String(unsafe.SliceData(respBody), len(respBody)),
			},
		}
		message, _ := json.Marshal(s)
		log.Info(string(message))
	}

}
