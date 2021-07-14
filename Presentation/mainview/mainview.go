package mainview

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type MainView struct {
	version   string
	window    *widgets.QMainWindow
	inLabel   *widgets.QLabel
	inButton  *widgets.QPushButton
	outLabel  *widgets.QLabel
	outButton *widgets.QPushButton
	cnvButton *widgets.QPushButton
}

func NewMainView(version string) *MainView {
	mv := &MainView{
		version: version,
	}
	mv.Initialize()
	return mv
}

func (m *MainView) Initialize() {
	m.createWindow()
	m.viewDesign()
	m.addStyle()
}

func (m *MainView) createWindow() {

	// 参考
	// https://saitodev.co/article/Go%E8%A8%80%E8%AA%9E%E3%81%A7Qt%E3%81%AEQFormLayout%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6%E3%81%BF%E3%82%8B/
	// https://saitodev.co/article/Go%E8%A8%80%E8%AA%9E%E3%81%A7Qt%E3%81%AEQBoxLayout%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6%E3%81%BF%E3%82%8B
	// https://qiita.com/shu32/items/53204832c1074ff7cd7f

	// Base Window 作成
	core.QCoreApplication_SetApplicationName("OpenReviewConverter")
	core.QCoreApplication_SetOrganizationName("company")
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	widgets.NewQApplication(len(os.Args), os.Args)

	m.window = widgets.NewQMainWindow(nil, 0)
	m.window.SetMinimumSize2(800, 150)
	m.window.SetWindowTitle("オープンレビュー変換 : " + m.version)

}

func (m *MainView) viewDesign() {

	//フレームワークに上記で作成したレイアウトをセットする
	baseWidget := widgets.NewQWidget(nil, 0)
	//フレームワークにQTBoxLayoutをはめ込む
	//第一引数で0は左から右、1は右から左、2は上から下、3は下から上
	vbox := widgets.NewQBoxLayout(2, nil)
	baseWidget.SetLayout(vbox)

	// NCデータファイルの指定
	label1 := widgets.NewQLabel2("NCデータファイル", nil, 0)
	(*m).inLabel = widgets.NewQLabel2("", nil, 0)
	(*m).inButton = widgets.NewQPushButton2("参照", nil)
	hbox1 := widgets.NewQHBoxLayout()
	hbox1.SetSpacing(5)
	hbox1.AddWidget(label1, 0, core.Qt__AlignLeft)
	hbox1.AddWidget((*m).inLabel, 0, core.Qt__AlignBaseline)
	hbox1.AddWidget((*m).inButton, 0, core.Qt__AlignRight)

	// 結合ファイルの保存先
	label2 := widgets.NewQLabel2("結合ファイルの保存先", nil, 0)
	(*m).outLabel = widgets.NewQLabel2("", nil, 0)
	(*m).outButton = widgets.NewQPushButton2("参照", nil)
	hbox2 := widgets.NewQHBoxLayout()
	hbox2.SetSpacing(5)
	hbox2.AddWidget(label2, 0, core.Qt__AlignLeft)
	hbox2.AddWidget(m.outLabel, 0, core.Qt__AlignBaseline)
	hbox2.AddWidget(m.outButton, 0, core.Qt__AlignRight)

	// コンバートの指定
	m.cnvButton = widgets.NewQPushButton2("実行", nil)
	m.cnvButton.SetMinimumHeight(30)
	m.cnvButton.SetMaximumHeight(50)

	m.window.SetCentralWidget(baseWidget)
	vbox.AddLayout(hbox1, 0)
	vbox.AddLayout(hbox2, 0)
	vbox.AddWidget(m.cnvButton, 0, core.Qt__AlignBaseline)

}

func (m *MainView) addStyle() {
	// (*m).inLabel.SetStyleSheet(
	// 	"border:1px solid black;",
	// )
}
