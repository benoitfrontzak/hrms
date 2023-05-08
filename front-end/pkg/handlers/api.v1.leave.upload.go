package handlers

import (
	"net/http"
	"path/filepath"
)

// API which upload to file server all posted files for one employee
// if files already exists, it archives them to file server
func (rep *Repository) UploadLeaveFiles(w http.ResponseWriter, r *http.Request) {
	// Get my context from middleware request
	myc := r.Context().Value(httpContext).(httpContextStruct)

	if myc.Auth {
		// collect information from payload
		a, err := extractAttachmentsInfo(r)
		if err != nil {
			rep.errorJSON(w, err)
			return
		}

		// create email/category directory
		fullpath := filepath.Join(path, a.Email, a.Filename, a.ApplicationID)
		err = createNewFolder(fullpath)
		if err != nil {
			rep.errorJSON(w, err)
			return
		}

		// check if directory already got uploaded files
		empty, err := isEmpty(fullpath)
		if err != nil {
			rep.errorJSON(w, err)
			return
		}

		// archive files when directory is not empty
		if !empty {
			err = archiveFiles(a.Email, a.Filename)
			if err != nil {
				rep.errorJSON(w, err)
				return
			}
		}

		// upload files to directory
		err = filesUploader(a.Files, a.Email, a.Filename, a.ApplicationID)
		if err != nil {
			rep.errorJSON(w, err)
			return
		}

		// redirect (TODO write response as API)
		http.Redirect(w, r, frontEnd+"myleave", http.StatusSeeOther)
	}
}
