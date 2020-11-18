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
	builder.WriteString(writeToSpan(val))
	builder.WriteString("</ul>")

	return builder.String()
}

func addShared() string {
	var builder strings.Builder

	builder.WriteString(" <table class=\"tfmt\">")
	builder.WriteString("<th>Имя</th><th>Путь</th><th>Описание</th><th>Наименование</th><th>Статус</th>")
	builder.WriteString("<tbody>")

	for _, unit := range soft.Shared {
		slice := getValue(unit)
		builder.WriteString("<tr>")
		builder.WriteString(writeToColfmt(slice))
		builder.WriteString("</tr>")
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

	for _, uint := range soft.Startup {
		slice := getValue(uint)
		builder.WriteString("<tr>")
		builder.WriteString(writeToColfmt(slice))
		builder.WriteString("</tr>")
	}

	builder.WriteString("</tbody>")
	builder.WriteString("</table>")

	return builder.String()
}

func addUpdate() string {
	var builder strings.Builder
	builder.WriteString(" <table class=\"tfmt\">")
	builder.WriteString("<th>ID</th><th>Описание</th><th>Пользователь</th><th>Дата установки</th><th>Комментарий</th><th>Статус</th><th>Наименование</th>")
	builder.WriteString("<tbody>")

	for _, unit := range soft.Updates {
		slice := getValue(unit)
		builder.WriteString("<tr>")
		builder.WriteString(writeToColfmt(slice))
		builder.WriteString("</tr>")
	}

	builder.WriteString("</tbody>")
	builder.WriteString("</table>")

	return builder.String()
}

func addPrograms() string {
	var builder strings.Builder

	builder.WriteString("<table class=\"tfmt\">")
	builder.WriteString(" <th>Наименование</th><th>Подпись</th><th>Описание</th><th>Дата установки</th><th>Каталог установки</th><th>Владелец</th><th>Компания</th><th>Производитель</th><th>Версия</th>")
	builder.WriteString("<tbody>")

	for _, unit := range soft.Programs {
		slice := getValue(unit)
		builder.WriteString("<tr>")
		builder.WriteString(writeToColfmt(slice))
		builder.WriteString("</tr>")
	}

	builder.WriteString("</tbody>")
	builder.WriteString("</table>")

	return builder.String()
}
