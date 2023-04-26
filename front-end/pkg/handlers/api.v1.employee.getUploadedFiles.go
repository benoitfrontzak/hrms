package handlers

import (
	"net/http"
	"path/filepath"
	"strings"
)

var wantedFiles = []string{"profile", "ic", "passport", "otherInformation"}

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
	myFiles := map[string][]string{}
	myArchive := map[string]map[string][]string{}

	for _, wf := range wantedFiles {
		myF, err := myActiveFiles(email, wf)
		if err != nil {
			return nil, err
		}
		myAF, err := myArchiveFiles(email, wf)
		if err != nil {
			return nil, err
		}
		myFiles[wf] = myF
		myArchive[wf] = myAF
	}

	myUploadedFiles := uploadedFiles{
		Files:   myFiles,
		Archive: myArchive,
	}

	return &myUploadedFiles, nil
}

func myActiveFiles(email, myFile string) ([]string, error) {
	// collect active profile files
	active, err := filesInDirectory(filepath.Join(path, email, myFile))
	if err != nil {
		return nil, err
	}

	return active, nil
}

func myArchiveFiles(email, myFile string) (map[string][]string, error) {
	// collect archive IC files (per timestamp directory)
	myArchiveDir, err := filesInDirectory(filepath.Join(path, email, "archive", myFile))
	if err != nil {
		return nil, err
	}
	myArchiveFiles := map[string][]string{}
	if len(myArchiveDir) > 0 {
		for _, dir := range myArchiveDir {
			myArchiveFiles[dir], err = filesInDirectory(filepath.Join(path, email, "archive", myFile, dir))
			if err != nil {
				return nil, err
			}
		}
	}

	return myArchiveFiles, nil
}
