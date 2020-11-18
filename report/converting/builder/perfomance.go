package builder

import (
	"report-maker-client/tools"
	"strings"

	"github.com/VladLeb13/report-maker-lib/datalib"
)

var (
	perf datalib.Perfomance
)

func Perfomance(ctx *tools.AppContex) string {
	report := ctx.Context.Value("report").(datalib.Report)
	perf = report.Perfomance
	var builder strings.Builder

	builder.WriteString("<p class=\"section\">Производительность</p>")

	builder.WriteString("<p>Средние показатели CPU</p>")
	builder.WriteString(addPerfomanceCPU())

	builder.WriteString("<p>Средние показатели RAM</p>")
	builder.WriteString(addPerfomanceRAM())

	builder.WriteString("<p>Средние показатели HDD</p>")
	builder.WriteString(addPerfomanceHDD())

	builder.WriteString("<p>Средние показатели по процессам</p>")
	builder.WriteString(addPerfomanceProcess())

	return builder.String()
}

func addPerfomanceCPU() string {
	var builder strings.Builder

	builder.WriteString("<ul class=\"list\">")

	for _, elem := range perf.CPU {
		val := getValue(elem)
		for _, m := range val {
			for key, value := range m {
				builder.WriteString("<li>")
				builder.WriteString("<a href=\"\">")
				builder.WriteString(key)
				builder.WriteString(":</a>")
				builder.WriteString("<span>")
				builder.WriteString(value)
				builder.WriteString("</span>")
				builder.WriteString("</li>")
			}
		}
	}
	builder.WriteString("</ul>")

	return builder.String()
}

func addPerfomanceRAM() string {
	var builder strings.Builder

	builder.WriteString("<ul class=\"list\">")

	for _, elem := range perf.RAM {
		val := getValue(elem)
		for _, m := range val {
			for key, value := range m {
				builder.WriteString("<li>")
				builder.WriteString("<a href=\"\">")
				builder.WriteString(key)
				builder.WriteString(":</a>")
				builder.WriteString("<span>")
				builder.WriteString(value)
				builder.WriteString("</span>")
				builder.WriteString("</li>")
			}
		}
	}
	builder.WriteString("</ul>")

	return builder.String()

}

func addPerfomanceHDD() string {
	var builder strings.Builder

	builder.WriteString("<ul class=\"list\">")
	for _, elem := range perf.HDD {
		val := getValue(elem)
		for _, m := range val {
			for key, value := range m {
				builder.WriteString("<li>")
				builder.WriteString("<a href=\"\">")
				builder.WriteString(key)
				builder.WriteString(":</a>")
				builder.WriteString("<span>")
				builder.WriteString(value)
				builder.WriteString("</span>")
				builder.WriteString("</li>")
			}
		}
	}
	builder.WriteString("</ul>")

	return builder.String()
}

func addPerfomanceProcess() string {
	var builder strings.Builder
	builder.WriteString(" <table class=\"tfmt\">")
	builder.WriteString(`<th>Операций на чтение в секунду</th>
						<th>Операций записи в секунду</th>
						<th>Название</th>
						<th>Использование процессора</th>
						<th>Использование потоков</th>
						<th>Использования виртуальной памяти</th>
						<th>ID процесса</th>`)
	builder.WriteString("<tbody>")

	for _, unit := range perf.Processes {
		slice := getValue(unit)
		builder.WriteString("<tr>")
		for _, elem := range slice {

			for _, value := range elem {
				builder.WriteString("<td class=\"colfmt\">")
				builder.WriteString(value)
				builder.WriteString("</td>")
			}

		}
		builder.WriteString("</tr>")
	}

	builder.WriteString("</tbody>")
	builder.WriteString("</table>")

	return builder.String()
}
