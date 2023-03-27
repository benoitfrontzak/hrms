const Common  = new MainHelpers(),
      Helpers = new EmployeeReadHelpers(),
      API     = new EmployeeReadAPI()
      
// set form's parameters (Required Input Fields...)
const myRIF = [ 'firstName', 'middleName', 'familyName', 'employeeCode', 
                'streetaddr1','streetaddr2', 'zip', 'city', 'state', 'country', 
                'nationality', 'residence',
                'primaryPhone', 'primaryEmail'],
      myForm         = 'createEmployeeForm',
      myFormSubmit   = 'employeeAddSubmit',
      openFormModal  = 'openCreateEmployee',
      myWarningClose = 'hideWarningMessage',
      myWarning      = 'warningMessageDiv'

// set card's parameters to enable show|hide function
const showCard = true,
      hideCard = false,
      myCards  = ['identity', 'access', 'contact']
   
// page redirection when form is successfully inserted
const sPage = "http://localhost/employee/update/"

// set data table parameters
const dt = 'employeeSummary',
      dtBody = 'employeeSummaryBody'

// set delete all employee parameter
const dAll = 'deleteAllEmployee',
      myDeleteWarningClose = 'hidedeleteWarningMessage',
      myDeleteWarning = 'deleteWarningMessageDiv',
      confirmDelete = 'confirmDelete',
      confirmDeleteBody = 'confirmDeleteBody',
      confirmDeleteSubmit = 'confirmDeleteSubmit',
      deleteCheckboxes = 'deleteCheckboxes',
      iconDeleteClass = 'deleteEmployee'

// set data view
const active = 'activeBtn',
      inactive = 'inactiveBtn',
      deleted = 'deletedBtn'

// When DOM is loaded
window.addEventListener('DOMContentLoaded', () => {
     // fetch all employees & update DOM & set delete by icon event
     API.getAllEmployees().then(resp => { 
        // display by default active employees
        Helpers.insertRows(dtBody, resp.data.Active, sPage)
        Helpers.triggerDT(dt, dtBody, sPage)

        // when active employees is requested
        document.querySelector('#'+active).addEventListener('click', () => {
            Helpers.insertRows(dtBody, resp.data.Active, sPage)
            Helpers.triggerDT(dt, dtBody, sPage)
        })

        // when inactive employees is requested
        document.querySelector('#'+inactive).addEventListener('click', () => {
            Helpers.insertRows(dtBody, resp.data.Inactive, sPage)
            Helpers.triggerDT(dt, dtBody, sPage)
        })

        // when deleted employee is requested
        document.querySelector('#'+deleted).addEventListener('click', () => {
            Helpers.insertRows(dtBody, resp.data.Deleted, sPage)
            Helpers.triggerDT(dt, dtBody, sPage)
        })

        // When one delete icon is clicked
        const myDelete = document.querySelectorAll('.'+iconDeleteClass)

        myDelete.forEach(element => {
            element.addEventListener('click', () => {
                const myCheckbox = element.parentNode.parentNode.firstElementChild
                myCheckbox.checked = true
                Helpers.populateConfirmDelete(confirmDeleteBody, '1')
                myConfirm.show()
            })
        })

    }) 

    // fetch all employee's config tables & update DOM
    API.getEmployeeCT().then(resp => { 
        Helpers.populateConfigTables(resp.data) 
    })  

    // when open form (add button) is clicked (clear warning message & field's values)
    const myModal = document.querySelector('#'+openFormModal)

    myModal.addEventListener('click', () =>{ 
        Common.clearForm(myForm, myRIF) 
    })

    // When form is submitted (save button)
    const mySubmit = document.querySelector('#'+myFormSubmit)
    mySubmit.addEventListener('click', () => {
        const error = Common.validateRequiredFields(myRIF)

        if (error == '0'){
            myData = Helpers.getForm(myForm)
            
            API.createEmployee(myData).then(resp => { 
                console.log(resp);
                if (! resp.error) window.location.href = sPage+resp.data.id 
            })            
        }

    })
    
    // initiate delete confirm modal
    const myConfirm = new bootstrap.Modal(document.getElementById(confirmDelete), { 
        keyboard: false 
    })

    // cleaned checked checkboxes when modal is close
    document.getElementById(confirmDelete).addEventListener('hidden.bs.modal', function (event) {
        const myDeleteCheckboxes = document.querySelectorAll('.'+deleteCheckboxes)

        myDeleteCheckboxes.forEach(element => {
            element.checked = false
        })

    })

    // When delete all is clicked
    const myDeleteAll = document.querySelector('#'+dAll)

    myDeleteAll.addEventListener('click', () => { 
        const checked = Helpers.selectedEmployee(myDeleteWarning)
       
        if (typeof checked != 'undefined' && checked.length > 0){
            Helpers.populateConfirmDelete(confirmDeleteBody, checked.length)
            myConfirm.show()
        }

    })

    // When confirm delete all is clicked
    const confirmedDelete = document.querySelector('#'+confirmDeleteSubmit)

    confirmedDelete.addEventListener('click', () => {
        const checked = Helpers.selectedEmployee(myDeleteWarning)

        API.softDeleteEmployeeByID(checked, connectedEmail).then(resp => {
            if (!resp.error){
                myConfirm.hide()
                location.reload()
            }
        })

    })

  
    // enable card's show|hide function
    myCards.forEach(card => {
        // hide card
        document.querySelector(`#${card}Hide`).addEventListener('click', () => { 
            Common.displayCardByID(card, hideCard) 
        })
        // show card
        document.querySelector(`#${card}Show`).addEventListener('click', () => { 
            Common.displayCardByID(card, showCard) 
        })
    })

    // close warning message
    const myWarningMessage = document.querySelector('#'+myWarningClose)
    myWarningMessage.addEventListener('click', () => {
        Common.hideDivByID(myWarning)
    })

     // close delete warning message
     const myDeleteWarningMessage = document.querySelector('#'+myDeleteWarningClose)
     myDeleteWarningMessage.addEventListener('click', () => {
         Common.hideDivByID(myDeleteWarning)
     })
})
