package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
)

/*
	goquery: go에서 쓸 수 있는 jQuery

	1. go get github.com/PuerkitoBio/goquery 로 설치

	이때 main.go 를 run하면 src에서 못찾는다는 둥 설치 제대로 안되어있음.
	go env -w GO111MODULE=off  (go 명령어 환경변수 변경)
	한 후에 다시 시도해본다. (정석적인 방법은 아닌듯..)

	-> workspace를 C:\Users\(name)\go에 두는 경우가 많더라구요.

	-> 환경변수에서 GOPATH를 현재 github.com/작성자 프로젝트명/ 으로 변경하거나
	   현재 작성중인 github.com폴더 자체를 통째로 c:/user/go/ 에다 옮겨준후
	   vscode 폴더를 옮겨준 폴더로 열어준다음 다시 go get ~~query부분부터 다시 하면 정상실행

	-> 나의 경우, go path setting(github readme.md에 올려둠)을 참고해서
	   GOROOT와 GOPATH를 재세팅하고 나서
	   go env -w GO111MODULE=off 한 후에야 드디어 됐다!

	2. indeed 페이지에서 각각의 페이지를 방문하고,
	   그 페이지로부터 job들을 추출

	3. 추출한 job들을 엑셀에 집어넣기

*/

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	totalPages := getPages() //총 페이지 수
	//fmt.Println(totalPages)

	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

//2.
func getPage(page int) {
	//strconv.Itoa == stringConversion + int to string
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting: ", pageURL)
}

//1. 페이지수를 가져온다
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
