package scanner

import (
	"fmt"
	"sync"
)

func workerPool(wg *sync.WaitGroup, jobs <-chan string) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println("Processing job:", job)
	}
}
