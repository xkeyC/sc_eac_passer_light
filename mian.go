package main

import (
	"github.com/sqweek/dialog"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	defer recoverError()
	if !isLaunchedInGamePath() {
		dialog.Message("Please put this program in the game LIVE folder!").Error()
		return
	}
	if !isInstalledScModding() {
		dialog.Message("Please make sure sc_eac_passer is installed correctly").Error()
		return
	}
	if len(os.Args) < 5 {
		dialog.Message("Please use RSI Launcher to launch the game").Error()
		return
	}
	dialog.Message("SC EAC PASSER form StarCitizen Chinese Community,it will disable EAC in order to load SC Modding Tools,Click 'ok' to start the game.\n" +
		"")
	startGame()
}

func startGame() {
	runtime.GC()

	start := time.Now()
	var newArgs []string
	newArgs = append(newArgs, os.Args[1:]...)

	_, gameExecErr := exec.Command("Bin64/StarCitizen.exe", newArgs...).Output()
	if gameExecErr != nil {
		// 暂且通过时长来区分游戏异常状态
		if time.Since(start) > time.Minute {
			dialog.Message("Game launch fail").Error()
		} else {
			dialog.Message("Game crashes").Error()
		}
		_ = exec.Command(`explorer`, `/select,`, `Game.log`).Start()
	}

}

func recoverError() {
	if e := recover(); e != nil {
		if r, ok := e.(error); ok && r != nil {
			dialog.Message("Game crashes").Error()
			panic(e)
		}
	}
}

func isLaunchedInGamePath() bool {
	return fileExists("Bin64/StarCitizen.exe")
}

func isInstalledScModding() bool {
	return (fileExists("Bin64/dbghelp.dll")) && fileExists("data/config.xml")
}
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
