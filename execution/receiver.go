package execution

type Receiver[T any] interface {
	SetValue(value T)
	SetStopped()
	SetError(err error)

	Value() chan T
	Stopped() chan bool
	Error() chan error

	StopToken() *StopToken
}

type UniReceiver[T any] struct {
	value     chan T
	err       chan error
	stopToken *StopToken
}

type MultiReceiver[T any] struct {
	values    []chan T
	errs      []chan error
	stopToken *StopToken
}

func makeReceiver[Output any](stopToken *StopToken) UniReceiver[Output] {
	if stopToken == nil {
		stopToken = makeStopToken()
	}
	return UniReceiver[Output]{
		value:     make(chan Output, 1),
		err:       make(chan error, 1),
		stopToken: stopToken,
	}
}

func makeMultiReceiver[Output any](stopToken *StopToken) *MultiReceiver[Output] {
	return &MultiReceiver[Output]{stopToken: stopToken}
}

func (r *UniReceiver[T]) SetValue(value T) {
	r.value <- value
}

func (r *UniReceiver[T]) SetStopped() {
	r.stopToken.Stop()
}

func (r *UniReceiver[T]) SetError(err error) {
	r.err <- err
}

func (r *UniReceiver[T]) Value() chan T {
	return r.value
}

func (r *UniReceiver[T]) Stopped() chan bool {
	return r.stopToken.StopChannel()
}

func (r *UniReceiver[T]) Error() chan error {
	return r.err
}

func (r *UniReceiver[T]) StopToken() *StopToken {
	return r.stopToken
}

func (r *MultiReceiver[T]) SetValue(value T) {
	for _, c := range r.values {
		c <- value
	}
}

func (r *MultiReceiver[T]) SetStopped() {
	r.stopToken.Stop()
}

func (r *MultiReceiver[T]) SetError(err error) {
	for _, c := range r.errs {
		c <- err
	}
}

func (r *MultiReceiver[T]) StopToken() *StopToken {
	return r.stopToken
}

func (r *MultiReceiver[T]) AppendReceiver() Receiver[T] {
	i := len(r.values)

	r.values = append(r.values, make(chan T, 1))
	r.errs = append(r.errs, make(chan error, 1))

	return &UniReceiver[T]{
		value:     r.values[i],
		err:       r.errs[i],
		stopToken: r.stopToken,
	}
}
