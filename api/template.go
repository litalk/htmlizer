package api

import (
	"html/template"
	"io"
	"strconv"
)

func Template() *template.Template {
	fmap := template.FuncMap{
		"formateFloatStr": FormateFloatStr,
		"isOdd":           IsOdd,
		"resultSymbol":    ResultSymbol,
	}

	t := template.New("test").Funcs(fmap)
	t = template.Must(t.ParseFiles("tmpl/summary.html"))
	t = template.Must(t.ParseFiles("tmpl/packages.html"))
	t = template.Must(t.ParseFiles("tmpl/testcase.html"))
	t = template.Must(t.ParseFiles("tmpl/failure_detail.html"))
	t = template.Must(t.ParseFiles("tmpl/maven_base.css"))
	t = template.Must(t.ParseFiles("tmpl/maven_theme.css"))
	t = template.Must(t.ParseFiles("tmpl/site.css"))
	t = template.Must(t.ParseFiles("tmpl/print.css"))
	t = template.Must(t.ParseFiles("tmpl/main.js"))
	t = template.Must(t.ParseFiles("tmpl/main.html"))

	return t
}

func ApplyTemplate(t *template.Template, wr io.Writer, report *Report) error {
	return t.ExecuteTemplate(wr, "main.html", report)
}

func FormateFloatStr(f float64) string {
	return strconv.FormatFloat(f, 'f', 3, 32)
}

func IsOdd(index int) bool {
	return index%2 == 1
}

func ResultSymbol(c collection) template.HTML {
	color := "#6ec258"
	result := c.Result()
	if result == ERROR {
		color = "#E4443B"
	} else if result == FAILURE || result == SKIPPED {
		color = "F39F3D"
	}
	return template.HTML(`<span style="width:13px;height:13px;display:inline-block;border-radius:13px;background-color:` + color + `;border:0px solid black;margin-left:1px;margin-right:1px"></span>`)
}

func tableTrClass(index int) string {
	if IsOdd(index) {
		return "a"
	}
	return "b"
}
