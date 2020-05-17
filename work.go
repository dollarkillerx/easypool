package easypool

type work struct {
	taskChan chan interface{}
	fn       func(interface{})
}

func newWork(taskChan chan interface{}, fn func(interface{})) *work {
	return &work{
		taskChan: taskChan,
		fn:       fn,
	}
}

func (w *work) Run() {
loop:
	for {
		select {
		case task, ok := <-w.taskChan:
			if !ok {
				break loop
			}
			w.fn(task)
		}
	}
}
