package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Println("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

}
