package snowflake

import (
	"fmt"
	"testing"
)

func TestSnowFlake(t *testing.T) {
	worker, err := New(1)
	if err != nil {
		return
	}

	ch := make(chan int64)
	defer close(ch)
	count := 10000
	for i := 0; i < count; i++ {
		go func() {
			id := worker.GetID()
			ch <- id
		}()
	}

	m := make(map[int64]int)
	for i := 0; i < count; i++ {
		id := <-ch
		fmt.Printf("%d\n", id)
		_, ok := m[id]
		if ok {
			t.Error("ID is not unique!\n")
			return
		}
		m[id] = i
	}
	fmt.Println("All", count, "snowflake ID Get successed!")
}
