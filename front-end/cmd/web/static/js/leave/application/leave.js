const Common  = new MainHelpers(),
      Helpers = new LeaveHelpers(),
      API     = new LeaveAPI()

// set form's parameters (Required Input Fields...)
const myRIF = [ 'name', 'description', 'category']

// store all employees (active, inactive & deleted) 
let allEmployees = new Map()
allEmployees.set(0, 'not defined')

// store empty row for leaves with no data (approved, rejected or pending)
const noData = {
    employeeid: 0,
    leaveDefinition: 0,
    leaveDefinitionCode: 'No data',
    leaveDefinitionName:'No data',
    description: 'No data',
    statusid: 0,
    status: 'No data',
    approvedAt: '0001-01-01',
    approvedBy: 0,
    rejectedReason: ''
}
const noDataAr = [noData]

// when DOM is loaded
window.addEventListener('DOMContentLoaded', () => {
    const noAction = false
     // fetch all employee information
     API.getAllEmployees().then(resp => {
        allEmployees = Common.updateEmployeeList(resp.data, allEmployees)

        // fetch all leaves & update DOM (data table)
        API.getAllLeaves().then(resp => {
            console.log(resp);
            // display by default pending request
            Helpers.insertRows(resp.data.Pending)
            // when approved is selected
            document.querySelector('#approvedBtn').addEventListener('click', () => {
                Helpers.insertRows(resp.data.Approved, noAction)
                document.querySelector('#leaveTitle').innerHTML = 'Approved Leaves'
            })

             // when rejected is selected
             document.querySelector('#rejectedBtn').addEventListener('click', () => {
                Helpers.insertRows(resp.data.Rejected, noAction)
                document.querySelector('#leaveTitle').innerHTML = 'Rejected Leaves'
            })

             // when pending is selected
             document.querySelector('#pendingBtn').addEventListener('click', () => {
                Helpers.insertRows(resp.data.Pending)
                document.querySelector('#leaveTitle').innerHTML = 'Pending Leaves'
            })
        })
    })

    // initiate confirm approval modal
    const myApprovalConfirm = new bootstrap.Modal(document.getElementById('approveModal'), { 
        keyboard: false 
    })
     // initiate confirm rejection modal
     const myRejectionConfirm = new bootstrap.Modal(document.getElementById('rejectModal'), { 
        keyboard: false 
    })

    // when approve button is clicked (open modal)
    document.querySelector('#approveBtn').addEventListener('click', () => {
       const checked = Helpers.selectedLeave()

       if (typeof checked != 'undefined' && checked.length > 0){
            Helpers.populateSelectedLeavesNumber(checked.length, 'selectedLeavesApproval', 'approve', '')
            myApprovalConfirm.show()
        }

    })
    // when confirm approval is clicked (submit)
    document.querySelector('#approveSubmit').addEventListener('click', () => {
        
        const myAmount = document.querySelector('#amount'),
              checked = Helpers.selectedLeave()

        API.approveLeaves(checked, connectedEmail, connectedID).then(resp => {
            if (!resp.error) location.reload()
        })

    })


    // when reject button is clicked (open modal)
    document.querySelector('#rejectBtn').addEventListener('click', () => {
       const checked = Helpers.selectedLeave()

       if (typeof checked !== 'undefined' && checked.length > 0){
            Helpers.populateSelectedLeavesNumber(checked.length, 'selectedLeaveRejection', 'reject', 'with this reason:')
            myRejectionConfirm.show()
        }

    })

    // when confirm rejection is clicked (submit)
    document.querySelector('#rejectSubmit').addEventListener('click', () => {
        const myReason = document.querySelector('#reject'),
              checked = Helpers.selectedLeave()

        if (myReason.value == '') Common.insertHTML('reject reason is required', 'rejectError')
        if (myReason.value != '') API.rejectLeaves(checked, myReason.value, connectedEmail, connectedID).then(resp => {
            if (!resp.error) location.reload()
        })

    })

    // close modal warning message
    document.querySelector('#hideModalWarningMessage').addEventListener('click', () => {
        Common.hideDivByID('modalWarningMessageDiv')
    })

    // cleaned checked checkboxes & modal values when modal is close
    document.getElementById('approveModal').addEventListener('hidden.bs.modal', function () {

        document.querySelectorAll('.leavesCheckboxes').forEach(element => {
             element.checked = false
         })

         Common.insertInputValue('', 'amount')
         Common.insertHTML('', 'amountError')

     })
    document.getElementById('rejectModal').addEventListener('hidden.bs.modal', function () {

        document.querySelectorAll('.leavesCheckboxes').forEach(element => {
             element.checked = false
         })

         Common.insertInputValue('', 'reject')
         Common.insertHTML('', 'rejectError')

     })

})
