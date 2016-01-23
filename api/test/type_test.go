package api_test

import (
	"github.com/wu8685/htmlizer.git/api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Type", func() {

	It("should support parsing xml file", func() {
		Expect(len(testSuite.Properties)).Should(Equal(58))
		Expect(testSuite.Properties[0].Name).Should(Equal("java.runtime.name"))
		Expect(testSuite.Properties[0].Value).Should(Equal("Java(TM) SE Runtime Environment"))

		Expect(len(testSuite.TestCases)).Should(Equal(57))
		Expect(api.FormateFloatStr(testSuite.TestCases[0].Time)).Should(Equal("3.325"))
		Expect(testSuite.TestCases[0].ClassName).Should(Equal("com.vip.paas.e2e.test.ci.cpms.GroupTest"))
		Expect(testSuite.TestCases[0].Name).Should(Equal("create"))

		Expect(testSuite.TestCases[len(testSuite.TestCases) - 6].Failure).ShouldNot(BeNil())
		failure := testSuite.TestCases[len(testSuite.TestCases) - 6].Failure
		Expect(failure.Message).Should(Equal("expected:<true> but was:<false>"))
		Expect(failure.Type).Should(Equal("java.lang.AssertionError"))
		Expect(len(failure.StackTrace)).ShouldNot(Equal(0))

		Expect(testSuite.TestCases[len(testSuite.TestCases) - 6].SystemOut).ShouldNot(BeNil())
		Expect(len(testSuite.TestCases[len(testSuite.TestCases) - 6].SystemOut.Log)).ShouldNot(Equal(0))
	})

	It("should support some method", func() {
		Expect(api.SuccessRate(testSuite.Tests, testSuite.Failures, testSuite.Errors, testSuite.Skipped)).Should(Equal("92.982%"))
	})

	It("should support aggregation", func() {
		Expect(report.Name).Should(Equal("com.vip.paas.e2e.suite.All"))
		Expect(api.FormateFloatStr(report.Time)).Should(Equal("540.554"))
		Expect(report.Tests).Should(Equal(57))
		Expect(report.Errors).Should(Equal(1))
		Expect(report.Failures).Should(Equal(1))
		Expect(report.Skipped).Should(Equal(2))
		Expect(len(report.Packages)).Should(Equal(3))
		
		var pkg *api.Package
		for _, p := range report.Packages {
			if p.Name == "com.vip.paas.e2e.test.ds.auth" {
				pkg = p
				break
			}
		}
		
		Expect(pkg.Name).Should(Equal("com.vip.paas.e2e.test.ds.auth"))
		Expect(api.FormateFloatStr(pkg.Time)).Should(Equal("11.566"))
		Expect(pkg.Tests).Should(Equal(16))
		Expect(pkg.Errors).Should(Equal(1))
		Expect(pkg.Failures).Should(Equal(1))
		Expect(pkg.Skipped).Should(Equal(2))
		Expect(len(pkg.Classes)).Should(Equal(4))
		
		var class *api.Class
		for _, c := range pkg.Classes {
			if c.Name == "com.vip.paas.e2e.test.ds.auth.OrgSpaceTest" {
				class = c
				break
			}
		}
		
		Expect(class.Name).Should(Equal("com.vip.paas.e2e.test.ds.auth.OrgSpaceTest"))
		Expect(api.FormateFloatStr(class.Time)).Should(Equal("0.891"))
		Expect(class.Tests).Should(Equal(6))
		Expect(class.Errors).Should(Equal(1))
		Expect(class.Failures).Should(Equal(0))
		Expect(class.Skipped).Should(Equal(2))
		Expect(len(class.Methods)).Should(Equal(6))
		
		var method *api.Method
		for _, m := range class.Methods {
			if m.Name == "removeOrganization" {
				method = m
				break
			}
		}
		
		Expect(method.Name).Should(Equal("removeOrganization"))
		Expect(api.FormateFloatStr(method.Time)).Should(Equal("0.064"))
		Expect(method.Tests).Should(Equal(1))
		Expect(method.Errors).Should(Equal(1))
		Expect(method.Failures).Should(Equal(0))
		Expect(method.Skipped).Should(Equal(0))
	})
})
