package api

func DELETE(path string) error {
	return do(nil, path, "DELETE", nil)
}
