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

	builder.WriteString(" <table class=\"tfmt\">")
	builder.WriteString(" <th>Пользователь</th><th>Файл лога</th><th>Сообщение</th><th>Категория</th><th>Имя ПК</th><th>Источник</th><th>Время записи в журнал</th><th>Пользователь</th>")
	builder.WriteString("<tbody>")

	for _, unit := range event.List {
		slice := getValue(unit)
		builder.WriteString("<tr>")
		builder.WriteString(writeToColfmt(slice))
		builder.WriteString("</tr>")
	}

	builder.WriteString("</tbody>")
	builder.WriteString("</table>")

	return builder.String()
}