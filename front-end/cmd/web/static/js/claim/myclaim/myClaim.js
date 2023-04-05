const Common  = new MainHelpers(),
      Helpers = new MyClaimHelpers(),
      API     = new MyClaimAPI()

// set form's parameters (Required Input Fields...)
const myRIF = ['claimDefinition', 'amount', 'description']

// store connected user information
// we update leave relative information with api call
const connectedUser = {
    id            : connectedID,
    email         : connectedEmail,
    code          : '', 
    fullname     : '',  
    nickname      : '', 
    joinDate      : '', 
    confirmDate   : '' 
}
// store all employee by id
let allEmployees = new Map()
allEmployees.set(0, 'not defined')

// when DOM is loaded
window.addEventListener('DOMContentLoaded', () => {
    // fetch all employee information
    API.getAllEmployees().then(resp => {
        allEmployees = Common.updateEmployeeList(resp.data, allEmployees)
        
        // fetch all 'my claims' & update DOM (data table)
        API.getAllMyClaim(connectedID, connectedEmail).then(resp => {
            Helpers.checkData(resp.data)
            // Helpers.makeEditable()
        })
    })    

    // fetch all claim's config table & update DOM (form)
    API.getClaimCT().then(resp => {
        Helpers.insertOptions('category',resp.data.Category)
    })

    // fetch all claim's definitions & update DOM (form)
    API.getAllClaimDefinition().then(resp => {
        Helpers.insertOptionsCD('claimDefinition', resp.data.Active)
        // when claim definition change (form)
        document.querySelector('#claimDefinition').addEventListener('change', (e) => {
            const claimDefinitionID = e.target.value
            console.log('selected CD: ' + claimDefinitionID);
            // if 'not defined' is selected 
                // display category & name (or hide)
                // make category & name required (or not)
            Helpers.switchFormFieldsView(claimDefinitionID)
            // populate hidden fields (for form validation: confirmation | seniority | docRequired | limitation)
            const myCDValidation = document.querySelector('#myclaimDefinition'+claimDefinitionID).dataset
            Helpers.populateHiddenFields(myCDValidation)
        })
    })

    // fetch connected user information and update connectedUser
    API.getUserInformation(connectedID, connectedEmail).then(resp => {
        Helpers.updateUserInformation(resp.data)
    })

       

    
    // when form is submitted (save button)
    document.querySelector('#claimFormSubmit').addEventListener('click', () => {
        const error = Common.validateRequiredFields(myRIF)
        myData = Common.getForm('claimForm', connectedID)
        console.log(myData);
        if (error == '0'){
            API.createClaim(myData).then(resp => {
                console.log(resp);
                if (! resp.error) location.reload()
            })
        }
    })
    // close warning message
    const myWarningMessage = document.querySelector('#hideWarningMessage')
    myWarningMessage.addEventListener('click', () => {
        Common.hideDivByID('warningMessageDiv')
    })
    // clear form (when open form is clicked)
    document.querySelector('#openCreateMyClaim').addEventListener('click', () =>{ 
        Common.clearForm('claimForm', myRIF) 
    })

    // initiate delete confirm modal
    const myConfirm = new bootstrap.Modal(document.getElementById('confirmDelete'), { 
        keyboard: false 
    })

    // cleaned checked checkboxes when modal is close
    document.getElementById('confirmDelete').addEventListener('hidden.bs.modal', function () {
       document.querySelectorAll('.deleteCheckboxes').forEach(element => {
            element.checked = false
        })
    })

    // When delete all is clicked
    document.querySelector('#deleteAllMyClaim').addEventListener('click', () => { 
        const checked = Helpers.selectedClaim()

        if (typeof checked != 'undefined' && checked.length > 0){
            Helpers.populateConfirmDelete(checked.length)
            myConfirm.show()
        }

    })

    // When confirm delete all is clicked
    const confirmedDelete = document.querySelector('#confirmDeleteSubmit')

    confirmedDelete.addEventListener('click', () => {
        const checked = Helpers.selectedClaim()

        API.softDeleteClaim(checked, connectedEmail).then(resp => {
            if (!resp.error){
                myConfirm.hide()
                location.reload()
            }
        })

    })

})
