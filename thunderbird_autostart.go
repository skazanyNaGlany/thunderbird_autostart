package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"github.com/skazanyNaGlany/thunderbird_autostart/windows/main_window"
)

func main() {
	changeCwd()

	if isRunThunderbirdCommand() {
		runThunderbird()
	} else {
		userInterface()
	}
}

func changeCwd() {
	os.Chdir(
		filepath.Dir(os.Args[0]))
}

func findThunderbirdExe(state *State) {
	for _, pathname := range THUNDERBIRD_PATHNAMES {
		if _, err := os.Stat(pathname); err == nil {
			state.SetThunderbirdExePathname(pathname)
			return
		}
	}
}

func isRunThunderbirdCommand() bool {
	return len(os.Args) == 3 && os.Args[1] == RUN_COMMAND && os.Args[2] != ""
}

func runThunderbird() {
	exec.Command(os.Args[2]).Start()
}

func userInterface() {
	app := app.New()

	state := State{}
	installer := AutostartInstaller{}

	installer.Init()

	mainWindow := main_window.MainWindow{}
	mainWindow.SetVersion(APP_VERSION)
	mainWindow.Init(app)

	if err := state.LoadStateFromFile(STATE_FILE_PATHNAME); err != nil {
		log.Println(err)
	}

	if state.GetThunderbirdExePathname() == "" {
		findThunderbirdExe(&state)
	}

	mainWindow.GetPathInput().SetText(
		state.GetThunderbirdExePathname())
	mainWindow.SetToggleAutostartButtonState(
		!installer.IsAutostartEnabled())

	if installer.IsAutostartEnabled() {
		mainWindow.GetPathInput().Disable()
		mainWindow.GetSetPathButton().Disable()
	}

	mainWindow.GetSetPathButton().OnTapped = func() {
		fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if reader == nil {
				return
			}

			exePathname := reader.URI().String()
			exePathname = strings.Replace(exePathname, "file://", "", 1)

			mainWindow.GetPathInput().SetText(exePathname)
			state.SetThunderbirdExePathname(exePathname)
		}, mainWindow.GetWindow())

		fileDialog.Show()
	}

	mainWindow.GetAboutButton().OnTapped = func() {
		dialog.ShowInformation(
			ABOUT_TITLE,
			fmt.Sprintf(ABOUT_MESSAGE, APP_VERSION),
			mainWindow.GetWindow())
	}

	mainWindow.GetExitButton().OnTapped = func() {
		app.Quit()
	}

	mainWindow.GetToggleAutostartButton().OnTapped = func() {
		if err := installer.EnableAutostart(
			!installer.IsAutostartEnabled(),
			state.GetThunderbirdExePathname()); err != nil {
			log.Println(err)
		}

		mainWindow.SetToggleAutostartButtonState(!installer.IsAutostartEnabled())

		if installer.IsAutostartEnabled() {
			mainWindow.GetPathInput().Disable()
			mainWindow.GetSetPathButton().Disable()
		} else {
			mainWindow.GetPathInput().Enable()
			mainWindow.GetSetPathButton().Enable()
		}
	}

	mainWindow.GetWindow().SetOnClosed(func() {
		if err := state.SaveStateToFile(STATE_FILE_PATHNAME); err != nil {
			log.Println(err)
		}
	})

	mainWindow.GetWindow().SetMaster()
	mainWindow.GetWindow().CenterOnScreen()
	mainWindow.GetWindow().ShowAndRun()
}

// <div class="streamStatus"><div class="layout horizontal"><div class="flex"></div><div class="pd8 round" style="background: rgba(0, 0, 0, 0.5);"><div class="fs08 uppercase op05">Connected</div></div><div class="flex"></div></div></div>
