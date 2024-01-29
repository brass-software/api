package api

import (
	"io"
)

func GET(w io.Writer, path string) error {
	return do(w, path, "GET", nil)
}
