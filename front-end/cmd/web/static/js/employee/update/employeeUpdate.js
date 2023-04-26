const Common  = new MainHelpers(),
      Helpers = new EmployeeUpdateHelpers(),
      API     = new EmployeeUpdateAPI()
      
// set form's parameters (Required Input Fields...)
const myRIF = [ 'fullName', 'employeeCode', 
                'streetAddressLine1','streetAddressLine2', 'zip', 'city', 'state', 'country', 
                'nationality', 'residence',
                'primaryPhone', 'primaryEmail']

// set card's parameters to enable show|hide function
const myCards  = ['identity', 'contact', 'bank', 'emergency', 'otherInformation',
                  'spouseIdentity', 'spouseWorking', 'spouseContact',
                  'payrollInfo', 'employmentInfo',
                  'epf', 'socso', 'incomeTax', 'others']

// set upload files button
const myUploadButtons = ['Profile', 'IC', 'Passport', 'OtherInformation']

// page redirection when form is successfully updated
const sPage = "http://localhost/employee/update/"

// store all employees (active, inactive & deleted) 
let allEmployees = new Map()
allEmployees.set(0, 'not defined')

// store current selected employee Employment & Statutory information
let cEmployment, cStatutory

// When DOM is loaded
window.addEventListener('DOMContentLoaded', () => {
    const path =  window.location.pathname.split( '/' ),
          eid = path[3]

    // fetch all requires information
    API.getAllRequiresInfo(eid, connectedID, connectedEmail).then(resp => {
        console.log('employee info');
        console.log(resp);
        if(!resp.error){
            const employeeEmail = resp.EmployeeInfo.Employee.primaryEmail

            cEmployment = resp.EmployeeInfo.Employment
            cStatutory = resp.EmployeeInfo.Statutory
            allEmployees = Common.updateEmployeeList(resp.AllEmployees, allEmployees)
            
            Helpers.populateConfigTables(resp.EmployeeCT)
            Helpers.insertEmployee('superior', resp.AllEmployees.Active)
            Helpers.insertEmployee('supervisor', resp.AllEmployees.Active)
            Helpers.populateFormData(resp.EmployeeInfo, eid)
            Helpers.populateArchivesDT(resp.EmployeeInfo.EmploymentArchive, resp.EmployeeInfo.StatutoryArchive)
            
            // fetch all employee's uploaded files & update DOM
            API.getUploadedFiles(employeeEmail).then(resp => {
                if(!resp.error) Helpers.populateUploadedFiles(resp.data, employeeEmail)
            })

        }
    }) 

    // When form update employee is submitted (save button)
    const mySubmit = document.querySelector('#updateEmployeeSubmit')
    mySubmit.addEventListener('click', () => {
        const error = Common.validateRequiredFields(myRIF)
        if (error == '0'){
            myData = Helpers.getForm('updateEmployeeForm', eid, connectedEmail, connectedID)
           
            
            console.log(myData);
            
            API.updateEmployee(myData).then(resp => {
                console.log(resp);
                if (! resp.error) window.location.href = sPage + resp.data.id
            })   

        }
    })
    
    // make menu clickable (personal | spouse | employment | statutory)
    const menuItems = ['personal', 'spouse', 'employment', 'statutory']
    menuItems.forEach(element => {
        const myMenu            = document.querySelector('#'+element + 'Menu'),
              myContent         = document.querySelector('#'+element + 'Content'),
              personalContent   = document.querySelector('#personalContent'),
              spouseContent     = document.querySelector('#spouseContent'),
              employmentContent = document.querySelector('#employmentContent'),
              statutoryContent  = document.querySelector('#statutoryContent'),
              personalMenu   = document.querySelector('#personalMenu'),
              spouseMenu     = document.querySelector('#spouseMenu'),
              employmentMenu = document.querySelector('#employmentMenu'),
              statutoryMenu  = document.querySelector('#statutoryMenu')

        myMenu.addEventListener('click', () => {
            // show/hide content
            personalContent.className   = 'hide'
            spouseContent.className     = 'hide'
            employmentContent.className = 'hide'
            statutoryContent.className  = 'hide'
            myContent.className         = 'show'
            // make menu link active
            personalMenu.className      = 'side-nav-link border-end-0'
            spouseMenu.className        = 'side-nav-link border-end-0'
            employmentMenu.className    = 'side-nav-link border-end-0'
            statutoryMenu.className     = 'side-nav-link border-end-0'
            myMenu.className            = 'side-nav-link border-end border-danger myTint6BG'
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
    });
    
    // // when upload ic button is clicked
    // const myICBtn = document.querySelector('#uploadICButton')
    // myICBtn.addEventListener('click', () => {
    //     const wanted = 'ic'
    //     Helpers.populateUploadFiles(wanted)
    //     uploadModal.toggle()
    // })

    // // when upload ic button passport clicked
    // const myPassportBtn = document.querySelector('#uploadPassportButton')
    // myPassportBtn.addEventListener('click', () => {
    //     const wanted = 'passport'
    //     Helpers.populateUploadFiles(wanted)
    //     uploadModal.toggle()
    // })

})
