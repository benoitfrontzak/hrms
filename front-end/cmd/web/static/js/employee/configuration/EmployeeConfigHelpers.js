class EmployeeConfigTablesHelpers{
    // make link sidebar active
    makeLinkActive(wanted){
        Common.insertInputValue(wanted,'ctName')
        // make all links inactive
        const myCT = document.querySelectorAll('.myConfigT')
        myCT.forEach(element => {
            element.className = 'side-nav-link myConfigT'
        })
        // make wanted actived
        document.querySelector('#'+wanted).className = 'side-nav-link myConfigT border-end border-danger myTint6BG'
    }

    // display proper CT since payroll item is different
    displayCT(isPI){
        if (isPI){
            // For Payroll Item table
            Common.showDivByID('piTable')
            Common.hideDivByID('configTable')
            Common.showDivByID('openModalPI')
            Common.hideDivByID('openModal')
        }else{
            // For all other config tables (only got name field)
            Common.hideDivByID('piTable')
            Common.showDivByID('configTable')
            Common.hideDivByID('openModalPI')
            Common.showDivByID('openModal')
        }
    }

    // populate config table information
    populateCT(data, title, id){
        // add config table name (title & hidden input)
        Common.insertHTML(title, 'configTableTitle')
        Common.insertInputValue(id,'ctName') // for delete purpose

        // manage case PayrollItem (not real CT: don't have only name field)
        if (id == 'PayrollItem'){
            this.displayCT(true)
            this.insertRowsPI(data)
        }else{
            this.displayCT(false)
            this.insertRows(data)
        }       
 
    }
    // create payroll item tooltips
    createTooltipsPI(element){
        let myTooltips = `<div class="row">
                            <div class="col-8 text-start">Pay EPF</div>  
                            <div class="col-4">${this.myBooleanIcons(element.payEPF)}</div>  
                          </div>
                          <div class="row">
                            <div class="col-8 text-start">Pay SOCSO</div>  
                            <div class="col-4">${this.myBooleanIcons(element.paySOCSO)}</div>  
                          </div>
                          <div class="row">
                            <div class="col-8 text-start">Pay EIF</div>  
                            <div class="col-4">${this.myBooleanIcons(element.paySOCSO)}</div>  
                          </div>
                          <div class="row">
                            <div class="col-8 text-start">Pay HRDF</div>  
                            <div class="col-4">${this.myBooleanIcons(element.payHRDF)}</div>  
                          </div>
                          <div class="row">
                            <div class="col-8 text-start">Pay Tax</div>  
                            <div class="col-4">${this.myBooleanIcons(element.payTax)}</div>  
                          </div>
                          <div class="row">
                            <div class="col-8 text-start">Is Fixed</div>  
                            <div class="col-4">${this.myBooleanIcons(element.isFixed)}</div>  
                          </div>`

        return myTooltips
    }
    // Insert PayrollItem to table 
    insertRowsPI(data){
        const target = document.querySelector('#piTableData')
        target.innerHTML = ''
        data.forEach(element => {
            let tooltips = this.createTooltipsPI(element)
            let row = document.createElement('div')
            row.id = element.id
            row.classList = 'row'
            row.innerHTML = `<div class="col pointer" data-bs-toggle="tooltip" data-bs-html="true" title='${tooltips}'><input type="text" class="form-control form-control-sm transparentInput" id="piType${element.id}" value="${element.type}" disabled /></div>
                             <div class="col showSave" data-id="${element.id}"><input type="text" class="pointer form-control form-control-sm transparentInput" id="piCode${element.id}" value="${element.code}" /></div>
                             <div class="col-5 showSave" data-id="${element.id}"><input type="text" class="pointer form-control form-control-sm transparentInput" id="piDescription${element.id}" value="${element.description}" /></div>
                             <div class="col showSave" data-id="${element.id}"><input type="number" step="0.01" class="pointer form-control form-control-sm transparentInput" id="piAmount${element.id}" value="${element.amount}" /></div>
                             <div class="col showSave" data-id="${element.id}"><input type="date" class="pointer form-control form-control-sm transparentInput" id="piStart${element.id}" value="${this.formatDate(element.start)}" /></div>
                             <div class="col showSave" data-id="${element.id}"><input type="date" class="pointer form-control form-control-sm transparentInput" id="piEnd${element.id}" value="${this.formatDate(element.end)}" /></div>
                             
                            <div class="col">
                                <div class="form-check" id="deleteDivPI${element.id}">
                                    <input class="form-check-input deleteCheckboxes"  type="checkbox" value="${element.id}" name="softDelete">
                                    <label class="form-check-label fw-lighter fst-italic smaller" for="softDelete"><i class="bi-trash2-fill largeIcon"></i></label>
                                </div>
                                <div class="hide" id="saveDivPI${element.id}"><a href="#" class="btn btn-danger saveValuePI" id="save${element.id}" data-id="${element.id}"><i class="bi-check-lg"></i> save</a></div>                           
                            </div>`
            target.appendChild(row)
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
    myBooleanIcons(value) {
        return (value == 1) ? '<i class="bi-check2-square"></i>' : '<i class="bi-x-square"></i>'
    }
    createRow(element){
        let row = document.createElement('div')
            row.id = element.ID
            row.classList = 'd-flex justify-content-between'
            row.innerHTML = `<div class="showValue pointer" id="showDiv${element.ID}" data-id="${element.ID}"> ${element.Name}</div>
                            <div class="hide" id="editDiv${element.ID}"> 
                                <input type="text" class="form-control form-control-sm" id="editValue${element.ID}" name="editValue${element.ID}" value="${element.Name}" />
                            </div>
                            <div>
                                <div class="form-check" id="deleteDiv${element.ID}">
                                    <input class="form-check-input deleteCheckboxes"  type="checkbox" value="${element.ID}" name="softDelete">
                                    <label class="form-check-label fw-lighter fst-italic smaller" for="softDelete"><i class="bi-trash2-fill largeIcon"></i></label>
                                </div>
                                <div class="hide" id="saveDiv${element.ID}"><a href="#" class="btn btn-danger saveValue" id="save${element.ID}" data-id="${element.ID}"><i class="bi-check-lg"></i> save</a></div>                           
                            </div>`
        return row
    }
    createRowDisabled(element){
        let row = document.createElement('div')
            row.id = element.ID
            row.classList = 'd-flex justify-content-between'
            row.innerHTML = `<div> ${element.Name}</div>
                             <div><i class="bi-lock-fill largeIcon"></i></div>`
        return row
    }
    // Insert to table one row per element of data
    insertRows(data){
        const target = document.querySelector('#configTableData')
        target.innerHTML = ''
        data.forEach(element => {
            if (element.ID == 0){
                let row = this.createRowDisabled(element)
                target.appendChild(row)
            }else{
                let row = this.createRow(element)
                target.appendChild(row)
            }
            
        })
    }
    // make CT row editable
    editableRow(rowID){
        const showDiv   = 'showDiv'  + rowID,
              editDiv   = 'editDiv'  + rowID,
              deleteDiv = 'deleteDiv'+ rowID,
              saveDiv   = 'saveDiv'  + rowID
        Common.hideDivByID(showDiv)
        Common.hideDivByID(deleteDiv)
        Common.showDivByID(editDiv)
        Common.showDivByID(saveDiv)
    }
    // select all checked checkbox
    selectedEntry(){
        const checked = document.querySelectorAll('input[name=softDelete]:checked')
        if (checked.length == 0) document.querySelector('#'+myDeleteWarning).style.display = 'block'
        if (checked.length > 0){
            let allChecked = []
            checked.forEach(element => {
                allChecked.push(element.value)
            })
      
            return allChecked
        }        
    }

    // clean PI form
    cleanForm(){
        Common.insertInputValue('', 'piCode')
        Common.insertInputValue('', 'piDescription')
        Common.insertInputValue('', 'piStart')
        Common.insertInputValue('', 'piEnd')
        Common.insertInputValue('', 'piAmount')
        document.getElementById("isFixed").checked = false;
        document.getElementById("payEPF").checked = false;
        document.getElementById("paySOCSO").checked = false;
        document.getElementById("payHRDF").checked = false;
        document.getElementById("payTax").checked = false;
    }

    // validate PI form
    validateForm(){
        let notValid = []
        
        // validate payroll item is selected
        if (document.querySelector('#piCode').value == ''){
            document.querySelector('#piCodeError').innerHTML = 'Code is required'
            notValid.push('Code is required')
        }

        // validate amount is set
        if (document.querySelector('#piAmount').value == ''){
            document.querySelector('#piAmountError').innerHTML = 'Amount is required'
            notValid.push('Amount is required')
        }

        return notValid.length
    }

    // get PI form
    getForm(){
        const form = document.querySelector(`#piForm`),
              data = new FormData(form),
              myjson = {}
        
        let myStart = '0001-01-01',
            myEnd = '0001-01-01'
        
        if (document.querySelector('#piStart').value != '') myStart = document.querySelector('#piStart').value
        if (document.querySelector('#piEnd').value != '') myEnd = document.querySelector('#piEnd').value

        myjson['type']          = (data.get('piType') == '0') ? '0' : '1'
        myjson['code']          = data.get('piCode')
        myjson['description']   = data.get('piDescription')
        myjson['start']         = myStart
        myjson['end']           = myEnd
        myjson['amount']        = document.querySelector('#piAmount').value
        myjson['isFixed']       = (data.get('isFixed') == null) ? '0' : '1'
        myjson['payEPF']        = (data.get('payEPF') == null) ? '0' : '1'
        myjson['paySOCSO']      = (data.get('paySOCSO') == null) ? '0' : '1'
        myjson['payHRDF']       = (data.get('payHRDF') == null) ? '0' : '1'
        myjson['payTax']        = (data.get('payTax') == null) ? '0' : '1'

        return JSON.stringify(myjson, function replacer(key, value) { return value})
  
    }

    // get PI row
    getRowPI(rowID){
        const myjson = {}
        
        let myStart = '0001-01-01',
            myEnd = '0001-01-01'
        
        if (document.querySelector('#piStart'+rowID).value != '') myStart = document.querySelector('#piStart'+rowID).value
        if (document.querySelector('#piEnd'+rowID).value != '') myEnd = document.querySelector('#piEnd'+rowID).value

        myjson['id']            = rowID
        myjson['code']          = document.querySelector('#piCode'+rowID).value
        myjson['description']   = document.querySelector('#piDescription'+rowID).value
        myjson['start']         = myStart
        myjson['end']           = myEnd
        myjson['amount']        = document.querySelector('#piAmount'+rowID).value

        return JSON.stringify(myjson, function replacer(key, value) { return value})
    }
}