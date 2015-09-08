package main

import (
	"fmt"
	"launchpad.net/xmlpath"
	"net/http"
	"sync"
)

func crawle(url string, quest chan string, wg *sync.WaitGroup) {
	// XPATH of quest name
	// '//table/tr[(position() mod 2) = 1]/td[1]/a[1]/text()'
	path := xmlpath.MustCompile("//table/tr/td[1]/a[1]/text()")
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
	count := 1
	for iter.Next() {
		name := iter.Node().String()
		if (count % 2) == 1 {
			quest <- name
		}
		count++
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	questList := map[string]string{
		"01":  "http://wiki.mh4g.org/data/1505.html", // 集會所 01
		"02":  "http://wiki.mh4g.org/data/1511.html", // 集會所 02
		"03":  "http://wiki.mh4g.org/data/1517.html", // 集會所 03
		"04":  "http://wiki.mh4g.org/data/1518.html", // 集會所 04
		"05":  "http://wiki.mh4g.org/data/1522.html", // 集會所 05
		"06":  "http://wiki.mh4g.org/data/1529.html", // 集會所 06
		"07":  "http://wiki.mh4g.org/data/1546.html", // 集會所 07
		"G1":  "http://wiki.mh4g.org/data/1719.html", // 集會所 G1
		"G2":  "http://wiki.mh4g.org/data/1720.html", // 集會所 G2
		"G3":  "http://wiki.mh4g.org/data/1758.html", // 集會所 G3
		"DLG": "http://wiki.mh4g.org/data/1724.html", // 下載Ｇ級
		"DLL": "http://wiki.mh4g.org/data/1567.html", // 下載上位
		"DLU": "http://wiki.mh4g.org/data/1490.html", // 下載下位
	}
	quest := make(chan string)
	go func() {
		wg.Add(len(questList))
		for _, url := range questList {
			go crawle(url, quest, &wg)
		}
		wg.Wait()
		close(quest)
	}()
	for c := range quest {
		fmt.Println(c)
	}
}
