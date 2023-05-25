class MyClaimHelpers {

    populateWarningMessage(notValid) {
        let msg
        (notValid.length == 1) ? msg = `Sorry there is an error with this field <ul>` :
            msg = `Sorry there is still few errors with those fields <ul>`

        notValid.forEach(field => {
            msg += `<li>${Common.replaceCamelCase(field)}</li>`
        })

        msg += `</ul>`

        return msg
    }

    // display main warning message
    displayWarningMessage(notValid) {
        const myWarningDivID = 'warningMessageDiv',
            myWarningBodyID = 'warningMessageBody'

        if (!notValid?.length) {
            Common.hideDivByID(myWarningDivID)
        } else {
            Common.showDivByID(myWarningDivID)
            Common.insertHTML(this.populateWarningMessage(notValid), myWarningBodyID)
        }

    }

    // populate form to edit entry
    populateFormEdit(data, active, rowID) {
        Common.insertHTML('Edit claim', 'createClaimDefinitionTitle')
        Common.insertInputValue('edit', 'formAction')
        Common.insertInputValue(rowID, 'rowid')
        Common.insertInputValue(data[0].innerHTML, 'name')
        Common.insertInputValue(data[1].innerHTML, 'description')
        Common.selectBoxOptionSelected(data[2].innerHTML, 'category')
        Common.checkRadio(data[3].innerHTML, 'confirmation')
        Common.insertInputValue(data[4].innerHTML, 'seniority')
        Common.checkRadio(active, 'active')
        Common.checkRadio(data[5].innerHTML, 'docRequired')
    }
    // populate and launch edit modal form
    makeEditable() {
        document.querySelectorAll('.editClaim').forEach(element => {
            element.addEventListener('click', (e) => {
                const rowID = e.target.dataset.id,
                    rowData = document.getElementById('CD' + rowID).querySelectorAll('.row-data'),
                    viewData = document.querySelector('#dataView').value
                Helpers.populateFormEdit(rowData, viewData, rowID)
                const myModalForm = new bootstrap.Modal(document.getElementById('createClaimDefinition'), {
                    keyboard: false
                })
                myModalForm.show()
            })
        })
    }

    // insert options into select tag
    insertOptions(id, data) {
        // remove first element (id=0; name='not defined')
        data.shift()
        const target = document.querySelector('#' + id)
        target.innerHTML = '<option selected hidden value=""></option>'
        data.forEach(element => {
            let opt = document.createElement('option')
            opt.value = element.ID
            opt.innerHTML = element.Name
            target.appendChild(opt)
        })
    }

    // insert Claim Definition options into select tag
    insertOptionsCD(id, data) {
        const target = document.querySelector('#' + id)
        target.innerHTML = '<option selected hidden value=""></option>'
        data.forEach(element => {
            let myID
            (typeof element.rowid == 'undefined') ? myID = 0 : myID = element.rowid;
            let opt = document.createElement('option')
            opt.classList = 'myclaimDefinition'
            opt.value = myID
            opt.innerHTML = element.name
            target.appendChild(opt)
        })
    }

    // populate upload files
    populateUploadFiles(wanted) {
        switch (wanted) {
            case 'docs':
                Common.insertHTML('claim', 'uploadedFilesTitle')
                Common.insertInputValue('claims', 'uploadedFilename')
                break;
            default:
                break;
        }
    }

    // populate all leave definition map
    populateClaimDefinitionMap(data, allClaimDefinition){
        data.forEach(element => {
            allClaimDefinition.set(element.rowid, element)
        })

        return allClaimDefinition
    }

    // populate my claim map (already requested amount by claim id)
    populateMyClaimMap(data, myClaim){
        data.forEach(element => {
            let previousAmount = 0.0;
            if (typeof myClaim.get(element.claimDefinitionID) != 'undefined') previousAmount = Number(myClaim.get(element.claimDefinitionID))
                 myClaim.set(element.claimDefinitionID,Number(element.amount)+previousAmount)
        })
        return myClaim
    }

    // populate all uploaded files map
    populateUploadedFilesMap(data, uploadedFiles){
        for (const [key, value] of Object.entries(data)) {
            uploadedFiles.set(key, value)
        }
        return uploadedFiles
    }
    // remove element from my required input fields array
    removeElementFromRIF(name) {
        const index = myRIF.indexOf(name);
        if (index > -1) { // splice array only when item is found
            myRIF.splice(index, 1); // 2nd parameter means remove one item only
        }
    }

    // display claim definition requirements 
    // (max amount allowed for selected claim definition)
    selectedClaimRequirements(claimDefinition, alreadyRequested, mySeniority){
        let maxAmount = 0.0, claimRequested = 0.0
        if (typeof alreadyRequested != 'undefined' && alreadyRequested != null) claimRequested = alreadyRequested

        // sort claims details by seniority desc
        let claimDetails = claimDefinition.details
        claimDetails.sort(function(a, b) {
            let keyA = Number(a.seniority),
                keyB = Number(b.seniority)
            // Compare the 2 seniority
            if (keyA < keyB) return -1
            if (keyA > keyB) return 1
            return 0
        })

        // calculate max entitlement for employee's seniority
        claimDetails.forEach(element => {
            if (mySeniority >= Number(element.seniority)){
                maxAmount      = Number(element.limitation)
                maxClaimAllow = maxAmount - claimRequested
                return
            }
        })

        // display selected leave requirements
        Common.insertHTML(`You are entitled for RM${maxAmount}, you already claimed RM${claimRequested}. Your balance is RM${maxClaimAllow} <br>`, 'myAuthorizationDiv')
        Common.showDivByID('myAuthorizationDiv')
    }

    // convert birthdate (remove timestamp)
    formatDate(t) {
        let b
        // get only the 10 first characters of the string
        const d = t.substring(0, 10)
        // the zero value of a date is 0001-01-01
        return (d == '0001-01-01') ? b = '' : b = d
    }

    // create attachments icon
    createAttachments(appID){
        return (myUploadedFiles.has(appID)) ? '<i class="bi-check-square pointer"></i>' : '<i class="bi-file-excel-fill not-allowed"></i>'
    }

    // populate attachments modal & open it
    populateAttachments(appID){
        if (myUploadedFiles.has(appID)){
            const myFiles = myUploadedFiles.get(appID),
                  target = document.querySelector('#uploadedFilesModalBody'),
                  myModal = new bootstrap.Modal(document.getElementById('uploadedFilesModal'), { 
                    backdrop: 'static',
                    keyboard: false 
                  })
            target.innerHTML = ''
            myFiles.forEach(element => {
                const oneAttachment = document.createElement('div')
                oneAttachment.classList = 'd-flex justify-content-center mb-3'
                oneAttachment.innerHTML = `<div>
                                                <div><img class="myPicture" src="/upload/${connectedEmail}/claims/${appID}/${element}" /></div>
                                                <div class="text-end"><a href="/upload/${connectedEmail}/claims/${appID}/${element}" class="myLink" download="${element}"><i class="bi-download pointer"></i> download</a></div>
                                           </div>`
                target.appendChild(oneAttachment)
            })
            myModal.show()
        }
    }
    // insert to datatable one row per element of data
    insertRows(data) {
        const target = document.querySelector('#myClaimBody')
        target.innerHTML = ''

        const myColumns = [
            {title: "Allowance"},
            {title: "Description"},
            {title: "Category"},
            {title: "Amount"},
            {title: "Status"},
            {title: "Approved At"},
            {title: "Approved By"},
            {title: "Approved Amount"},
            {title: "Reason"},
            {title: "Attachment"},
            {title: "Created By"},
            {title: "Created At"},
            {title: "Updated By"},
            {title: "Updated At"},
            {title: "Action"}
        ]

        // when no data
        if (data == null){
          DT.initiateMyTable('myClaimTable', myColumns)
          $('#myClaimTable').DataTable().clear().draw()
          return
        }

        data.forEach(element => {
            // since omitempty has been set golang side...
            let approvedBy, approvedAmount;
            (typeof element.approvedAmount == 'undefined') ? approvedAmount = 0 : approvedAmount = element.approvedAmount;
            (typeof element.approvedBy == 'undefined') ? approvedBy = 0 : approvedBy = Number(element.approvedBy);

            let attachments = this.createAttachments(element.rowid)

            let row = document.createElement('tr')
            row.id = 'myClaim' + element.rowid
            row.innerHTML = `<td class="row-data" data-id="${element.claimDefinitionID}">${element.claimDefinition}</td>
                             <td class="row-data">${element.description}</td>
                             <td class="row-data">${element.category}</td>
                             <td class="row-data">${element.amount}</td>
                             <td class="row-data" data-id="${element.statusID}">${element.status}</td>
                             <td class="row-data">${this.formatDate(element.approvedAt)}</td>
                             <td class="row-data">${allEmployees.get(approvedBy)}</td>
                             <td class="row-data">${approvedAmount}</td>
                             <td class="row-data">${element.rejectedReason}</td>
                             <td class="row-data myAttachments text-center" data-id=${element.rowid}>${attachments}</td>
                             <td class="row-data">${allEmployees.get(Number(element.createdBy))}</td>
                             <td class="row-data">${this.formatDate(element.createdAt)}</td>
                             <td class="row-data">${allEmployees.get(Number(element.updatedBy))}</td>
                             <td class="row-data">${this.formatDate(element.updatedAt)}</td>`
            // if status is pending
            if (element.statusid == 2){
                row.innerHTML += `<td>
                                    <div class="row">
                                        <div class="col text-end">
                                            <div class="form-check">
                                                <label class="form-check-label fw-lighter fst-italic smaller" for="softDelete"><i class="bi-trash2-fill largeIcon pointer deleteLeave"></i></label>
                                                <input class="form-check-input deleteCheckboxes"  type="checkbox" value="${element.rowid}" name="softDelete">                                    
                                            </div> 
                                        </div>
                                    </div>                                                 
                                </td>`
            }else{
                row.innerHTML += `<td class="text-end"><i class="bi-lock"></i></td>`
            } 
            target.appendChild(row)
        })
        
        // initiate datatable & show/hide columns
        DT.initiateMyTable('myClaimTable', myColumns)
    }

    // returns list of selected claim id to be deleted
    selectedClaim() {
        const checked = document.querySelectorAll('input[name=softDelete]:checked')
        if (checked.length == 0) document.querySelector('#deleteWarningMessageDiv').style.display = 'block'
        if (checked.length > 0) {
            let allChecked = []
            checked.forEach(element => {
                allChecked.push(element.value)
            })
            
            return allChecked
        }
    }
    // populate confirm detele message
    populateConfirmDelete(nb) {
        let msg
        (nb == 1) ? msg = 'Do you really want to delete this claim?' : msg = `Do you really want to delete ${nb} claims?`

        const myBody = document.querySelector('#confirmDeleteBody')
        myBody.innerHTML = msg
    }

    // generate warning message
    populateWarningMessage(notValid){
        let msg = `Please check the following error(s) <ul>`

        notValid.forEach(err => {
            msg += `<li>${err}</li>`
        })

        msg += `</ul>`

        return msg
    }

    // display main warning message
    displayWarningMessage(notValid){
        const myWarningDivID  = 'warningMessageDiv',
              myWarningBodyID = 'warningMessageBody'

        if (!notValid?.length) {
            Common.hideDivByID(myWarningDivID)
        }else{
            Common.showDivByID(myWarningDivID)
            Common.insertHTML(this.populateWarningMessage(notValid), myWarningBodyID)
        }
      
    }

    validateApplication() {
        let notValid = []

        // validate if amount requested is allowed (<= maxClaimAllow )
        const amountRequested = document.querySelector('#amount')
        if (parseFloat(amountRequested.value) > maxClaimAllow){
            Common.insertHTML(`You can'\'t select more dates than RM${maxClaimAllow}`, 'amountError')
            notValid.push('The maximun allowed days for this leave application is exceeded')
        }else{
            Common.insertHTML('', 'amountError')
        }

        // validate when attachment is required if file is uploaded
        if (attachmentRequired == 1){
            const uploadedFiles = document.querySelector('#uploadedFiles')
            if (uploadedFiles.value == ''){
                Common.insertHTML(`You must upload at least one attachment`, 'uploadedFilesError')
                notValid.push('No attachment was uploaded')
            }else{
                Common.insertHTML('', 'uploadedFilesError')
            }
        }

        // display main warning message
        this.displayWarningMessage(notValid)

        return notValid.length
    }

    // send attachment 
    SendAttachment(leaveApplicationID, connectedEmail, connectedID){
        Common.insertInputValue(leaveApplicationID, 'applicationID')
        Common.insertInputValue('claims', 'uploadedFilename')
        Common.insertInputValue(connectedEmail, 'employeeEmail')
        Common.insertInputValue(connectedID, 'employeeID')
 
        document.getElementById("uploadedFilesForm").submit()
    }
}