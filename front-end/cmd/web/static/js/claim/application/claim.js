const Common    = new MainHelpers(),
      DT        = new DataTableFeatures(),
      Draggable = new DraggableModal(),
      Helpers   = new ClaimHelpers(),
      API       = new ClaimAPI()

// set form's parameters (Required Input Fields...)
const myRIF = ['name', 'description', 'category']

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
            const approved = resp.data.Approved,
                  rejected = resp.data.Rejected,
                  pending  = resp.data.Pending

            // display by default pending request
            Helpers.insertRows(pending)
            // create listener when data source changed (pending | approved | rejected)
            const dSource = ['pending', 'approved', 'rejected']
            dSource.forEach(element => {
                // create capitalize function
                const capitalize = element => (element && element[0].toUpperCase() + element.slice(1)) || ""

                document.querySelector('#'+element+'Btn').addEventListener('click', () => {
                    $('#claimApplicationsTable').DataTable().destroy()
                    let ds
                    switch (element) {
                        case 'pending':
                            ds = pending
                            action = true
                            break;
                        case 'approved':
                            ds = approved
                            action = false
                            break;
                        case 'rejected':
                            ds = rejected
                            action = false
                            break;                    
                    } 
                    Helpers.insertRows(ds, action)
                    document.querySelector('#claimTitle').innerHTML = capitalize(element) + ' Claim Applications'
                })
            })

            // when attachments is clicked
            $('#claimApplicationsTable').on('click', '.myAttachments', function (e) {
                const appID = e.currentTarget.dataset.appid,
                      email = e.currentTarget.dataset.email,
                      entries = e.currentTarget.dataset.entries,
                      icon = e.currentTarget.innerHTML
                Helpers.populateAttachments(appID, email, entries, icon)
            })
        })
    })

    // initiate confirm approval modal
    const myApprovalConfirm = new bootstrap.Modal(document.getElementById('approveModal'), {
        backdrop: 'static',
        keyboard: false
    })
    // initiate confirm rejection modal
    const myRejectionConfirm = new bootstrap.Modal(document.getElementById('rejectModal'), {
        backdrop: 'static',
        keyboard: false
    })

    // when approve button is clicked (open modal)
    document.querySelector('#approveBtn').addEventListener('click', () => {
        const checked = Helpers.selectedClaim()

        if (typeof checked != 'undefined' && checked.length > 0) {
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

        if (typeof checked !== 'undefined' && checked.length > 0) {
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

    // make modals draggable
    Draggable.draggableModal('approveModal')
    Draggable.draggableModal('rejectModal')
    Draggable.draggableModal('uploadedFilesModal')
})
