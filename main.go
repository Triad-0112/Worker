package main

import (
	"flag"
	"fmt"
	"strconv"
	"sync"
	"time"
	
	"github.com/Triad-0112/Worker/color"
	"github.com/Triad-0112/Worker/worker"
)

//WORKER POOL TEST
type WorkTool interface {
}
type Pool struct {
	Jobs       []*Jobs
	TWorker    int
	JobChannel chan *Jobs
	wg         sync.WaitGroup
}
type Jobs struct {
	year int
	dir  string
}

// FACTORY PATTERN
func NewPool(jobs []*Jobs, tworker int) *Pool {
	return &Pool{
		Jobs:       jobs,
		TWorker:    tworker,
		JobChannel: make(chan *Jobs),
	}
}
func (p *Pool) Run() {
	for i := 0; i < p.TWorker; i++ {
		go p.Work(i)
	}
	p.wg.Add(len(p.Jobs))
	for _, jobs := range p.Jobs {
		p.JobChannel <- jobs
	}
	close(p.JobChannel)
	p.wg.Wait()
}
func (p *Pool) Work(id int) {
	defer fmt.Printf("%s\n\n", colortext.Notificationcolor("Worker %d DEPLOYED to Work Site", id+1))
	//notificationcolor("Worker %d Rest...", id+1)
	fmt.Printf("%s\n\n", colortext.Notificationcolor("Worker %d DEPLOYED to Work Site", id+1))
	//fmt.Printf("%s\n\n", notificationcolor("Worker %d Deployed to Working Site\n\n", id+1))
	for jobs := range p.JobChannel {
		jobs.Run(&p.wg, id)
	}
}
func (j *Jobs) Run(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	defer fmt.Printf("%s %s\n\n", colortext.Workercolor("[Worker %d] :", id+1), colortext.Textcolor("Finished working on data-%s", colortext.Filenamecolor("%d.csv", j.year)))
	fmt.Printf("%s %s\n\n", colortext.Workercolor("[Worker %d] :", id+1), colortext.Textcolor("Starting to work on data-%s", colortext.Filenamecolor("%d.csv", j.year)))
	worker.CreateFile(&j.dir, strconv.Itoa(j.year)+".csv", worker.Fetcher(j.year, id), id) //THIS
}
func NewJobs(year int, dir string) *Jobs {
	return &Jobs{
		year: year,
		dir:  dir,
	}
}

//WORKER WORK

func main() {
	totalworker := flag.Int("concurrent_limit", 2, "Input total worker")
	dir := flag.String("output", "D:/Test/", "Destination Output file")
	from := flag.Int("from", 1993, "Range start to Fetch and Create its file")
	until := flag.Int("until", 2014, "Range end to Fetch and Create its file")
	flag.Parse()
	jobs := []*Jobs{}
	for i := *from; i <= *until; i++ {
		jobs = append(jobs, NewJobs(i, *dir))
	}
	p := NewPool(jobs, *totalworker)
	p.Run()
	fmt.Printf("%s", colortext.Timecolor("%s", time.Since(colortext.Now)))
}

// EFFECTIVELY TIME : 1s to work
