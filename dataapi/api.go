package api
const Baseurl = "https://data.gov.sg/api/action/datastore_search?resource_id=eb8b932c-503c-41e7-b513-114cffbe2338&q="

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