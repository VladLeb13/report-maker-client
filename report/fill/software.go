package fill

import (
	"report-maker-client/tools"

	"github.com/VladLeb13/report-maker-lib/datalib"
	"github.com/VladLeb13/report-maker-lib/wmilib/software"
)

var (
	softwareData   map[string]tools.Data
	softwareReport datalib.Software
)

func Software() datalib.Software {

	getSoftwareData()

	fillOS()
	fillPrograms()
	fillShared()
	fillStartup()
	fillUpdates()

	return softwareReport
}
func getSoftwareData() {
	softwareData = map[string]tools.Data{
		"os":       new(software.OS),
		"programs": new(software.Programs),
		"startup":  new(software.Startup),
		"updates":  new(software.Updates),
		"shared":   new(software.Shared),
	}
	for _, v := range softwareData {
		tools.GetValues(v)
	}
}

func fillOS() {
	OS := datalib.OS{}

	for _, v := range softwareData["os"].(*software.OS).WindowsInfo {

		OS.BuildNumber = v.BuildNumber
		OS.Caption = v.Caption
		OS.InstallDate = v.InstallDate
		OS.Name = v.Name
		OS.OSArchitecture = v.OSArchitecture
		OS.SerialNumber = v.SerialNumber
		OS.Version = v.Version
	}
	for _, v := range softwareData["os"].(*software.OS).SystemInfo {
		OS.UserName = v.UserName
		OS.Domain = v.Domain
	}

	softwareReport.OS = OS
}

func fillPrograms() {
	Programs := []datalib.Program{}
	for _, v := range softwareData["programs"].(*software.Programs).List {
		unit := datalib.Program{}

		unit.InstallDate = v.InstallDate
		unit.InstallLocation = v.InstallLocation
		unit.Name = v.Name
		unit.Vendor = v.Vendor
		unit.Version = v.Version

		Programs = append(Programs, unit)
	}
	softwareReport.Programs = Programs

}

func fillShared() {
	Shared := []datalib.Shared{}
	for _, v := range softwareData["shared"].(*software.Shared).Resource {
		unit := datalib.Shared{}

		unit.Caption = v.Caption
		unit.Description = v.Description
		unit.Name = v.Name
		unit.Path = v.Path
		unit.Status = v.Status

		Shared = append(Shared, unit)
	}
	softwareReport.Shared = Shared
}

func fillStartup() {
	Startup := []datalib.Startup{}
	for _, v := range softwareData["startup"].(*software.Startup).Commands {
		unit := datalib.Startup{}

		unit.Caption = v.Caption
		unit.Command = v.Command
		unit.Location = v.Location
		unit.Name = v.Name
		unit.User = v.User

		Startup = append(Startup, unit)
	}
	softwareReport.Startup = Startup

}

func fillUpdates() {
	Update := []datalib.Update{}
	for _, v := range softwareData["updates"].(*software.Updates).List {
		unit := datalib.Update{}

		unit.Description = v.Description
		unit.HotFixID = v.HotFixID
		unit.InstalledBy = v.InstalledBy
		unit.InstalledOn = v.InstalledOn

		Update = append(Update, unit)
	}
	softwareReport.Updates = Update
}
