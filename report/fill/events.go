package fill

import (
	"report-maker-client/tools"

	"github.com/VladLeb13/report-maker-lib/datalib"
	"github.com/VladLeb13/report-maker-lib/wmilib/events"
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

		unit.LogFile = v.LogFile
		unit.Message = v.Message
		unit.TimeWritten = v.TimeWritten
		unit.User = v.User

		Events = append(Events, unit)
	}
	eventsReport.List = Events
}
