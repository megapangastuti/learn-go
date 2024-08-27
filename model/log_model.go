package model

import (
	"log"
	"time"
)

type LogModel struct {
	AccessTime time.Time
	Latency    time.Duration
	ClientIP   string
	Method     string
	Code       int
	Path       string
	UserAgent  string
	Level      string
}

//	clientIP,
//
// t,
// level,
// method,
// path,
// statusCode,
// latency,
// userAgent
func SendLogRequest(request LogModel) {
	switch {
	case request.Code >= 500:
		request.Level = "error"
	case request.Code >= 400:
		request.Level = "warning"
	default:
		request.Level = "info"
	}

	log.Printf("[LOG] %s - [%v] level = %s \"%s %s %d %v \"%s\"\n",
		request.ClientIP,
		request.AccessTime,
		request.Level,
		request.Method,
		request.Path,
		request.Code,
		request.Latency,
		request.UserAgent,
	)
}
