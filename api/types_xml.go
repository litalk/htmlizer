package api

import "encoding/xml"

type TestSuite struct {
	XMLName xml.Name `xml:"testsuite"`

	Failures int     `xml:"failures,attr"`
	Time     float64 `xml:"time,attr"`
	Errors   int     `xml:"errors,attr"`
	Skipped  int     `xml:"skipped,attr"`
	Tests    int     `xml:"tests,attr"`
	Name     string  `xml:"name,attr"`

	Properties []*Property `xml:"properties>property"`
	TestCases  []*TestCase `xml:"testcase"`
}

type Property struct {
	XMLName xml.Name `xml:"property"`

	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type TestCase struct {
	XMLName xml.Name `xml:"testcase"`

	Time      float64 `xml:"time,attr"`
	ClassName string  `xml:"classname,attr"`
	Name      string  `xml:"name,attr"`

	Skipped   *Skipped   `xml:"skipped"`
	Error     *Error     `xml:"error"`
	Failure   *Failure   `xml:"failure"`
	SystemOut *SystemOut `xml:"system-out"`
}

type Skipped struct {
	XMLName xml.Name `xml:"skipped"`
}

type Error struct {
	XMLName xml.Name `xml:"error"`

	Type       string `xml:"type,attr"`
	StackTrace string `xml:",chardata"`
}

type Failure struct {
	XMLName xml.Name `xml:"failure"`

	Message    string `xml:"message,attr"`
	Type       string `xml:"type,attr"`
	StackTrace string `xml:",chardata"`
}

type SystemOut struct {
	XMLName xml.Name `xml:"system-out"`

	Log string `xml:",chardata"`
}

func (t *TestCase) isFail() bool {
	return t.Failure != nil
}

func (t *TestCase) isError() bool {
	return t.Error != nil
}

func (t *TestCase) isSkip() bool {
	return t.Skipped != nil
}
