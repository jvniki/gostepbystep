package main
import (
	"fmt"
	"runtime"
	"sync"
)
func main() {
	//i: 10, j: 0-9
	runtime.GOMAXPROCS(1)
	// runtime.GOMAXPROCS(2)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("j: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
