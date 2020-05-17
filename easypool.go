package easypool

type EasyPoolEntity struct {
	works    []*work
	taskChan chan interface{}
}

func New(size int, pf func(interface{})) *EasyPoolEntity {
	pool := &EasyPoolEntity{}
	pool.taskChan = make(chan interface{}, 0)
	pool.works = make([]*work, size)
	for i := 0; i < size; i++ {
		wo := newWork(pool.taskChan, pf)
		go wo.Run()
		pool.works = append(pool.works, wo)
	}

	return pool
}

func (e *EasyPoolEntity) Close() {
	close(e.taskChan)
	e.works = make([]*work, 0)
}

func (e *EasyPoolEntity) Invoke(params interface{}) {
	e.taskChan <- params
}
