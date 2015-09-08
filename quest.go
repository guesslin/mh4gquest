package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func randQuest(quests List) string {
	num := rand.Int31n(int32(len(quests.Quests)))
	return fmt.Sprintf("%s\t%s", quests.Quests[num].Rank, quests.Quests[num].Name)
}

func readQuests(name string) (List, error) {
	file, err := os.Open(name)
	if err != nil {
		return List{}, err
	}
	content, err := ioutil.ReadAll(file)
	var questList List
	json.Unmarshal(content, &questList)
	return questList, nil
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	input := flag.String("file", "quests.json", "Quests json location")
	flag.Parse()
	questList, err := readQuests(*input)
	if err != nil {
		fmt.Println("Read Quest Error!!")
		return
	}
	fmt.Println(randQuest(questList))
}
