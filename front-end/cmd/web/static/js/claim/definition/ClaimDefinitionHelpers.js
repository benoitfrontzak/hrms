class ClaimDefinitionHelpers {
    // populate form to edit entry
    populateFormEdit(data, rowID) {
        Common.insertHTML('Edit claim', 'createClaimDefinitionTitle')
        Common.insertInputValue('edit', 'formAction')
        Common.insertInputValue(rowID, 'rowid')
        Common.insertInputValue(data[0].innerHTML, 'name')
        Common.insertInputValue(data[1].innerHTML, 'description')
        Common.selectBoxOptionSelected(data[2].innerHTML, 'category')
        Common.checkRadio(1, 'active')
        Common.checkRadio(data[5].innerHTML, 'docRequired')
        this.insertExistingDetailsRow(data);
    }
    // populate and launch edit modal form
    makeEditable() {
        document.querySelectorAll('.editClaim').forEach(element => {
            element.addEventListener('click', (e) => {
                const rowID = e.target.dataset.id,
                    rowData = document.getElementById('CD' + rowID).querySelectorAll('.row-data')
                const myModalForm = new bootstrap.Modal(document.getElementById('createClaimDefinition'), {
                    keyboard: false
                })
                myModalForm.show()
                Helpers.removeRowFromDivsByClass('detailRemovable')
                Helpers.populateFormEdit(rowData, rowID)
                const toggleCheckboxes = document.querySelectorAll(".toggle-checkbox");

                toggleCheckboxes.forEach(function (checkbox) {
                    checkbox.addEventListener("change", function () {
                        const row = checkbox.closest(".row");
                        const inputBoxes = row.querySelectorAll(".input-row");
                        if (checkbox.checked) {
                            row.classList.add("disabled");
                        } else {
                            row.classList.remove("disabled");
                        }
                    });
                });
            });
        });
    }



    // insert one new row details
    insertExistingDetailsRow(data) {
        // create new row
        const target = document.querySelector('#existingDetailRows')
        let amountOfExistingDetails = data[3].childNodes.length
        for (let elem = 0; elem < amountOfExistingDetails; elem++) {
            let row = document.createElement('div')
            row.id = 'existingDetail'
            row.classList = 'row mb-3 detailRemovable'
            let existingSeniorityID = data[3].childNodes[elem].id
            let existingLimitationID = data[4].childNodes[elem].id
            //extract id only
            let idOnly = this.extractIDfromRowID(existingLimitationID)
            let seniorityExist = data[3].childNodes[elem].innerText.replace('From year ', '')
            row.innerHTML = `<div class='existingDetailsRow row' id="${idOnly}">
                                <div class="col-sm-4 existingDetails">
                                    <div class="form-floating">
                                        <input type="text" class="form-control input-row existingSeniority" id="${existingSeniorityID}" name="${existingSeniorityID}" value="${seniorityExist}"/>
                                            <label for="seniority"><i class="bi-shield-fill-exclamation"></i> seniority</label>
                                    </div>
                                    <div class="fw-bolder text-danger fst-italic smaller existingSeniorityError" id="${existingSeniorityID}Error">
                                    </div>
                                </div>
                                <div class="col-sm-4 existingDetails">
                                    <div class="form-floating">
                                        <input type="text" class="form-control input-row existingLimitation" id="${existingLimitationID}" name="${existingLimitationID}" value="${data[4].childNodes[elem].innerText}" />
                                            <label for="limitation"><i class="bi-shield-fill-exclamation"></i> limitation</label>
                                    </div>
                                    <div class="fw-bolder text-danger fst-italic smaller existingLimitationError" id="${existingLimitationID}Error">
                                    </div>
                                </div> 
                                <div class="col-sm-4 existingDetails">
                                    <input type="checkbox" class="toggle-checkbox" /><i class="bi-trash2-fill largeIcon pointer"></i>
                                </div>
                            </div>`
            target.appendChild(row)


            // set new elements as required
            myRIF.push(existingSeniorityID)
            myRIF.push(existingLimitationID)
        }
    }

        // insert one new row details
    insertDefaultDetailsRow() {
        // create new row
        const target = document.querySelector('#defaultDetailRow')
        let row = document.createElement('div')
        row.classList = 'row mb-3 detailRemovable'
        row.innerHTML =`<!-- Claim definition details Seniority -->
                        <div class="col-sm-4">
                            <div class="form-floating">
                                <input type="text" class="form-control" id="seniority0" name="seniority0" value="0" />
                                <label for="seniority"><i class="bi-shield-fill-exclamation"></i> Seniority</label>
                            </div>
                            <div class="fw-bolder text-danger fst-italic smaller" id="seniority0Error"></div>
                            </div>
                            <!-- Claim definition details Limitation -->
                            <div class="col-sm-4">
                            <div class="form-floating">
                                <input type="text" class="form-control" id="limitation0" name="limitation0" />
                                <label for="limitation"><i class="bi-shield-fill-exclamation"></i> Limitation</label>
                            </div>
                            <div class="fw-bolder text-danger fst-italic smaller" id="limitation0Error"></div>
                            </div>
                        <div class="col-sm-4"></div>
                        `
        target.appendChild(row)
        myRIF.push('limitation0')
        myRIF.push('seniority0')
    }


    // insert one new row details
    insertDetailsRow() {
        // actual rowID number to be inserted
        const n = rowDetailsNumber

        // create new row
        const target = document.querySelector('#extraDetailRows')
        let row = document.createElement('div')
        
        row.id = 'newDetails' + n
        row.classList = 'row mb-3 newRowDetails detailRemovable'
        row.innerHTML = `<div class='newDetailsRow row'>
                            <div class="col-sm-4 detailRemovable">
                                <div class="form-floating">
                                    <input type="text" class="form-control input-row newSeniority" id="seniority${n}" 
                                        name="seniority${n}" />
                                    <label for="seniority"><i class="bi-shield-fill-exclamation"></i> seniority</label>
                                </div>
                                <div class="fw-bolder text-danger fst-italic smaller newSeniorityError" id="seniority${n}Error">
                                </div>
                            </div>
                            <div class="col-sm-4 detailRemovable">
                                <div class="form-floating">
                                    <input type="text" class="form-control input-row newLimitation" id="limitation${n}"
                                        name="limitation${n}" />
                                    <label for="limitation"><i class="bi-shield-fill-exclamation"></i> limitation</label>
                                </div>
                                <div class="fw-bolder text-danger fst-italic smaller newLimitationError" id="limitation${n}Error">
                                </div>
                            </div>
                            <div class="col-sm-4 detailRemovable">
                                <i class="bi-trash2-fill largeIcon pointer deleteDetailRowsButton" data-id="${n}" id="deleteDetailsRow${n}">
                                </i>
                            </div>
                        </div>`
        target.appendChild(row)

        // set new elements as required
        myRIF.push('seniority' + n)
        myRIF.push('limitation' + n)

        // update rowID number
        rowDetailsNumber++

        // create event listener for delete row
        //TODO: discuss with Ben whether to change to class instead of n ... quite confuse on how this works
        document.querySelector('#deleteDetailsRow' + n).addEventListener('click', (e) => {

            //e.currentTarget.id will capture the clicked button id.
            let result = Helpers.extractIDfromRowID(e.currentTarget.id);
            let clickID = result ? result : n; //get the id from click.... technically won't be n but just place it for now.

            const myRow = document.querySelector('#newDetails' + clickID)

            // remove DOM element   
            myRow.remove()

            // remove all extra details rows from RIF
            this.unrequireAllDetailsRows(rowDetailsNumber)

            // update rows number
            rowDetailsNumber--

            // recreate extra details rows to required list
            this.recreateRequiredField(rowDetailsNumber)

            // reindex all row's elements
            this.reindexByClassname('newRowDetails', 'newDetails')
            this.reindexByClassname('newSeniority', 'seniority')
            this.reindexErrorByClassname('newSeniorityError', 'seniority')
            this.reindexByClassname('newLimitation', 'limitation')
            this.reindexErrorByClassname('newLimitationError', 'limitation')
            this.reindexByClassname('deleteDetailRowsButton', 'deleteDetailsRow')

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

    // reindex input name & id from classname
    reindexByClassname(classname, inputname, indexStart = 1) {
        let i = indexStart
        document.querySelectorAll('.' + classname).forEach(element => {
            element.id = inputname + i
            element.name = inputname + i
            i++
        })
    }

    // reindex error div from classname
    reindexErrorByClassname(classname, inputname, indexStart = 1) {
        let i = indexStart
        document.querySelectorAll('.' + classname).forEach(element => {
            element.id = inputname + i + 'Error'
            i++
        })
    }

    // remove all added details rows from RIF
    unrequireAllDetailsRows(n) {
        for (let index = 1; index < n; index++) {
            this.removeElementFromRIF('seniority' + index)
            this.removeElementFromRIF('limitation' + index)
        }
    }

    // remove element from my required input fields array
    removeElementFromRIF(name) {
        const index = myRIF.indexOf(name);
        if (index > -1) { // splice array only when item is found
            myRIF.splice(index, 1); // 2nd parameter means remove one item only
        }
    }

    // recreate extra details rows to RIF
    recreateRequiredField(n) {
        for (let index = 1; index < n; index++) {
            myRIF.push('seniority' + index)
            myRIF.push('limitation' + index)
        }
    }

    // customize boolean (0|1) with icons
    myBooleanIcons(value) {
        return (value == 1) ? '<i class="bi-check2-square"></i> true' : '<i class="bi-x-square"></i> false'
    }


    // create rows leave definition details (modal)
    createSeniorityRows(entries) {
        let details = ``

        if (entries.length > 1) {
            entries.forEach(e => {
                details += `<div id='existingSeniority${e.rowid}'>From year ${e.seniority}</div>`
            })
        } else {
            details = `<div>N/A</div>`
        }

        return details
    }
    createLimitationRows(entries) {
        let details = ``

        if (entries.length > 1) {
            entries.forEach(e => {
                details += `<div id='existingLimitation${e.rowid}'>${e.limitation}</div>`
            })
        } else {
            details = `<div id='existingLimitation${entries[0].rowid}'>${entries[0].limitation}</div>`
        }

        return details
    }

    // remove timestamp from date yyyy-mm-dd
    removeTimestamp(t){
        let myDate
        // get only the 10 first characters of the string
        const d = t.substring(0,10)
  
        // the zero value of a date is 0001-01-01
        return (d == '0001-01-01') ? myDate = '' : myDate = d
    }

    // insert to datatable one row per element of data
    insertRows(data) {
        const target = document.querySelector('#claimDefinitionBody')
        target.innerHTML = ''

        const myColumns = [
            {title: "Name"},
            {title: "Description"},
            {title: "Category"},
            {title: "Seniority"},
            {title: "Limitation"},
            {title: "Requires document"},
            {title: "Created By"},
            {title: "Created At"},
            {title: "Updated By"},
            {title: "Updated At"},
            {title: "Action"}
        ]

        // when no data
        if (data == null){
          DT.initiateMyTable('claimDefinitionTable', myColumns)
          $('#claimDefinitionTable').DataTable().clear().draw()
          return
        }

        data.forEach(element => {
            let seniority = this.createSeniorityRows(element.details),
                limitation = this.createLimitationRows(element.details)
            let row = document.createElement('tr')
            row.id = 'CD' + element.rowid
            row.innerHTML = `<td class="row-data">${element.name}</td>
                             <td class="row-data">${element.description}</td>
                             <td class="row-data">${element.category}</td>
                             <td class="row-data">${seniority}</td>
                             <td class="row-data">${limitation}</td>
                             <td class="row-data">${this.myBooleanIcons(element.docRequired)}</td>
                             <td class="row-data">${allEmployees.get(Number(element.createdBy))}</td>
                             <td class="row-data">${this.removeTimestamp(element.createdAt)}</td>
                             <td class="row-data">${allEmployees.get(Number(element.updatedBy))}</td>
                             <td class="row-data">${this.removeTimestamp(element.updatedAt)}</td>
                             <td>
                             <div class="row">
                                <div class="col">
                                    <div class="form-check">
                                        <label class="form-check-label fw-lighter fst-italic smaller"><i class="bi-pencil-fill largeIcon pointer editClaim" data-id=${element.rowid}></i></label>
                                    </div>
                                </div>
                                <div class="col">
                                    <div class="form-check">
                                        <input class="form-check-input deleteCheckboxes"  type="checkbox" value="${element.rowid}" name="softDelete">
                                        <label class="form-check-label fw-lighter fst-italic smaller" for="softDelete"><i class="bi-trash2-fill largeIcon pointer deleteClaim"></i></label>
                                    </div>
                                </div>
                             </div>                                                           
                            </td>`
            target.appendChild(row)
        })

        // initiate datatable & show/hide columns
        DT.initiateMyTable('claimDefinitionTable', myColumns)
        
    }

    // remove element from my required input fields array
    removeElementFromRIF(name) {
        const index = myRIF.indexOf(name);
        if (index > -1) { // splice array only when item is found
            myRIF.splice(index, 1); // 2nd parameter means remove one item only
        }
    }

    // returns list of selected claim definition id to be deleted
    selectedClaimDefinition() {
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

    //extract the id from the row(Typically it would be a button beside the row)
    extractIDfromRowID(rowId) {
        let result = rowId.match(/\D+(\d+)$/);
        return result ? result[1] : null;
    }

    //remove rows from div by class
    removeRowFromDivsByClass(className) {
        let divs = document.querySelectorAll('.' + className)
        divs.forEach(element => {
            element.remove()
        })
    }

    // extract detail from form 
    extractDetailsFromForm(myData) {
        let dataObj = JSON.parse(myData)
        let details = []
        let detailObj = {seniority: '', limitation: ''}
    
        //get all dataObj keys with seniority(n) and limitation(n)
        for (const [key, value] of Object.entries(dataObj)) {
            if (key.includes('seniority')) {
                detailObj.seniority = value
            }
            if (key.includes('limitation')) {
                detailObj.limitation = value
            }
            // push only when both are not empty}
            if (detailObj.seniority != '' && detailObj.limitation != ''){ 
                details.push(detailObj)
                detailObj = {seniority: '', limitation: ''}
            }

        }
        dataObj.details = details
        return JSON.stringify(dataObj)
    }

    // add update details and delete row id array to form
    addUpdateDetailsAndDeleterowIDToForm(myData, updateDetails, deleteRowID) {
        let dataObj = JSON.parse(myData)
        dataObj.detailsUpdate= updateDetails
        dataObj.detailsDeleted = deleteRowID
        return JSON.stringify(dataObj)
    }

}