package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	quests List
)

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

func root(c *gin.Context) {
	output := randQuest()
	c.Writer.WriteString(output)
	log.Println(output)
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	log.SetOutput(os.Stderr)
}

func main() {
	var input string
	var httpPtr bool
	var ip string
	var port int
	flag.StringVar(&input, "file", "quests.json", "Quests json location")
	flag.BoolVar(&httpPtr, "http", false, "Open http server")
	flag.StringVar(&ip, "ip", "0.0.0.0", "http server ip")
	flag.IntVar(&port, "port", 8080, "http server port")
	flag.Parse()
	err := readQuests(input)
	if err != nil {
		log.Fatal("Read Quest Error!!")
	}
	if httpPtr == false {
		log.Println(randQuest())
		return
	}
	server := gin.New()
	server.Use(gin.LoggerWithWriter(os.Stderr), gin.Recovery())
	server.GET("/", root)
	server.Run(fmt.Sprintf("%s:%d", ip, port))
}
