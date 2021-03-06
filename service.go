package service

import (
	"context"
	"fmt"
	"os"

	"github.com/google/uuid"
	pbmicro "github.com/unistack-org/micro-logger-service/v3/micro"
	pb "github.com/unistack-org/micro-logger-service/v3/proto"
	"github.com/unistack-org/micro/v3/client"
	"github.com/unistack-org/micro/v3/logger"
	"github.com/unistack-org/micro/v3/store"
)

type serviceLogger struct {
	opts    logger.Options
	service string
	client  pbmicro.LoggerClient
	store   store.Store
	fields  map[string]interface{}
}

func (l *serviceLogger) Init(opts ...logger.Option) error {
	for _, o := range opts {
		o(&l.opts)
	}

	var cli client.Client
	if l.opts.Context != nil {
		if v, ok := l.opts.Context.Value(clientKey{}).(client.Client); ok && v != nil {
			cli = v
		}
		if v, ok := l.opts.Context.Value(serviceKey{}).(string); ok && v != "" {
			l.service = v
		}
	}

	if l.service == "" {
		return fmt.Errorf("missing Service option")
	}

	if cli == nil {
		return fmt.Errorf("missing Client option")
	}

	l.client = pbmicro.NewLoggerClient(l.service, cli)

	return nil
}

func (l *serviceLogger) Fields(fields map[string]interface{}) logger.Logger {
	nl := &serviceLogger{opts: l.opts, client: l.client, store: l.store}
	nl.fields = make(map[string]interface{}, len(fields))
	for k, v := range fields {
		nl.fields[k] = v
	}
	return nl
}

func (l *serviceLogger) V(level logger.Level) bool {
	return l.opts.Level >= level
}

func (l *serviceLogger) Info(ctx context.Context, args ...interface{}) {
	l.Log(ctx, logger.InfoLevel, args...)
}

func (l *serviceLogger) Error(ctx context.Context, args ...interface{}) {
	l.Log(ctx, logger.ErrorLevel, args...)
}

func (l *serviceLogger) Warn(ctx context.Context, args ...interface{}) {
	l.Log(ctx, logger.WarnLevel, args...)
}

func (l *serviceLogger) Debug(ctx context.Context, args ...interface{}) {
	l.Log(ctx, logger.DebugLevel, args...)
}

func (l *serviceLogger) Trace(ctx context.Context, args ...interface{}) {
	l.Log(ctx, logger.TraceLevel, args...)
}

func (l *serviceLogger) Fatal(ctx context.Context, args ...interface{}) {
	l.Log(ctx, logger.FatalLevel, args...)
	os.Exit(1)
}

func (l *serviceLogger) Infof(ctx context.Context, msg string, args ...interface{}) {
	l.Logf(ctx, logger.InfoLevel, msg, args...)
}

func (l *serviceLogger) Errorf(ctx context.Context, msg string, args ...interface{}) {
	l.Logf(ctx, logger.ErrorLevel, msg, args...)
}

func (l *serviceLogger) Warnf(ctx context.Context, msg string, args ...interface{}) {
	l.Logf(ctx, logger.WarnLevel, msg, args...)
}

func (l *serviceLogger) Debugf(ctx context.Context, msg string, args ...interface{}) {
	l.Logf(ctx, logger.DebugLevel, msg, args...)
}

func (l *serviceLogger) Tracef(ctx context.Context, msg string, args ...interface{}) {
	l.Logf(ctx, logger.TraceLevel, msg, args...)
}

func (l *serviceLogger) Fatalf(ctx context.Context, msg string, args ...interface{}) {
	l.Logf(ctx, logger.FatalLevel, msg, args...)
	os.Exit(1)
}

func (l *serviceLogger) Log(ctx context.Context, level logger.Level, args ...interface{}) {
	msg := l.newMessage(level, "", args...)
	_, err := l.client.Log(ctx, msg)
	if err != nil {
		l.storeMessage(ctx, msg)
	}
}

func (l *serviceLogger) Logf(ctx context.Context, level logger.Level, format string, args ...interface{}) {
	msg := l.newMessage(level, "", args...)
	_, err := l.client.Log(ctx, msg)
	if err != nil {
		l.storeMessage(ctx, msg)
	}
}

func (l *serviceLogger) String() string {
	return "service"
}

func (l *serviceLogger) Options() logger.Options {
	return l.opts
}

// NewLogger builds a new logger based on options
func NewLogger(opts ...logger.Option) logger.Logger {
	options := logger.NewOptions(opts...)
	l := &serviceLogger{opts: options}
	return l
}

func (l *serviceLogger) newMessage(level logger.Level, format string, args ...interface{}) *pb.Message {
	msg := &pb.Message{Level: int32(level), Format: format}
	return msg
}

func (l *serviceLogger) storeMessage(ctx context.Context, msg *pb.Message) error {
	if l.store == nil {
		return nil
	}
	uid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	return l.store.Write(ctx, uid.String(), msg)
}
