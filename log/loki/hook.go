package loki

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type PromtailHook struct {
	logrus.Hook
	client Client
}

//
// Initializes a Promtail hook for Logrus logger.
//	- lokiAddress - address of Grafana Loki server to push logs to (e.g. loki:3100)
//	- labels - is kinda tags for grepping in Loki's and Grafana's queries
func NewPromtailHook(lokiURL string, labels map[string]string) (*PromtailHook, error) {
	var (
		hook = &PromtailHook{}
		err  error
	)

	hook.client, err = NewJSONv1Client(lokiURL, labels)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize Promtail client: %s", err.Error())
	}

	return hook, err
}

func (rcv *PromtailHook) Fire(entry *logrus.Entry) error {
	if entry == nil {
		return fmt.Errorf("log entry is nil")
	}

	//line, err := entry.String()
	//if err != nil {
	//	return fmt.Errorf("unable to read log entry: %s", err)
	//}
	rcv.client.Logf(time.Now().Format("2006-01-02 15:04:05")+" "+entry.Message,entry)

	return nil
}

func (rcv *PromtailHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (rcv *PromtailHook) LokiHealthCheck() error {
	_, err := rcv.client.Ping()
	if err != nil {
		return err
	}

	return nil
}

// Compile time validation
var _ logrus.Hook = (*PromtailHook)(nil)
