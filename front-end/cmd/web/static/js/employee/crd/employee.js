const Common            = new MainHelpers(),
      DT                = new DataTableFeatures(),
      Draggable         = new DraggableModal(),
      Employee          = new EmployeeRead(),
      CT                = new EmployeeCT(),
      Helpers           = new EmployeeHelpers(),
      API               = new EmployeeReadAPI()
      
// set form's parameters (Required Input Fields...)
const myRIF = [ 'fullName', 'nickName', 'employeeCode', 
                'streetaddr1','streetaddr2', 'zip', 'city', 'state', 'country', 
                'nationality', 'residence',
                'primaryPhone', 'primaryEmail']

// set card's parameters to enable show|hide function
const showCard = true,
      hideCard = false,
      myCards  = ['identity', 'access', 'contact']

// store all employee by id
let allEmployees = new Map()
allEmployees.set(0, 'not defined')

// When DOM is loaded
window.addEventListener('DOMContentLoaded', () => {
    // fetch all employee's config tables & update create new employee form
    API.getEmployeeCT().then(resp => {
        CT.populateConfigTables(resp.data) 
    })

    // fetch all employees & insert rows to datatable
    API.getAllEmployees().then(resp => {

        allEmployees = Common.updateEmployeeList(resp.data, allEmployees) 
        const active = resp.data.Active,
              inactive = resp.data.Inactive,
              deleted = resp.data.Deleted

        // display by default active employees
        Employee.insertRows(active)
        
        // create listener when data source changed (active | inactive | deleted)
        Helpers.dataSourceListener(active, inactive, deleted)

    })   

    
    // when form (add button) is clicked (clear warning message & field's values)
    document.querySelector('#openCreateEmployee').addEventListener('click', () =>{ 
        Common.clearForm('createEmployeeForm', myRIF) 
    })

    // When form is submitted (save button)
    document.querySelector('#employeeAddSubmit').addEventListener('click', () => {
        // Check for errors and display it
        const error = Common.validateRequiredFields(myRIF)

        if (error == '0'){
            myData = Helpers.getForm('createEmployeeForm')
            API.createEmployee(myData).then(resp => {
                if (! resp.error) window.open('/employee/update/' + resp.data.id, "_self") 
            })            
        }

    })
    
    // initiate confirm delete modal
    const myConfirm = new bootstrap.Modal(document.getElementById('confirmDelete'), { 
        backdrop: 'static',
        keyboard: false 
    })

    // Clear selected employee when confirm delete modal is closed
    Helpers.clearSelectedEmployee()
    

    // When delete employees is clicked (delete button)
   document.querySelector('#deleteAllEmployee').addEventListener('click', () => { 
        const checked = Helpers.selectedEmployee('deleteWarningMessageDiv')
       
        if (typeof checked != 'undefined' && checked.length > 0){
            Helpers.populateConfirmDelete('confirmDeleteBody', checked.length)
            myConfirm.show()
        }

    })

    // When confirm delete employees is clicked (confirm button from modal)
    const confirmedDelete = document.querySelector('#confirmDeleteSubmit')

    confirmedDelete.addEventListener('click', () => {
        const checked = Helpers.selectedEmployee('deleteWarningMessageDiv')

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
    const myWarningMessage = document.querySelector('#hideWarningMessage')
    myWarningMessage.addEventListener('click', () => {
        Common.hideDivByID('warningMessageDiv')
    })

     // close delete warning message
     const myDeleteWarningMessage = document.querySelector('#hidedeleteWarningMessage')
     myDeleteWarningMessage.addEventListener('click', () => {
         Common.hideDivByID('deleteWarningMessageDiv')
     })

    // make modals draggable
    Draggable.draggableModal('createEmployee')
    Draggable.draggableModal('confirmDelete')
})
