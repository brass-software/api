package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func do(w io.Writer, path string, method string, body io.Reader) error {
	req, err := http.NewRequest(method, URL+path, body)
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
		return fmt.Errorf("unexpected api response: %s: %s", res.Status, s)
	}
	_, err = io.Copy(w, res.Body)
	if err != nil {
		return err
	}
	return nil
}
