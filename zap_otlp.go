package zap_otlp

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SpanCtx(ctx context.Context) zap.Field {
	span := trace.SpanFromContext(ctx)
	sctx := span.SpanContext()

	if !sctx.IsValid() {
		return zap.Skip()
	}

	return zap.Inline(zapSpan{
		traceID: sctx.TraceID().String(),
		spanID:  sctx.SpanID().String(),
	})
}

type zapSpan struct {
	traceID string
	spanID  string
}

func (z zapSpan) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("trace_id", z.traceID)
	enc.AddString("span_id", z.spanID)
	return nil
}
