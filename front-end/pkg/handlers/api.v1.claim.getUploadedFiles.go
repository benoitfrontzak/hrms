package handlers

import (
	"net/http"
	"path/filepath"
	"strings"
)

// API which fetchs and returns all uploaded files of one employee
func (rep *Repository) GetUploadedClaimFiles(w http.ResponseWriter, r *http.Request) {

	// Get my context from middleware request
	myc := r.Context().Value(httpContext).(httpContextStruct)

	if myc.Auth {
		// get employee email from url
		email := strings.TrimPrefix(r.URL.Path, "/api/v1/claim/getUploadedFiles/")

		// get all employee's uploaded files
		myFiles, err := myClaimUploadedFiles(email)
		if err != nil {
			rep.errorJSON(w, err)
			return
		}

		// response to be sent
		answer := jsonResponse{
			Error:   false,
			Message: "claim's uploaded files'",
			Data:    myFiles,
		}
		rep.writeJSON(w, http.StatusAccepted, answer)
	}
}

// collect all employee's uploaded files from one employee (email)
func myClaimUploadedFiles(email string) (*uploadedFiles, error) {
	myFiles := map[string][]string{}
	// myArchive := map[string]map[string][]string{}

	// fetch application IDs
	myIDs, err := filesInDirectory(filepath.Join(path, email, "claims"))
	if err != nil {
		return nil, err
	}

	// fetch files for each appID
	for _, appID := range myIDs {
		// fetch appID's files
		myF, err := filesInDirectory(filepath.Join(path, email, "claims/"+appID))
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
