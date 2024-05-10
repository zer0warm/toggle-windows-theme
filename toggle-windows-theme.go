package main

import "golang.org/x/sys/windows/registry"

const (
	AppName = "AppsUseLightTheme"
	SysName = "SystemUsesLightTheme"
)

func main() {
	key, err := registry.OpenKey(
		registry.CURRENT_USER,
		`SOFTWARE\Microsoft\Windows\CurrentVersion\Themes\Personalize`,
		registry.QUERY_VALUE|registry.SET_VALUE,
	)
	if err != nil {
		panic(err)
	}
	defer key.Close()

	oldSysData, _, err := key.GetIntegerValue(SysName)
	if err != nil {
		panic(err)
	}

	oldAppData, _, err := key.GetIntegerValue(AppName)
	if err != nil {
		panic(err)
	}

	newSysData, newAppData := uint32(oldSysData)^1, uint32(oldAppData)^1

	err = key.SetDWordValue(SysName, newSysData)
	if err != nil {
		panic(err)
	}

	err = key.SetDWordValue(AppName, newAppData)
	if err != nil {
		panic(err)
	}
}
