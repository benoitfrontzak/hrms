const Common    = new MainHelpers(),
      Draggable = new DraggableModal(),
      Helpers   = new EmployeeUpdateHelpers(),
      API       = new EmployeeUpdateAPI()
      
// set form's parameters (Required Input Fields...)
const myRIF = [ 'fullName', 'employeeCode', 
                'streetAddressLine1','streetAddressLine2', 'zip', 'city', 'state', 'country', 
                'nationality', 'residence',
                'primaryPhone', 'primaryEmail']

// set sidenav items
const menuItems = ['personal', 'spouse', 'employment', 'statutory', 'payrollItem']

// set card's parameters to enable show|hide function
const myCards  = ['identity', 'contact', 'bank', 'emergency', 'otherInformation',
                  'spouseIdentity', 'spouseWorking', 'spouseContact',
                  'payrollInfo', 'employmentInfo',
                  'epf', 'socso', 'incomeTax', 'others',
                  'payrollItem']

// set upload files button
const myUploadButtons = ['Profile', 'IC', 'Passport', 'OtherInformation']

// page redirection when form is successfully updated
const sPage = "http://localhost/employee/update/"

// store all employees (active, inactive & deleted) 
let allEmployees = new Map()
allEmployees.set(0, 'not defined')

// store all payroll items by id
let allPI = new Map()

// store current selected employee Employment & Statutory information (to save only when data change)
let cEmployment, cStatutory

