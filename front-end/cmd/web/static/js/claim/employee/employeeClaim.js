const Common    = new MainHelpers(),
      DT        = new DataTableFeatures(),
      Draggable = new DraggableModal(),
      Helpers   = new EmployeeClaimHelpers(),
      API       = new EmployeeClaimAPI()

// store all employee by id
let allEmployees = new Map()
allEmployees.set(0, 'not defined')

// store all uploaded files by leave application id
let myUploadedFiles = new Map()

// when DOM is loaded
window.addEventListener('DOMContentLoaded', () => {
    // fetch all employees  
    API.getAllEmployees().then(resp => {
        Helpers.populateEmployeeList(resp.data.Active)

        // when employee is selected from search input
        const searchInput = document.querySelector('#searchEmployee')        
        searchInput.addEventListener('input', function() {
            const selectedOption = document.querySelector('#datalistOptions option[value="' + searchInput.value + '"]')
            if (selectedOption) {
                // employee id
                const eid = selectedOption.getAttribute('id')
                document.querySelector('#employeeClaimDiv').style.display = 'block'

                // fetch & populate requested employee claims (total)
                API.getEmployeeClaimByID(eid).then(resp => {
                    Helpers.populateMyClaims(resp.data)        
                    google.charts.setOnLoadCallback(Helpers.pieChartClaims(resp.data))

                    API.findEmployeeEmail(eid).then(r => {
                        employeeEmail = r.data

                        // fetch uploaded attachment for employee email
                        API.getUploadedFiles(employeeEmail).then(resp => {
                            if (!resp.error)
                                if (Object.keys(resp.data.Files).length > 0) myUploadedFiles = Helpers.populateUploadedFilesMap(resp.data.Files, myUploadedFiles)            
                        })
                        
                        // we use connectedEmail for log-service
                        API.getEmployeeClaims(eid, connectedEmail).then(resp => {
                            allEmployees = Common.updateEmployeeList(resp.AllEmployees, allEmployees)                           
                            Helpers.insertRows(resp.MyClaims)

                            // when attachments is clicked
                            $('#claimEmployeeTable').on('click', '.myAttachments', function (e) {
                                const appID = e.currentTarget.dataset.id
                                Helpers.populateAttachments(appID, employeeEmail)
                            })
                        })
                    })

                })

            }
        })
    })

    // make modals draggable
    Draggable.draggableModal('uploadedFilesModal')
})
