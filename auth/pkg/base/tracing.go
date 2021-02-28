package base

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type (
	// Trace description
	Trace struct {
		timestamp time.Time
		level     string
		data      interface{}
	}
)

type (
	// Tracer description
	Tracer struct {
		mutex  sync.Mutex
		level  string
		traces chan Trace
		quit   chan struct{}
	}
)

const (
	noneTrace  = "none"
	debugTrace = "debug"
	infoTrace  = "info"
	errorTrace = "error"
)

const (
	// NOTE: Make this values confifigurable?
	size         = 1000
	confortIndex = 0.9
	purgeEvery   = 100 // milliseconds
)

// NewTracer build and returns a new Tracer instance
func NewTracer(level string) *Tracer {
	t := &Tracer{
		level:  level,
		traces: make(chan Trace, size),
		quit:   make(chan struct{}),
	}

	go t.startHealer()

	return t
}

// Timestamp returns trace timestamp
func (t Trace) Timestamp() time.Time {
	return t.timestamp
}

// Level returns trace level
func (t Trace) Level() string {
	return t.level
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
func (t *Tracer) push(trace Trace) {
	t.traces <- trace
}

func (t *Tracer) pull() (trace Trace) {
	return <-t.traces
}

func (t *Tracer) currentTraces() []Trace {
	traces := []Trace{}

	for i := 0; i < len(t.traces); i++ {
		traces = append(traces, <-t.traces)
	}

	return traces
}

// Tracer

// EnableTracing enables tracing mode
func (t *Tracer) EnableTracing() {
	t.mutex.Lock()
	t.level = infoTrace
	t.traces = make(chan Trace, size)
	t.quit = make(chan struct{})
	t.startHealer()
	t.mutex.Unlock()
}

// DisableTracing disables tracing mode
func (t *Tracer) DisableTracing() {
	t.mutex.Lock()
	t.quit <- struct{}{}
	close(t.quit)
	t.level = noneTrace
	t.traces = make(chan Trace, size)
	t.mutex.Unlock()
}

func (t *Tracer) SendDebug(data interface{}) {
	t.SendTrace(debugTrace, data)
}

func (t *Tracer) SendDebugf(format string, data ...interface{}) {
	t.SendDebug(fmt.Sprintf(format, data...))
}

func (t *Tracer) SendInfo(data interface{}) {
	t.SendTrace(infoTrace, data)
}

func (t *Tracer) SendInfof(format string, data ...interface{}) {
	t.SendInfo(fmt.Sprintf(format, data...))
}

func (t *Tracer) SendError(data interface{}) {
	t.SendTrace(errorTrace, data)
}

func (t *Tracer) SendErrorf(format string, data ...interface{}) {
	t.SendError(fmt.Sprintf(format, data...))
}

func (t *Tracer) SendTrace(level string, data interface{}, tags ...string) {
	if !t.IsTracingEnabled() {
		return
	}

	// TODO: Make concurrent
	t.SaveTrace(
		Trace{
			timestamp: time.Now(),
			level:     level,
			data:      data,
		})
}

func (t *Tracer) IsTracingEnabled() bool {
	return t.level != noneTrace
}

func (t *Tracer) SaveTrace(trace Trace) {
	if !t.IsTracingEnabled() {
		return
	}

	select {
	case t.traces <- trace:
		// Sent is enough
	default:
		// Nothing to do
	}
}

func (t *Tracer) PrintTracerStack() {
	for _, t := range t.currentTraces() {
		log.Printf("%+v\n", t)
	}
}

func (t *Tracer) startHealer() {
	max := roundDown(float64(size) * float64(confortIndex))
	qtyToDequeue := size - max

	ticker := time.NewTicker(purgeEvery * time.Millisecond)

	for {
		select {
		case <-ticker.C:
			if len(t.traces) >= max {
				for i := 0; i == qtyToDequeue; i++ {

					select {
					case <-t.traces:
						// Receive is enough
					default:
						// Nothing to do
					}

				}
			}
		case <-t.quit:
			return
		}
	}
}

func roundDown(val float64) int {
	if val < 0 {
		return int(val - 1.0)
	}
	return int(val)
}
