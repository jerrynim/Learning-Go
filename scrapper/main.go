package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	var totalJobs []extractedJob
	//? 모든 페이지의 jobs 를 모은 채널
	channelForTotalJobs := make(chan []extractedJob)

	totalPages := getPages()

	//* channelForTotalJobs 에 extractedJobs를 채운다.
	for i := 0; i < totalPages; i++ {
		go getPage(i, channelForTotalJobs)
	}

	for i := 0; i < totalPages; i++ {
		pagejobs := <-channelForTotalJobs
		totalJobs = append(totalJobs, pagejobs...)
	}
	fmt.Println("Done, extracted", len(totalJobs))
	writeJobs(totalJobs)
}

//* .csv파일에 extractedJobs[]를 저장
func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Link", "Title", "Location", "Salary", "Summary"}

	wErr := writer.Write(headers)
	checkErr(wErr)

	var totalJobSlices []string

	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}
		totalJobSlices = append(totalJobSlices, jobSlice...)
	}
	jwErr := writer.Write(totalJobSlices)
	checkErr(jwErr)

	// channelForWrite := make(chan []string)

	// for _, job := range jobs {
	// 	go writeJob(writer, job, channelForWrite)
	// }
}

func writeJob(writer *csv.Writer, job extractedJob, channelForWrite chan<- []string) {
	jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}
	jwErr := writer.Write(jobSlice)
	checkErr(jwErr)
}

func getPage(page int, channelForTotalJobs chan<- []extractedJob) {
	var pagejobs []extractedJob

	//? 페이지의 job을 모으는 채널
	channelForEachPage := make(chan extractedJob)

	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)

	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)
	//? res.Body = bytes. 메모리 누수 방지
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, channelForEachPage)
	})
	fmt.Println(page, searchCards.Length())

	for i := 0; i < searchCards.Length(); i++ {
		cardJobs := <-channelForEachPage
		pagejobs = append(pagejobs, cardJobs)
	}
	channelForTotalJobs <- pagejobs
}

//* 채용 카드의 직업 정보를 추출
func extractJob(card *goquery.Selection, channelForEachPage chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find(".title>a").Text())
	location := cleanString(card.Find(".sjcl").Text())
	salary := cleanString(card.Find(".salaryText").Text())
	summary := cleanString(card.Find(".summary").Text())

	channelForEachPage <- extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary}
}

//* 조회할 페이지 갯수 검색
func getPages() int {
	pagesCount := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)
	//? res.Body = bytes. 메모리 누수 방지
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each((func(i int, s *goquery.Selection) {
		pagesCount = s.Find("a").Length()
	}))

	return pagesCount
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
