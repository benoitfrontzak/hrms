package handlers

import (
	"net/http"
	"path/filepath"
	"strings"
)

// API which fetchs and returns all uploaded files of one employee
func (rep *Repository) GetUploadedFiles(w http.ResponseWriter, r *http.Request) {

	// Get my context from middleware request
	myc := r.Context().Value(httpContext).(httpContextStruct)

	if myc.Auth {
		// get employee email from url
		email := strings.TrimPrefix(r.URL.Path, "/api/v1/employee/getUploadedFiles/")

		// get all employee's uploaded files
		myFiles, err := myUploadedFiles(email)
		if err != nil {
			rep.errorJSON(w, err)
			return
		}

		// response to be sent
		answer := jsonResponse{
			Error:   false,
			Message: "employee's uploaded files'",
			Data:    myFiles,
		}
		rep.writeJSON(w, http.StatusAccepted, answer)
	}
}

// collect all employee's uploaded files from one employee (email)
func myUploadedFiles(email string) (*uploadedFiles, error) {
	// collect active IC files
	myIC, err := filesInDirectory(filepath.Join(path, email, "ic"))
	if err != nil {
		return nil, err
	}
	// collect active passport files
	myPassport, err := filesInDirectory(filepath.Join(path, email, "passport"))
	if err != nil {
		return nil, err
	}

	// collect archive IC files (per timestamp directory)
	myArchiveDirIC, err := filesInDirectory(filepath.Join(path, email, "archive", "ic"))
	if err != nil {
		return nil, err
	}
	myArchiveICFiles := map[string][]string{}
	if len(myArchiveDirIC) > 0 {
		for _, dir := range myArchiveDirIC {
			myArchiveICFiles[dir], err = filesInDirectory(filepath.Join(path, email, "archive", "ic", dir))
			if err != nil {
				return nil, err
			}
		}
	}

	// collect archive passport files (per timestamp directory)
	myArchiveDirPassport, err := filesInDirectory(filepath.Join(path, email, "archive", "passport"))
	if err != nil {
		return nil, err
	}
	myArchivePassportFiles := map[string][]string{}
	if len(myArchiveDirPassport) > 0 {
		for _, dir := range myArchiveDirPassport {
			myArchivePassportFiles[dir], err = filesInDirectory(filepath.Join(path, email, "archive", "passport", dir))
			if err != nil {
				return nil, err
			}
		}
	}

	myFiles := map[string][]string{}
	myFiles["ic"] = myIC
	myFiles["passport"] = myPassport

	myArchive := map[string]map[string][]string{}
	myArchive["ic"] = myArchiveICFiles
	myArchive["passport"] = myArchivePassportFiles

	myUploadedFiles := uploadedFiles{
		Files:   myFiles,
		Archive: myArchive,
	}

	return &myUploadedFiles, nil
}
