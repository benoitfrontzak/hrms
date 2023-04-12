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

    // create rows leave requested dates (modal)
    createRowsLRD(entries){
        let details =``

        if (entries.length > 1){
            entries.forEach(e => {
                details += `<div class="d-flex justify-content-between">
                            <div>From year ${e.requestedDate}</div>
                            <div>${e.isHalf} days</div>
                         </div>`
            })
        }else{
            details = `<div class="d-flex justify-content-between">
                        <div>No condition</div>
                        <div>${entries[0].requestedDate} (${entries[0].isHalf} )</div>
                    </div>`
        }

        return details
    }
        
    // insert to datatable one row per element of data
    insertRows(data, action=true) {
        $('#myLeaveTable').DataTable().destroy
        const target = document.querySelector('#myLeaveBody')
        target.innerHTML = ''
        if(data != null){
            data.forEach(element => { 
                let details = this.createRowsLRD(element.details)
                let row = document.createElement('tr')
                row.id = 'leave' + element.rowid
                if (action){
                    Common.showDivByID('approveBtn')
                    Common.showDivByID('rejectBtn')
                    row.innerHTML = `<td class="row-data">${allEmployees.get(Number(element.employeeid))}</td>
                                 <td class="row-data" data-id="${element.leaveDefinition}">${element.leaveDefinitionCode} - ${element.leaveDefinitionName}</td>
                                 <td class="row-data">${element.description}</td>
                                 <td class="row-data" data-id="${element.statusid}">${element.status}</td>
                                 <td class="row-data">${this.formatDate(element.approvedAt)}</td>
                                 <td class="row-data">${allEmployees.get(Number(element.approvedBy))}</td>
                                 <td class="row-data">${element.rejectedReason}</td>
                                 <td class="row-data" data-entries=${JSON.stringify(element.details, function replacer(key, value) { return value})} data-bs-toggle="modal" data-bs-target="#details${element.rowid}">dates</td>
                                 <td>
                                    <div class="d-flex justify-content-between">
                                        <div class="form-check">
                                            <label class="form-check-label" for="leaves"><i class="bi-person-check-fill largeIcon"></i></label>
                                            <input class="form-check-input leavesCheckboxes"  type="checkbox" value="${element.rowid}" name="leaves">
                                        </div>
                                    </div>                                                            
                                </td>
                                <div class="modal fade" id="details${element.rowid}" tabindex="-1">
                                    <div class="modal-dialog">
                                        <div class="modal-content">
                                        <div class="modal-header">
                                            <h5 class="modal-title">Leave reuqested dates ${element.leaveDefinitionCode} - ${element.leaveDefinitionName}</h5>
                                            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                                        </div>
                                        <div class="modal-body">
                                            ${details}
                                        </div>
                                        <div class="modal-footer">
                                            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">close</button>
                                        </div>
                                        </div>
                                    </div>
                                </div>`
                }else{
                    Common.hideDivByID('approveBtn')
                    Common.hideDivByID('rejectBtn')
                    row.innerHTML = `<td class="row-data">${allEmployees.get(Number(element.employeeid))}</td>
                                     <td class="row-data" data-id="${element.leaveDefinition}">${element.leaveDefinitionCode} - ${element.leaveDefinitionName}</td>
                                     <td class="row-data">${element.description}</td>
                                     <td class="row-data" data-id="${element.statusid}">${element.status}</td>
                                     <td class="row-data">${this.formatDate(element.approvedAt)}</td>
                                     <td class="row-data">${allEmployees.get(Number(element.approvedBy))}</td>
                                     <td class="row-data">${element.rejectedReason}</td>
                                     <td><i class="bi-lock"></i></td>`
                }
                
                target.appendChild(row)
            })
            $('#myLeaveTable').DataTable()
        }else{
            $('#myLeaveTable').DataTable()
            $('#myLeaveTable').DataTable().clear().draw()
        }
        
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