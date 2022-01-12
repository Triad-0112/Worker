package main

import (
	"flag"
	"fmt"
	"time"
	
	"github.com/Triad-0112/Worker/worker/method"
)
//WORKER WORK

func main() {
	totalworker := flag.Int("concurrent_limit", 2, "Input total worker")
	dir := flag.String("output", "D:/Test/", "Destination Output file")
	from := flag.Int("from", 1993, "Range start to Fetch and Create its file")
	until := flag.Int("until", 2014, "Range end to Fetch and Create its file")
	flag.Parse()
	jobs := []*Jobs{}
	for i := *from; i <= *until; i++ {
		jobs = append(jobs, wmethod.NewJobs(i, *dir))
	}
	p := wmethod.NewPool(jobs, *totalworker)
	p.Run()
	fmt.Printf("%s", colortext.Timecolor("%s", time.Since(colortext.Now)))
}

// EFFECTIVELY TIME : 1s to work
