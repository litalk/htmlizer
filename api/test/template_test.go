package api_test

import (
	"bytes"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/wu8685/htmlizer.git/api"
)

var _ = Describe("Template", func() {

	It("should support parse to summary template", func() {
		html, err := executeTemplate("summary.html", report)
		Expect(err).Should(BeNil())
		applied(html, "57")
		applied(html, "1")
		applied(html, "1")
		applied(html, "2")
		applied(html, "92.982%")
		applied(html, "540.554<")
	})

	It("should support parse to package template", func() {
		html, err := executeTemplate("packages.html", report)
		Expect(err).Should(BeNil())
		
		applied(html, "com.vip.paas.e2e.test.ds.cpms")
		applied(html, "com.vip.paas.e2e.test.ds.auth")
		applied(html, "com.vip.paas.e2e.test.ci.cpms")
		
		applied(html, ">GroupTest<")
	})
	
	It("should support parse to testcase template", func() {
		html, err := executeTemplate("testcase.html", report)
		Expect(err).Should(BeNil())
		
		applied(html, "java.lang.NullPointerException")
	})
	
	It("should support parse to failure detail template", func() {
		html, err := executeTemplate("failure_detail.html", report)
		Expect(err).Should(BeNil())
		
		applied(html, "java.lang.NullPointerException")
	})
})

func executeTemplate(file string, report *api.Report) (summary string, err error) {
	buff := bytes.NewBufferString("")
	err = t.ExecuteTemplate(buff, file, report)
	summary = buff.String()
	return
}

func applied(html, key string) {
	Expect(strings.Index(html, key)).ShouldNot(Equal(-1))
}
