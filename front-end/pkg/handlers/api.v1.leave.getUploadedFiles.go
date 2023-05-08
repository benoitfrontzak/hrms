package handlers

import (
	"net/http"
	"path/filepath"
	"strings"
)

// API which fetchs and returns all uploaded files of one employee
func (rep *Repository) GetLeaveUploadedFiles(w http.ResponseWriter, r *http.Request) {

	// Get my context from middleware request
	myc := r.Context().Value(httpContext).(httpContextStruct)

	if myc.Auth {
		// get employee email from url
		email := strings.TrimPrefix(r.URL.Path, "/api/v1/leave/getUploadedFiles/")

		// get all employee's uploaded files
		myFiles, err := myLeaveUploadedFiles(email)
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
func myLeaveUploadedFiles(email string) (*uploadedFiles, error) {
	myFiles := map[string][]string{}
	// myArchive := map[string]map[string][]string{}

	// fetch application IDs
	myIDs, err := filesInDirectory(filepath.Join(path, email, "leaves"))
	if err != nil {
		return nil, err
	}

	// fetch files for each appID
	for _, appID := range myIDs {
		// fetch appID's files
		myF, err := filesInDirectory(filepath.Join(path, email, "leaves/"+appID))
		if err != nil {
			return nil, err
		}
		myFiles[appID] = myF
	}

	myUploadedFiles := uploadedFiles{
		Files: myFiles,
	}

	return &myUploadedFiles, nil
}
