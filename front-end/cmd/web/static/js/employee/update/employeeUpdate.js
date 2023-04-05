const Common  = new MainHelpers(),
      Helpers = new EmployeeUpdateHelpers(),
      API     = new EmployeeUpdateAPI()
      
// set form's parameters (Required Input Fields...)
const myRIF = [ 'fullName', 'employeeCode', 
                'streetAddressLine1','streetAddressLine2', 'zip', 'city', 'state', 'country', 
                'nationality', 'residence',
                'primaryPhone', 'primaryEmail']

// set card's parameters to enable show|hide function
const showCard = true,
      hideCard = false,
      myCards  = ['identity', 'contact', 'bank', 'emergency', 'otherInformation',
                  'spouseIdentity', 'spouseWorking', 'spouseContact',
                  'payrollInfo', 'employmentInfo',
                  'epf', 'socso', 'incomeTax', 'others']
   
// page redirection when form is successfully updated
const sPage = "http://localhost/employee/update/"

// store all employees (active, inactive & deleted) 
let allEmployees = new Map()
allEmployees.set(0, 'not defined')

// When DOM is loaded
window.addEventListener('DOMContentLoaded', () => {
    const path =  window.location.pathname.split( '/' ),
          eid = path[3]

    // fetch all employee list for supervisor, superior & update DOM
    API.getActiveEmployees().then(resp => {
        if(!resp.error){
            Helpers.insertEmployee('superior', resp.data)
            Helpers.insertEmployee('supervisor', resp.data)
        } 
    })
    // fetch all employee's config tables & update DOM
    API.getEmployeeCT().then(resp => {
        const CT = resp.data
        Helpers.populateConfigTables(CT)
        // fetch all employee's information & update DOM        
        API.getEmployeeInfo(eid).then(resp => {
            console.log('resp employee info');
            console.log(resp);
            const employee = resp.data,
                  employmentArchive = resp.data.EmploymentArchive,
                  statutoryArchive = resp.data.StatutoryArchive

            if(!resp.error) {
                // populate form
                Helpers.populateFormData(employee, eid)

                // fetch all employee information
                API.getAllEmployees().then(resp => {
                    allEmployees = Common.updateEmployeeList(resp.data, allEmployees)
                    // populate archives DT
                    Helpers.populateArchivesDT(employmentArchive, statutoryArchive)
                    // todo replace ids by values
                })
                

                const employeeEmail = resp.data.Employee.primaryEmail
                // fetch all employee's uploaded files & update DOM
                API.getUploadedFiles(employeeEmail).then(resp => {
                    if(!resp.error) Helpers.populateUploadedFiles(resp.data, employeeEmail)
                })
            }
        })
        
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
    
    // enable card's show|hide function
    myCards.forEach(card => {
        document.querySelector(`#${card}Hide`).addEventListener('click', () => { Common.displayCardByID(card, hideCard) })
        document.querySelector(`#${card}Show`).addEventListener('click', () => { Common.displayCardByID(card, showCard) })
    })

    // close warning message
    const myWarningMessage = document.querySelector('#hideWarningMessage')
    myWarningMessage.addEventListener('click', () => {
        Common.hideDivByID('warningMessageDiv')
    })

    // initiate upload modal
    const uploadModal = new bootstrap.Modal(document.getElementById('uploadedFilesModal'), {
        keyboard: false
    })

    // when upload ic button is clicked
    const myICBtn = document.querySelector('#uploadICButton')
    myICBtn.addEventListener('click', () => {
        const wanted = 'ic'
        Helpers.populateUploadFiles(wanted)
        uploadModal.toggle()
    })

    // when upload ic button passport clicked
    const myPassportBtn = document.querySelector('#uploadPassportButton')
    myPassportBtn.addEventListener('click', () => {
        const wanted = 'passport'
        Helpers.populateUploadFiles(wanted)
        uploadModal.toggle()
    })

})
