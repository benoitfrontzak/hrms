const Common    = new MainHelpers(),
      DT        = new DataTableFeatures(),
      Draggable = new DraggableModal(),
      Helpers   = new EmployeeLeaveHelpers(),
      API       = new EmployeeLeaveAPI()

// store all employee by id
let allEmployees = new Map()
allEmployees.set(0, 'not defined')

// store all uploaded files by leave application id
let myUploadedFiles = new Map()
 
// when DOM is loaded
window.addEventListener('DOMContentLoaded', () => {
    // fetch all employees  
    API.getAllEmployees().then(resp => {
        const active = resp.data.Active
        Helpers.populateEmployeeList(active)
        allEmployees = Common.updateEmployeeList(resp.data, allEmployees)

        // when employee is selected from search input
        const searchInput = document.querySelector('#searchEmployee')

        searchInput.addEventListener('input', function() {
            const selectedOption = document.querySelector('#datalistOptions option[value="' + searchInput.value + '"]')
            if (selectedOption) {
                const eid = selectedOption.getAttribute('id')

                // fetch & populate requested employee leaves
                API.getEmployeeLeaveDetailsByID(eid).then(resp => {                    
                    Common.showDivByID('leaveEmployeeForm')
                    Common.showDivByID('employeeLeaveDiv')
                    Helpers.populateMyLeaveDetails(resp.data)
                    Helpers.populateMyLeaveDetailsProgress(resp.data)
                    
                    API.findEmployeeEmail(eid).then(r => {
                        employeeEmail = r.data

                        // fetch uploaded attachment for employee email
                        API.getUploadedFiles(employeeEmail).then(resp => {
                            if (!resp.error)
                                if (Object.keys(resp.data.Files).length > 0) myUploadedFiles = Helpers.populateUploadedFilesMap(resp.data.Files, myUploadedFiles)            
                        })

                        // fetch all employee's leaves (we use connected email for log-service)
                        API.getEmployeeLeaves(eid, connectedEmail).then(resp => {
                            // insert rows master list table: myLeaves
                            Helpers.insertRows(resp.MyLeaves)

                            // when attachments is clicked
                            $('#employeeLeaveTable').on('click', '.myAttachments', function (e) {
                                const appID = e.currentTarget.dataset.id
                                Helpers.populateAttachments(appID, employeeEmail)
                            })
                        })
                    })
                    
                })

            }
        })
    })

    // when form is submitted
    document.querySelector('#updateEmployeeLeave').addEventListener('click', () => {
        const data = Helpers.getForm(connectedID, connectedEmail)
        API.updateEmployeeLeaves(data).then(resp => {
            if (!resp.error){
                Common.showDivByID('successMessageDiv')
                setTimeout(function() { Common.hideDivByID('successMessageDiv'); }, 5000);
            }
        })
    })

    // make modals draggable
    Draggable.draggableModal('uploadedFilesModal')
})
