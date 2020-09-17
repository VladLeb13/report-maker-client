package report

import (
	"fmt"

	"github.com/VladLeb13/report-maker-lib/src/datalib"
	"github.com/VladLeb13/report-maker-lib/src/wmilib/events"
	"github.com/VladLeb13/report-maker-lib/src/wmilib/hardware"
	"github.com/VladLeb13/report-maker-lib/src/wmilib/perfomance"
	"github.com/VladLeb13/report-maker-lib/src/wmilib/software"
)

type data interface {
	Get()
}

var (
	hardwareData   map[string]data
	softwareData   map[string]data
	perfomanceData map[string]data
	eventsData     data
)

func Create() {

	//getHardwareData()
	//hardareReport := new(datalib.Hardware)
	//	fillHardware(hardareReport.)
	//	fmt.Println(report)

	// getSoftwareData()
	// softwareReport := new(datalib.Software)
	// fillSoftware(softwareReport)
	// fmt.Println(softwareReport)

	// getPerfomanceData()
	// perfomanceReport := new(datalib.Perfomance)
	// fillPerfomance(perfomanceReport)
	// fmt.Println(perfomanceReport)

	getEventsData()
	eventReport := new(datalib.Events)
	fillEvents(eventReport)
	fmt.Println(eventReport)

}

func fillHardware(report *datalib.Hardware) {

	cpu := []datalib.CPU{}
	for _, v := range hardwareData["cpu"].(*hardware.CPU).Units {
		cpuUnit := datalib.CPU{}
		cpuUnit.Caption = v.Caption
		cpuUnit.Description = v.Description
		cpuUnit.Manufacturer = v.Manufacturer
		cpuUnit.MaxClockSpeed = v.MaxClockSpeed
		cpuUnit.Name = v.Name
		cpuUnit.NumberOfCores = v.NumberOfCores
		cpuUnit.ThreadCount = v.ThreadCount
		cpu = append(cpu, cpuUnit)
	}
	report.CPU = cpu

	board := datalib.Board{}
	for _, v := range hardwareData["board"].(*hardware.Board).Bios {
		board.BIOS.Manufacturer = v.Manufacturer
		board.BIOS.Name = v.Name
		board.BIOS.Version = v.SMBIOSBIOSVersion
	}
	for _, v := range hardwareData["board"].(*hardware.Board).Сomponents {
		board.Manufacturer = v.Manufacturer
		board.Model = v.Model
		board.PartNumber = v.PartNumber
		board.Product = v.Product
		board.SerialNumber = v.SerialNumber
		board.Version = v.Version
	}

	hdd := []datalib.HDD{}
	for _, v := range hardwareData["hdd"].(*hardware.HDD).Units {
		unit := datalib.HDD{}
		unit.InterfaceType = v.InterfaceType
		unit.Model = v.Model
		unit.Name = v.Name
		unit.Partitions = v.Partitions
		unit.SerialNumber = v.SerialNumber
		unit.Size = v.Size

		hdd = append(hdd, unit)
	}
	report.HDDs = hdd

	nic := []datalib.NIC{}
	for _, v := range hardwareData["nic"].(*hardware.NIC).Units {
		unit := datalib.NIC{}
		unit.DHCPEnabled = v.DHCPEnabled
		unit.DHCPServer = v.DHCPServer
		unit.DNSServerSearchOrder = v.DNSServerSearchOrder
		unit.DefaultIPGateway = v.DefaultIPGateway
		unit.Description = v.Description
		unit.IPAddress = v.IPAddress
		nic = append(nic, unit)
	}
	report.NICs = nic

	ram := []datalib.RAM{}
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

		ram = append(ram, unit)
	}
	report.RAMs = ram

	volume := []datalib.Volume{}
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
		volume = append(volume, unit)
	}
	report.Volumes = volume
}

func fillSoftware(report *datalib.Software) {
	os := datalib.OS{}
	for _, v := range softwareData["os"].(*software.OS).WindowsInfo {

		os.BuildNumber = v.BuildNumber
		os.Caption = v.Caption
		os.InstallDate = v.InstallDate
		os.Name = v.Name
		os.OSArchitecture = v.OSArchitecture
		os.SerialNumber = v.SerialNumber
		os.Version = v.Version
	}
	for _, v := range softwareData["os"].(*software.OS).SystemInfo {
		os.UserName = v.UserName
		os.Domain = v.Domain
	}
	report.OS = os

	programs := []datalib.Program{}
	for _, v := range softwareData["programs"].(*software.Programs).List {
		Unit := datalib.Program{}
		Unit.Caption = v.Caption
		Unit.Description = v.Description
		Unit.InstallDate = v.InstallDate
		Unit.InstallLocation = v.InstallLocation
		Unit.Name = v.Name
		Unit.RegCompany = v.RegCompany
		Unit.RegOwner = v.RegOwner
		Unit.Vendor = v.Vendor
		Unit.Version = v.Version
		programs = append(programs, Unit)
	}
	report.Programs = programs

	shared := []datalib.Shared{}
	for _, v := range softwareData["shared"].(*software.Shared).Resource {
		Unit := datalib.Shared{}
		Unit.Caption = v.Caption
		Unit.Description = v.Description
		Unit.Name = v.Name
		Unit.Path = v.Path
		Unit.Status = v.Status
		shared = append(shared, Unit)
	}
	report.Shared = shared

	startup := []datalib.Startup{}
	for _, v := range softwareData["startup"].(*software.Startup).Commands {
		Unit := datalib.Startup{}
		Unit.Caption = v.Caption
		Unit.Command = v.Command
		Unit.Location = v.Location
		Unit.Name = v.Name
		Unit.User = v.User
		startup = append(startup, Unit)
	}
	report.Startup = startup

	update := []datalib.Update{}
	for _, v := range softwareData["updates"].(*software.Updates).List {
		Unit := datalib.Update{}
		Unit.Description = v.Description
		Unit.FixComments = v.FixComments
		Unit.HotFixID = v.HotFixID
		Unit.InstalledBy = v.InstalledBy
		Unit.InstalledOn = v.InstalledOn
		Unit.Name = v.Name
		Unit.Status = v.Status
		update = append(update, Unit)
	}
	report.Updates = update

}

