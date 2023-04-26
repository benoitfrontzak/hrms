class LeaveDefinitionHelpers {
    // populate form to edit entry
    populateFormEdit(data, rowID){
        // hide add button
        Common.hideDivByID('addDetailsBtn')

        // update title & form action & rowID
        Common.insertHTML('Edit Leave', 'createLeaveDefinitionTitle')
        Common.insertInputValue('edit', 'formAction')
        Common.insertInputValue(rowID, 'rowid')

        // update default values
        Common.insertInputValue(data[0].innerHTML, 'code')
        Common.insertInputValue(data[1].innerHTML, 'description')
        Common.checkRadio(data[2].dataset.id, 'unpaid')
        Common.checkRadio(data[3].dataset.id, 'replacement')
        Common.checkRadio(data[4].dataset.id, 'attachment')
        Common.selectBoxOptionSelected(data[5].dataset.id, 'expiry')
        Common.selectBoxOptionSelected(data[6].dataset.id, 'gender')
        Common.selectBoxOptionSelected(data[7].dataset.id, 'limitation')
        Common.selectBoxOptionSelected(data[8].dataset.id, 'calculation')

        // reset default details row
        Common.insertHTML('', 'defaultDetailRow')

        // set details rows number to 1
        rowDetailsNumber = 0

        // insert existing details rows
        const myJson = this.parseJson(data[8].dataset.entries)
        myJson.forEach(element => {
            this.insertExistingDetailsRow(element.seniority, element.entitled, element.rowid)
        })

 
    }
    // populate and launch edit modal form
    makeEditable(){
        document.querySelectorAll('.editLeave').forEach(element => {
            element.addEventListener('click', (e) => {
                const rowID = e.target.dataset.id,
                      rowData = document.getElementById('LD'+rowID).querySelectorAll('.row-data')
   
                Helpers.populateFormEdit(rowData, rowID)
                const myModalForm = new bootstrap.Modal(document.getElementById('createLeaveDefinition'), { 
                    keyboard: false 
                })
                myModalForm.show()
            })
        })
    }
    
    // convert stringify json to object
    parseJson(stringifiedJSON){
        let myObj = JSON.parse(stringifiedJSON, (key, value) => { return value })
        return myObj
    }

    // insert one existing row details
    insertExistingDetailsRow(seniority, entitled, rowID){
        // actual rowID number to be inserted
        const n = rowDetailsNumber
        
        // create new row
        const target = document.querySelector('#existingDetailRows')
        let row = document.createElement('div')
        row.id = 'existingDetails'+n
        row.dataset.rowID = rowID
        row.classList = 'row mb-3 existingRowDetails'
        row.innerHTML = `<div class="col-sm-4">
                            <div class="form-floating">
                            <input type="text" class="form-control existingSeniority" id="seniority${n}" name="seniority${n}" value="${seniority}" />
                            <label for="seniority"><i class="bi-shield-fill-exclamation"></i> seniority</label>
                            </div>
                            <div class="fw-bolder text-danger fst-italic smaller existingSeniorityError" id="seniority${n}Error"></div>
                        </div>
                        <div class="col-sm-4">
                            <div class="form-floating">
                            <input type="text" class="form-control existingEntitled" id="entitled${n}" name="entitled${n}" value="${entitled}" />
                            <label for="entitled"><i class="bi-shield-fill-exclamation"></i> entitled</label>
                            </div>
                            <div class="fw-bolder text-danger fst-italic smaller existingEntitledError" id="entitled${n}Error"></div>
                        </div> 
                        <div class="col-sm-4"><i class="bi-trash2-fill largeIcon pointer" data-id="${n}" id="deleteExistingDetailsRow${n}"></i></div>`
        target.appendChild(row)

        // set new elements as required
        myRIF.push('seniority'+n)
        myRIF.push('entitled'+n)

        // update rowID number
        rowDetailsNumber++

        // create event listener for delete row
        document.querySelector('#deleteExistingDetailsRow'+n).addEventListener('click', (e) => {
            const myRow = document.querySelector('#existingDetails'+n)
            
            // store rowID to delete
            rowDetailsDelete.push(myRow.dataset.rowID)

            // remove DOM element   
            myRow.remove()

            // remove all extra details rows from RIF
            this.unrequireAllDetailsRows(rowDetailsNumber)

            // update rows number
            rowDetailsNumber--

            // recreate extra details rows to required list
            this.recreateRequiredField(rowDetailsNumber)

            // reindex all row's elements
            const indexStart = 0
            this.reindexByClassname('existingRowDetails', 'existingDetails', indexStart)
            this.reindexByClassname('existingSeniority', 'seniority', indexStart)
            this.reindexErrorByClassname('existingSeniorityError', 'seniority', indexStart)
            this.reindexByClassname('existingEntitled', 'entitled', indexStart)
            this.reindexErrorByClassname('existingEntitledError', 'entitled', indexStart)

        })
    }

    // insert one new row details
    insertDetailsRow(){
        // actual rowID number to be inserted
        const n = rowDetailsNumber
        
        // create new row
        const target = document.querySelector('#extraDetailRows')
        let row = document.createElement('div')
        row.id = 'newDetails'+n
        row.classList = 'row mb-3 newRowDetails'
        row.innerHTML = `<div class="col-sm-4">
                            <div class="form-floating">
                            <input type="text" class="form-control newSeniority" id="seniority${n}" name="seniority${n}" />
                            <label for="seniority"><i class="bi-shield-fill-exclamation"></i> seniority</label>
                            </div>
                            <div class="fw-bolder text-danger fst-italic smaller newSeniorityError" id="seniority${n}Error"></div>
                        </div>
                        <div class="col-sm-4">
                            <div class="form-floating">
                            <input type="text" class="form-control newEntitled" id="entitled${n}" name="entitled${n}" />
                            <label for="entitled"><i class="bi-shield-fill-exclamation"></i> entitled</label>
                            </div>
                            <div class="fw-bolder text-danger fst-italic smaller newEntitledError" id="entitled${n}Error"></div>
                        </div> 
                        <div class="col-sm-4"><i class="bi-trash2-fill largeIcon pointer" data-id="${n}" id="deleteDetailsRow${n}"></i></div>`
        target.appendChild(row)

        // set new elements as required
        myRIF.push('seniority'+n)
        myRIF.push('entitled'+n)

        // update rowID number
        rowDetailsNumber++

        // create event listener for delete row
        document.querySelector('#deleteDetailsRow'+n).addEventListener('click', (e) => {
            const myRow = document.querySelector('#newDetails'+n)
            
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
            this.reindexByClassname('newEntitled', 'entitled')
            this.reindexErrorByClassname('newEntitledError', 'entitled')

        })
    }

    // reindex input name & id from classname
    reindexByClassname(classname, inputname, indexStart = 1){
        let i = indexStart
        document.querySelectorAll('.'+classname).forEach(element => {
            element.id = inputname+i
            element.name = inputname+i
            i++
        })
    }

    // reindex error div from classname
    reindexErrorByClassname(classname, inputname, indexStart = 1){
        let i = indexStart
        document.querySelectorAll('.'+classname).forEach(element => {
            element.id = inputname+i+'Error'
            i++
        })
    }

    // remove all added details rows from RIF
    unrequireAllDetailsRows(n){
        for (let index = 1; index < n; index++) {
            this.removeElementFromRIF('seniority'+index)
            this.removeElementFromRIF('entitled'+index)            
        }
    }

    // remove element from my required input fields array
    removeElementFromRIF(name){
        const index = myRIF.indexOf(name);
            if (index > -1) { // splice array only when item is found
                myRIF.splice(index, 1); // 2nd parameter means remove one item only
            }
    }

    // recreate extra details rows to RIF
    recreateRequiredField(n){
        for (let index = 1; index < n; index++) {
            myRIF.push('seniority'+index)           
            myRIF.push('entitled'+index)           
        }        
    }

    // insert options into select tag
    insertOptions(id, data) {
        const target = document.querySelector('#' + id)
        target.innerHTML = '<option selected hidden value="">Please select a value</option>'
        data.forEach(element => {
            let opt = document.createElement('option')
            opt.value = element.ID
            opt.innerHTML = element.Name
            target.appendChild(opt)
        })
    }

    // customize boolean (0|1) with icons
    myBooleanIcons(value){
        return (value == 1) ? '<i class="bi-check-square"></i>' : '<i class="bi-square"></i>'
    }

    // convert month digit
    myMonth(d){
        let month
        switch (d) {
            case '0':
                month = ''
                break;
            case '1':
                month = 'January'
                break;
            case '2':
                month = 'February'
                break;
            case '3':
                month = 'March'
                break;
            case '4':
                month = 'April'
                break;
            case '5':
                month = 'May'
                break;
            case '6':
                month = 'June'
                break;
            case '7':
                month = 'July'
                break;
            case '8':
                month = 'August'
                break;
            case '9':
                month = 'September'
                break;
            case '10':
                month = 'October'
                break;
            case '11':
                month = 'November'
                break;
            case '12':
                month = 'December'
                break;
            default:
                break;
        }
        return month
    }
    
    // create rows leave definition details (modal)
    createSeniorityRows(entries){
        let details =`<div>N/A</div>`

        if (entries.length > 1){
            entries.forEach(e => {
                details += `<div>From year ${e.seniority}</div>`
            })
        }

        return details
    }
    createEntitledRows(entries){
        let details =`<div>${entries[0].entitled} days</div>`

        if (entries.length > 1){
            entries.forEach(e => {
                details += `<div>${e.entitled} days</div>`
            })
        }

        return details
    }
    // insert to datatable one row per element of data
    insertRows(data) {
        const target = document.querySelector('#leaveDefinitionBody')
        target.innerHTML = ''
        
        if (data != null){
            data.forEach(element => {
                let seniority = this.createSeniorityRows(element.details),
                    entitled = this.createEntitledRows(element.details)
                let row = document.createElement('tr')
                row.id = 'LD' + element.rowid
                row.innerHTML = `<td class="row-data">${element.code}</td>
                                 <td class="row-data">${element.description}</td>
                                 <td class="row-data" data-id=${element.unpaid}>${this.myBooleanIcons(element.unpaid)}</td>
                                 <td class="row-data" data-id=${element.replacementRequired}>${this.myBooleanIcons(element.replacementRequired)}</td>
                                 <td class="row-data" data-id=${element.docRequired}>${this.myBooleanIcons(element.docRequired)}</td>
                                 <td class="row-data" data-id=${element.expiry}>${this.myMonth(element.expiry)}</td>
                                 <td class="row-data" data-id=${element.genderid}>${element.gender}</td>
                                 <td class="row-data" data-id=${element.limitationid}>${element.limitation}</td>
                                 <td class="row-data" data-id=${element.calculationid}>${element.calculation}</td>
                                 <td class="row-data">${seniority}</td>
                                 <td class="row-data">${entitled}</td>
                                 <td class="row-data">${element.createdBy}</td>
                                 <td class="row-data">${element.createdAt}</td>
                                 <td class="row-data">${element.updatedBy}</td>
                                 <td class="row-data">${element.updatedAt}</td>
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
                                </td>`
                target.appendChild(row)
            })
       
            const table = $('#leaveDefinitionTable').DataTable({
                columns:[
                    {title: "Code"},
                    {title: "Description"},
                    {title: "Is Unpaid"},
                    {title: "Requires Replacement"},
                    {title: "Requires Attachment"},
                    {title: "Expiry"},
                    {title: "Gender"},
                    {title: "Limitation"},
                    {title: "Entitlement Calculation Method"},
                    {title: "Requires Seniority"},
                    {title: "Max Entitled"},
                    {title: "Created By"},
                    {title: "Created At"},
                    {title: "Updated By"},
                    {title: "Updated At"},
                    {title: "Action"}
                  ],
                scrollX:        true,
                scrollCollapse: true,
                fixedColumns:   {
                    left: 1,
                    right: 1
                },        
                colReorder: true,
                lengthMenu: [
                    [10, 25, 50, -1],
                    ['10 rows', '25 rows', '50 rows', 'Show all']
                ],
                "bInfo" : false,
                "initComplete": () => {
                    $('#processing').remove()
                }
            })
            //Default columns to display
            const defaultConf = ['Code','Description','Is_Unpaid','Requires_Replacement','Requires_Attachment','Expiry','Gender','Limitation','Entitlement_Calculation Method','Requires_Seniority','Max_Entitled','Created_By','Created_At','Updated_By','Updated_At', 'Action']
            DT.showHideColumnsDT(table, defaultConf, 'columnsLeaveDefinition', 'columnMenu')
            // DT.draggableColumn('leaveDefinitionTable')
        }else{
            $('#leaveDefinitionTable').DataTable()
            $('#leaveDefinitionTable').DataTable().clear().draw()
        }
        
    } 

    // returns list of selected claim definition id to be deleted
    selectedLeaveDefinition(){
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

    // get form elements
    getForm(formID){
        const form   = document.querySelector(`#${formID}`),
              data   = new FormData(form),
              myjson = {}

        myjson['rowid']         = data.get('rowid')
        myjson['code']          = data.get('code')
        myjson['description']   = data.get('description')
        myjson['expiry']        = data.get('expiry')
        myjson['genderid']      = data.get('gender')
        myjson['limitationid']  = data.get('limitation')
        myjson['calculationid'] = data.get('calculation')
        myjson['unpaid']        = data.get('unpaid')
        myjson['replacement']   = data.get('replacement')
        myjson['attachment']    = data.get('attachment')
        myjson['details']       = this.getDetailsRowsValue()
        myjson['detailsDeleted']= rowDetailsDelete
        myjson['createdBy']     = connectedID
        myjson['UpdatedBy']     = connectedID
        
        return JSON.stringify(myjson, function replacer(key, value) { return value})
    } 

    getDetailsRowsValue(){
        const myRows = []
        for (let i = 0; i < rowDetailsNumber; i++) {
            const seniority = document.querySelector('#seniority'+i),
                  entitled  = document.querySelector('#entitled'+i),
                  myRow     = document.querySelector('#existingDetails'+i)
            let entry
            if (typeof myRow !== 'undefined' && myRow !== null){
                entry     = {seniority : seniority.value, entitled : entitled.value, rowid : myRow.dataset.rowID}
            }else{
                entry     = {seniority : seniority.value, entitled : entitled.value}
            }
                  
            myRows.push(entry)            
        }
        return myRows
    }

}