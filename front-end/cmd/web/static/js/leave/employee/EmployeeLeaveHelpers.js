class EmployeeLeaveHelpers {

    // populate option list with active employee
    populateEmployeeList(data){
        const target = document.querySelector('#datalistOptions')
        target.innerHTML = ''
        data.forEach(element => {
            const row = document.createElement('option')
            row.id = element.ID
            row.value = element.Fullname
            target.appendChild(row)
        })
    }

    // update connected user leaves (details)
    populateMyLeaveDetails(data){
        // populate main card with leave details
        const target = document.querySelector('#leaveDetailsBody')
        target.innerHTML = ''

        data.forEach(element => {
            const balance = Number(this.cleanDecimal(element.entitled)) + Number(element.credits) - Number(element.taken)
            let row = document.createElement('div')
            row.classList = 'row justify-content-md-center'
            row.innerHTML = `<div class="col-4 text-start"><input type="text" class="form-control form-control-sm transparentInput" value="${element.leaveDefinitionCode} - ${element.leaveDefinitionName}" disabled /></div>
                             <div class="col-2 text-start"><input type="text" class="form-control form-control-sm transparentInput" value="${this.cleanDecimal(element.entitled)}" disabled /></div>
                             <div class="col-2 text-start"><input type="number" step="0.01" class="form-control form-control-sm transparentInput allCredits" data-rowid="${element.rowid}" value="${element.credits}" /></div>
                             <div class="col-2 text-start"><input type="text" class="form-control form-control-sm transparentInput" value="${element.taken}" disabled /></div>
                             <div class="col-2 text-start"><input type="text" class="form-control form-control-sm transparentInput" value="${balance}" disabled /></div>`

            target.appendChild(row)
        })
    }

     // update connected user leaves (progress details)
     populateMyLeaveDetailsProgress(data){
        // populate main card with leave details
        const target = document.querySelector('#chartLeaves')
        target.innerHTML = ''

        data.forEach(element => {
            const balance = Number(this.cleanDecimal(element.entitled)) + Number(element.credits) - Number(element.taken),
                  max = Number(this.cleanDecimal(element.entitled)) + Number(element.credits)
     
            let percentage
            (element.taken != 0)? percentage = Math.round(Number(element.taken)*100/max) : percentage = 0

            let row = document.createElement('div')
            row.classList = 'mb-3 p-3 border borderRadiusTop borderRadiusBottom'
            row.innerHTML = `<div class="row mb-3">
                                <div class="col">${element.leaveDefinitionCode} - ${element.leaveDefinitionName}</div>
                                <div class="col text-end">${element.taken}/${max} days</div>
                            </div>
                            <div class="row">
                                <div class="col">
                                    <div class="progress">
                                        <div class="progress-bar myHarmonyBlue" role="progressbar" style="width: ${percentage}%" aria-valuenow="${percentage}" aria-valuemin="0" aria-valuemax="${max}">${percentage}%</div>
                                    </div>
                                </div>
                             </div>`

            target.appendChild(row)
        }) 
    }

    // clean decimal value to follow rule: only .5 is allowed
    cleanDecimal(v){
        let cleaned = v
        // check if string contains decimal value (ie got a .)
        if (v.includes(".")){
            const n = Number(v),
                  int = Number(v.split('.')[0]),
                  fract = n%1
                  
            let cFract = 0
            if (fract.toFixed(4) >= 0.5) cFract = 0.5
            cleaned = int + cFract
        }
        return cleaned.toString()
    }
    
    // get form data
    getForm(connectedID, connectedEmail){
        const allCredits = []
        // fetch all credits of employee
        const rows = document.querySelectorAll('.allCredits')
        rows.forEach(credits => {
            const oneCredits = {
                rowid:credits.dataset.rowid,
                credits:credits.value,
                connectedUser:connectedEmail,
                userID:connectedID
            }
            allCredits.push(oneCredits)
        })
        
        return JSON.stringify(allCredits, function replacer(key, value) { return value})
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
    populateAttachments(appID, employeeEmail){
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
                                                <div><img class="myPicture" src="/upload/${employeeEmail}/leaves/${appID}/${element}" /></div>
                                                <div class="text-end"><a href="/upload/${employeeEmail}/leaves/${appID}/${element}" class="myLink" download="${element}"><i class="bi-download pointer"></i> download</a></div>
                                           </div>`
                target.appendChild(oneAttachment)
            })
            myModal.show()
        }
    }

    // insert to datatable one row per element of data
    insertRows(data) {
        const target = document.querySelector('#employeeLeaveBody')
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
            {title: "Updated By"}
        ]
        
        // can initiate DT only one time
        if ( $.fn.dataTable.isDataTable( '#employeeLeaveTable' ) ) {
            $('#employeeLeaveTable').DataTable().destroy()
        }

        if (data == null){
            DT.initiateMyTable('employeeLeaveTable', myColumns)
            $('#employeeLeaveTable').DataTable().clear().draw()
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
                                       
            target.appendChild(row)
        })

        DT.initiateMyTable('employeeLeaveTable', myColumns)
    }

    // convert timestamp (remove timestamp)
    formatDate(t){
        let b
        // get only the 10 first characters of the string
        const d = t.substring(0,10)
        // the zero value of a date is 0001-01-01
        return (d == '0001-01-01') ? b = '' : b = d
    }

    // populate all uploaded files map
    populateUploadedFilesMap(data, uploadedFiles){
        for (const [key, value] of Object.entries(data)) {
            uploadedFiles.set(key, value)
        }
        return uploadedFiles
    }
}