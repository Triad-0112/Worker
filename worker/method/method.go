package wmethod
import (
	"github.com/Triad-0112/Worker/worker"
)

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
