const Common  = new MainHelpers(),
      Helpers = new MyLeaveHelpers(),
      API     = new MyLeaveAPI()

// set form's parameters (Required Input Fields...)
const myRIF = ['leaveDefinition', 'description', 'requestedDates']

// store all employee by id
let allEmployees = new Map()
allEmployees.set(0, 'not defined')

// store all requested dates
const myRequestedDates = []

// when DOM is loaded
window.addEventListener('DOMContentLoaded', () => {
    // fetch all needed informations  
    API.getAllInformationsMyLeave(connectedID, connectedEmail).then(resp => {
        console.log(resp);
        allEmployees = Common.updateEmployeeList(resp.AllEmployees, allEmployees)
        Helpers.generateDT(resp.MyLeaves)
        Helpers.insertOptions('leaveDefinition', resp.AllLeaveDefinitions)
    })
      
    // when leave definition change
    document.querySelector('#leaveDefinition').addEventListener('change', () => {
        // todo fetch employee authorization (how many entitled days...)
        Common.insertHTML('You are entitled for xx days, you already took xx days. Your balance is xx days <br><hr>todo...', 'myAuthorizationDiv')
        Common.showDivByID('myAuthorizationDiv')
    })

    
    // when form is submitted (save button)
    document.querySelector('#leaveFormSubmit').addEventListener('click', () => {
        const error = Common.validateRequiredFields(myRIF)
        myData = Helpers.getForm('leaveForm')

        if (error == '0'){
            API.createLeave(myData).then(resp => {
                console.log(resp);
                if (! resp.error) location.reload()
            })
        }
    })
    // close warning message
    const myWarningMessage = document.querySelector('#hideWarningMessage')
    myWarningMessage.addEventListener('click', () => {
        Common.hideDivByID('warningMessageDiv')
    })
    // clear form (when open form is clicked)
    document.querySelector('#openCreateMyLeave').addEventListener('click', () =>{ 
        Common.clearForm('leaveForm', myRIF) 
        Common.hideDivByID('myAuthorizationDiv')
    })

    // initiate delete confirm modal
    const myConfirm = new bootstrap.Modal(document.getElementById('confirmDelete'), { 
        keyboard: false 
    })

    // cleaned checked checkboxes when modal is close
    document.getElementById('confirmDelete').addEventListener('hidden.bs.modal', function () {
       document.querySelectorAll('.deleteCheckboxes').forEach(element => {
            element.checked = false
        })
    })

    // When delete all is clicked
    document.querySelector('#deleteAllMyLeave').addEventListener('click', () => { 
        const checked = Helpers.selectedLeave()

        if (typeof checked != 'undefined' && checked.length > 0){
            Helpers.populateConfirmDelete(checked.length)
            myConfirm.show()
        }

    })

    // When confirm delete all is clicked
    const confirmedDelete = document.querySelector('#confirmDeleteSubmit')

    confirmedDelete.addEventListener('click', () => {
        const checked = Helpers.selectedLeave()

        API.softDeleteLeave(checked, connectedEmail).then(resp => {
            if (!resp.error){
                myConfirm.hide()
                location.reload()
            }
        })

    })

    // initiate date picker
    $('#datepicker').datepicker({
        format: 'yyyy/mm/dd',
        multidate: true,
        daysOfWeekDisabled: [0, 6],
        clearBtn: true,
        todayHighlight: true,
        daysOfWeekHighlighted: [1, 2, 3, 4, 5]
    })
    // populate selected dates
    $('#datepicker').on('changeDate', function() {
        $('#requestedDates').val(
            $('#datepicker').datepicker('getFormattedDate')
        )
        const myStringDates = document.querySelector('#requestedDates').value
        Helpers.populateSelectedDates(myStringDates.split(','))
    })
})
