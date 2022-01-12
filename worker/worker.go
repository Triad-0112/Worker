package worker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/"
)

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
