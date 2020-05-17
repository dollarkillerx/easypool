# easypool
easypool

### install
``` 
go get github.com/dollarkillerx/easypool
```

### easy
```go
func TestPool(t *testing.T) {
	task := &testItem{
		name: "Hello Rust",
		Resp: make(chan string,0),
	}
	entity := New(1000, func(i interface{}) {
		resp := i.(*testItem)
		resp.Resp <- resp.name + " Go!!!"
	})
	defer entity.Close()
	entity.Invoke(task)

	resp := <-task.Resp
	fmt.Println(resp)
}

type testItem struct {
	name string
	Resp chan string
}
```