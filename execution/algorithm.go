package execution

// Just creates a sender that sends the given values as its output.
func Just[Output any](value Output) *Sender[Output] {
	return &Sender[Output]{worker: func() Receiver[Output] {
		output := makeReceiver[Output](nil)
		go func() {
			output.SetValue(value)
		}()
		return &output
	}, stopToken: makeStopToken()}
}

// Then creates a sender that reads from "input", processes it with "worker" and
// then writes the result as output. The worker executes asynchronously
func Then[Input, Output any](input *Sender[Input], worker func(Input) Output) *Sender[Output] {
	return &Sender[Output]{worker: func() Receiver[Output] {
		output := makeReceiver[Output](input.stopToken)
		go eventHandler[Input, Output](input, &output, worker, nil, nil)
		return &output
	}, stopToken: input.stopToken}
}

// ThenWithCancel creates a sender that reads from "input", processes it with "worker" and
// then writes the result as output. The worker executes asynchronously.
func ThenWithCancel[Input, Output any](input *Sender[Input], worker func(*StopToken, Input) Output) *Sender[Output] {
	return &Sender[Output]{worker: func() Receiver[Output] {
		output := makeReceiver[Output](input.stopToken)
		go eventHandlerWithCancel[Input, Output](input, &output, worker, nil, nil)
		return &output
	}, stopToken: input.stopToken}
}
