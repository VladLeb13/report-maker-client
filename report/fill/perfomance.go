package fill

import (
	"report-maker-client/tools"

	"github.com/VladLeb13/report-maker-lib/datalib"
	"github.com/VladLeb13/report-maker-lib/wmilib/perfomance"
)

var (
	perfomanceData   map[string]tools.Data
	perfomanceReport datalib.Perfomance
)

func Perfomance() datalib.Perfomance {
	getPerfomanceData()

	fillCPU()
	fillHDD()
	fillRAM()
	fillProcess()

	return perfomanceReport

}

func getPerfomanceData() {
	perfomanceData = map[string]tools.Data{
		"cpu":     new(perfomance.CPU),
		"hdd":     new(perfomance.HDD),
		"ram":     new(perfomance.RAM),
		"process": new(perfomance.Process),
	}
	for _, v := range perfomanceData {
		tools.GetValues(v)
	}
}

func fillCPU() {
	Perf_cpu := []datalib.Perf_cpu{}
	for _, v := range perfomanceData["cpu"].(*perfomance.CPU).Data {
		data := datalib.Perf_cpu{}

		data.PercentProcessorPerformance = v.PercentProcessorPerformance
		data.PercentProcessorUtility = v.PercentProcessorUtility
		data.ProcessorFrequency = v.ProcessorFrequency

		Perf_cpu = append(Perf_cpu, data)
	}
	perfomanceReport.CPU = Perf_cpu
}

func fillHDD() {
	Perf_hdd := []datalib.Perf_hdd{}
	for _, v := range perfomanceData["hdd"].(*perfomance.HDD).Data {
		data := datalib.Perf_hdd{}

		data.AvgDiskQueueLength = v.AvgDiskQueueLength
		data.AvgDisksecPerTransfer = v.AvgDisksecPerTransfer
		data.DiskReadBytesPersec = v.DiskReadBytesPersec
		data.DiskWriteBytesPersec = v.DiskWriteBytesPersec
		data.PercentDiskTime = v.PercentDiskTime

		Perf_hdd = append(Perf_hdd, data)
	}
	perfomanceReport.HDD = Perf_hdd
}

func fillProcess() {
	Perf_process := []datalib.Perf_process{}
	for _, v := range perfomanceData["process"].(*perfomance.Process).Data {
		data := datalib.Perf_process{}

		data.IDProcess = v.IDProcess
		data.IOReadOperationsPersec = v.IOReadOperationsPersec
		data.IOWriteOperationsPersec = v.IOWriteOperationsPersec
		data.Name = v.Name
		data.PercentProcessorTime = v.PercentProcessorTime
		data.ThreadCount = v.ThreadCount
		data.VirtualBytesPeak = v.VirtualBytesPeak

		Perf_process = append(Perf_process, data)
	}
	perfomanceReport.Processes = Perf_process
}

func fillRAM() {

	Perf_ram := []datalib.Perf_ram{}
	for _, v := range perfomanceData["ram"].(*perfomance.RAM).Data {
		data := datalib.Perf_ram{}

		data.AvailableMBytes = v.AvailableMBytes
		data.PercentCommittedBytesInUse = v.PercentCommittedBytesInUse

		Perf_ram = append(Perf_ram, data)
	}
	perfomanceReport.RAM = Perf_ram
}
