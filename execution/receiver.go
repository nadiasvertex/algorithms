package execution

type Receiver[T any] interface {
	SetValue(value T)
	SetStopped()
	SetError(err error)

	Value() chan T
	Stopped() chan bool
	Error() chan error
}

type UniReceiver[T any] struct {
	value   chan T
	stopped chan bool
	err     chan error
}

type MultiReceiver[T any] struct {
	values  []chan T
	stopped []chan bool
	errs    []chan error
}

func makeReceiver[Output any]() UniReceiver[Output] {
	return UniReceiver[Output]{
		value:   make(chan Output, 1),
		stopped: make(chan bool, 1),
		err:     make(chan error, 1),
	}
}

func makeMultiReceiver[Output any]() *MultiReceiver[Output] {
	return &MultiReceiver[Output]{}
}

func (r *UniReceiver[T]) SetValue(value T) {
	r.value <- value
}

func (r *UniReceiver[T]) SetStopped() {
	close(r.stopped)
}

func (r *UniReceiver[T]) SetError(err error) {
	r.err <- err
}

func (r *UniReceiver[T]) Value() chan T {
	return r.value
}

func (r *UniReceiver[T]) Stopped() chan bool {
	return r.stopped
}

func (r *UniReceiver[T]) Error() chan error {
	return r.err
}

func (r *MultiReceiver[T]) SetValue(value T) {
	for _, c := range r.values {
		c <- value
	}
}

func (r *MultiReceiver[T]) SetStopped() {
	for _, c := range r.stopped {
		close(c)
	}
}

func (r *MultiReceiver[T]) SetError(err error) {
	for _, c := range r.errs {
		c <- err
	}
}

func (r *MultiReceiver[T]) AppendReceiver() Receiver[T] {
	i := len(r.values)

	r.values = append(r.values, make(chan T, 1))
	r.stopped = append(r.stopped, make(chan bool, 1))
	r.errs = append(r.errs, make(chan error, 1))

	return &UniReceiver[T]{
		value:   r.values[i],
		stopped: r.stopped[i],
		err:     r.errs[i],
	}
}
