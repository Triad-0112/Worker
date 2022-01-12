package wmethod

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/Triad-0112/Worker/color"
	"github.com/Triad-0112/Worker/worker"
)

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
func NewPool(jobs []*Jobs, tworker int) *Pool {
	return &Pool{
		Jobs:       jobs,
		TWorker:    tworker,
		JobChannel: make(chan *Jobs),
	}
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