func fillPerfomance(report *datalib.Perfomance) {
	perf_cpu := []datalib.Perf_cpu{}
	for _, v := range perfomanceData["cpu"].(*perfomance.CPU).Data {
		data := datalib.Perf_cpu{}
		data.PercentProcessorPerformance = v.PercentProcessorPerformance
		data.PercentProcessorUtility = v.PercentProcessorUtility
		data.ProcessorFrequency = v.ProcessorFrequency
		perf_cpu = append(perf_cpu, data)
	}
	report.CPU = perf_cpu

	perf_hdd := []datalib.Perf_hdd{}
	for _, v := range perfomanceData["hdd"].(*perfomance.HDD).Data {
		data := datalib.Perf_hdd{}
		data.AvgDiskQueueLength = v.AvgDiskQueueLength
		data.AvgDisksecPerTransfer = v.AvgDisksecPerTransfer
		data.DiskReadBytesPersec = v.DiskReadBytesPersec
		data.DiskWriteBytesPersec = v.DiskWriteBytesPersec
		data.PercentDiskTime = v.PercentDiskTime
		perf_hdd = append(perf_hdd, data)
	}
	report.HDD = perf_hdd

	perf_process := []datalib.Perf_process{}
	for _, v := range perfomanceData["process"].(*perfomance.Process).Data {
		data := datalib.Perf_process{}
		data.IDProcess = v.IDProcess
		data.IOReadOperationsPersec = v.IOReadOperationsPersec
		data.IOWriteOperationsPersec = v.IOWriteOperationsPersec
		data.Name = v.Name
		data.PercentProcessorTime = v.PercentProcessorTime
		data.ThreadCount = v.ThreadCount
		data.VirtualBytesPeak = v.VirtualBytesPeak
		perf_process = append(perf_process, data)
	}
	report.Processes = perf_process

	perf_ram := []datalib.Perf_ram{}
	for _, v := range perfomanceData["ram"].(*perfomance.RAM).Data {
		data := datalib.Perf_ram{}
		data.AvailableMBytes = v.AvailableMBytes
		data.PercentCommittedBytesInUse = v.PercentCommittedBytesInUse
		perf_ram = append(perf_ram, data)
	}
	report.RAM = perf_ram

}

func fillEvents(report *datalib.Events) {
	evnts := []datalib.Event{}
	for _, v := range eventsData.(*events.Events).List {
		Unit := datalib.Event{}
		Unit.CategoryString = v.CategoryString
		Unit.ComputerName = v.ComputerName
		Unit.Data = v.Data
		Unit.LogFile = v.LogFile
		Unit.Message = v.Message
		Unit.SourceName = v.SourceName
		Unit.TimeWritten = v.TimeWritten
		Unit.User = v.User
		evnts = append(evnts, Unit)
	}
	report.List = evnts
}

func getHardwareData() {
	hardwareData = map[string]data{
		"board":  new(hardware.Board),
		"cpu":    new(hardware.CPU),
		"hdd":    new(hardware.HDD),
		"nic":    new(hardware.NIC),
		"ram":    new(hardware.RAM),
		"volume": new(hardware.Volume),
	}

	for _, v := range hardwareData {
		getValues(v)
	}
}

func getSoftwareData() {
	softwareData = map[string]data{
		"os":       new(software.OS),
		"programs": new(software.Programs),
		"shared":   new(software.Shared),
		"startup":  new(software.Startup),
		"updates":  new(software.Updates),
	}
	for _, v := range softwareData {
		getValues(v)
	}
}

func getPerfomanceData() {
	perfomanceData = map[string]data{
		"cpu":     new(perfomance.CPU),
		"hdd":     new(perfomance.HDD),
		"ram":     new(perfomance.RAM),
		"process": new(perfomance.Process),
	}
	for _, v := range perfomanceData {
		getValues(v)
	}
}

func getEventsData() {
	eventsData = new(events.Events)
	getValues(eventsData)
}

func getValues(src data) {
	src.Get()

}
