package api

import (
	"bytes"
	"encoding/json"
)

func Send(path, method string, args map[string]string) (map[string]string, error) {
	b, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	in := bytes.NewReader(b)
	r := bytes.NewBuffer(nil)
	err = do(r, path, method, in)
	if err != nil {
		return nil, err
	}
	out := map[string]string{}
	err = json.NewDecoder(r).Decode(&out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
