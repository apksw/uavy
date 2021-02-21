package base

import (
	"sync"
	"time"
)

type (
	Trace struct {
		timestamp time.Time
		level     string
		traceType string
		data      interface{}
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

func newTraces() *traces {
	return &traces{
		allTraces: []Trace{},
	}
}

// Trace interface

// Timestamp returns trace timestamp
func (t Trace) Timestamp() time.Time {
	return t.timestamp
}

// Level returns trace level
func (t Trace) Level() string {
	return t.level
}

// Type return trace type
func (t Trace) Type() string {
	return t.traceType
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
	t.tracing = true
}

// DisableTracing disables tracing mode
func (t *Tracer) DisableTracing() {
	t.traces = newTraces()
	t.tracing = false
}

func (t *Tracer) saveTracingData(data, level, tracingType string) {
	if t.tracing {
		t.SaveTrace(
			Trace{
				timestamp: time.Now(),
				level:     level,
				traceType: tracingType,
				data:      data,
			})
	}
}

func (t *Tracer) IsTracingEnabled() bool {
	return t.tracing
}

func (t *Tracer) SaveTrace(trace Trace) {
	t.traces.push(trace)
}

func (t *Tracer) LastEntry() Trace {
	return t.traces.last()
}

func (t *Tracer) LastEntries() []Trace {
	return t.traces.all()
}
