package execution

func SyncWait[Output any](snd *Sender[Output]) (Output, error) {
	var v Output
	r := snd.Eval()
	select {
	case v = <-r.Value():
		return v, nil
	case e := <-r.Error():
		return v, e
	case <-r.Stopped():
		return v, nil
	}
}
