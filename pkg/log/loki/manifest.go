package loki

import "github.com/sirupsen/logrus"

type Client interface {
	Logf(format string, e *logrus.Entry)
	LogfWithLabels(labels map[string]string, format string)

	Ping() (*PongResponse, error)
	Close()
}

type PongResponse struct {
	IsReady bool
}
