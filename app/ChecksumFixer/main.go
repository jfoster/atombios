package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"

	"github.com/jfoster/atombios/atom"

	"github.com/jfoster/go-dir"
)

const (
	name    = "ChecksumFixer"
	version = "0.0.1"
)

var (
	err error
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

	var file *dir.File

	button := widgets.NewQPushButton2("Open File", nil)
	button.ConnectClicked(func(bool) {
		fp := widgets.QFileDialog_GetOpenFileName(nil, "Open File...", core.QDir_CurrentPath(), "ROM files (*.rom);;BIN files (*.bin);;All files (*.*)", "", 0)
		if fp != "" && loadFile(fp) {
			window.StatusBar().ShowMessage(fmt.Sprintf("Opened %v", core.QDir_ToNativeSeparators(fp)), 0)

			file = dir.NewFilePath(fp)
			window.StatusBar().ShowMessage(file.String(), 0)

			err = file.Read()
			if err != nil {
				window.StatusBar().ShowMessage(err.Error(), 0)
				return
			}

			var checksum = file.Content[atom.AtomRomChecksumOffset]
			var expected = atom.CalcChecksum(file.Content)

			window.StatusBar().ShowMessage(fmt.Sprintf("Current: 0x%X Expected: 0x%X", checksum, expected), 0)
		} else {
			window.StatusBar().ShowMessage(fmt.Sprintf("Could not open %v", core.QDir_ToNativeSeparators(file.FullName())), 0)
		}
	})
	widget.Layout().AddWidget(button)

	receiveButton := widgets.NewQPushButton2("Fix Checksum", nil)
	receiveButton.ConnectClicked(func(bool) {
		if len(file.Content) > 0 {
			atom.FixChecksum(&file.Content)
			err := file.Write()
			if err != nil {
				window.StatusBar().ShowMessage(err.Error(), 0)
				return
			}
			window.StatusBar().ShowMessage(fmt.Sprintf("Written to %s", core.QDir_ToNativeSeparators(file.Path())), 0)
		}
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
