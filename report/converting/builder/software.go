package builder

import (
	"report-maker-client/tools"
	"strings"

	"github.com/VladLeb13/report-maker-lib/datalib"
)

var (
	soft datalib.Software
)

func Software(ctx *tools.AppContex) string {
	report := ctx.Context.Value("report").(datalib.Report)
	soft = report.Software
	var builder strings.Builder

	builder.WriteString("<p class=\"section\">Программное обеспечение</p>")

	builder.WriteString("<p>Операционная система:</p>")
	builder.WriteString(addOS())

	builder.WriteString("<p>Общие ресурсы:</p>")
	builder.WriteString(addShared())

	builder.WriteString("<p>Автозагрузка:</p>")
	builder.WriteString(addStartup())

	builder.WriteString("<p>Обновления ОС:</p>")
	builder.WriteString(addUpdate())

	builder.WriteString("<p>Установленные программы:</p>")
	builder.WriteString(addPrograms())

	return builder.String()
}

func addOS() string {
	var builder strings.Builder

	builder.WriteString("<ul class=\"list\">")

	val := getValue(soft.OS)
	lines := val.writeInLine()
	builder.WriteString(lines)

	builder.WriteString("</ul>")

	return builder.String()
}

func addShared() string {
	var builder strings.Builder

	builder.WriteString(" <table class=\"tfmt\">")
	builder.WriteString("<th>Имя</th><th>Путь</th><th>Описание</th><th>Наименование</th><th>Статус</th>")
	builder.WriteString("<tbody>")

	for _, unit := range soft.Shared {
		val := getValue(unit)
		columns := val.writeInCol()
		builder.WriteString(columns)
	}
	builder.WriteString("</tbody>")
	builder.WriteString("</table>")

	return builder.String()
}

func addStartup() string {
	var builder strings.Builder
	builder.WriteString(" <table class=\"tfmt\">")
	builder.WriteString("<th>Название</th><th>Команда</th><th>Пользователь</th><th>Наименование</th><th>Расположение</th>")
	builder.WriteString("<tbody>")

	for _, unit := range soft.Startup {
		val := getValue(unit)
		columns := val.writeInCol()
		builder.WriteString(columns)
	}

	builder.WriteString("</tbody>")
	builder.WriteString("</table>")

	return builder.String()
}

func addUpdate() string {
	var builder strings.Builder
	builder.WriteString(" <table class=\"tfmt\">")
	builder.WriteString("<th>ID</th><th>Описание</th><th>Пользователь</th><th>Дата установки</th>")
	builder.WriteString("<tbody>")

	for _, unit := range soft.Updates {
		val := getValue(unit)
		columns := val.writeInCol()
		builder.WriteString(columns)
	}

	builder.WriteString("</tbody>")
	builder.WriteString("</table>")

	return builder.String()
}

func addPrograms() string {
	var builder strings.Builder

	builder.WriteString("<table class=\"tfmt\">")
	builder.WriteString(" <th>Наименование</th><th>Дата установки</th><th>Каталог установки</th><th>Производитель</th><th>Версия</th>")
	builder.WriteString("<tbody>")

	for _, unit := range soft.Programs {
		val := getValue(unit)
		columns := val.writeInCol()
		builder.WriteString(columns)
	}

	builder.WriteString("</tbody>")
	builder.WriteString("</table>")

	return builder.String()
}
