const Common  = new MainHelpers(),
      Helpers = new EmployeeReadHelpers(),
      API     = new EmployeeReadAPI()
      
// set form's parameters (Required Input Fields...)
const myRIF = [ 'fullName', 'employeeCode', 
                'streetaddr1','streetaddr2', 'zip', 'city', 'state', 'country', 
                'nationality', 'residence',
                'primaryPhone', 'primaryEmail']

// set card's parameters to enable show|hide function
const showCard = true,
      hideCard = false,
      myCards  = ['identity', 'access', 'contact']
   
// page redirection when form is successfully inserted
const sPage = "http://localhost/employee/update/"

// When DOM is loaded
window.addEventListener('DOMContentLoaded', () => {
    
     // fetch all employees & update DOM & set delete by icon event
     API.getAllEmployees().then(resp => { 
        console.log(resp);
        // display by default active employees
        Helpers.insertRows(resp.data.Active, sPage)

        // when active employees is requested
        document.querySelector('#activeBtn').addEventListener('click', () => {
            Helpers.insertRows(resp.data.Active, sPage)
            document.querySelector('#employeeTitle').innerHTML = 'Active Employees'
        })

        // when inactive employees is requested
        document.querySelector('#inactiveBtn').addEventListener('click', () => {
            Helpers.insertRows(resp.data.Inactive, sPage)
            document.querySelector('#employeeTitle').innerHTML = 'Inactive Employees'
        })

        // when deleted employee is requested
        document.querySelector('#deletedBtn').addEventListener('click', () => {
            Helpers.insertRows(resp.data.Deleted, sPage)
            document.querySelector('#employeeTitle').innerHTML = 'Deleted Employees'
        })

        // When one delete icon is clicked
        const myDelete = document.querySelectorAll('.deleteEmployee')

        myDelete.forEach(element => {
            element.addEventListener('click', () => {
                const myCheckbox = element.parentNode.parentNode.firstElementChild
                myCheckbox.checked = true
                Helpers.populateConfirmDelete('confirmDeleteBody', '1')
                myConfirm.show()
            })
        })

    }) 

    // fetch all employee's config tables & update DOM
    API.getEmployeeCT().then(resp => { 
        Helpers.populateConfigTables(resp.data) 
    })  

    // when open form (add button) is clicked (clear warning message & field's values)
    const myModal = document.querySelector('#openCreateEmployee')

    myModal.addEventListener('click', () =>{ 
        Common.clearForm('createEmployeeForm', myRIF) 
    })

    // When form is submitted (save button)
    const mySubmit = document.querySelector('#employeeAddSubmit')
    mySubmit.addEventListener('click', () => {
        const error = Common.validateRequiredFields(myRIF)

        if (error == '0'){
            myData = Helpers.getForm('createEmployeeForm')
            API.createEmployee(myData).then(resp => {
                if (! resp.error) window.location.href = sPage+resp.data.id 
            })            
        }

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

    // When delete all is clicked
    const myDeleteAll = document.querySelector('#deleteAllEmployee')

    myDeleteAll.addEventListener('click', () => { 
        const checked = Helpers.selectedEmployee('deleteWarningMessageDiv')
       
        if (typeof checked != 'undefined' && checked.length > 0){
            Helpers.populateConfirmDelete('confirmDeleteBody', checked.length)
            myConfirm.show()
        }

    })

    // When confirm delete all is clicked
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

     
})
