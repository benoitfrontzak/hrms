class LeaveHelpers {
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
    formatDate(t){
        let b
        // get only the 10 first characters of the string
        const d = t.substring(0,10)
        // the zero value of a date is 0001-01-01
        return (d == '0001-01-01') ? b = '' : b = d
    }

    // convert isHalf
    leaveDay(isHalf){
        let result
        return (isHalf == 1) ? result = 'Full <i class="bi-square-fill"></i>' : result = 'Half <i class="bi-square-half"></i>'
    }
    // create rows leave requested dates (modal)
    createRowsLRD(entries){
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
        
    // insert to datatable one row per element of data
    insertRows(data, action=true) {
        // $('#myLeaveTable').DataTable().destroy
        const target = document.querySelector('#myLeaveBody')
        target.innerHTML = ''

        const myColumns = [
            {title: "Employee"},
            {title: "Leave"},
            {title: "Description"},
            {title: "Approved At"},
            {title: "Approved By"},
            {title: "Reason"},
            {title: "Dates"},
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
            let details = this.createRowsLRD(element.details)
            let row = document.createElement('tr')
            row.id = 'leave' + element.rowid
            row.innerHTML = `<td class="row-data">${allEmployees.get(Number(element.employeeid))}</td>
                             <td class="row-data" data-id="${element.leaveDefinition}">${element.leaveDefinitionCode} - ${element.leaveDefinitionName}</td>
                             <td class="row-data">${element.description}</td>
                             <td class="row-data">${this.formatDate(element.approvedAt)}</td>
                             <td class="row-data">${allEmployees.get(Number(element.approvedBy))}</td>
                             <td class="row-data">${element.rejectedReason}</td>
                             <td class="row-data">${details}</td>
                             <td class="row-data">${this.formatDate(element.createdAt)}</td>
                             <td class="row-data">${allEmployees.get(Number(element.createdBy))}</td>
                             <td class="row-data">${this.formatDate(element.updatedAt)}</td>
                             <td class="row-data">${allEmployees.get(Number(element.createdBy))}</td>`
            if (action){
                Common.showDivByID('approveBtn')
                Common.showDivByID('rejectBtn')
                row.innerHTML += `
                             <td>
                                <div class="d-flex justify-content-between">
                                    <div class="form-check">
                                        <label class="form-check-label" for="leaves"><i class="bi-person-check-fill largeIcon"></i></label>
                                        <input class="form-check-input leavesCheckboxes"  type="checkbox" value="${element.rowid}" name="leaves">
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
        DT.initiateMyTable('myLeaveTable', myColumns)
        
    } 

    // returns list of selected leave id to be deleted
    selectedLeave(){
        const checked = document.querySelectorAll('input[name=leaves]:checked')
        if (checked.length == 0) document.querySelector('#modalWarningMessageDiv').style.display = 'block'
        if (checked.length > 0){
            let allChecked = []
            checked.forEach(element => {
                allChecked.push(element.value)
            })
            return allChecked
        }        
    }
    // populate confirm detele message
    populateSelectedLeavesNumber(nb, id, action, comment){
        let msg

        (nb == 1)? msg = `You're about to ${action} one leave ${comment}` : msg = `You're about to ${action} ${nb} leave ${comment}`

        const myBody = document.querySelector('#'+id)
        myBody.innerHTML = msg
    }
}