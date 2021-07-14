package openreview

import (
	"bufio"
	"os"
	"reflect"
	"testing"
)

func readInputTestFile() []string {
	fp, _ := os.Open("testdata\\testAllScript")
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	var r []string
	for scanner.Scan() {
		r = append(r, scanner.Text())
	}
	return r
}

func readOutputTestFile() []string {
	fp, _ := os.Open("testdata\\testVerifyScript")
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	var r []string
	for scanner.Scan() {
		r = append(r, scanner.Text())
	}
	return r
}

func TestConvertedOpenReview_Convert(t *testing.T) {
	type args struct {
		sources []string
	}
	tests := []struct {
		name    string
		c       *ConvertedOpenReview
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "正常系_オープンレビュー変換できること",
			c:    NewConvertedOpenReview(),
			args: args{
				sources: readInputTestFile(),
			},
			want:    readOutputTestFile(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Convert(tt.args.sources)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertedOpenReview.Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			f, _ := os.Create("testdata/testResult" + tt.name)
			defer f.Close()
			for _, line := range got {
				if _, err := f.WriteString(line + "\r\n"); err != nil {
					panic(err)
				}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertedOpenReview.Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
