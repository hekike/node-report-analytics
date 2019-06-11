package model

import (
	"strconv"
	"time"
)

// DiagnosticReport Node.js Diagnostic Report
type DiagnosticReport struct {
	Hash            string
	Path            string
	Header          DiagnosticReportHeader `json:"header"`
	JavaScriptStack JavaScriptStack        `json:"javascriptStack"`
}

// JavaScriptStack error message with stack trace
type JavaScriptStack struct {
	Message string   `json:"message"`
	Stack   []string `json:"stack"`
}

// DiagnosticReportHeader of the report
type DiagnosticReportHeader struct {
	Event                string            `json:"event"`
	Trigger              string            `json:"trigger"`
	Filename             string            `json:"filename"`
	DumpEventTime        time.Time         `json:"dumpEventTime"`
	DumpEventTimeStamp   UnixTimestamp     `json:"dumpEventTimeStamp"`
	ProcessID            int               `json:"processId"`
	Cwd                  string            `json:"cwd"`
	CommandLine          []string          `json:"commandLine"`
	NodejsVersion        string            `json:"nodejsVersion"`
	GlibcVersionRuntime  string            `json:"glibcVersionRuntime"`
	GlibcVersionCompiler string            `json:"glibcVersionCompiler"`
	WordSize             int               `json:"wordSize"`
	Arch                 string            `json:"arch"`
	Platform             string            `json:"platform"`
	OsName               string            `json:"osName"`
	OsRelease            string            `json:"osRelease"`
	OsVersion            string            `json:"osVersion"`
	OsMachine            string            `json:"osMachine"`
	Host                 string            `json:"host"`
	Release              Release           `json:"release"`
	ComponentVersions    ComponentVersions `json:"componentVersions"`
}

// Release node release
type Release struct {
	Name string `json:"name"`
}

// ComponentVersions components and their version
type ComponentVersions struct {
	Node       string `json:"node"`
	V8         string `json:"v8"`
	Uv         string `json:"uv"`
	Zlib       string `json:"zlib"`
	Ares       string `json:"ares"`
	Modules    string `json:"modules"`
	Nghttp2    string `json:"nghttp2"`
	Napi       string `json:"napi"`
	Llhttp     string `json:"llhttp"`
	HTTPParser string `json:"http_parser"`
	Openssl    string `json:"openssl"`
}

// UnixTimestamp JavaScript timestamp
type UnixTimestamp struct {
	time.Time
}

// UnmarshalJSON for JavaScript timestamp
func (sd *UnixTimestamp) UnmarshalJSON(input []byte) error {
	str := string(input)
	i, err := strconv.ParseInt(str[1:len(str)-1], 10, 64)
	if err != nil {
		panic(err)
	}
	sd.Time = time.Unix(i/1000, 0)
	return nil
}
