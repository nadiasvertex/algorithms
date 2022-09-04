package execution

import "sync"

// Split creates a sender that reads from "input", and sends it to worker1 and worker2. The input
// is evaluated only once, and both chains share the same stop token. That means a cancel anywhere
// along the route will cause _both_ chains to stop.
func Split[Input, Output any](input Sender[Input], worker1, worker2 func(Input) Output) (*Sender[Output],
	*Sender[Output]) {

	mr := makeMultiReceiver[Input](input.stopToken)
	var once sync.Once
	inputWorker := func() {
		in := input.Eval()
		select {
		case v := <-in.Value():
			mr.SetValue(v)
		case e := <-in.Error():
			mr.SetError(e)
		case <-in.Stopped():
			mr.SetStopped()
		}
	}

	return &Sender[Output]{worker: func() Receiver[Output] {
			output := makeReceiver[Output](mr.stopToken)
			in := mr.AppendReceiver()
			go func() {
				once.Do(inputWorker)
				select {
				case v := <-in.Value():
					output.SetValue(worker1(v))
				case e := <-in.Error():
					output.SetError(e)
				case <-in.Stopped():
					output.SetStopped()
				}
			}()
			return &output
		}, stopToken: mr.StopToken()},
		&Sender[Output]{worker: func() Receiver[Output] {
			output := makeReceiver[Output](mr.stopToken)
			in := mr.AppendReceiver()
			go func() {
				once.Do(inputWorker)
				select {
				case v := <-in.Value():
					output.SetValue(worker2(v))
				case e := <-in.Error():
					output.SetError(e)
				case <-in.Stopped():
					output.SetStopped()
				}
			}()
			return &output
		}, stopToken: mr.stopToken}
}
