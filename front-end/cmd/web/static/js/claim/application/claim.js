const Common  = new MainHelpers(),
      Helpers = new ClaimHelpers(),
      API     = new ClaimAPI()

// set form's parameters (Required Input Fields...)
const myRIF = [ 'name', 'description', 'category']

// store all employees (active, inactive & deleted) 
let allEmployees = new Map()
allEmployees.set(0, 'not defined')

// when DOM is loaded
window.addEventListener('DOMContentLoaded', () => {
    const noAction = false
     // fetch all employee information
     API.getAllEmployees().then(resp => {
        allEmployees = Common.updateEmployeeList(resp.data, allEmployees)
        
        // fetch all claims & update DOM (data table)
        API.getAllClaims().then(resp => {
            // display by default pending request
            Helpers.insertRows(resp.data.Pending)
            // when approved is selected
            document.querySelector('#approvedBtn').addEventListener('click', () => {
                Helpers.insertRows(resp.data.Approved, noAction)
                document.querySelector('#claimTitle').innerHTML = 'Approved Claims'
            })

             // when rejected is selected
             document.querySelector('#rejectedBtn').addEventListener('click', () => {
                Helpers.insertRows(resp.data.Rejected, noAction)
                document.querySelector('#claimTitle').innerHTML = 'Rejected Claims'
            })

             // when pending is selected
             document.querySelector('#pendingBtn').addEventListener('click', () => {
                Helpers.insertRows(resp.data.Pending)
                document.querySelector('#claimTitle').innerHTML = 'Pending Claims'
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
       const checked = Helpers.selectedClaim()

       if (typeof checked != 'undefined' && checked.length > 0){
            Helpers.populateSelectedClaimsNumber(checked.length, 'selectedClaimsApproval', 'approve', 'amount')
            myApprovalConfirm.show()
        }

    })
    // when confirm approval is clicked (submit)
    document.querySelector('#approveSubmit').addEventListener('click', () => {
        
        const myAmount = document.querySelector('#amount'),
              checked = Helpers.selectedClaim()

        if (myAmount.value == '') Common.insertHTML('amount is required', 'amountError')
        if (myAmount.value != '') API.approveClaims(checked, myAmount.value, connectedEmail, connectedID).then(resp => {
            if (!resp.error) location.reload()
        })

    })


    // when reject button is clicked (open modal)
    document.querySelector('#rejectBtn').addEventListener('click', () => {
       const checked = Helpers.selectedClaim()

       if (typeof checked !== 'undefined' && checked.length > 0){
            Helpers.populateSelectedClaimsNumber(checked.length, 'selectedClaimsRejection', 'reject', 'reason')
            myRejectionConfirm.show()
        }

    })

    // when confirm rejection is clicked (submit)
    document.querySelector('#rejectSubmit').addEventListener('click', () => {
        const myReason = document.querySelector('#reject'),
              checked = Helpers.selectedClaim()

        if (myReason.value == '') Common.insertHTML('reject reason is required', 'rejectError')
        if (myReason.value != '') API.rejectClaims(checked, myReason.value, connectedEmail, connectedID).then(resp => {
            if (!resp.error) location.reload()
        })

    })

    // close modal warning message
    document.querySelector('#hideModalWarningMessage').addEventListener('click', () => {
        Common.hideDivByID('modalWarningMessageDiv')
    })

    // cleaned checked checkboxes & modal values when modal is close
    document.getElementById('approveModal').addEventListener('hidden.bs.modal', function () {

        document.querySelectorAll('.claimsCheckboxes').forEach(element => {
             element.checked = false
         })

         Common.insertInputValue('', 'amount')
         Common.insertHTML('', 'amountError')

     })
    document.getElementById('rejectModal').addEventListener('hidden.bs.modal', function () {

        document.querySelectorAll('.claimsCheckboxes').forEach(element => {
             element.checked = false
         })

         Common.insertInputValue('', 'reject')
         Common.insertHTML('', 'rejectError')

     })

})
