class MyClaimHelpers {
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
            (typeof element.rowid == 'undefined') ? myID = 0 : myID = element.rowid
            let opt = document.createElement('option')
            opt.classList = 'myclaimDefinition'
            opt.id = 'myclaimDefinition' + myID
            opt.value = myID
            opt.innerHTML = element.name
            opt.dataset.confirmation = element.confirmation
            opt.dataset.seniority = element.seniority
            opt.dataset.docRequired = element.docRequired
            opt.dataset.limitation = element.limitation
            target.appendChild(opt)
        })
    }

    // add user information to connectedUser
    updateUserInformation(data){
        connectedUser.code          = data.Code
        connectedUser.fullname     = data.Fullname
        connectedUser.nickname      = data.Nickname
        connectedUser.joinDate      = data.JoinDate
        connectedUser.confirmDate   = data.ConfirmDate
    }

    // remove element from my required input fields array
    removeElementFromRIF(name){
        const index = myRIF.indexOf(name);
            if (index > -1) { // splice array only when item is found
                myRIF.splice(index, 1); // 2nd parameter means remove one item only
            }
    }

    // show | hide category, name from form
    // add | remove category, name from required list
    switchFormFieldsView(id){
        if (id == 0 ) {
            Common.showDivByID('categoryDiv')
            Common.showDivByID('nameDiv')
            myRIF.push('category', 'name')
        }else{
            Common.hideDivByID('categoryDiv')
            Common.hideDivByID('nameDiv')
            this.removeElementFromRIF('category')
            this.removeElementFromRIF('name')
        }
    }
    // populate hidden field (form: confirmation | seniority | docRequired | limitation)
    populateHiddenFields(data){
        Common.insertInputValue(data.confirmation, 'confirmation')
        Common.insertInputValue(data.seniority, 'seniority')
        Common.insertInputValue(data.docRequired, 'docRequired')
        Common.insertInputValue(data.limitation, 'limitation')
    }

    // convert birthdate (remove timestamp)
    formatDate(t){
        let b
        // get only the 10 first characters of the string
        const d = t.substring(0,10)
        // the zero value of a date is 0001-01-01
        return (d == '0001-01-01') ? b = '' : b = d
    }
    // insert to datatable one row per element of data
    insertRows(data) {
        const target = document.querySelector('#myClaimBody')
        target.innerHTML = ''
        data.forEach(element => { 
            let opt = document.createElement('tr')
            opt.id = 'myClaim' + element.rowid
            opt.innerHTML = `<td class="row-data" data-id="${element.claimDefinitionID}">${element.claimDefinition}</td>
                             <td class="row-data" data-id="${element.categoryID}">${element.category}</td>
                             <td class="row-data">${element.name}</td>
                             <td class="row-data">${element.description}</td>
                             <td class="row-data">${element.amount}</td>
                             <td class="row-data" data-id="${element.statusID}">${element.status}</td>
                             <td class="row-data">${this.formatDate(element.approvedAt)}</td>
                             <td class="row-data">${allEmployees.get(Number(element.approvedBy))}</td>
                             <td class="row-data">${element.approvedAmount}</td>
                             <td class="row-data">${element.approvedReason}</td>
                             <td>
                                <div class="d-flex justify-content-center">
                                    <div class="form-check">
                                        <input class="form-check-input deleteCheckboxes"  type="checkbox" value="${element.rowid}" name="softDelete">
                                        <label class="form-check-label fw-lighter fst-italic smaller" for="softDelete"><i class="bi-trash2-fill largeIcon pointer deleteClaim"></i></label>
                                    </div>
                                    <div class="form-check">
                                        <label class="form-check-label fw-lighter fst-italic smaller"><i class="bi-pencil-fill largeIcon pointer editClaim" data-id=${element.rowid}></i></label>
                                    </div>
                                </div>                                                            
                            </td>`
            target.appendChild(opt)
        })
        $('#myClaimTable').DataTable()
    } 

    // returns list of selected claim id to be deleted
    selectedClaim(){
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
        (nb == 1)? msg = 'Do you really want to delete this claim?' : msg = `Do you really want to delete ${nb} claims?`

        const myBody = document.querySelector('#confirmDeleteBody')
        myBody.innerHTML = msg
    }
}