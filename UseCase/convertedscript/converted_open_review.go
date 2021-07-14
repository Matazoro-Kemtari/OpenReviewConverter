package convertedscript

import (
	"OpenReviewConverter/Domain/alterationncscript"
	"OpenReviewConverter/Domain/openreview"
	"fmt"
	"log"
)

type ConvertedOpenReviewUseCase struct {
	converter  *openreview.ConvertedOpenReview
	fileReader *alterationncscript.FileReader
	fileWriter *alterationncscript.FileWriter
}

func NewConvertedOpenReviewUseCase(
	converter *openreview.ConvertedOpenReview,
	fileReader *alterationncscript.FileReader,
	fileWriter *alterationncscript.FileWriter,
) *ConvertedOpenReviewUseCase {
	return &ConvertedOpenReviewUseCase{
		converter:  converter,
		fileReader: fileReader,
		fileWriter: fileWriter,
	}
}

func (c *ConvertedOpenReviewUseCase) ConvertedOpenReview(inPath string, outPath string) error {
	// ファイルの読み込み
	log.Println("info:ファイル存在確認", inPath)
	if !(*c.fileReader).FileExist(inPath) {
		log.Println("info: ファイルが存在しません", inPath)
		return fmt.Errorf("ファイルが存在しません %s", inPath)
	}

	log.Printf("info: ファイルを読み込みます file: %s\n", inPath)
	lines, err := (*c.fileReader).ReadAll(inPath)
	if err != nil {
		return err
	}

	// 変換
	log.Printf("info: 変換します %v\n", lines)
	conversions, err := (*c.converter).Convert(lines)
	if err != nil {
		return err
	}

	// ファイルを保存
	log.Printf("info: 変換ファイルを保存します %v\n", conversions)
	if err := (*c.fileWriter).WriteAll(outPath, conversions); err != nil {
		return fmt.Errorf("変換ファイルの保存に失敗しました")
	}

	return nil
}
