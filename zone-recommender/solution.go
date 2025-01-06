package main

// https://quera.org/problemset/237886

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var TagMap map[string]*Tag

type Creative struct {
	name string
	cpc  int8
	tags []Tag
}

type Zone struct {
	name string
	cpc  int8
	tags []Tag
}

type Tag struct {
	id   int
	name string
}

type Advertisement struct {
	creative *Creative
	zone     *Zone
}

type Instruction interface {
	execute() string
}

type AddTagCommand struct {
	tagName string
}

func (c *AddTagCommand) execute() string {
	_, ok := TagMap[c.tagName]

	if ok {
		return "Error: Tag already exists"
	}

	id := len(TagMap) + 1
	TagMap[c.tagName] = &Tag{id: id, name: c.tagName}

	return fmt.Sprintf("Done: Tag id is %d", id)
}

type ListTagQuery struct{}

func (q *ListTagQuery) execute() string {
	tagList := make([]string, len(TagMap))

	for _, tag := range TagMap {
		tagList[tag.id-1] = tag.name
	}

	tagListStr := strings.Join(tagList, " ")
	return fmt.Sprintf("TAGS: %s", tagListStr)
}

func parseInstruction(instStr string) (instruction Instruction) {
	parts := strings.Split(instStr, " ")

	switch parts[0] {
	case "ADD-TAG":
		return &AddTagCommand{tagName: parts[1]}
	case "TAG-LIST":
		return &ListTagQuery{}
	default:
		panic("unexpected instruction")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	TagMap = make(map[string]*Tag)

	for scanner.Scan() {
		line := scanner.Text()
		instruction := parseInstruction(line)
		fmt.Println(instruction.execute())
	}
}
