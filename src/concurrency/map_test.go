package concurrency

import (
	"strconv"
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	count := 1_000
	wg := sync.WaitGroup{}
	m := NewMap[string, int]()

	wg.Add(count)
	for i := 0; i < count; i++ {
		val := i
		go func() {
			m.Write(strconv.Itoa(val), val)
			wg.Done()
		}()
	}
	wg.Wait()

	wg.Add(count)
	res := make(chan int, count)
	for i := 0; i < count; i++ {
		val := i
		go func() {
			res <- m.Read(strconv.Itoa(val))
			wg.Done()
		}()
	}
	wg.Wait()
	close(res)

	if len(res) != count {
		t.Fatalf("expected 10 elements, got %d", len(res))
	}
}
