del /s /q /f .\*.exe
go build -ldflags "-s -w -H=windowsgui"  .
upx sc_eac_passer_light.exe --best -k -o StarCitizen_Launcher.exe