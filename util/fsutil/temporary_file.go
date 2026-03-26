package fsutil

import "os"

func TemporaryFile() (string, func(), error) {
	file, err := os.CreateTemp("", "dollar-tool-")
	if err != nil {
		return "", nil, err
	}

	cleanup := func() {
		_ = os.Remove(file.Name())
	}

	return file.Name(), cleanup, nil
}
