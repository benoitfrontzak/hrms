class MyLeaveHelpers {
    
    // populate form to edit entry
    populateFormEdit(data, active, rowID){
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
    makeEditable(){
        document.querySelectorAll('.editClaim').forEach(element => {
            element.addEventListener('click', (e) => {
                const rowID = e.target.dataset.id,
                      rowData = document.getElementById('CD'+rowID).querySelectorAll('.row-data'),
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
    insertOptions(id, data, myGender) {
        const target = document.querySelector('#' + id)
        target.innerHTML = '<option selected hidden value=""></option>'
        data.forEach(element => {
            if((element.genderid == 2) || (element.genderid == myGender)){
                let opt = document.createElement('option')
                opt.value = element.rowid
                opt.innerHTML = element.code + ' - ' + element.description
                target.appendChild(opt)
            }
        })
    }

    // convert timestamp (remove timestamp)
    formatDate(t){
        let b
        // get only the 10 first characters of the string
        const d = t.substring(0,10)
        // the zero value of a date is 0001-01-01
        return (d == '0001-01-01') ? b = '' : b = d
    }
    // customize boolean (0|1) with icons
    myBooleanIcons(value){
        return (value == 1) ? '<i class="bi-check2-square"></i> true' : '<i class="bi-x-square"></i> false'
    }

    // convert isHalf
    leaveDay(isHalf){
        let result
        return (isHalf == 1) ? result = 'Half <i class="bi-square-half"></i>' : result = 'Full <i class="bi-square-fill"></i>'
    }
    // create rows leave details
    createRowsLD(entries){
        let details = `<div class="d-flex justify-content-between">
                            <div>${this.formatDate(entries[0].requestedDate)}</div>
                            <div>&nbsp;</div>
                            <div>${this.leaveDay(entries[0].isHalf)}</div>
                        </div>`

        if (entries.length > 1){
            details =``
            entries.forEach(e => {
                details += `<div class="d-flex justify-content-between">
                              <div>${this.formatDate(e.requestedDate)}</div>
                              <div>&nbsp;</div>
                              <div>${this.leaveDay(e.isHalf)}</div>
                            </div>`
            })
        }

        return details
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
                                                <div><img class="myPicture" src="/upload/${connectedEmail}/leaves/${appID}/${element}" /></div>
                                                <div class="text-end"><a href="/upload/${connectedEmail}/leaves/${appID}/${element}" class="myLink" download="${element}"><i class="bi-download pointer"></i> download</a></div>
                                           </div>`
                target.appendChild(oneAttachment)
            })
            myModal.show()
        }
    }

    // insert to datatable one row per element of data
    insertRows(data) {
        const target = document.querySelector('#myLeaveBody')
        target.innerHTML = ''
        
        const myColumns = [
            {title: "Leave"},
            {title: "Description"},
            {title: "Status"},
            {title: "Approved At"},
            {title: "Approved By"},
            {title: "Reason"},
            {title: "Requested dates"},
            {title: "Attachment"},
            {title: "Created At"},
            {title: "Created By"},
            {title: "Updated At"},
            {title: "Updated By"},
            {title: "Action"}
        ]
        
        if (data == null){
            DT.initiateMyTable('myLeaveTable', myColumns)
            $('#myLeaveTable').DataTable().clear().draw()
            return
        }

        data.forEach(element => { 
            let details = this.createRowsLD(element.details)
            let attachments = this.createAttachments(element.rowid)
            let row = document.createElement('tr')
            row.id = 'myLeave' + element.rowid
            row.innerHTML = `<td class="row-data" data-id=${element.leaveDefinition}>${element.leaveDefinitionCode} - ${element.leaveDefinitionName}</td>
                             <td class="row-data">${element.description}</td>
                             <td class="row-data" data-id=${element.statusid}>${element.status}</td>
                             <td class="row-data">${this.formatDate(element.approvedAt)}</td>
                             <td class="row-data" data-id=${element.approvedBy}>${allEmployees.get(Number(element.approvedBy))}</td>
                             <td class="row-data">${element.rejectedReason}</td>
                             <td class="row-data pointer" data-entries=${JSON.stringify(element.details, function replacer(key, value) { return value})}>${details}</td>
                             <td class="row-data myAttachments text-center" data-id=${element.rowid}>${attachments}</td>
                             <td class="row-data">${this.formatDate(element.createdAt)}</td>
                             <td class="row-data">${allEmployees.get(Number(element.createdBy))}</td>
                             <td class="row-data">${this.formatDate(element.updatedAt)}</td>
                             <td class="row-data">${allEmployees.get(Number(element.updatedBy))}</td>`
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

        DT.initiateMyTable('myLeaveTable', myColumns)
    } 

    // returns list of selected leave id to be deleted
    selectedLeave(){
        const checked = document.querySelectorAll('input[name=softDelete]:checked')
        if (checked.length == 0) document.querySelector('#deleteWarningMessageDiv').style.display = 'block'
        if (checked.length > 0){
            let allChecked = []
            checked.forEach(element => {
                allChecked.push(element.value)
            })
            
            return allChecked
        }        
    }
    // populate confirm detele message
    populateConfirmDelete(nb){
        let msg
        (nb == 1)? msg = 'Do you really want to delete this leave?' : msg = `Do you really want to delete ${nb} leaves?`

        const myBody = document.querySelector('#confirmDeleteBody')
        myBody.innerHTML = msg
    }

    // populate each selected dates
    populateSelectedDates(myDates){
        const target = document.querySelector('#selectedDatesDiv')
        target.innerHTML = ''
        let i = 0
        myDates.forEach(element => { 
            let row = document.createElement('div')
            row.classList = 'row'
            row.innerHTML = `<div class="col">${element}</div>
                             <div class="col">
                                <div class="form-check">
                                    <input class="form-check-input allSelectedDates" type="checkbox" value="${element}" id="isHalf${i}">
                                </div>
                             </div>`
            target.appendChild(row)
            i++
        })
        
    }

    // replace camelCase to camel case
    replaceCamelCase(str){
        let result
        (Boolean(str.match(/[A-Z]/))) ? result = str.replace(/([a-z])([A-Z])/g, '$1 $2').toLowerCase() : result = str
        return result
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

    // validate form 
    validateForm(myRequired){
        const myRequestedDates = document.querySelector('#requestedDates').value,
              totalRequested = Object.keys(myRequestedDates.split(',')).length,
              myDates = document.querySelectorAll('.allSelectedDates')
        
        let halfDays = 0.0,
            notValid = []

        // loop over required fields and check wether it's empty or not
        myRequired.forEach(field => {
            let myField = document.querySelector(`#${field}`)
            if (myField.value == '') {
                Common.insertHTML('This field is required', `${field}Error`)
                notValid.push(`${this.replaceCamelCase(field)} is required`)
            }else{
                Common.insertHTML('', `${field}Error`)
            }
        })

        // calculate total half days requested
        myDates.forEach(element => {
            const isHalf = this.convertMyBool(element.checked)
            if (isHalf == 1) {
                halfDays += 0.5
            }
        })

        // validate if total requested days is allowed (<= maxEntitledAllow )
        const finalRequestedDays = Number(totalRequested)-parseFloat(halfDays)
        if (finalRequestedDays > maxEntitledAllow){
            Common.insertHTML(`You can'\'t select more dates than ${maxEntitledAllow} days`, 'requestedDatesError')
            notValid.push('The maximun allowed days for this leave application is exceeded')
        }else{
            Common.insertHTML('', 'requestedDatesError')
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

    // get form
    getForm(){
        const form   = document.querySelector('#leaveForm'),
              data   = new FormData(form),
              myjson = {}

        myjson['leaveDefinition']   = data.get('leaveDefinition')
        myjson['description']       = data.get('description')
        myjson['employeeid']        = data.get('employeeid')
        myjson['details']           = this.getAllSelectedDates()

        return JSON.stringify(myjson, function replacer(key, value) { return value})
    }

    // get uploaded attachment
    getUploadedFiles(){
        var formData = []
        var fileInput = document.getElementById('uploadedFiles'); // Get the file input element
        var files = fileInput.files; // Get the selected files from the file input
        for (var i = 0; i < files.length; i++) {
          formData.push(files[i]); // Append each file to the FormData object with the field name 'files[]'
        }
        
        return formData
    }

    // populate all requested dates from calendar
    getAllSelectedDates(){
        const myDates = document.querySelectorAll('.allSelectedDates'),
              details = []

        myDates.forEach(element => {
            const entry = {
                requestedDate : element.value,
                isHalf : `${this.convertMyBool(element.checked)}`,
            }
            details.push(entry)
        })

        return details
    }

    // convert boolean to int
    convertMyBool(entry){
        let myVal
        return (entry)? myVal = 1 : myVal = 0
    }

    // populate all leave definition map
    populateLeaveDefinitionMap(data, allLeaveDefinition){
        data.forEach(element => {
            allLeaveDefinition.set(element.rowid, element)
        })

        return allLeaveDefinition
    }

    // populate my leave entitled map
    populateMyEntitledLeaveMap(data, myEntitled){
        data.forEach(element => {
            myEntitled.set(element.leaveDefinition, element)
        })

        return myEntitled
    }

    // populate all uploaded files map
    populateUploadedFilesMap(data, uploadedFiles){
        for (const [key, value] of Object.entries(data)) {
            uploadedFiles.set(key, value)
        }
        return uploadedFiles
    }

    // sort array
    sortArray(arr){
        arr.sort(function(a,b) {
            a = a.split('/').reverse().join('')
            b = b.split('/').reverse().join('')
            return a.localeCompare(b)
            // return a > b ? 1 : a < b ? -1 : 0; // <-- alternative   
        })
        return arr
    }

    // populate dates for disable calendar (already requested dates: pending & approved)
    populateDatesForDisable(myLeaves){
        let myLeaveDatesTaken = []

        // get requested dates pending & approved (to disable it)
        myLeaves.forEach(element => {
            if (element.statusid != 3){
                // requested dates
                let rd = element.details
                rd.forEach(e => {
                    myLeaveDatesTaken.push(this.formatDate(e.requestedDate))
                });
            }
        })

        // sort dates
        return this.sortArray(myLeaveDatesTaken) 
    }
    // populate all public holidays array (and sort it)
    populatePublicHolidays(data, ph){
         // get ph dates
        data.forEach(element => {
            ph.push(this.formatDate(element.date))
        });

        // sort dates
        return this.sortArray(ph)
    }

    // define gender connected user
    myGender(employeeid, employees){
        employees.forEach(element => {
            if (element.ID == employeeid){
                return element.Gender
            }
        });
    }

    // display leave definition requirements 
    // (attachment is required & max days allowed for selected leave definition)
    selectedLeaveRequirements(leaveDefinition, entitlement, mySeniority){        
        // check wether attachment is required or not
        this.attachmentRequired(leaveDefinition.docRequired)
        
        // calculate max days allowed for selected leave
        this.maxDaysAllowed(leaveDefinition, entitlement, mySeniority)
    }

    // display selected leave requirements regarding attachment
    attachmentRequired(required){
        attachmentRequired = required

        if (required == 1){
            Common.showDivByID('myAttachmentDiv')
        }else{
            Common.hideDivByID('myAttachmentDiv')
        }
    }

    // display selected leave requirements regarding max days allowed
    maxDaysAllowed(leaveDefinition, entitlement, mySeniority){
        let maxEntitled = 0.0
        // if limitationid != 3 (limited without advance) => selected leave limitation is either unlimited(1) or limited with advance (2)
        // => max allow = max leave definition entitled for respective seniority
        if (leaveDefinition.limitationid != 3){

            // sort leave details by seniority desc
            let leaveDetails = leaveDefinition.details
            leaveDetails.sort(function(a, b) {
                let keyA = Number(a.seniority),
                    keyB = Number(b.seniority)
                // Compare the 2 seniority
                if (keyA > keyB) return -1
                if (keyA < keyB) return 1
                return 0
            })

            // calculate max entitlement for employee's seniority
            leaveDetails.forEach(element => {
                if (Number(element.seniority) >= mySeniority){
                    maxEntitled      = Number(element.entitled) + parseFloat(entitlement.credits)
                    maxEntitledAllow = maxEntitled - parseFloat(entitlement.taken)
                    return
                }
            })

        }else{
            // calculate max entitlement for leave type limited without advance
            // => max = earned + credits
            maxEntitled      = parseFloat(entitlement.entitled) + parseFloat(entitlement.credits)
            maxEntitledAllow = maxEntitled - parseFloat(entitlement.taken)
        }

        // display selected leave requirements
        Common.insertHTML(`You are entitled for ${maxEntitled} days, you already took ${entitlement.taken} days. Your balance is ${maxEntitledAllow} days <br>`, 'myAuthorizationDiv')
        Common.showDivByID('myAuthorizationDiv')
    }

    // send attachment 
    SendAttachment(leaveApplicationID, connectedEmail, connectedID){
        Common.insertInputValue(leaveApplicationID, 'applicationID')
        Common.insertInputValue('leaves', 'uploadedFilename')
        Common.insertInputValue(connectedEmail, 'employeeEmail')
        Common.insertInputValue(connectedID, 'employeeID')
 
        document.getElementById("uploadedFilesForm").submit()
    }
}