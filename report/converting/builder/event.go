package builder

import (
	"report-maker-client/tools"
	"strings"

	"github.com/VladLeb13/report-maker-lib/datalib"
)

var (
	event datalib.Events
)

func Event(ctx *tools.AppContex) string {
	report := ctx.Context.Value("report").(datalib.Report)
	event = report.Events
	var builder strings.Builder

	builder.WriteString("<p class=\"section\">Отчет об ошибках</p>")
	builder.WriteString(addEvents())

	return builder.String()
}

func addEvents() string {
	var builder strings.Builder

	builder.WriteString(" <table class=\"tfmt2\">")
	builder.WriteString("<thead class=\"tfmt3\"><tr> <td>Пользователь</td><td>Файл лога</td><td class=\"eventMsg\" width=\"400px\">Сообщение</td><td>Время записи в журнал</td></tr></thead>")
	builder.WriteString("<tbody>")

	for _, unit := range event.List {
		val := getValue(unit)
		columns := val.writeInCol()
		builder.WriteString(columns)
	}

	builder.WriteString("</tbody>")
	builder.WriteString("</table>")

	return builder.String()
}
