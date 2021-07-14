package mainview

import (
	"OpenReviewConverter/UseCase/convertedscript"
	"fmt"
	"log"

	"github.com/therecipe/qt/widgets"
)

type MainViewController struct {
	version   string
	converter *convertedscript.ConvertedOpenReviewUseCase
	inPath    string
	outPath   string
	mainView  *MainView
}

func NewMainViewController(version string, converter *convertedscript.ConvertedOpenReviewUseCase) *MainViewController {
	return &MainViewController{
		version:   version,
		converter: converter,
		mainView:  NewMainView(version),
	}
}

func (v *MainViewController) Initialize() {
	// IN参照
	v.mainView.inButton.ConnectClicked(func(checked bool) {
		// ファイルを開くダイアログ表示
		p := widgets.QFileDialog_GetOpenFileName(
			nil,
			"NCファイルを指定してください",
			"C:\\",
			"すべて(*)",
			"すべて(*)",
			widgets.QFileDialog__DontUseCustomDirectoryIcons,
		)
		if len(p) > 0 {
			(*v).inPath = p
			v.mainView.inLabel.SetText(p)
			log.Println("入力ファイル", p)
		}
		v.setConvertButtonEnabled()
	})

	// OUT参照
	v.mainView.outButton.ConnectClicked(func(checked bool) {
		// 変換ファイルを保存ダイアログ表示
		p := widgets.QFileDialog_GetSaveFileName(
			nil,
			"オープンレビューファイル名を指定してください",
			"C:\\",
			"すべて(*)",
			"すべて(*)",
			widgets.QFileDialog__DontUseCustomDirectoryIcons,
		)
		if len(p) > 0 {
			(*v).outPath = p
			v.mainView.outLabel.SetText(p)
			log.Println(v.outPath)
		}
		v.setConvertButtonEnabled()
	})

	// 変換ボタン
	v.mainView.cnvButton.ConnectClicked(func(checked bool) {
		v.mainView.cnvButton.SetEnabled(!v.mainView.cnvButton.IsEnabled())
		if err := v.converter.ConvertedOpenReview(v.inPath, v.outPath); err != nil {
			// エラーメッセージ
			log.Println("error:", "オープンレビュー変換ボタンでエラー:", err)

			widgets.QMessageBox_Warning(
				v.mainView.window,
				"エラー情報",
				fmt.Sprintf(
					"深刻なエラーが発生しました:\n"+
						"進捗リストが送信できていない可能性があります\n"+
						"再実行しても回復しない場合は管理者へ連絡してください\n"+
						"%v",
					err,
				),
				widgets.QMessageBox__Ok,
				widgets.QMessageBox__Ok,
			)

			v.mainView.cnvButton.SetEnabled(!v.mainView.cnvButton.IsEnabled())
			return
		}

		// 終了メッセージ
		widgets.QMessageBox_Information(
			v.mainView.window,
			"正常終了",
			"変換処理が正常に終了しました",
			widgets.QMessageBox__Ok,
			widgets.QMessageBox__Ok,
		)

		v.mainView.inLabel.Clear()
		v.mainView.outLabel.Clear()
		v.inPath = ""
		v.outPath = ""
		v.setConvertButtonEnabled()
	})

	v.setConvertButtonEnabled()

	v.mainView.window.Show()
	widgets.QApplication_Exec()
}

/* 変換ボタンの有効化 */
func (v *MainViewController) setConvertButtonEnabled() {
	v.mainView.cnvButton.SetEnabled(len(v.inPath) > 0 && len(v.outPath) > 0)
}
