const Common    = new MainHelpers(),
      Helpers   = new EmployeeLeaveHelpers(),
      API       = new EmployeeLeaveAPI()

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
                const id = selectedOption.getAttribute('id')
                // fetch & populate requested employee leaves
                API.getEmployeeLeaveDetailsByID(id).then(resp => {                    
                    Common.showDivByID('leaveEmployeeForm')
                    Helpers.populateMyLeaveDetails(resp.data)
                    Helpers.populateMyLeaveDetailsProgress(resp.data)
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
})
