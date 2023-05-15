const Common    = new MainHelpers(),
      DT        = new DataTableFeatures(),
      Draggable = new DraggableModal(),
      Helpers   = new LeaveHelpers(),
      API       = new LeaveAPI()

// set form's parameters (Required Input Fields...)
const myRIF = [ 'name', 'description', 'category']

// store all employees (active, inactive & deleted) 
let allEmployees = new Map()
allEmployees.set(0, 'not defined')

// when DOM is loaded
window.addEventListener('DOMContentLoaded', () => {
    let action = false
     // fetch all employee information
     API.getAllEmployees().then(resp => {
        allEmployees = Common.updateEmployeeList(resp.data, allEmployees)

        // fetch all leaves & update DOM (data table)
        API.getAllLeaves().then(resp => {
            const approved = resp.data.Approved,
                  rejected = resp.data.Rejected,
                  pending  = resp.data.Pending

            // display by default pending request
            Helpers.insertRows(resp.data.Pending)

            // create listener when data source changed (pending | approved | rejected)
            const dSource = ['pending', 'approved', 'rejected']
            dSource.forEach(element => {
                // create capitalize function
                const capitalize = element => (element && element[0].toUpperCase() + element.slice(1)) || ""

                document.querySelector('#'+element+'Btn').addEventListener('click', () => {
                    $('#leaveApplicationTable').DataTable().destroy()
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
                    document.querySelector('#leaveTitle').innerHTML = capitalize(element) + ' Leave Applications'
                })
            })
  
            // when attachments is clicked
            $('#leaveApplicationTable').on('click', '.myAttachments', function (e) {
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
       const checked = Helpers.selectedLeave()

       if (typeof checked != 'undefined' && checked.length > 0){
            Helpers.populateSelectedLeavesNumber(checked.length, 'selectedLeavesApproval', 'approve', '')
            myApprovalConfirm.show()
        }

    })
    // when confirm approval is clicked (submit)
    document.querySelector('#approveSubmit').addEventListener('click', () => {
        
        const checked = Helpers.selectedLeave()

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

     })
    document.getElementById('rejectModal').addEventListener('hidden.bs.modal', function () {

        document.querySelectorAll('.leavesCheckboxes').forEach(element => {
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
