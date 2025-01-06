package main

// https://quera.org/problemset/237886

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var TagMap map[string]*Tag
var CreativeMap map[string]*Creative
var ZoneMap map[string]*Zone
var AdvertisementMap map[string]*Advertisement

type Creative struct {
	id   int
	name string
	cpc  int16
	tags []string
}

func checkTagsExist(tags []string) bool {
	for _, tag := range tags {
		if _, exists := TagMap[tag]; !exists {
			return false
		}
	}
	return true
}

func getDistinctTags(tags []string) []string {
	existMap := make(map[string]bool)
	distinctTags := make([]string, 0, len(tags))
	for _, tg := range tags {
		if _, exists := existMap[tg]; exists {
			continue
		}
		distinctTags = append(distinctTags, tg)
		existMap[tg] = true
	}
	return distinctTags
}

func addCreative(name string, cpc int16, tagNames []string) string {

	if allExist := checkTagsExist(tagNames); !allExist {
		return "Error: Tag not found"
	}

	distinctTags := getDistinctTags(tagNames)

	if _, exists := CreativeMap[name]; exists {
		return "Error: Ad already exists"
	}

	id := len(CreativeMap) + 1

	CreativeMap[name] = &Creative{id: id, name: name, cpc: cpc, tags: distinctTags}

	return fmt.Sprintf("Done: Ads id is %d", id)
}

func listCreatives() string {
	creativeList := make([]string, len(CreativeMap))

	for _, creative := range CreativeMap {
		creativeList[creative.id-1] = creative.name
	}

	creativeListStr := strings.Join(creativeList, " ")
	return fmt.Sprintf("ADSs: %s", creativeListStr)
}

type Zone struct {
	id   int
	name string
	cpc  int16
	tags []string
}

func addZone(name string, cpc int16, tagNames []string) string {
	if allExist := checkTagsExist(tagNames); !allExist {
		return "Error: Tag not found"
	}

	distinctTags := getDistinctTags(tagNames)

	if _, exists := ZoneMap[name]; exists {
		return "Error: Place already exists"
	}

	id := len(ZoneMap) + 1

	ZoneMap[name] = &Zone{id: id, name: name, cpc: cpc, tags: distinctTags}

	return fmt.Sprintf("Done: Place id is %d", id)
}

func listZones() string {
	zonesList := make([]string, len(ZoneMap))

	for _, zone := range ZoneMap {
		zonesList[zone.id-1] = zone.name
	}

	zoneListStr := strings.Join(zonesList, " ")
	return fmt.Sprintf("PLACEs: %s", zoneListStr)
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

func parseInstruction(insParts []string, keywords []string) [][]string {
	instructionElements := make([][]string, len(keywords)+1)

	keywordMap := make(map[string]bool)
	for _, word := range keywords {
		keywordMap[word] = true
	}

	insPartId := 0
	for _, word := range insParts {
		if _, exists := keywordMap[word]; exists {
			insPartId += 1
			instructionElements[insPartId] = make([]string, 0, len(insParts))
			continue
		} else {
			instructionElements[insPartId] = append(instructionElements[insPartId], word)
		}
	}
	return instructionElements
}

type AddCreativeCommand struct {
	name string
	cpc  int16
	tags []string
}

func (c *AddCreativeCommand) execute() string {
	return addCreative(c.name, c.cpc, c.tags)
}

type ListCreativesQuery struct{}

func (q *ListCreativesQuery) execute() string {
	return listCreatives()
}

type AddPlaceCommand struct {
	name string
	cpc  int16
	tags []string
}

func (c *AddPlaceCommand) execute() string {
	return addZone(c.name, c.cpc, c.tags)
}

type ListPlaceQuery struct{}

func (q *ListPlaceQuery) execute() string {
	return listZones()
}

func makeInstruction(instStr string) (instruction Instruction) {
	parts := strings.Split(instStr, " ")
	switch parts[0] {
	case "ADD-TAG":
		return &AddTagCommand{tagName: parts[1]}
	case "TAG-LIST":
		return &ListTagQuery{}
	case "ADD-ADS":
		parsed := parseInstruction(parts[1:], []string{"-name", "-cpc", "-tags"})[1:]
		cpc, _ := strconv.ParseInt(parsed[1][0], 10, 16)
		return &AddCreativeCommand{name: parsed[0][0], cpc: int16(cpc), tags: parsed[2]}
	case "ADS-LIST":
		return &ListCreativesQuery{}
	case "ADD-PLACE":
		parsed := parseInstruction(parts[1:], []string{"-name", "-cpc", "-tags"})[1:]
		cpc, _ := strconv.ParseInt(parsed[1][0], 10, 16)
		return &AddPlaceCommand{name: parsed[0][0], cpc: int16(cpc), tags: parsed[2]}
	case "PLACE-LIST":
		return &ListPlaceQuery{}
	default:
		panic("unexpected instruction")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	TagMap = make(map[string]*Tag)
	CreativeMap = make(map[string]*Creative)
	ZoneMap = make(map[string]*Zone)

	for scanner.Scan() {
		line := scanner.Text()
		instruction := makeInstruction(line)
		fmt.Println(instruction.execute())
	}
}
