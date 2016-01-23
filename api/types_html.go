package api

import "strings"

var (
	SUCCEED string = "s"
	FAILURE string = "f"
	ERROR   string = "e"
	SKIPPED string = "i"
)

type collection interface {
	Result() string
}

type testCollection struct {
	Name     string
	Time     float64
	Tests    int
	Failures int
	Errors   int
	Skipped  int
}

type Report struct {
	*testCollection

	Packages       []*Package
	FailureMethods []*Method
	Date           string
}

type Package struct {
	*testCollection

	Classes []*Class
}

type Class struct {
	*testCollection

	Methods []*Method
}

type Method struct {
	*testCollection

	ClassName  string
	Message    string
	Type       string
	StackTrace string
	Log        string
}

func (c *testCollection) merge(other *testCollection) {
	c.Time = c.Time + other.Time
	c.Tests = c.Tests + other.Tests
	c.Failures = c.Failures + other.Failures
	c.Errors = c.Errors + other.Errors
	c.Skipped = c.Skipped + other.Skipped
}

func (c *testCollection) SuccessRate() string {
	return SuccessRate(c.Tests, c.Failures, c.Errors, c.Skipped)
}

func (c *testCollection) Result() string {
	if c.Errors > 0 {
		return ERROR
	} else if c.Failures > 0 {
		return FAILURE
	} else if c.Skipped > 0 {
		return SKIPPED
	}
	return SUCCEED
}

func (c *Class) PackageName() string {
	lidx := strings.LastIndex(c.Name, ".")
	if lidx == -1 {
		return ""
	}
	return c.Name[0:lidx]
}

func (c *Class) SimpleName() string {
	lidx := strings.LastIndex(c.Name, ".")
	if lidx == -1 {
		return c.Name
	}
	return c.Name[lidx+1:]
}

func (p *Package) mergeTestCase(testCase *TestCase) {
	p.Tests = p.Tests + 1
	p.Time = p.Time + testCase.Time

	if testCase.isError() {
		p.Errors = p.Errors + 1
	}

	if testCase.isFail() {
		p.Failures = p.Failures + 1
	}

	if testCase.isSkip() {
		p.Skipped = p.Skipped + 1
	}
}

func (p *Package) SuccessRate() string {
	return SuccessRate(p.Tests, p.Failures, p.Errors, p.Skipped)
}

func (m *Method) Exception() string {
	if m.Errors > 0 {
		return strings.TrimSpace(Line(m.StackTrace, 1))
	}
	return ""
}

func (m *Method) ErrorLocation() string {
	if m.StackTrace == "" {
		return ""
	}

	fullName := m.ClassName + "." + m.Name
	lines := strings.Split(m.StackTrace, "\n")
	for _, line := range lines {
		if strings.Index(line, fullName) != -1 {
			return line
		}
	}
	return ""
}
