package main

import (
	"fmt"
	"github.com/bamohmohamed/go-cncr-ex/model"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]model.Manga{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {

	wg := &sync.WaitGroup{}
	mu := &sync.RWMutex{}
	chC := make(chan model.Manga)
	chDb := make(chan model.Manga)

	for i := 0; i < 9; i++ {
		id := rnd.Intn(9) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, mu *sync.RWMutex, ch chan<- model.Manga) {
			if m, ok := queryCache(id, mu); ok {
				//fmt.Println("from cache", m)
				chC <- m
			}
			wg.Done()
		}(id, wg, mu, chC)
		go func(id int, wg *sync.WaitGroup, mu *sync.RWMutex, ch chan<- model.Manga) {
			if m, ok := queryDb(id); ok {
				//fmt.Println("from DB", m)
				mu.Lock()
				cache[id] = m
				mu.Unlock()
				chDb <- m
			}
			wg.Done()
		}(id, wg, mu, chDb)
		go func(ch1, ch2 <-chan model.Manga) {
			select {
			case m := <-ch1:
				fmt.Println("from Cache", m)
				<-ch2
			case m := <-ch2:
				fmt.Println("from DB", m)
			}
		}(chC, chDb)

		time.Sleep(250 * time.Millisecond)
	}

	wg.Wait()

}

func queryCache(id int, mu *sync.RWMutex) (model.Manga, bool) {
	mu.RLock()
	m, ok := cache[id]
	mu.RUnlock()
	return m, ok
}

func queryDb(id int) (model.Manga, bool) {
	time.Sleep(200 * time.Millisecond)
	for _, m := range model.Mangas {
		if m.Id == id {
			return m, true
		}
	}
	return model.Manga{}, false
}
