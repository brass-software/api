package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func do(out io.Writer, path string, method string, in io.Reader) error {
	req, err := http.NewRequest(method, URL+path, in)
	if err != nil {
		return err
	}
	req.Header.Add("BRASS_TOKEN", token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		b, _ := io.ReadAll(res.Body)
		s := strings.TrimSpace(string(b))
		return fmt.Errorf("%s: %s", res.Status, s)
	}
	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}
	return nil
}
