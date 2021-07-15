package openreview

import (
	"fmt"
	"log"
	"regexp"
)

type ConvertedOpenReview struct{}

func NewConvertedOpenReview() *ConvertedOpenReview {
	return &ConvertedOpenReview{}
}

func (c *ConvertedOpenReview) Convert(sources []string) ([]string, error) {
	if len(sources) == 0 {
		return nil, fmt.Errorf("変換対象がありません")
	}

	var canSkipLine bool
	rexM00 := regexp.MustCompile(`^M00$`)
	rexM01 := regexp.MustCompile(`^M01$`)
	rexM30orM99 := regexp.MustCompile(`^\(M(30|99)\)$`)
	regPercent := regexp.MustCompile(`^%$`)
	var res []string
	for i, line := range sources {
		log.Println("line:", line, "canSkipLine:", canSkipLine)
		if rexM00.MatchString(line) {
			continue
		}

		if !canSkipLine && rexM01.MatchString(line) {
			res = append(res, line)
			canSkipLine = true
		} else if canSkipLine && rexM30orM99.MatchString(line) {
			res = append(res, line)
			canSkipLine = false
		} else if !canSkipLine {
			if i > 0 && regPercent.MatchString(line) {
				res = append(res, "M30")
			}
			res = append(res, line)
		}
	}
	return res, nil
}
