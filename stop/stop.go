package stop

// Done is the variable sent on the Chan to indicate
// that something has stopped.
var Done = struct{}{}

// Chan is a stop channel.
type Chan chan struct{}

// Stopper represents types that implement
// the stop channel pattern.
type Stopper interface {
	Stop() Chan
}

// MakeChan makes a new Chan used to indicate when
// stopping has finished.
func MakeChan() Chan {
	return make(chan struct{})
}

// All stops all Stopper types and returns another Chan
// on which Done will be sent once all things have
// finished stopping.
func All(stoppers ...Stopper) Chan {
	all := MakeChan()
	go func() {
		var allChans []Chan
		for _, stopper := range stoppers {
			allChans = append(allChans, stopper.Stop())
		}
		for _, ch := range allChans {
			<-ch
		}
		all <- Done
	}()
	return all
}