// When DOM is loaded
window.addEventListener('DOMContentLoaded', () => {
    const path =  window.location.pathname.split( '/' ),
          eid = path[3]

    // fetch all requires information
    API.getAllRequiresInfo(eid, connectedID, connectedEmail).then(resp => {
        console.log('resp');
        console.log(resp);
        if(!resp.error){
            const employeeEmail = resp.EmployeeInfo.Employee.primaryEmail

            cEmployment = resp.EmployeeInfo.Employment
            cStatutory = resp.EmployeeInfo.Statutory
            allEmployees = Common.updateEmployeeList(resp.AllEmployees, allEmployees)
            allPI = Helpers.updatePayrollItemsMap(resp.EmployeeCT.PayrollItem, allPI)
    
            Helpers.populateConfigTables(resp.EmployeeCT)
            Helpers.insertEmployee('superior', resp.AllEmployees.Active)
            Helpers.insertEmployee('supervisor', resp.AllEmployees.Active)
            Helpers.populateFormData(resp.EmployeeInfo, eid)
            Helpers.populateArchivesDT(resp.EmployeeInfo.EmploymentArchive, resp.EmployeeInfo.StatutoryArchive)
     
            // fetch all employee's uploaded files & update DOM
            API.getUploadedFiles(employeeEmail).then(resp => {
                if(!resp.error) Helpers.populateUploadedFiles(resp.data, employeeEmail)
            })

            // initiate tooltips
            const tooltipTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="tooltip"]'));
            tooltipTriggerList.map(function (tooltipTriggerEl) {
                return new bootstrap.Tooltip(tooltipTriggerEl);
            })

        }
    }) 

    // clear payroll item form (when open form is clicked)
    document.querySelector('#openPayrollItem').addEventListener('click', () =>{
        Helpers.clearPI()
    })

    // when payroll item is selected
    document.querySelector('#payrollItem').addEventListener('change', (e) =>{
        const selectedID = e.target.value,
              piSelected = allPI.get(selectedID)
        Helpers.populateFormPI(piSelected);
    })

    // when form add payroll item is submitted (save button)
    document.querySelector('#savePI').addEventListener('click', () => {
        const errors = Helpers.checkFormPI()
        if (errors == 0){
            const data = Helpers.getFormPI(eid)
            API.addPayrollItem(data).then(resp => {
                if (!resp.error) location.reload()
            })            
        }
    })

    // When form update employee is submitted (save button)
    document.querySelector('#updateEmployeeSubmit').addEventListener('click', () => {
        const error = Common.validateRequiredFields(myRIF)
        if (error == '0'){
            myData = Helpers.getForm('updateEmployeeForm', eid, connectedEmail, connectedID)
            
            API.updateEmployee(myData).then(resp => {
                if (! resp.error) window.location.href = sPage + resp.data.id
            })   

        }
    })
    
    // make menu clickable (personal | spouse | employment | statutory | payrollItem)    
    menuItems.forEach(element => {
        const myMenu              = document.querySelector('#'+element + 'Menu'),
              myContent           = document.querySelector('#'+element + 'Content'),
              personalContent     = document.querySelector('#personalContent'),
              spouseContent       = document.querySelector('#spouseContent'),
              employmentContent   = document.querySelector('#employmentContent'),
              statutoryContent    = document.querySelector('#statutoryContent'),
              payrollItemContent  = document.querySelector('#payrollItemContent'),
              personalMenu        = document.querySelector('#personalMenu'),
              spouseMenu          = document.querySelector('#spouseMenu'),
              employmentMenu      = document.querySelector('#employmentMenu'),
              statutoryMenu       = document.querySelector('#statutoryMenu'),
              payrollItemMenu     = document.querySelector('#payrollItemMenu')

        myMenu.addEventListener('click', () => {
            // show/hide content
            personalContent.className    = 'hide'
            spouseContent.className      = 'hide'
            employmentContent.className  = 'hide'
            statutoryContent.className   = 'hide'
            payrollItemContent.className = 'hide'
            myContent.className          = 'show'
            // make menu link active
            personalMenu.className       = 'side-nav-link border-end-0'
            spouseMenu.className         = 'side-nav-link border-end-0'
            employmentMenu.className     = 'side-nav-link border-end-0'
            statutoryMenu.className      = 'side-nav-link border-end-0'
            payrollItemMenu.className    = 'side-nav-link border-end-0'
            myMenu.className             = 'side-nav-link border-end border-danger myTint6BG'
        })

    })
    
    // enable card's show|hide function
    myCards.forEach(card => {
        document.querySelector(`#${card}Accordion`).addEventListener('click', () => {
            const isDisplayed = document.querySelector(`#${card}BodyCard`).style.display
            let display
            (isDisplayed == '' || isDisplayed == 'block') ?  display = false : display = true;
            Common.displayCardByID(card, display)
        })
    })

    // close warning message
    const myWarningMessage = document.querySelector('#hideWarningMessage')
    myWarningMessage.addEventListener('click', () => {
        Common.hideDivByID('warningMessageDiv')
    })

    // initiate upload modal
    const uploadModal = new bootstrap.Modal(document.getElementById('uploadedFilesModal'), {
        backdrop: 'static',
        keyboard: false
    })

    // event listener for each upload buttons
    myUploadButtons.forEach(element => {
        const wanted = element.toLowerCase(),
              myButton = document.querySelector(`#upload${element}Button`)
        myButton.addEventListener('click', () => {
            Helpers.populateUploadFiles(wanted)
            uploadModal.toggle()
        })
    })

    // initiate confirm delete modal
    const myConfirm = new bootstrap.Modal(document.getElementById('confirmDelete'), { 
        backdrop: 'static',
        keyboard: false 
    })

    // Clear selected employee when confirm delete modal is closed
    Helpers.clearSelectedEmployee()
    

    // When delete employees is clicked (delete button)
   document.querySelector('#deleteAllPayrollItem').addEventListener('click', () => { 
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

        API.softDeletePayrollItem(checked, connectedEmail).then(resp => {
            if (!resp.error){
                myConfirm.hide()
                location.reload()
            }
        })

    })

    // close delete warning message
    const myDeleteWarningMessage = document.querySelector('#hidedeleteWarningMessage')
    myDeleteWarningMessage.addEventListener('click', () => {
        Common.hideDivByID('deleteWarningMessageDiv')
    })

    // make modals draggable 
    Draggable.draggableModal('payrollItems')
    Draggable.draggableModal('archivedProfile')
    Draggable.draggableModal('uploadedIC')
    Draggable.draggableModal('archivedIC')
    Draggable.draggableModal('uploadedPassport')
    Draggable.draggableModal('archivedPassport')
    Draggable.draggableModal('uploadedOtherInformation')
    Draggable.draggableModal('archivedOtherInformation')
    Draggable.draggableModal('payrollArchived')
    Draggable.draggableModal('employmentArchived')
    Draggable.draggableModal('epfArchived')
    Draggable.draggableModal('socsoArchived')
    Draggable.draggableModal('taxArchived')
    Draggable.draggableModal('othersArchived')
    Draggable.draggableModal('confirmDelete')

})
