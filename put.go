package api

import "io"

func PUT(path string, body io.Reader) error {
	return do(nil, path, "PUT", nil)
}
