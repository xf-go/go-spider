package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var url = "https://www.qb5.tw/book_114945/%s.html"

func main() {
	startTime := time.Now()

	start := 47119721
	end := 47119721 + 100

	ch := make(chan int)
	for i := start; i < end; i++ {
		go run(i, ch)
	}

	for i := start; i < end; i++ {
		n := <-ch
		fmt.Println("n: ", n)
	}

	elapsed := time.Since(startTime)

	fmt.Println("elapsed: ", elapsed)
}

func run(page int, ch chan int) {
	data, err := getData(fmt.Sprintf(url, strconv.Itoa(page)))
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	title, err := getTitle(data)
	// res, err := reg(string(bs))
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	writeFileLine("D:\\goworkspace\\go-spider\\cont.txt", title)

	ch <- page
}

func getData(reqURL string) ([]byte, error) {
	bs, err := CURLGet(reqURL)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func getTitle(content []byte) ([]byte, error) {
	reg, err := regexp.Compile("<h1>(.+?)</h1>")
	if err != nil {
		return nil, err
	}

	find := reg.FindSubmatch(content)

	bs, err := GbkToUtf8(find[1])
	if err != nil {
		return nil, err
	}

	return bs, nil
}
