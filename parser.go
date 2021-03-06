package gobot

import (
	"fmt"
	"regexp"
	"strconv"
)

func Parse(input string) (Executable, error) {
	move := regexp.MustCompile(`^\s*MOVE\s*$`)
	left := regexp.MustCompile(`^\s*LEFT\s*$`)
	right := regexp.MustCompile(`^\s*RIGHT\s*$`)
	report := regexp.MustCompile(`^\s*REPORT\s*$`)
	place := regexp.MustCompile(`^\s*PLACE\s(\d+),(\d+),(NORTH|EAST|SOUTH|WEST)$`)
	switch {
	case move.MatchString(input):
		return *new(Move), nil
	case left.MatchString(input):
		return *new(Left), nil
	case right.MatchString(input):
		return *new(Right), nil
	case place.MatchString(input):
		var groups []string = place.FindStringSubmatch(input)
		return placeFromString(groups[1:]), nil
	case report.MatchString(input):
		return *new(Report), nil
	default:
		return *new(Move), fmt.Errorf("Unknown command %s", input)
	}
}

var directionTokenLookup = map[string]Direction{
	"NORTH": NORTH,
	"EAST":  EAST,
	"SOUTH": SOUTH,
	"WEST":  WEST,
}

func placeFromString(groups []string) Place {
	x, _ := strconv.Atoi(groups[0])
	y, _ := strconv.Atoi(groups[1])
	facing := directionTokenLookup[groups[2]]
	return Place{x, y, facing}
}
