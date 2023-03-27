const Common  = new MainHelpers(),
      Helpers = new EmployeeUpdateHelpers(),
      API     = new EmployeeUpdateAPI()
      
// set form's parameters (Required Input Fields...)
const myRIF = [ 'firstName', 'middleName', 'familyName', 'employeeCode', 
                'streetAddressLine1','streetAddressLine2', 'zip', 'city', 'state', 'country', 
                'nationality', 'residence',
                'primaryPhone', 'primaryEmail'],
      myForm         = 'updateEmployeeForm',
      myFormSubmit   = 'updateEmployeeSubmit',
      myWarningClose = 'hideWarningMessage',
      myWarning      = 'warningMessageDiv'

// set card's parameters to enable show|hide function
const showCard = true,
      hideCard = false,
      myCards  = ['identity', 'contact', 'bank', 'emergency', 'otherInformation',
                  'spouseIdentity', 'spouseWorking', 'spouseContact',
                  'payrollInfo', 'employmentInfo',
                  'epf', 'socso', 'incomeTax', 'others']
   
// page redirection when form is successfully updated
const sPage = "http://localhost/employee/update/"

// set upload modal parameters
const myModal = 'uploadedFilesModal',
      myModalTitle = 'uploadedFilesTitle',
      uploadedFilename = 'uploadedFilename',
      uploadICButton = 'uploadICButton',
      uploadPassportButton = 'uploadPassportButton'


// When DOM is loaded
window.addEventListener('DOMContentLoaded', () => {
    const path =  window.location.pathname.split( '/' ),
          eid = path[3]

    // fetch all employee list for supervisor, superior & update DOM
    API.getActiveEmployees().then(resp => {
        console.log('employee list:',resp);
        if(!resp.error){
            Helpers.insertEmployee('superior', resp.data)
            Helpers.insertEmployee('supervisor', resp.data)
        } 
    })
    // fetch all employee's config tables & update DOM
    API.getEmployeeCT().then(resp => { 
        Helpers.populateConfigTables(resp.data)
        // fetch all employee's information & update DOM        
        API.getEmployeeInfo(eid).then(resp => { 
            if(!resp.error) {
                Helpers.populateFormData(resp.data, eid)
                const employeeEmail = resp.data.Employee.primaryEmail
                // fetch all employee's uploaded files & update DOM
                API.getUploadedFiles(employeeEmail).then(resp => {
                    console.log(resp);
                    if(!resp.error) Helpers.populateUploadedFiles(resp.data, employeeEmail)
                })
            }
        })
        
    }) 

    // When form update employee is submitted (save button)
    const mySubmit = document.querySelector('#'+myFormSubmit)
    mySubmit.addEventListener('click', () => {
        const error = Common.validateRequiredFields(myRIF)
        if (error == '0'){
            myData = Helpers.getForm(myForm, eid, connectedEmail, connectedID)
            console.log(myData);
            
            API.updateEmployee(myData).then(resp => {
                console.log(resp);
                if (! resp.error) window.location.href = sPage + resp.data.ID
            })   

        }
    })
    
    // enable card's show|hide function
    myCards.forEach(card => {
        document.querySelector(`#${card}Hide`).addEventListener('click', () => { Common.displayCardByID(card, hideCard) })
        document.querySelector(`#${card}Show`).addEventListener('click', () => { Common.displayCardByID(card, showCard) })
    })

    // close warning message
    const myWarningMessage = document.querySelector('#'+myWarningClose)
    myWarningMessage.addEventListener('click', () => {
        Common.hideDivByID(myWarning)
    })

    // initiate upload modal
    const uploadModal = new bootstrap.Modal(document.getElementById(myModal), {
        keyboard: false
    })

    // when upload ic button is clicked
    const myICBtn = document.querySelector('#'+uploadICButton)
    myICBtn.addEventListener('click', () => {
        const wanted = 'ic'
        Helpers.populateUploadFiles(wanted)
        uploadModal.toggle()
    })

    // when upload ic button is clicked
    const myPassportBtn = document.querySelector('#'+uploadPassportButton)
    myPassportBtn.addEventListener('click', () => {
        const wanted = 'passport'
        Helpers.populateUploadFiles(wanted)
        uploadModal.toggle()
    })

})
