package wstructint

import "sync"

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
