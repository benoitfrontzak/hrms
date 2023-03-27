package handlers

import (
	"errors"
	"io"
	"os"
	"strings"
	"time"
)

// create new directory
func createNewFolder(fullpath string) error {
	// create recursively new directory if doesn't exists
	if _, err := os.Stat(fullpath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(fullpath, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

// format timestamp to use as folder name (yyyy-mm-dd_hh-mm-ss)
func myTimestamp() string {
	// replace space from time.Now to underscore
	nowF := strings.Replace(time.Now().Format("2006-01-02 15:04:05"), " ", "_", 1)
	// replace colon symbol to dash
	myNow := strings.Replace(nowF, ":", "-", 2)

	return myNow
}

// check if directory is empty
func isEmpty(name string) (bool, error) {
	// open directory
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	// try to read he contents of the directory
	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil // directory is empty
	}

	return false, err // directory is either not empty or error
}

// collect all filenames from directory
func filesInDirectory(path string) ([]string, error) {
	var files []string

	// check if directory exists
	exist, err := exists(path)
	if err != nil {
		return nil, err
	}

	if exist {
		// returns directory entries
		entries, err := os.ReadDir(path)
		if err != nil {
			return nil, err
		}
		// append directory entries
		if len(entries) > 0 {
			for _, e := range entries {
				files = append(files, e.Name())
			}
		}
	}

	return files, nil
}

// check if directory exists or not
func exists(path string) (bool, error) {
	// returns file info
	_, err := os.Stat(path)
	if err == nil {
		return true, nil // file exists
	}
	// check if file doesn't exists
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
