package main

const ABOUT_TITLE = "Thunderbird Autostart"
const ABOUT_MESSAGE = "Start Thunderbird with the system\n\nVersion %v\n"
const APP_VERSION = "0.1"
const THUNDERBIRD_PATHNAME_0 = "C:\\Program Files\\Mozilla Thunderbird\\thunderbird.exe"
const THUNDERBIRD_PATHNAME_1 = "/usr/bin/thunderbird"
const STATE_FILE_PATHNAME = "./thunderbird_autostart.state.txt"
const RUN_COMMAND = "--run"

var THUNDERBIRD_PATHNAMES []string = []string{
	THUNDERBIRD_PATHNAME_0,
	THUNDERBIRD_PATHNAME_1,
}
