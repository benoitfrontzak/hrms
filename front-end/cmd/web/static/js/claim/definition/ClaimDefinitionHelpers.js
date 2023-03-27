class ClaimDefinitionHelpers {
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

    // customize boolean (0|1) with icons
    myBooleanIcons(value){
        return (value == 1) ? '<i class="bi-check2-square"></i> true' : '<i class="bi-x-square"></i> false'
    }
    // insert to datatable one row per element of data
    insertRows(id, data) {
        const target = document.querySelector('#' + id)
        target.innerHTML = ''
        data.forEach(element => { 
            let opt = document.createElement('tr')
            opt.id = 'CD' + element.rowid
            opt.innerHTML = `<td class="row-data">${element.name}</td>
                             <td class="row-data">${element.description}</td>
                             <td class="row-data">${element.category}</td>
                             <td class="row-data">${this.myBooleanIcons(element.confirmation)}</td>
                             <td class="row-data">${element.seniority}</td>
                             <td class="row-data">${element.limitation}</td>
                             <td class="row-data">${this.myBooleanIcons(element.docRequired)}</td>
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
    } 
    // trigger datatable and row click event
    triggerDT(dt, dtBody) {
        const table = $('#' + dt).DataTable()
    }
    // generate data table
    generateDT(data, active){
        this.insertRows('claimDefinitionBody', data)
        this.triggerDT('claimDefinitionTable', 'claimDefinitionBody')
        Common.insertInputValue(active, 'dataView') // set dataView to active
    }
    // check wether data is null or not
    checkData(data, active){
        if (data == null ){
            document.querySelector('#gotData').style.display = 'none'
            document.querySelector('#noData').style.display = 'block'
        } else {
            document.querySelector('#gotData').style.display = 'block'
            document.querySelector('#noData').style.display = 'none'
            this.generateDT(data, active)
            this.makeEditable()
        }
    }

    // remove element from my required input fields array
    removeElementFromRIF(name){
        const index = myRIF.indexOf(name);
            if (index > -1) { // splice array only when item is found
                myRIF.splice(index, 1); // 2nd parameter means remove one item only
            }
    }

    // returns list of selected claim definition id to be deleted
    selectedClaimDefinition(){
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
        (nb == 1)? msg = 'Do you really want to delete this claim?' : msg = `Do you really want to delete ${nb} claims?`

        const myBody = document.querySelector('#confirmDeleteBody')
        myBody.innerHTML = msg
    }

}