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
	sem := make(chan struct{}, pool)
	res1 := make([]user, n)

	var i int64
	for ; i < n; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(j int64) {
			// curUser := getOne(j)
			// res = append(res, curUser)
			res1[j] = getOne(j)
			<-sem
			wg.Done()
		}(i)
	}
	wg.Wait()
	return res1
}
