package execution

type Sender[T any] struct {
	worker    func() Receiver[T]
	stopToken *StopToken
}

// Eval executes the sender
func (s *Sender[T]) Eval() Receiver[T] {
	return s.worker()
}

// Stop indicates that the sender chain should stop. This is guaranteed to
// terminate the chain, but any work in progress may continue to execute unless
// it checks the stop token.
func (s *Sender[T]) Stop() {
	s.stopToken.Stop()
}

// Stopped indicates if the chain was stopped.
func (s *Sender[T]) Stopped() bool {
	return s.stopToken.Stopped()
}

// NewChild creates a new sender as a child of the current sender. The new child
// shares the same stop token as the parent, but the two senders have no causal
// relationship. use algorithms like Then to create an execution relationship
// between two senders.
func (s *Sender[T]) NewChild(childWorker func() Receiver[T]) *Sender[T] {
	return &Sender[T]{worker: childWorker, stopToken: s.stopToken}
}

// OnStopped adds an event handler to the chain that is executed when the
// stop signal is detected at this point in the chain.
func (s *Sender[T]) OnStopped(handler func()) *Sender[T] {
	return s.NewChild(func() Receiver[T] {
		output := makeReceiver[T](s.stopToken)
		go eventHandler[T, T](s, &output, func(in T) T {
			return in
		}, nil, handler)
		return &output
	})
}

// OnError adds an event handler to the chain that is executed when an
// error is returned at this point in the chain.
func (s *Sender[T]) OnError(handler func(error)) *Sender[T] {
	return s.NewChild(func() Receiver[T] {
		output := makeReceiver[T](s.stopToken)
		go eventHandler[T, T](s, &output, func(in T) T {
			return in
		}, handler, nil)
		return &output
	})
}

// eventHandler encapsulates common code used to evaluate and propagate values, errors,
// and cancellation events along an execution chain.
func eventHandler[Input, Output any](input *Sender[Input], output Receiver[Output], onValue func(Input) Output,
	onError func(err error), onStopped func()) {
	in := input.Eval()
	select {
	case v := <-in.Value():
		output.SetValue(onValue(v))
	case e := <-in.Error():
		if onError != nil {
			onError(e)
		}
		output.SetError(e)
	case <-in.Stopped():
		if onStopped != nil {
			onStopped()
		}
		output.SetStopped()
	}
}

// eventHandlerWithCancel encapsulates common code used to evaluate and propagate values, errors,
// and cancellation events along an execution chain. This variation supports workers that
// want access to the chain's stop token.
func eventHandlerWithCancel[Input, Output any](input *Sender[Input], output Receiver[Output],
	onValue func(*StopToken, Input) Output,
	onError func(err error), onStopped func()) {
	in := input.Eval()
	select {
	case v := <-in.Value():
		output.SetValue(onValue(input.stopToken, v))
	case e := <-in.Error():
		if onError != nil {
			onError(e)
		}
		output.SetError(e)
	case <-in.Stopped():
		if onStopped != nil {
			onStopped()
		}
		output.SetStopped()
	}
}
