package ncfile

import (
	"OpenReviewConverter/Domain/alterationncscript"
	"bufio"
	"fmt"
	"os"
)

type ReadableNcScriptFile struct{}

func NewReadableNcScriptFile() *alterationncscript.FileReader {
	var obj alterationncscript.FileReader = &ReadableNcScriptFile{}
	return &obj
}

func (n *ReadableNcScriptFile) ReadAll(path string) ([]string, error) {
	if len(path) == 0 {
		return nil, fmt.Errorf("引数が空です")
	}

	fp, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("ファイルの読み込みに失敗しました error:%v", err)
	}
	defer fp.Close()

	s := bufio.NewScanner(fp)
	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return lines, nil
}

func (n *ReadableNcScriptFile) FileExist(file string) bool {
	if len(file) <= 0 {
		return false
	} else if f, err := os.Stat(file); os.IsNotExist(err) || f.IsDir() {
		return false
	}
	return true
}
