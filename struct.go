package main

type Quest struct {
	Rank string `json:"rank"`
	Name string `json:"name"`
}

type List struct {
	Quests []Quest `json:"quests"`
}
