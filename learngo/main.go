package main

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

/*
	일 정보를 추출하는 함수는 getPage()와 channel 을 이용해서 정보를 주고 받음
	getPage()는 main()과 channel을 이용해서 소통함

	channel은 2개
	1. main() <-> getPage()
	2. getPage() <-> extractJob()


	1. extractJob()을 goroutine으로 만든다
	   -> 1. getPage()에 channel을 만든다
	   -> 2. extractJob()에서 channel을 받는다
	   c := make(chan extractedJob)

*/

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

//gets all the jobs
func main() {
	var jobs []extractedJob //jobs라는 빈 배열

	totalPages := getPages() //총 페이지 수
	//fmt.Println(totalPages)

	for i := 0; i < totalPages; i++ { //각 페이지별로 getPage함수 호출
		extractedJobs := getPage(i) //한 페이지당 50개 정보가 있음
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	fmt.Println("Done, extracted: ", len(jobs))
}

//4. 일자리를 csv 파일로 저장
func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file) //파일을 새로 만듦
	defer w.Flush()          //defer: 함수가 끝나는 시점에 // flush(): 파일에 데이터를 입력하는 함수

	headers := []string{"Link", "Title", "Location", "Salary", "Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	//write 함수에는 []string 형태가 입력되어야함
	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	} // 이게 끝나면 defer가 실행되고, 데이터가 파일에 입력됨

}

//2. 각 페이지에 있는 일자리를 모두 반환
func getPage(page int) []extractedJob {
	var jobs []extractedJob //jobs 라는 빈 배열

	c := make(chan extractedJob)

	//2.1 필요한 주소를 만듦 + 정보요청
	//strconv.Itoa == stringConversion + int to string
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting: ", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body) //res.Body = byte(IO (입출력)) -> close해줘야함
	checkErr(err)

	//2.2 모든 카드를 찾은 후 각 카드에서 일자리 정보를 가져옴
	searchCards := doc.Find(".job_seen_beacon")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c) //extractedJob struct를 job에 저장
		//jobs = append(jobs, job)   //job이 추출될 때마다 jobs에 추가
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	return jobs //main으로 돌아감
}

// 3. extractedJob struct를 반환
func extractJob(card *goquery.Selection, c chan<- extractedJob) extractedJob {
	id_path := card.Find(".jcs-JobTitle")
	id, _ := id_path.Attr("data-jk") //data-jk 라는 속성 추출
	title := cleanString(id_path.Find("a>span").Text())
	location := cleanString(card.Find(".companyLocation").Text())
	salary := cleanString(card.Find(".attribute_snippet").Text())
	summary := cleanString(card.Find(".job-snippet").Text())
	//fmt.Println(id, title, location, salary, summary)

	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary}
}

func cleanString(str string) string {
	//Fields(): strings를 나눔 -> 텍스트로만 이루어진 배열을 리턴
	//TrimSpace(): 양쪽의 공백을 없앰
	//return strings.Fields(strings.TrimSpace(str))

	//join: 배열을 seperator로 합침
	return strings.Join(strings.Fields(strings.TrimSpace(str)), "  ")
}

//1. 총 페이지수를 가져온다
func getPages() int {

	pages := 0
	/* Get()
	func Get(url string) (resp *Response, err error) {
		return DefaultClient.Get(url)
	}
	*/
	//resp *Response, err error
	res, err := http.Get(baseURL)

	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	//goquery document 가져오기
	//Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body) //res.Body = byte(IO (입출력)) -> close해줘야함
	checkErr(err)

	//class name="pagination"인 것을 찾기
	//
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		//selection.Html() -> gets the HTML contents of the first element in the set of matched elements. It includes text and comment nodes.
		//fmt.Println(s.Html())

		//pagination 안에 몇개의 링크가 있는지(페이지 수)
		pages = s.Find("a").Length()

	}) //find한 각각의 것
	//fmt.Println(doc)

	return pages //전체페이지수
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}
}
