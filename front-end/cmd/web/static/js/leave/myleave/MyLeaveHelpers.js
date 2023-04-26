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
                console.log(e.target);
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
    insertOptions(id, data) {
        const target = document.querySelector('#' + id)
        target.innerHTML = '<option selected hidden value=""></option>'
        data.forEach(element => {
            let opt = document.createElement('option')
            opt.value = element.rowid
            opt.innerHTML = element.code + ' - ' + element.description
            target.appendChild(opt)
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
    // create rows leave details (modal)
    createRowsLD(entries){
        let details =``

        entries.forEach(e => {
            details += `<div class="d-flex justify-content-between">
                        <div>${this.formatDate(e.requestedDate)}</div>
                        <div>${this.myBooleanIcons(e.isHalf)}</div>
                     </div>`
        })

        return details
    }
    // insert to datatable one row per element of data
    insertRows(id, data) {
        const target = document.querySelector('#' + id)
        target.innerHTML = ''
        data.forEach(element => { 
            let details = this.createRowsLD(element.details)
            let row = document.createElement('tr')
            row.id = 'myLeave' + element.rowid
            row.innerHTML = `<td class="row-data" data-id=${element.leaveDefinition}>${element.leaveDefinitionCode} - ${element.leaveDefinitionName}</td>
                             <td class="row-data">${element.description}</td>
                             <td class="row-data" data-id=${element.statusid}>${element.status}</td>
                             <td class="row-data">${this.formatDate(element.approvedAt)}</td>
                             <td class="row-data" data-id=${element.approvedBy}>${allEmployees.get(Number(element.approvedBy))}</td>
                             <td class="row-data">${element.rejectedReason}</td>
                             <td class="row-data pointer" data-entries=${JSON.stringify(element.details, function replacer(key, value) { return value})} data-bs-toggle="modal" data-bs-target="#details${element.rowid}">details</td>
                             <td>
                                <div class="row">
                                    <div class="col">
                                        <div class="form-check">
                                            <label class="form-check-label fw-lighter fst-italic smaller"><i class="bi-pencil-fill largeIcon pointer editLeave" data-id=${element.rowid}></i></label>
                                        </div>
                                    </div>
                                    <div class="col">
                                        <div class="form-check">
                                            <label class="form-check-label fw-lighter fst-italic smaller" for="softDelete"><i class="bi-trash2-fill largeIcon pointer deleteLeave"></i></label>
                                            <input class="form-check-input deleteCheckboxes"  type="checkbox" value="${element.rowid}" name="softDelete">                                    
                                        </div> 
                                    </div>
                                </div>                                                 
                            </td>
                            <div class="modal fade" id="details${element.rowid}" tabindex="-1">
                                <div class="modal-dialog">
                                    <div class="modal-content">
                                    <div class="modal-header">
                                        <h5 class="modal-title">Leave details ${element.leaveDefinitionCode} - ${element.leaveDefinitionName}</h5>
                                        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                                    </div>
                                    <div class="modal-body">
                                        <div class="d-flex justify-content-between">
                                            <div>Dates</div>
                                            <div>Half day</div>
                                        </div>
                                        ${details}
                                    </div>
                                    <div class="modal-footer">
                                        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">close</button>
                                    </div>
                                    </div>
                                </div>
                            </div>`
            target.appendChild(row)
        })
    } 
    // trigger datatable and row click event
    triggerDT(dt) {
        const table = $('#' + dt).DataTable()
    }
    // generate data table
    generateDT(data){
        this.insertRows('myLeaveBody', data)
        this.triggerDT('myLeaveTable')
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
            console.log('allChecked: '+allChecked);
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

    getAllSelectedDates(){
        const myDates = document.querySelectorAll('.allSelectedDates'),
              details = []

        myDates.forEach(element => {
            const entry = {
                requestedDate : element.value,
                isHalf : this.convertMyBool(element.checked),
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
}