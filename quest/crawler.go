package main

import (
	"fmt"
	"launchpad.net/xmlpath"
	"net/http"
	"sync"
)

func crawle(rank, url string, quest chan string, wg *sync.WaitGroup) {
	// XPATH of quest name
	// '//table/tr[(position() mod 2) = 1]/td[1]/a[1]/text()'
	path := xmlpath.MustCompile("//table/tr/td[1]/span/../a[1]/text()")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while fetching", url)
		return
	}
	defer resp.Body.Close()
	root, err := xmlpath.ParseHTML(resp.Body)
	if err != nil {
		fmt.Println("Error while parsing", url)
		return
	}
	iter := path.Iter(root)
	for iter.Next() {
		quest <- fmt.Sprintf("{\"rank\": \"%s\", \"name\": \"%s\"},", rank, iter.Node().String())
	}
	wg.Done()
}

func main() {
	questList := map[string]string{
		"集會所下位一星": "http://wiki.mh4g.org/data/1505.html", // 集會所 01
		"集會所下位二星": "http://wiki.mh4g.org/data/1511.html", // 集會所 02
		"集會所下位三星": "http://wiki.mh4g.org/data/1517.html", // 集會所 03
		"集會所上位四星": "http://wiki.mh4g.org/data/1518.html", // 集會所 04
		"集會所上位五星": "http://wiki.mh4g.org/data/1522.html", // 集會所 05
		"集會所上位六星": "http://wiki.mh4g.org/data/1529.html", // 集會所 06
		"集會所上位七星": "http://wiki.mh4g.org/data/1546.html", // 集會所 07
		"集會所Ｇ級一":  "http://wiki.mh4g.org/data/1719.html", // 集會所 G1
		"集會所Ｇ級二":  "http://wiki.mh4g.org/data/1720.html", // 集會所 G2
		"集會所Ｇ級三":  "http://wiki.mh4g.org/data/1758.html", // 集會所 G3
		"集會所下載Ｇ級": "http://wiki.mh4g.org/data/1724.html", // 下載Ｇ級
		"集會所下載上位": "http://wiki.mh4g.org/data/1567.html", // 下載上位
		"集會所下載下位": "http://wiki.mh4g.org/data/1490.html", // 下載下位
		"故事任務":    "http://wiki.mh4g.org/data/1785.html", // 故事任務
	}
	quest := make(chan string)
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(questList))
		for rank, url := range questList {
			go crawle(rank, url, quest, &wg)
		}
		wg.Wait()
		close(quest)
	}()
	fmt.Println("{\"quests\": [")
	for c := range quest {
		fmt.Println(c)
	}
	fmt.Println("]}")
}
