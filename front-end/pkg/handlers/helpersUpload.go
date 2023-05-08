package handlers

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// extract form data from payload
func extractAttachmentsInfo(r *http.Request) (*attachments, error) {
	// Parse our multipart form, 10 << 20 specifies a maximum upload of 10 MB files.
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return nil, err
	}
	// read the Form data
	formdata := r.MultipartForm

	// collect form information:
	files := formdata.File["uploadedFiles"]
	filename := r.FormValue("uploadedFilename")
	email := r.FormValue("employeeEmail")
	ID := r.FormValue("employeeID")

	// check wether application ID is post
	aid := r.PostForm.Has("applicationID")
	// get/set application ID
	applicationID := "0"
	if aid {
		applicationID = r.FormValue("applicationID")
	}

	// populate attachment's members
	a := attachments{
		Files:         files,
		Filename:      filename,
		ApplicationID: applicationID,
		Email:         email,
		ID:            ID,
	}

	return &a, nil
}

// archive files
func archiveFiles(email, filename string) error {
	// create archive folder (upload/email/archive/category/timestamp)
	now := myTimestamp()
	archive := filepath.Join(path, email, "archive", filename, now)
	err := createNewFolder(archive)
	if err != nil {
		return err
	}

	// select all files from actual directory
	actual := filepath.Join(path, email, filename)
	entries, err := os.ReadDir(actual)
	if err != nil {
		return err
	}

	// move all files to archive folder
	for _, e := range entries {
		err = os.Rename(filepath.Join(actual, e.Name()), filepath.Join(archive, e.Name()))
		if err != nil {
			return err
		}
	}

	return nil
}

// upload files to server
func filesUploader(files []*multipart.FileHeader, email, filename, applicationID string) error {
	// loop through the files one by one
	for i, f := range files {
		// store file number
		number := strconv.Itoa(i)

		// store uploaded file
		file, err := files[i].Open()
		if err != nil {
			return err
		}
		defer file.Close()

		// store authorized uploaded file extension
		ext := getExtension(f.Header.Values("Content-Type")[0])
		if ext == "lock" {
			return errors.New("unauthorized file extension")
		}

		// rename filename (must give full path with file extension)
		myFile := filename + number + "." + ext
		fullpath := filepath.Join(path, email, filename, myFile)
		if applicationID != "" {
			fullpath = filepath.Join(path, email, filename, applicationID, myFile)
		}

		// create new empty file
		out, err := os.Create(fullpath)
		if err != nil {
			return err
		}
		defer out.Close()

		// copy stored uploaded file to new empty file
		_, err = io.Copy(out, file) // file not files[i] !
		if err != nil {
			return err
		}

	}

	return nil
}

// get extension of file content type header
// We accept only "jpg", "gif", "png", "bmp", "pdf"
func getExtension(ct string) (fileext string) {
	switch ct {
	case "image/jpeg", "image/jpg":
		fileext = "jpg"

	case "image/gif":
		fileext = "gif"

	case "image/bmp":
		fileext = "bmp"

	case "image/png":
		fileext = "png"

	case "application/pdf":
		fileext = "pdf"

	default:
		fileext = "lock"
	}

	return
}
