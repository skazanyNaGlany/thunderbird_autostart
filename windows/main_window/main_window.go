package main_window

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type MainWindow struct {
	window                     fyne.Window
	vContainer                 *fyne.Container
	pathLabel                  *widget.Label
	pathInput                  *widget.Entry
	setPathButton              *widget.Button
	aboutButton                *widget.Button
	exitButton                 *widget.Button
	toggleAutostartButton      *widget.Button
	toggleAutostartButtonState bool
	version                    string
}

func (mw *MainWindow) Init(app fyne.App) {
	mw.window = app.NewWindow(fmt.Sprintf(TITLE, mw.version))
	mw.vContainer = container.NewVBox()
	mw.pathLabel = widget.NewLabel(PATHNAME_LABEL)
	mw.pathInput = widget.NewEntry()
	mw.setPathButton = widget.NewButton(PATHNAME_BUTTON, func() {})
	mw.aboutButton = widget.NewButton(ABOUT_BUTTON, func() {})
	mw.exitButton = widget.NewButton(EXIT_BUTTON, func() {})
	mw.toggleAutostartButton = widget.NewButton(TOGGLE_AUTOSTART_BUTTON_DISABLE, func() {})

	mw.vContainer.Add(mw.pathLabel)
	mw.vContainer.Add(mw.pathInput)
	mw.vContainer.Add(widget.NewSeparator())
	mw.vContainer.Add(mw.setPathButton)
	mw.vContainer.Add(mw.toggleAutostartButton)
	mw.vContainer.Add(mw.aboutButton)
	mw.vContainer.Add(mw.exitButton)

	mw.window.SetContent(mw.vContainer)
	mw.window.SetPadded(true)

	mw.window.Resize(fyne.NewSize(
		WINDOW_WIDTH,
		WINDOW_HEIGHT,
	))
}

func (mw *MainWindow) GetWindow() fyne.Window {
	return mw.window
}

func (mw *MainWindow) GetSetPathButton() *widget.Button {
	return mw.setPathButton
}

func (mw *MainWindow) SetToggleAutostartButtonState(enable bool) {
	if enable {
		mw.toggleAutostartButton.SetText(TOGGLE_AUTOSTART_BUTTON_ENABLE)
	} else {
		mw.toggleAutostartButton.SetText(TOGGLE_AUTOSTART_BUTTON_DISABLE)
	}

	mw.toggleAutostartButtonState = enable
}

func (mw *MainWindow) GetToggleAutostartButtonState() bool {
	return mw.toggleAutostartButtonState
}

func (mw *MainWindow) GetToggleAutostartButton() *widget.Button {
	return mw.toggleAutostartButton
}

func (mw *MainWindow) GetAboutButton() *widget.Button {
	return mw.aboutButton
}

func (mw *MainWindow) GetExitButton() *widget.Button {
	return mw.exitButton
}

func (mw *MainWindow) SetVersion(version string) {
	mw.version = version
}

func (mw *MainWindow) GetPathInput() *widget.Entry {
	return mw.pathInput
}
