package execution

import "sync"

// Just creates a sender that sends the given values as its output.
func Just[Output any](value Output) Sender[Output] {
	return func() UniReceiver[Output] {
		output := makeReceiver[Output]()
		go func() {
			output.SetValue(value)
		}()
		return output
	}
}

// Then creates a sender that reads from "input", processes it with "worker" and
// then writes the result as output. The worker executes asynchronously
func Then[Input, Output any](input Sender[Input], worker func(Input) Output) Sender[Output] {
	return func() UniReceiver[Output] {
		output := makeReceiver[Output]()
		go func() {
			in := input()
			select {
			case v := <-in.Value():
				output.SetValue(worker(v))
			case e := <-in.Error():
				output.SetError(e)
			case <-in.Stopped():
				output.SetStopped()
			}
		}()
		return output
	}
}

// Split creates a sender that reads from "input", and sends it to worker1 and worker2.
func Split[Input, Output any](input Sender[Input], worker1, worker2 func(Input) Output) (Sender[Output],
	Sender[Output]) {

	mr := makeMultiReceiver[Input]()
	var once sync.Once
	inputWorker := func() {
		in := input()
		select {
		case v := <-in.Value():
			mr.SetValue(v)
		case e := <-in.Error():
			mr.SetError(e)
		case <-in.Stopped():
			mr.SetStopped()
		}
	}

	return func() UniReceiver[Output] {
			output := makeReceiver[Output]()
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
			return output
		}, func() UniReceiver[Output] {
			output := makeReceiver[Output]()
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
			return output
		}
}
