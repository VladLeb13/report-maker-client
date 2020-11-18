package fill

import (
	"report-maker-client/tools"

	"github.com/VladLeb13/report-maker-lib/datalib"
	"github.com/VladLeb13/report-maker-lib/wmilib/hardware"
)

var (
	hardwareData   map[string]tools.Data
	hardwareReport datalib.Hardware
)

func Hardware() datalib.Hardware {
	getHardwareData()

	fillBoard()
	fillCPUs()
	fillHDDs()
	fillNICs()
	fillRAMs()
	fillVolumes()

	return hardwareReport

}

func getHardwareData() {
	hardwareData = map[string]tools.Data{
		"board":  new(hardware.Board),
		"cpu":    new(hardware.CPU),
		"hdd":    new(hardware.HDD),
		"nic":    new(hardware.NIC),
		"ram":    new(hardware.RAM),
		"volume": new(hardware.Volume),
	}

	for _, v := range hardwareData {
		tools.GetValues(v)
	}
}

func fillCPUs() {
	CPUs := []datalib.CPU{}
	for _, v := range hardwareData["cpu"].(*hardware.CPU).Units {
		unit := datalib.CPU{}

		unit.Caption = v.Caption
		unit.Description = v.Description
		unit.Manufacturer = v.Manufacturer
		unit.MaxClockSpeed = v.MaxClockSpeed
		unit.Name = v.Name
		unit.NumberOfCores = v.NumberOfCores
		unit.ThreadCount = v.ThreadCount

		CPUs = append(CPUs, unit)
	}
	hardwareReport.CPUs = CPUs
}

func fillBoard() {
	Board := datalib.Board{}
	for _, v := range hardwareData["board"].(*hardware.Board).Bios {
		Board.BIOS.Manufacturer = v.Manufacturer
		Board.BIOS.Name = v.Name
		Board.BIOS.Version = v.SMBIOSBIOSVersion
	}
	for _, v := range hardwareData["board"].(*hardware.Board).Сomponents {
		Board.Manufacturer = v.Manufacturer
		Board.Model = v.Model
		Board.PartNumber = v.PartNumber
		Board.Product = v.Product
		Board.SerialNumber = v.SerialNumber
		Board.Version = v.Version
	}

	hardwareReport.Board = Board
}

func fillHDDs() {
	HDDs := []datalib.HDD{}
	for _, v := range hardwareData["hdd"].(*hardware.HDD).Units {
		unit := datalib.HDD{}

		unit.InterfaceType = v.InterfaceType
		unit.Model = v.Model
		unit.Name = v.Name
		unit.Partitions = v.Partitions
		unit.SerialNumber = v.SerialNumber
		unit.Size = v.Size

		HDDs = append(HDDs, unit)
	}
	hardwareReport.HDDs = HDDs
}

func fillNICs() {
	NICs := []datalib.NIC{}
	for _, v := range hardwareData["nic"].(*hardware.NIC).Units {
		unit := datalib.NIC{}

		unit.DHCPEnabled = v.DHCPEnabled
		unit.DHCPServer = v.DHCPServer
		unit.DNSServerSearchOrder = v.DNSServerSearchOrder
		unit.DefaultIPGateway = v.DefaultIPGateway
		unit.Description = v.Description
		unit.IPAddress = v.IPAddress

		NICs = append(NICs, unit)
	}
	hardwareReport.NICs = NICs
}

func fillRAMs() {
	RAMs := []datalib.RAM{}
	for _, v := range hardwareData["ram"].(*hardware.RAM).Units {
		unit := datalib.RAM{}

		unit.Capacity = v.Capacity
		unit.DeviceLocator = v.DeviceLocator
		unit.FormFactor = v.FormFactor
		unit.Manufacturer = v.Manufacturer
		unit.Model = v.Model
		unit.Name = v.Name
		unit.OtherIdentifyingInfo = v.OtherIdentifyingInfo
		unit.PartNumber = v.PartNumber
		unit.Speed = v.Speed

		RAMs = append(RAMs, unit)
	}
	hardwareReport.RAMs = RAMs
}

func fillVolumes() {
	Volumes := []datalib.Volume{}
	for _, v := range hardwareData["volume"].(*hardware.Volume).Disks {
		unit := datalib.Volume{}

		unit.Caption = v.Caption
		unit.Description = v.Description
		unit.DeviceID = v.DeviceID
		unit.FileSystem = v.FileSystem
		unit.FreeSpace = v.FreeSpace
		unit.Name = v.Name
		unit.Size = v.Size
		unit.VolumeName = v.VolumeName

		Volumes = append(Volumes, unit)
	}
	hardwareReport.Volumes = Volumes
}
