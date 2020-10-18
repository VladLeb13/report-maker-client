package builder

import (
	"strings"
	"tools"

	"github.com/VladLeb13/report-maker-lib/src/datalib"
)

var (
	hardware datalib.Hardware
)

func Hardware(ctx *tools.AppContex) string {
	report := ctx.Context.Value("report").(datalib.Report)
	hardware = report.Hardware
	var builder strings.Builder

	builder.WriteString("<p class=\"section\">Аппаратное обеспечение</p>")

	builder.WriteString("<p>Процессор:</p>")
	builder.WriteString(addCPUs())

	builder.WriteString("<p>Материнская плата:</p>")
	builder.WriteString(addBoard())

	builder.WriteString("<p>ОЗУ:</p>")
	builder.WriteString(addRAMs())

	builder.WriteString("<p>Запоминающие устройства</p>")
	builder.WriteString(addHDDs())

	builder.WriteString("<p>Логические диски</p>")
	builder.WriteString(addVolumes())

	builder.WriteString("<p>Сетевые устройства</p>")
	builder.WriteString(addNICs())

	return builder.String()
}

func addCPUs() string {
	var builder strings.Builder

	builder.WriteString("<ul class=\"list\">")

	for _, elem := range hardware.CPUs {
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

func addBoard() string {
	var builder strings.Builder

	builder.WriteString("<ul class=\"list\">")

	val := getValue(hardware.Board)
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
	builder.WriteString("</ul>")

	return builder.String()
}

func addRAMs() string {
	var builder strings.Builder

	builder.WriteString(" <table class=\"tfmt\">")
	builder.WriteString(" <th>Объем</th><th>Скорость</th><th>Шина</th><th>Номер</th><th>Призводитель</th><th>Модель</th><th>Наименование</th><th>Форм фактор</th><th>Прочее</th>")
	builder.WriteString("<tbody>")

	for _, unit := range hardware.RAMs {
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

func addHDDs() string {
	var builder strings.Builder

	builder.WriteString(" <table class=\"tfmt\">")
	builder.WriteString("<th>Серийный номер</th><th>Тип интерфейса</th><th>Объем</th><th>Количество разделов</th><th>Модель</th><th>Название</th>")
	builder.WriteString("<tbody>")

	for _, unit := range hardware.HDDs {
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

func addVolumes() string {
	var builder strings.Builder

	builder.WriteString(" <table class=\"tfmt\">")
	builder.WriteString("<th>Метка тома</th><th>Описание</th><th>ID</th><th>Файловая система</th><th>Свободное место</th><th>Название</th><th>Объем</th><th>Название тома</th>")
	builder.WriteString("<tbody>")

	for _, unit := range hardware.Volumes {
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

func addNICs() string {
	var builder strings.Builder

	builder.WriteString(" <table class=\"tfmt\">")
	builder.WriteString("<th>Описание</th><th>Активность DHCP</th><th>DHCP сервер</th><th>IP</th><th>Шлюз</th><th>DNS</th>")
	builder.WriteString("<tbody>")

	for _, unit := range hardware.NICs {
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
