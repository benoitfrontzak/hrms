const Common    = new MainHelpers(),
      DT        = new DataTableFeatures(),
      Draggable = new DraggableModal(),
      Helpers   = new MyLeaveHelpers(),
      API       = new MyLeaveAPI()

// set form's parameters (Required Input Fields...)
const myRIF = ['leaveDefinition', 'description', 'requestedDates']

// store all employee by id
let allEmployees = new Map()
allEmployees.set(0, 'not defined')

// store all requested dates
// const myRequestedDates = []

// store all leave definition by id (form add|edit when leave definition is selected...)
let allLeaveDefinition = new Map()

// store all my entitled leave by leave definition id (form add|edit when leave definition is selected...)
let myEntitled = new Map()

// store all uploaded files by leave application id
let myUploadedFiles = new Map()

// store dates to be disabled from calendar (all pending & approved leave applications)
let datesForDisable = []

// store public holidays 
let ph = []

// store requested leave max entitled days allowed
let maxEntitledAllow = 0.0

// store if requested leave attachment is required or not
let attachmentRequired = 0


// when DOM is loaded
window.addEventListener('DOMContentLoaded', () => {
    // fetch uploaded files
    API.getUploadedFiles(connectedEmail).then(resp => {
        if (!resp.error)
            if (Object.keys(resp.data.Files).length > 0) myUploadedFiles = Helpers.populateUploadedFilesMap(resp.data.Files, myUploadedFiles)            
    })
    
    // fetch all public holidays
    API.getAllPublicHolidays().then(resp => {
        Helpers.populatePublicHolidays(resp.data, ph)
    })

    // fetch all needed informations  
    API.getAllInformationsMyLeave(connectedID, connectedEmail).then(resp => {
        // update variables
        allEmployees        = Common.updateEmployeeList(resp.AllEmployees, allEmployees)
        allLeaveDefinition  = Helpers.populateLeaveDefinitionMap(resp.AllLeaveDefinitions, allLeaveDefinition)
        myEntitled          = Helpers.populateMyEntitledLeaveMap(resp.MyEntitled, myEntitled)
        myLeaveDatesTaken   = Helpers.populateDatesForDisable(resp.MyLeaves)
        datesForDisable     = Helpers.sortArray([...ph, ...myLeaveDatesTaken])

        // connected employee information
        const mySeniority   = Number(resp.MySeniority),
              myGender      = Helpers.myGender(Number(connectedID), resp.AllEmployees.Active)

        // initiate date picker (TODO: Add Public Holidays)
        $('#datepicker').datepicker({
            format: 'yyyy/mm/dd',
            multidate: true,
            daysOfWeekDisabled: [0, 6],
            clearBtn: true,
            todayHighlight: true,
            daysOfWeekHighlighted: [1, 2, 3, 4, 5],
            datesDisabled: datesForDisable
        })
        
        // insert rows master list table: myLeaves
        Helpers.insertRows(resp.MyLeaves)

        // insert select options to form: add new leave request (leave definition selectbox)
        Helpers.insertOptions('leaveDefinition', resp.AllLeaveDefinitions, myGender)

        // when leave definition change
        document.querySelector('#leaveDefinition').addEventListener('change', (e) => {
            const selectedIndex    = e.target.value,
                  leaveDefinition  = allLeaveDefinition.get(selectedIndex),
                  entitlement      = myEntitled.get(selectedIndex)
            Helpers.selectedLeaveRequirements(leaveDefinition, entitlement, mySeniority)
            Common.showDivByID('calendarDatePicker')
            Common.insertHTML('', 'selectedDatesDiv')    
        })

        // when attachments is clicked
        $('#myLeaveTable').on('click', '.myAttachments', function (e) {
            const appID = e.currentTarget.dataset.id
            Helpers.populateAttachments(appID)
        })
    })
 
    // when form is submitted (save button)
    document.querySelector('#leaveFormSubmit').addEventListener('click', () => {
        const checkError = Helpers.validateForm(myRIF)

        if (checkError == 0){
            const myData = Helpers.getForm('leaveForm')

            API.createLeave(myData).then(resp => {
                if (! resp.error) {
                    // check if got uploaded files
                    const uFiles = document.querySelector('#uploadedFiles')
                    if (uFiles.value == '') location.reload() 
                    if (uFiles.value != ''){
                        const leaveApplicationID = resp.data
                        Helpers.SendAttachment(leaveApplicationID, connectedEmail, connectedID)
                    }
                }
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
        Common.hideDivByID('myAttachmentDiv')
        Common.hideDivByID('calendarDatePicker')
        Common.insertHTML('', 'selectedDatesDiv')
        Common.insertInputValue('', 'uploadedFiles')
    })

    // initiate delete confirm modal
    const myConfirm = new bootstrap.Modal(document.getElementById('confirmDelete'), { 
        backdrop: 'static',
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

        API.softDeleteLeave(checked, connectedEmail, connectedID).then(resp => {
            if (!resp.error){
                myConfirm.hide()
                location.reload()
            }
        })

    })

    // populate selected dates
    $('#datepicker').on('changeDate', function() {
        $('#requestedDates').val(
            $('#datepicker').datepicker('getFormattedDate')
        )
        const myStringDates = document.querySelector('#requestedDates').value
        Helpers.populateSelectedDates(myStringDates.split(','))
    })

    // TODO : Fixed bug when click on clear dates (it don't remove first row)
    // $('#datepicker').on('clearDate', function(e) {
    //     document.querySelector('#selectedDatesDiv').innerHTML = ''
    //     // $('#selectedDatesDiv')[0].html('')
    // })

    // make modals draggable
    Draggable.draggableModal('confirmDelete')
    Draggable.draggableModal('myLeave')
    Draggable.draggableModal('uploadedFilesModal')
})
