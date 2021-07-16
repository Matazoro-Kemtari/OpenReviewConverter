package openreview

import (
	"fmt"
	"log"
	"regexp"
	"sort"
)

type ConvertedOpenReview struct{}

func NewConvertedOpenReview() *ConvertedOpenReview {
	return &ConvertedOpenReview{}
}

func (c *ConvertedOpenReview) Convert(sources []string) ([]string, error) {
	if len(sources) == 0 {
		return nil, fmt.Errorf("変換対象がありません")
	}

	idx := sort.Search(len(sources), func(i int) bool {
		return sources[i] >= "%"
	})
	fmt.Print(idx)

	var canSkipLine bool
	rexM00 := regexp.MustCompile(`^M00$`)
	rexM01 := regexp.MustCompile(`^M01$`)
	rexM30orM99 := regexp.MustCompile(`^\(M(30|99)\)$`)
	rexM30 := regexp.MustCompile(`^M30$`)
	canAppendFinallyM30 := true
	regPercentOrBlank := regexp.MustCompile(`^%?$`)
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
			if i > 0 && canAppendFinallyM30 && rexM30.MatchString(line) {
				// M30が見つかったら追記しなくても良いかも
				canAppendFinallyM30 = false
			} else if i > 0 && !canAppendFinallyM30 && !regPercentOrBlank.MatchString(line) {
				// '?' '空行' 以外が見つかったら最後のコマンドじゃなかった
				canAppendFinallyM30 = true
			}

			if i > 0 && canAppendFinallyM30 && regPercent.MatchString(line) {
				// M30はファイルの最後の%の1つ前に追記する
				res = append(res, "M30")
			}
			res = append(res, line)
		}
	}
	return res, nil
}
