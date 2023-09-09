package cronJob

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"time"
)

// 2
func hello(name string) {
	message := fmt.Sprintf("%v", name)
	fmt.Println(message)
}

func RunCronJobs() {
	count := 0
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Minute().Do(func() {
		count++
		hello(fmt.Sprintf("Call %d", count))
	})
	s.StartBlocking()
}
