package converting

import (
	"context"
	"report/converting/builder"
	"strings"
	"tools"
)

var (
	HTMLreport strings.Builder
)

func HTML(ctx *tools.AppContex) {

	openHead()
	addStyle()
	closeHead()
	openBody()

	addSoftware(ctx)
	addHardware(ctx)
	addEvent(ctx)
	addPerfomance(ctx)

	closeBody()

	ctx.Context = context.WithValue(ctx.Context, "HTMLReport", HTMLreport.String())
}

func openHead() {
	head := "<html><head><META http-equiv=\"Content-Type\" content=\"text/html; charset=UTF-8\">"
	HTMLreport.WriteString(head)
}

func addStyle() {
	style := `
	<style type="text/css">
			.lw { font-size: 60px; }
			p.section { 
				background-color:#2c54e5;
				font-size: 25;
				height: 50;
				width: 100%;}
				padding: 15px 15px;
				padding-top: 15px 15px;
				padding-bottom: 3px;
				ul.list  {transition: .4s linear;}
				ul.list {
				background: white;
				list-style-type: circle;
				list-style-position: inside;
				padding: 0 10px;
				margin: 0;
				width: 48%;
			}
		ul.list li {
			font-family: "Trebuchet MS", "Lucida Sans";
			border-bottom: 1px solid #efefef;
			padding: 5px 0;
		}
			
		ul.list a {
			text-decoration: none;
			color: black;
			font-size: 10px;
		}
			
		ul.list li span {
			float: right;
			display: inline-block;
			border: 1px solid #efefef;
			padding: 0 5px;
			font-size: 10px;
			color: #000000;
		}
			
		ul.list li:hover a {color: #c93961;}
		ul.list li:hover span {
			color: #c93961;
			border: 1px solid #c93961;
		}  
			
		p {
			color:  black;
			font-size: 20;
			font-family: "Trebuchet MS", "Lucida Sans";
			background-color: #6da2de;
			width: 100%; 
		}
			
		table {
			font-family: "Trebuchet MS", "Lucida Sans";
			font-size: 10px;
			background: white;
			max-width: 100%;
			width: 100%;
			border-collapse: collapse;
			text-align: left;
		}
			
		th {
			font-weight: normal;
			font-size: 10px;
			color: #039;
			word-break: break-all;
			border-bottom: 2px solid #6678b1;
			padding: 8px 8px;
		}
			
		td {
			border-bottom: 1px solid #ccc;
			color: #669;
			padding: 2px 5px;
			transition: .3s linear;
			word-break: break-all;
		}
			
		h3 {
			font-size:10px;
			font-family: "Trebuchet MS", "Lucida Sans";
		}
			tr:hover td {color: #6699ff;}
			div {max-width:100%}
	</style>
	`
	HTMLreport.WriteString(style)
}

func closeHead() {
	HTMLreport.WriteString("</head>")
}

func openBody() {
	HTMLreport.WriteString("<body> <div>")
}
func closeBody() {
	HTMLreport.WriteString("</div></body>")
}

func addSoftware(ctx *tools.AppContex) {
	report_str := builder.Software(ctx)
	HTMLreport.WriteString(report_str)
}

func addHardware(ctx *tools.AppContex) {
	report_str := builder.Hardware(ctx)
	HTMLreport.WriteString(report_str)
}

func addEvent(ctx *tools.AppContex) {
	report_str := builder.Event(ctx)
	HTMLreport.WriteString(report_str)
}

func addPerfomance(ctx *tools.AppContex) {
	report_str := builder.Perfomance(ctx)
	HTMLreport.WriteString(report_str)
}
