package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"

	"github.com/jfoster/atombiostool/atom"
)

const (
	name    = "AtomBiosTool"
	version = "0.0.1"
)

func main() {
	// App
	app := widgets.NewQApplication(len(os.Args), os.Args)
	app.SetApplicationName(name)
	app.SetApplicationVersion(version)

	// Main Window
	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(300, 200)
	window.SetWindowTitle(name)
	window.SetStatusBar(widgets.NewQStatusBar(window))

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQGridLayout2())
	window.SetCentralWidget(widget)

	button := widgets.NewQPushButton2("Open File", nil)
	button.ConnectClicked(func(bool) {
		fileName := widgets.QFileDialog_GetOpenFileName(nil, "Open File...", core.QDir_HomePath(), "ROM files (*.rom);;BIN files (*.bin);;All files (*.*)", "", 0)
		if fileName != "" {
			if loadFile(fileName) {
				window.StatusBar().ShowMessage(fmt.Sprintf("Opened %v", core.QDir_ToNativeSeparators(fileName)), 0)

				dat, err := ioutil.ReadFile(fileName)
				if err != nil {
					fmt.Print(err)
				}

				var checksum byte = dat[atom.AtomRomChecksumOffset]
				var size int = int(dat[atom.AtomRomHeaderSizeOffset]) * 512

				var check byte
				for _, v := range dat[0:size] {
					check += v
				}

				var expected = checksum - check

				window.StatusBar().ShowMessage(fmt.Sprintf("Current: 0x%X Expected: 0x%X", checksum, expected), 0)

			} else {
				window.StatusBar().ShowMessage(fmt.Sprintf("Could not open %v", core.QDir_ToNativeSeparators(fileName)), 0)
			}
		}
	})
	widget.Layout().AddWidget(button)

	receiveButton := widgets.NewQPushButton2("Fix Checksum", nil)
	receiveButton.ConnectClicked(func(bool) {

	})
	widget.Layout().AddWidget(receiveButton)

	// Run App
	widgets.QApplication_SetStyle2("fusion")
	window.Show()
	app.Exec()
}

func loadFile(fileName string) bool {
	if !(core.QFile_Exists(fileName)) {
		return false
	}
	var file = core.NewQFile2(fileName)
	if !file.Open(core.QIODevice__ReadOnly) {
		return false
	}
	defer file.Close()

	return true
}
