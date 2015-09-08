package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var quests List

func randQuest() string {
	num := rand.Int31n(int32(len(quests.Quests)))
	return fmt.Sprintf("%s\t%s", quests.Quests[num].Rank, quests.Quests[num].Name)
}

func readQuests(name string) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	content, err := ioutil.ReadAll(file)
	json.Unmarshal(content, &quests)
	return nil
}

func root(rw http.ResponseWriter, req *http.Request) {
	output := randQuest()
	io.WriteString(rw, output+"\n")
	fmt.Println(output)
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	input := flag.String("file", "quests.json", "Quests json location")
	httpPtr := flag.Bool("http", false, "Open http server")
	ip := flag.String("ip", "0.0.0.0", "http server ip")
	port := flag.Int("port", 8080, "http server port")
	flag.Parse()
	err := readQuests(*input)
	if err != nil {
		fmt.Println("Read Quest Error!!")
		return
	}
	if *httpPtr == false {
		fmt.Println(randQuest())
		return
	}
	http.Handle("/root/", http.HandlerFunc(root))
	err = http.ListenAndServe(fmt.Sprintf("%s:%d", *ip, *port), nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
