package report

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/hekike/node-report-analytics/pkg/model"
)

// ListDir reads the diagnostic report files from a specific folder
func ListDir(dir string) ([]string, error) {
	var files []string
	r, _ := regexp.Compile("report\\.[0-9]{8}\\.[0-9\\.]+json")

	err := filepath.Walk(dir, func(
		path string,
		info os.FileInfo,
		err error,
	) error {
		if r.MatchString(path) {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// ReadFile reads the diagnostic report file
func ReadFile(path string) (*model.DiagnosticReport, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	report, err := FromJSON(string(file))
	if err != nil {
		return nil, err
	}

	report.Path = path

	return report, err
}

// ReadDir reads the diagnostic report file
func ReadDir(dir string) ([]*model.DiagnosticReport, error) {
	files, err := ListDir(dir)
	var reports []*model.DiagnosticReport

	if err != nil {
		return nil, err
	}

	for _, path := range files {
		report, err := ReadFile(path)
		if err != nil {
			return nil, err
		}
		if report != nil {
			reports = append(reports, report)
		}

	}

	return reports, nil
}

// DeleteFile removes a report from the disk
func DeleteFile(path string) error {
	err := os.Remove(path)
	return err
}
