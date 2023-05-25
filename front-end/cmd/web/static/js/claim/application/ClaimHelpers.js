class ClaimHelpers {

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

    // convert birthdate (remove timestamp)
    formatDate(t) {
        let b
        // get only the 10 first characters of the string
        const d = t.substring(0, 10)
        // the zero value of a date is 0001-01-01
        return (d == '0001-01-01') ? b = '' : b = d
    }


    // insert to datatable one row per element of data
    insertRows(data, action = true) {        
        const target = document.querySelector('#myClaimBody')
        target.innerHTML = ''

        const myColumns = [
            {title: "Employee"},
            {title: "Allowance"},
            {title: "Category"},
            {title: "Description"},
            {title: "Amount"},
            {title: "Approved At"},
            {title: "Approved By"},
            {title: "Approved Amount"},
            {title: "Reason"},
            {title: "Attachments"},
            {title: "Created At"},
            {title: "Created By"},
            {title: "Updated At"},
            {title: "Updated By"},
            {title: "Action"}
        ]

        if (data == null){
            DT.initiateMyTable('claimApplicationsTable', myColumns)
            $('#claimApplicationsTable').DataTable().clear().draw()
            return
        }
        
        data.forEach(element => {
            // since omitempty has been set golang side...
            let approvedBy, approvedAmount;
            (typeof element.approvedAmount == 'undefined') ? approvedAmount = 0 : approvedAmount = element.approvedAmount;
            (typeof element.approvedBy == 'undefined') ? approvedBy = 0 : approvedBy = Number(element.approvedBy);

            let row = document.createElement('tr')
            row.id = 'leave' + element.rowid
            row.innerHTML = `<td class="row-data">${allEmployees.get(Number(element.employeeid))}</td>
                             <td class="row-data" data-id="${element.claimDefinitionID}">${element.claimDefinition}</td>
                             <td class="row-data" data-id="${element.categoryID}">${element.category}</td>
                             <td class="row-data">${element.description}</td>
                             <td class="row-data">${element.amount}</td>
                             <td class="row-data">${this.formatDate(element.approvedAt)}</td>
                             <td class="row-data">${allEmployees.get(approvedBy)}</td>
                             <td class="row-data">${approvedAmount}</td>
                             <td class="row-data">${element.rejectedReason}</td>
                             <td class="row-data myAttachments" data-appid=${element.rowid}>${element.employeeid}</td>
                             <td class="row-data">${this.formatDate(element.createdAt)}</td>
                             <td class="row-data">${allEmployees.get(Number(element.createdBy))}</td>
                             <td class="row-data">${this.formatDate(element.updatedAt)}</td>
                             <td class="row-data">${allEmployees.get(Number(element.updatedBy))}</td>`
            if (action){
                Common.showDivByID('approveBtn')
                Common.showDivByID('rejectBtn')
                row.innerHTML += `
                             <td>
                                <div class="d-flex justify-content-between">
                                    <div class="form-check">
                                        <label class="form-check-label" for="claims"><i class="bi-person-check-fill largeIcon"></i></label>
                                        <input class="form-check-input claimsCheckboxes"  type="checkbox" value="${element.rowid}" name="claims">
                                    </div>
                                </div>                                                            
                            </td>`
            }else{
                Common.hideDivByID('approveBtn')
                Common.hideDivByID('rejectBtn')
                row.innerHTML += `<td><i class="bi-lock"></i></td>`
            }

            target.appendChild(row)
        })
        DT.initiateMyTable('claimApplicationsTable', myColumns)
        
        // update attachments
        const table = $('#claimApplicationsTable').DataTable(),
              cells = table.cells('.myAttachments');

        // loop over all cells
        cells.every(function() {
            const cell = this,
                    employeeID = cell.data(),
                    appID = $(cell.node()).attr('data-appid')
            // fetch employee email
            API.getEmployeeEmailByID(employeeID).then(resp => {
                const email = resp.data
                // set email attribute
                $(cell.node()).attr('data-email', email);
                API.getUploadedFiles(email).then(r => {
                    let icon
                    if (r.data.Files[appID] !== undefined) {
                        // set entries attribute
                        $(cell.node()).attr('data-entries', JSON.stringify(r.data.Files[appID], function replacer(key, value) { return value}))
                        icon='<i class="bi-check-square pointer"></i>'
                    }else{
                        icon='<i class="bi-file-excel-fill not-allowed"></i>'
                    } 
                    cell.data(icon)
                })
                
            })
        });
    }

    // populate attachments modal & open it
    populateAttachments(appID, email, entries, icon){
        if (icon != '<i class="bi-file-excel-fill not-allowed"></i>'){
            const target = document.querySelector('#uploadedFilesModalBody'),
                  myModal = new bootstrap.Modal(document.getElementById('uploadedFilesModal'), { 
                    backdrop: 'static',
                    keyboard: false 
                  })
            target.innerHTML = ''
            JSON.parse(entries).forEach(element => {
                const oneAttachment = document.createElement('div')
                oneAttachment.classList = 'd-flex justify-content-center mb-3'
                oneAttachment.innerHTML = `<div>
                                                <div><img class="myPicture" src="/upload/${email}/claims/${appID}/${element}" /></div>
                                                <div class="text-end"><a href="/upload/${email}/claims/${appID}/${element}" class="myLink" download="${element}"><i class="bi-download pointer"></i> download</a></div>
                                           </div>`
                target.appendChild(oneAttachment)
            })
            myModal.show()
        }
    }

    // populate uploaded files
    populateUploadedFiles(data, emailAndClaimID) {
        let email = emailAndClaimID.split('/')[0]
        let claimID = emailAndClaimID.split('/')[1]
        if (data.Files.claimDoc != null && data.Files.claimDoc.length > 0) {
            this.insertAttachments(data.Files.claimDoc, email, 'uploadedClaimAttachmentBody', 'claimDoc', claimID)
        }
        // if (data.Archive.ic != null && Object.keys(data.Archive.ic).length > 0) {
        //     this.insertArchiveAttachments(data.Archive.ic, email, 'uploadedClaimAttachmentBody', 'archive/ic', data.ClaimID)
        // }
    }

    // create attachment download link
    insertAttachments(data, email, id, category, claimID) {
        const target = document.querySelector('#' + id)
        target.innerHTML = ''

        data.forEach(element => {
            let row = document.createElement('div')
            row.classList = 'd-flex justify-content-center mb-3'
            row.innerHTML = `<div>
                                    <div><img class="myPicture" src="/upload/${email}/${category}/claims/${element}" /></div>
                                    <div class="text-end"><a href="/upload/${email}/${category}/claims/${element}" class="myLink" download="${element}"><i class="bi-download pointer"></i> download</a></div>
                                 </div>`
            target.appendChild(row)
        })

    }


    // returns list of selected claim id to be deleted
    selectedClaim() {
        const checked = document.querySelectorAll('input[name=claims]:checked')
        if (checked.length == 0) document.querySelector('#modalWarningMessageDiv').style.display = 'block'
        if (checked.length > 0) {
            let allChecked = []
            checked.forEach(element => {
                allChecked.push(element.value)
            })
            return allChecked
        }
    }
    // populate confirm detele message
    populateSelectedClaimsNumber(nb, id, action, verb) {
        let msg

        (nb == 1) ? msg = `You're about to ${action} one claim with this ${verb}:` : msg = `You're about to ${action} ${nb} claims with this ${verb}:`

        const myBody = document.querySelector('#' + id)
        myBody.innerHTML = msg
    }
}