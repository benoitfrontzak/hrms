const Common  = new MainHelpers(),
      Helpers = new ClaimDefinitionHelpers(),
      API     = new ClaimDefinitionAPI()

// set form's parameters (Required Input Fields...)
const myRIF = [ 'name', 'description', 'category','limitation', 'seniority']

// when DOM is loaded
window.addEventListener('DOMContentLoaded', () => {

    // fetch all claim's definition & update DOM 
    API.getAllClaimDefinition().then(resp => {
        const active = 1,
              inactive = 0
        // display by default active claim definition
        Helpers.checkData(resp.data.Active, active)
        // when active claim definition is requested
        document.querySelector('#activeBtn').addEventListener('click', () => {
            Helpers.checkData(resp.data.Active, active)
        })
        // when inactive claim definition is requested
        document.querySelector('#inactiveBtn').addEventListener('click', () => {
            Helpers.checkData(resp.data.Inactive, inactive)
        })            
        // when deleted claim definition is requested
        document.querySelector('#deletedBtn').addEventListener('click', () => {
            Helpers.checkData(resp.data.Deleted, inactive)
        })        
    })

    // fetch claim category & update DOM (form dropdown)
    API.getClaimCT().then(resp => {
        Helpers.insertOptions('category', resp.data.Category)
    })

    // when open form (add button) is clicked (clear warning message & field's values)
    document.querySelector('#openCreateClaimDefinition').addEventListener('click', () =>{ 
        Common.clearForm('createClaimDefinitionForm', myRIF)
        Common.insertInputValue('create', 'formAction')
    })

    // when confirmation required is clicked
    document.querySelector('#confirmation1').addEventListener('click', (e) => {
        Common.showDivByID('seniorityDiv')
        Common.showDivByID('seniorityError')
        myRIF.push('seniority')
    })

    // when confirmation is not required is clicked
    document.querySelector('#confirmation2').addEventListener('click', (e) => {
        Common.hideDivByID('seniorityDiv')
        Common.hideDivByID('seniorityError')
        Common.insertInputValue('0', 'seniority')
        Helpers.removeElementFromRIF('seniority')
    })

    // when form is submitted (save button)
    document.querySelector('#createClaimDefinitionSubmit').addEventListener('click', () => {
        const error = Common.validateRequiredFields(myRIF),
              formAction = document.querySelector('#formAction').value

        // create form   
        if (error == '0' && formAction == 'create'){
            myData = Common.getForm('createClaimDefinitionForm', connectedID)
            console.log(myData);
            API.createClaimDefinition(myData).then(resp => { 
                if (! resp.error) location.reload()
            })            
        }

        // edit form
        if (error == '0' && formAction == 'edit'){
            myData = Common.getForm('createClaimDefinitionForm', connectedID)
            
            API.updateClaimDefinition(myData).then(resp => {
                if (! resp.error) location.reload()
            })            
        }
    })

    // close warning message
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
    const myDeleteAll = document.querySelector('#deleteAllClaimDefinition')

    myDeleteAll.addEventListener('click', () => { 
        const checked = Helpers.selectedClaimDefinition()
       
        if (typeof checked != 'undefined' && checked.length > 0){
            Helpers.populateConfirmDelete(checked.length)
            myConfirm.show()
        }

    })

    // when confirm delete all is clicked
    const confirmedDelete = document.querySelector('#confirmDeleteSubmit')

    confirmedDelete.addEventListener('click', () => {
        const checked = Helpers.selectedClaimDefinition()

        API.softDeleteClaimDefinition(checked, connectedEmail).then(resp => {
            if (!resp.error){
                myConfirm.hide()
                location.reload()
            }
        })

    })
    
})
