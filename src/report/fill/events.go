package fill

import (
	"tools"

	"github.com/VladLeb13/report-maker-lib/src/datalib"
	"github.com/VladLeb13/report-maker-lib/src/wmilib/events"
)

var (
	eventsData   tools.Data
	eventsReport datalib.Events
)

func Events() datalib.Events {
	getEventsData()

	fillEvents()

	return eventsReport
}

func getEventsData() {
	eventsData = new(events.Events)
	tools.GetValues(eventsData)
}

func fillEvents() {
	Events := []datalib.Event{}
	for _, v := range eventsData.(*events.Events).List {
		unit := datalib.Event{}

		unit.CategoryString = v.CategoryString
		unit.ComputerName = v.ComputerName
		unit.Data = v.Data
		unit.LogFile = v.LogFile
		unit.Message = v.Message
		unit.SourceName = v.SourceName
		unit.TimeWritten = v.TimeWritten
		unit.User = v.User

		Events = append(Events, unit)
	}
	eventsReport.List = Events
}
