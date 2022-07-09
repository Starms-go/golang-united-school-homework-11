package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {

	var wg sync.WaitGroup
	// var mx sync.Mutex

	sem := make(chan struct{}, pool)
	var i int64
	for ; i < n; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(j int64) {
			// mx.Lock()
			curUser := getOne(j)
			res = append(res, curUser)
			// mx.Unlock()
			<-sem
			wg.Done()
		}(i)
	}
	wg.Wait()
	return
}
