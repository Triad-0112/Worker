package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/fatih/color"
)

const baseurl = "https://data.gov.sg/api/action/datastore_search?resource_id=eb8b932c-503c-41e7-b513-114cffbe2338&q="

//Data TYPE DONT CHANGE
type Graduate struct {
	Success bool   `json:"success"`
	Result  Result `json:"result"`
}

type Result struct {
	Resource_id string    `json:"resource_id"`
	Fields      []Fields  `json:"fields"`
	Records     []Records `json:"records"`
}

type Fields struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

type Records struct {
	Ide    int    `json:"_id"`
	Sex    string `json:"sex"`
	No     string `json:"no_of_graduates"`
	Course string `json:"type_of_course"`
	Year   string `json:"year"`
}

var textcolor = color.New(color.FgHiWhite, color.Bold).SprintfFunc()
var workercolor = color.New(color.FgHiCyan, color.Bold).SprintfFunc()
var directorycolor = color.New(color.FgHiYellow, color.Bold, color.Italic).SprintfFunc()
var filenamecolor = color.New(color.FgHiGreen, color.BlinkRapid, color.Bold).SprintfFunc()
var timecolor = color.New(color.FgHiMagenta, color.Bold).SprintfFunc()
var notificationcolor = color.New(color.FgHiRed, color.Bold).SprintfFunc()
var now = time.Now()

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
	defer fmt.Printf("%s\n\n", notificationcolor("Worker %d Rest...", id+1))
	fmt.Printf("%s\n\n", notificationcolor("Worker %d DEPLOYED to Work Site", id+1))
	//fmt.Printf("%s\n\n", notificationcolor("Worker %d Deployed to Working Site\n\n", id+1))
	for jobs := range p.JobChannel {
		jobs.Run(&p.wg, id)
	}
}
func (j *Jobs) Run(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	defer fmt.Printf("%s %s\n\n", workercolor("[Worker %d] :", id+1), textcolor("Finished working on data-%s", filenamecolor("%d.csv", j.year)))
	fmt.Printf("%s %s\n\n", workercolor("[Worker %d] :", id+1), textcolor("Starting to work on data-%s", filenamecolor("%d.csv", j.year)))
	CreateFile(&j.dir, strconv.Itoa(j.year)+".csv", fetcher(j.year, id), id)
}
func NewJobs(year int, dir string) *Jobs {
	return &Jobs{
		year: year,
		dir:  dir,
	}
}

//WORKER WORK
func fetcher(year int, id int) [][]string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	defer fmt.Printf("%s %s %s\n\n", workercolor("[Worker %d] :", id+1), textcolor("Finished collecting data of %s", filenamecolor("%d", year)), textcolor("from API"))
	fmt.Printf("%s %s", workercolor("[Worker %d] :", id+1), textcolor("Starting to fetch data of %s\n\n", filenamecolor("%d.csv", year)))
	url := baseurl + strconv.Itoa(year)
	m := make(map[string][][]string)
	spaceClient := http.Client{
		Timeout: time.Second * 2,
	}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	response, getErr := spaceClient.Do(request)
	if getErr != nil {
		log.Fatal(getErr)
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	record := Graduate{}
	jsonErr := json.Unmarshal(body, &record)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	convert := strconv.Itoa(year)
	for i := range record.Result.Records {
		temp := []string{
			strconv.Itoa(record.Result.Records[i].Ide),
			record.Result.Records[i].Sex,
			record.Result.Records[i].Course,
			record.Result.Records[i].Year,
		}
		m[convert] = append(m[convert], temp)
	}
	return m[convert]
}
func CreateFile(dir *string, filename string, a [][]string, id int) {
	defer fmt.Printf("%s %s", workercolor("[Worker %d]:", id+1), textcolor("Finished Creating %s %s %s\n\n", filenamecolor("%s", filename), textcolor("at"), directorycolor("%s", *dir)))
	fmt.Printf("%s %s", workercolor("[Worker %d]:", id+1), textcolor("Creating %s %s %s\n\n", filenamecolor("%s", filename), textcolor("at"), directorycolor("%s", *dir)))
	filepath, err := filepath.Abs(*dir + filename)
	if err != nil {
		log.Fatalln("Invalid path")
	}
	f, err := os.Create(filepath)
	if err != nil {

		log.Fatalln("failed to open file", err)
	}
	//value := <-records
	w := csv.NewWriter(f)
	err = w.WriteAll(a) // calls Flush internally
	if err != nil {
		log.Fatal(err)
	}
}
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
	fmt.Printf("%s", timecolor("%s", time.Since(now)))
}

// EFFECTIVELY TIME : 1s to work
