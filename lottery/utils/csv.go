package utils

import (
	"bytes"
	"encoding/csv"
	"net/http"

	"github.com/gin-gonic/gin/render"
	"github.com/pkg/errors"
)

// CSV.
type CSV struct {
	Content [][]string
	Title   string
}

var _ render.Render = CSV{}

// WriteContentType.
func (j CSV) WriteContentType(w http.ResponseWriter) {
	// 此处跟gin的render不相容
}

// Render.
func (j CSV) Render(w http.ResponseWriter) (err error) {
	bs, err := formatCSV(j.Content)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	if _, err = w.Write(bs); err != nil {
		err = errors.WithStack(err)
	}
	return
}

// formatCSV format to csv data
func formatCSV(records [][]string) (data []byte, err error) {
	buf := new(bytes.Buffer)
	// windows bom头，解决ms excel乱码问题
	buf.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(buf)
	if err = w.WriteAll(records); err != nil {
		return
	}
	data = buf.Bytes()
	return
}

