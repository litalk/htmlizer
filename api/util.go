package api

import (
	"encoding/xml"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func ParseJunitXML(report string) (testSuite *TestSuite, err error) {
	data, err := ioutil.ReadFile(report)
	if err != nil {
		return
	}

	testSuite = &TestSuite{Properties: make([]*Property, 0)}
	err = xml.Unmarshal(data, testSuite)
	return
}

func Aggregate(testSuite *TestSuite) (report *Report) {
	methods := []*Method{}
	failureMethods := []*Method{}

	for _, tc := range testSuite.TestCases {
		m := GenerateMethod(tc)
		methods = append(methods, m)

		if m.Errors != 0 || m.Failures != 0 || m.Skipped != 0 {
			failureMethods = append(failureMethods, m)
		}
	}

	classes := aggregateMethod(methods)
	packages := aggregateClass(classes)
	report = aggregatePackage(packages)

	report.Name = testSuite.Name
	report.FailureMethods = failureMethods
	report.Date = time.Now().Format("2006-01-02 15:04:05")

	return report
}

func aggregateMethod(methods []*Method) (result []*Class) {
	classMap := make(map[string]*Class)
	for _, m := range methods {
		var class *Class
		if c, ok := classMap[m.ClassName]; ok {
			class = c
		} else {
			class = &Class{&testCollection{}, []*Method{}}
			class.Name = m.ClassName
			classMap[m.ClassName] = class
		}

		class.testCollection.merge(m.testCollection)
		class.Methods = append(class.Methods, m)
	}

	result = []*Class{}
	for _, class := range classMap {
		result = append(result, class)
	}
	return
}

func aggregateClass(classes []*Class) (result []*Package) {
	packageMap := make(map[string]*Package)
	for _, c := range classes {
		var pkg *Package
		if p, ok := packageMap[c.PackageName()]; ok {
			pkg = p
		} else {
			pkg = &Package{&testCollection{}, []*Class{}}
			pkg.Name = c.PackageName()
			packageMap[pkg.Name] = pkg
		}

		pkg.testCollection.merge(c.testCollection)
		pkg.Classes = append(pkg.Classes, c)
	}

	result = []*Package{}
	for _, pkg := range packageMap {
		result = append(result, pkg)
	}
	return
}

func aggregatePackage(packages []*Package) (report *Report) {
	report = &Report{}
	report.testCollection = &testCollection{}
	report.Packages = []*Package{}
	for _, p := range packages {
		report.testCollection.merge(p.testCollection)
		report.Packages = append(report.Packages, p)
	}

	return
}

func GenerateMethod(tc *TestCase) *Method {
	m := &Method{}
	m.testCollection = &testCollection{}
	m.ClassName = tc.ClassName

	m.Name = tc.Name
	m.Time = tc.Time
	m.Tests = 1

	if tc.Error != nil {
		m.Errors = 1

		m.StackTrace = tc.Error.StackTrace
		m.Type = tc.Error.Type
		m.Message = ""
		if tc.SystemOut != nil {
			m.Log = tc.SystemOut.Log
		}
	} else if tc.Failure != nil {
		m.Failures = 1

		m.StackTrace = tc.Failure.StackTrace
		m.Type = tc.Failure.Type
		m.Message = tc.Failure.Message
		if tc.SystemOut != nil {
			m.Log = tc.SystemOut.Log
		}
	} else if tc.Skipped != nil {
		m.Skipped = 1
	}
	return m
}

func SuccessRate(tests, failures, errors, skipped int) string {
	succeed := float64(tests - failures - errors - skipped)
	succeedRate := strconv.FormatFloat(succeed/float64(tests)*100, 'f', 3, 32)
	return succeedRate + "%"
}

func Line(content string, line int) string {
	lines := strings.Split(content, "\n")
	if len(lines) >= line {
		return lines[line-1]
	}
	return ""
}
