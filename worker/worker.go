package worker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/Triad-0112/Worker/color"
)

func Fetcher(year int, id int) [][]string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	defer fmt.Printf("%s %s %s\n\n", colortext.Workercolor("[Worker %d] :", id+1), colortext.Textcolor("Finished collecting data of %s", colortext.Filenamecolor("%d", year)), colortext.Textcolor("from API"))
	fmt.Printf("%s %s", colortext.Workercolor("[Worker %d] :", id+1), colortext.Textcolor("Starting to fetch data of %s\n\n", colortext.Filenamecolor("%d.csv", year)))
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
