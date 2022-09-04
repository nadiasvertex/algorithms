package execution

type Sender[T any] func() UniReceiver[T]

func (s Sender[Output]) OnStopped(handler func()) Sender[Output] {
	return func() UniReceiver[Output] {
		output := makeReceiver[Output]()
		go func() {
			in := s()
			select {
			case v := <-in.Value():
				output.SetValue(v)
			case e := <-in.Error():
				output.SetError(e)
			case <-in.Stopped():
				handler()
				output.SetStopped()
			}
		}()
		return output
	}
}

func (s Sender[Output]) OnError(handler func(error)) Sender[Output] {
	return func() UniReceiver[Output] {
		output := makeReceiver[Output]()
		go func() {
			in := s()
			select {
			case v := <-in.Value():
				output.SetValue(v)
			case e := <-in.Error():
				handler(e)
				output.SetError(e)
			case <-in.Stopped():
				output.SetStopped()
			}
		}()
		return output
	}
}
