package api_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"

	"html/template"
	"testing"

	"github.com/wu8685/htmlizer.git/api"
)

var (
	testSuite *api.TestSuite
	report    *api.Report
	t         *template.Template
)

func TestApi(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "api Suite", []Reporter{junitReporter})
}

var _ = BeforeSuite(func() {
	testSuiteTmp, err := api.ParseJunitXML("sample_junit.xml")
	testSuite = testSuiteTmp
	Expect(err).Should(BeNil())

	report = api.Aggregate(testSuite)

	fmap := template.FuncMap{
		"formateFloatStr": api.FormateFloatStr,
		"isOdd":           api.IsOdd,
		"resultSymbol":    api.ResultSymbol,
	}

	t = template.New("test").Funcs(fmap)
	t = template.Must(t.ParseFiles("../../tmpl/summary.html"))
	t = template.Must(t.ParseFiles("../../tmpl/packages.html"))
	t = template.Must(t.ParseFiles("../../tmpl/testcase.html"))
	t = template.Must(t.ParseFiles("../../tmpl/failure_detail.html"))
})
