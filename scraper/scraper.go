package scraper

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Danny-Dasilva/CycleTLS/cycletls"
	"github.com/PuerkitoBio/goquery"
)

const (
	userAgent string = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0"
	ja3       string = "771,4865-4867-4866-49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-51-57-47-53-10,0-23-65281-10-11-35-16-5-51-43-13-45-28-21,29-23-24-25-256-257,0"
)

type jobDetails struct {
	id       string
	title    string
	company  string
	location string
	salary   string
	summary  string
	link     string
}

// Scrape scrapes job listings with term
func Scrape(term string) {
	url := "https://www.indeed.com/jobs?q=" + term + "&start="
	totalPages := getPages(url)
	c := make(chan []jobDetails)
	for i := 0; i < totalPages; i++ {
		go getPage(url, i, c)
	}
	file, err := os.Create("jobs.csv")
	w := csv.NewWriter(file)
	defer w.Flush()
	headers := []string{"Link", "Title", "Location", "Salary", "Summary"}
	err = w.Write(headers)
	checkErr(err)

	checkErr(err)
	jobCounter := 0
	for i := 0; i < totalPages; i++ {
		jobs := <-c
		jobCounter += len(jobs)
		writeCSV(jobs, w)
	}
	fmt.Println("Total", jobCounter, "job(s) exported")
}

func getPage(url string, pageNum int, mainChan chan<- []jobDetails) {
	c := make(chan jobDetails)
	pageURL := url + fmt.Sprint(pageNum*10)
	fmt.Println("Requesting", pageURL)
	resp, err := getRespWithTLS(pageURL)
	checkErr(err)
	checkCode(resp.Status)
	body := strings.NewReader(resp.Body)
	doc, err := goquery.NewDocumentFromReader(body)
	checkErr(err)

	jobs := []jobDetails{}
	blocks := doc.Find(".jobsearch-LeftPane").Find(".job_seen_beacon")
	blocks.Each(func(i int, s *goquery.Selection) {
		go extractJob(s, c)
	})
	for i := 0; i < blocks.Length(); i++ {
		jobs = append(jobs, <-c)
	}
	mainChan <- jobs
}

// Request page with get method with tls credentials to avoid 403 error
func getPages(url string) int {
	pages := 0
	resp, err := getRespWithTLS(url)
	checkErr(err)
	checkCode(resp.Status)
	body := strings.NewReader(resp.Body)
	doc, err := goquery.NewDocumentFromReader(body)
	checkErr(err)
	pages = doc.Find(`nav[role="navigation"]`).Find("a").Length()
	return pages
}

// How to avoid request failed with status code: 403
func getRespWithTLS(baseURL string) (*cycletls.Response, error) {
	client := cycletls.Init()
	res, err := client.Do(baseURL, cycletls.Options{
		Body:      "",
		Ja3:       ja3,
		UserAgent: userAgent,
	}, "GET")
	// returning the pointer helps avoid copying the whole struct
	return &res, err
}

func extractJob(s *goquery.Selection, c chan<- jobDetails) {
	id, _ := s.Find(".jobTitle>a").Attr("id")
	id = strings.Split(id, "_")[1]
	title := s.Find(".jobTitle>a").Text()
	// link, _ := s.Find(".jobTitle>a").Attr("href")
	company := s.Find(".companyName").Text()
	location := s.Find(".companyLocation").Text()
	salary := s.Find(`.metadata[class*="salary"]`).Text()
	if salary == "" {
		salary = "Not Specified"
	}
	summary := s.Find(".job-snippet>ul").Text()
	c <- jobDetails{
		id:       id,
		title:    CleanString(title),
		company:  CleanString(company),
		location: CleanString(location),
		salary:   CleanString(salary),
		summary:  CleanString(summary),
		link:     "https://www.indeed.com/viewjob?jk=" + id,
	}
}

func writeCSV(jobs []jobDetails, w *csv.Writer) {
	for _, job := range jobs {
		jobSlice := []string{job.link, job.title, job.location, job.salary, job.summary}
		err := w.Write(jobSlice)
		checkErr(err)
	}
}

// CleanString cleanses the string
func CleanString(str string) string {
	// TrimSpace removes leading and trailing spaces
	// Join and Split removes duplicate spaces
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res int) {
	if res != 200 {
		log.Fatalln("Request failed with status code:", res)
	}
}
