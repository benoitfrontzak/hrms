class ClaimConfigTablesHelpers{
    // populate CT list to data view dropdown
    populateViewData(data){
        // remove country CT from data
        delete data['Country']
        // populate each ct name to view dropdown
        const target = document.querySelector('#viewList')
        target.innerHTML = ''
        for (let key in data){
            let li = document.createElement('li')
            li.innerHTML = `<a class="dropdown-item myConfigT" href="#">${key}</a>`
            target.appendChild(li)
        }
    }

    // populate config table information
    populateCT(data, title){
        // remove first element (id=0; name='not defined')
        data.shift()
        // add config table name (title & hidden input)
        Common.insertHTML(title, 'configTableTitle')
        Common.insertInputValue(title,'ctName')
        // populate config table rows
        this.insertRows(data)
        // show action button (from main nav)
        Common.showDivByID('openModal')
        Common.showDivByID('deleteCT')
    }
    // Insert to table one row per element of data
    insertRows(data){
        const target = document.querySelector('#configTableData')
        target.innerHTML = ''
        data.forEach(element => {
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
            target.appendChild(row)
        })
    }

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
}