package base

import (
	"fmt"
	"sync"
	"time"
)

type (
	Trace struct {
		timestamp time.Time
		level     string
		data      interface{}
		tags      []string
	}

	traces struct {
		sync.Mutex
		lastTrace Trace
		allTraces []Trace
	}
)

type (
	Tracer struct {
		traces  *traces
		tracing bool
	}
)

const (
	debugTrace = "debug"
	infoTrace  = "info"
	errorTrace = "error"
)

func newTraces() *traces {
	return &traces{
		allTraces: []Trace{},
	}
}

// Timestamp returns trace timestamp
func (t Trace) Timestamp() time.Time {
	return t.timestamp
}

// Level returns trace level
func (t Trace) Level() string {
	return t.level
}

// Tags return trace type
func (t Trace) Tags() []string {
	return t.tags
}

// String return a string representation of the trace
func (t Trace) String() string {
	return fmt.Sprintf("[%s] %s", t.level, t.data)
}

// Data returns trace data
func (t Trace) Data() interface{} {
	return t.data
}

// Tracing queue
func (traces *traces) last() Trace {
	return traces.lastTrace
}

func (traces *traces) push(trace Trace) {
	traces.Lock()
	defer traces.Unlock()

	traces.lastTrace = trace
	traces.allTraces = append(traces.allTraces, trace)
}

func (traces *traces) pull() (trace Trace) {
	traces.Lock()
	defer traces.Unlock()

	all := traces.allTraces
	l := len(all) - 1

	trace, traces.allTraces = all[l], all[:l]

	return trace
}

func (traces *traces) all() []Trace {
	return traces.allTraces
}

// Tracer

// EnableTracing enables tracing mode
func (t *Tracer) EnableTracing() {
	t.traces = newTraces()
	t.tracing = true
}

// DisableTracing disables tracing mode
func (t *Tracer) DisableTracing() {
	t.traces = newTraces()
	t.tracing = false
}

func (t *Tracer) SendDebug(data interface{}, tags ...string) {
	t.SendTrace(debugTrace, data, tags...)
}

func (t *Tracer) SendInfo(data interface{}, tags ...string) {
	t.SendTrace(infoTrace, data, tags...)
}

func (t *Tracer) SendError(data interface{}, tags ...string) {
	t.SendTrace(errorTrace, data, tags...)
}

func (t *Tracer) SendTrace(level string, data interface{}, tags ...string) {
	if !t.tracing {
		return
	}

	// TODO: Make concurrent
	t.SaveTrace(
		Trace{
			timestamp: time.Now(),
			level:     level,
			data:      data,
			tags:      tags,
		})
}

func (t *Tracer) IsTracingEnabled() bool {
	return t.tracing
}

func (t *Tracer) SaveTrace(trace Trace) {
	if !t.tracing {
		return
	}

	t.ensureTraces()
	t.traces.push(trace)
}

func (t *Tracer) LastEntry() Trace {
	t.ensureTraces()
	return t.traces.last()
}

func (t *Tracer) LastEntries() []Trace {
	t.ensureTraces()
	return t.traces.all()
}

func (t *Tracer) ensureTraces() {
	t.traces.Lock()
	defer t.traces.Unlock()

	if t.traces == nil {
		t.traces = newTraces()
	}
}
