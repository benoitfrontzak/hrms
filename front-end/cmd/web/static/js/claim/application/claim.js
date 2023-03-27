const Common  = new MainHelpers(),
      Helpers = new ClaimHelpers(),
      API     = new ClaimAPI()

// set form's parameters (Required Input Fields...)
const myRIF = [ 'name', 'description', 'category']

// store all employees (active, inactive & deleted) 
const allEmployees = new Map()
allEmployees.set(0, 'not defined')

// store empty row for claims with no data (approved, rejected or pending)
const noData = {
    employeeid: 0,
    claimDefinitionID: 0,
    claimDefinition: 'No data',
    categoryID: 0,
    category: 'No data',
    name: 'No data',
    description: 'No data',
    amount: 0,
    statusID: 0,
    status: 'No data',
    approvedAt: '0001-01-01',
    approvedBy: 0,
    approvedAmount: 0,
    approvedReason: ''
}
const noDataAr = [noData]

// when DOM is loaded
window.addEventListener('DOMContentLoaded', () => {
    const noAction = false
     // fetch all employee information
     API.getAllEmployees().then(resp => {
        // update allEmployee active
        resp.data.Active.forEach(element => {
            allEmployees.set(element.ID, element.Fullname)
        })
        // update allEmployee inactive
        resp.data.Inactive.forEach(element => {
            allEmployees.set(element.ID, element.Fullname)
        })
        // update allEmployee deleted
        resp.data.Deleted.forEach(element => {
            allEmployees.set(element.ID, element.Fullname)
        })
        // fetch all claims & update DOM (data table)
        API.getAllClaims().then(resp => {
            console.log(resp);
            // display by default pending request
            if (resp.data.Pending !== null && typeof resp.data.Pending !== undefined) Helpers.generateDT(resp.data.Pending)
            if (resp.data.Pending === null || typeof resp.data.Pending === undefined) Helpers.generateDT(noData)
            // when approved is selected
            document.querySelector('#approvedBtn').addEventListener('click', () => {
                if (resp.data.Approved !== null && typeof resp.data.Approved !== undefined) Helpers.generateDT(resp.data.Approved, noAction)
                if (resp.data.Approved === null || typeof resp.data.Approved === undefined) Helpers.generateDT(noData)
            })

             // when rejected is selected
             document.querySelector('#rejectedBtn').addEventListener('click', () => {
                if (resp.data.Rejected !== null && typeof resp.data.Rejected !== undefined) Helpers.generateDT(resp.data.Rejected, noAction)
                if (resp.data.Rejected === null || typeof resp.data.Rejected === undefined) Helpers.generateDT(noData)
            })

             // when pending is selected
             document.querySelector('#pendingBtn').addEventListener('click', () => {
                if (resp.data.Pending !== null && typeof resp.data.Pending !== undefined) Helpers.generateDT(resp.data.Pending)
                if (resp.data.Pending === null || typeof resp.data.Pending === undefined) Helpers.generateDT(noData)
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
