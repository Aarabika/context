package context

import (
	"context"
	"github.com/sirupsen/logrus"
)

type requestId struct{}

type log struct{}

// RequestId takes request id from context.
func RequestId(c context.Context) string {
	return c.Value(requestId{}).(string)
}

// WithRequestId adds request id to context.
func SetRequestId(c context.Context, id string) context.Context {
	return context.WithValue(c, requestId{}, id)
}

func Log(c context.Context) *logrus.Entry  {
	return c.Value(log{}).(*logrus.Entry)
}

func SetLog(c context.Context, l *logrus.Entry) context.Context  {
	return context.WithValue(c, log{}, l)
}