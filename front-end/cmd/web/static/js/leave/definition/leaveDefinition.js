const Common  = new MainHelpers(),
      DT      = new DataTableFeatures(),
      Helpers = new LeaveDefinitionHelpers(),
      API     = new LeaveDefinitionAPI()

// set form's parameters (Required Input Fields...)
const myRIF = [ 'code', 'description', 'expiry', 'gender', 'limitation', 'calculation', 'seniority0', 'entitled0']

// set row details number to 1
let rowDetailsNumber = 1

// store rowID of existing row details to be deleted
const rowDetailsDelete = []


// when DOM is loaded
window.addEventListener('DOMContentLoaded', () => {

    // fetch all leave's definition & update DOM 
    API.getAllLeaveDefinition().then(resp => {
        console.log(resp.data);
        Helpers.insertRows(resp.data)
        // when edit icon is clicked
        Helpers.makeEditable()
    })

    // fetch leave CT & update DOM (form dropdown)
    API.getLeaveCT().then(resp => {
        Helpers.insertOptions('limitation', resp.data.Limitation)
        Helpers.insertOptions('calculation', resp.data.Calculation)
    })

    // when add details row is clicked
    document.querySelector('#addDetails').addEventListener('click', () => {
        Helpers.insertDetailsRow()
    })
    
    // when open form (add button) is clicked (clear warning message & field's values)
    document.querySelector('#openCreateLeaveDefinition').addEventListener('click', () =>{ 
        Common.clearForm('createLeaveDefinitionForm', myRIF) 
    })

    // when form is submitted (save button)
    document.querySelector('#createLeaveDefinitionSubmit').addEventListener('click', () => {
        const error = Common.validateRequiredFields(myRIF),
              formAction = document.querySelector('#formAction').value

        // create form   
        if (error == '0' && formAction == 'create'){
            myData = Helpers.getForm('createLeaveDefinitionForm')
            API.createLeaveDefinition(myData).then(resp => { 
                if (! resp.error) location.reload()
            })            
        }

        // edit form
        if (error == '0' && formAction == 'edit'){
            myData = Helpers.getForm('createLeaveDefinitionForm')
            API.updateLeaveDefinition(myData).then(resp => {
                console.log(resp);
                if (! resp.error) location.reload()
            })            
        }
    })

    // close warning message for delete modal
    document.querySelector('#hidedeleteWarningMessage').addEventListener('click', () => {
        Common.hideDivByID('deleteWarningMessageDiv')
    })

    // close warning message for form modal
    document.querySelector('#hideWarningMessage').addEventListener('click', () => {
        Common.hideDivByID('warningMessageDiv')
    })
    
    // initiate delete confirm modal
    const myConfirm = new bootstrap.Modal(document.getElementById('confirmDelete'), { 
        keyboard: false 
    })

    // cleaned checked checkboxes when modal is close
    document.getElementById('confirmDelete').addEventListener('hidden.bs.modal', function (event) {
        const myDeleteCheckboxes = document.querySelectorAll('.deleteCheckboxes')

        myDeleteCheckboxes.forEach(element => {
            element.checked = false
        })

    })

    // when delete all is clicked
    const myDeleteAll = document.querySelector('#deleteAllLeaveDefinition')

    myDeleteAll.addEventListener('click', () => { 
        const checked = Helpers.selectedLeaveDefinition()
       
        if (typeof checked != 'undefined' && checked.length > 0){
            Helpers.populateConfirmDelete(checked.length)
            myConfirm.show()
        }

    })

    // when confirm delete all is clicked
    const confirmedDelete = document.querySelector('#confirmDeleteSubmit')

    confirmedDelete.addEventListener('click', () => {
        const checked = Helpers.selectedLeaveDefinition()

        API.softDeleteLeaveDefinition(checked, connectedEmail).then(resp => {
            if (!resp.error){
                myConfirm.hide()
                location.reload()
            }
        })

    })
    
})
