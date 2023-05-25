const Common    = new MainHelpers(),
      DT        = new DataTableFeatures(),
      Draggable = new DraggableModal(),
      Helpers   = new MyClaimHelpers(),
      API       = new MyClaimAPI()

// set form's parameters (Required Input Fields...)
const myRIF = ['claimDefinition', 'amount', 'description']

// store all employee by id
let allEmployees = new Map()
allEmployees.set(0, 'not defined')

// store all claim definition by id (form add|edit when claim definition is selected...)
let allClaimDefinition = new Map()

// store all my claims by claim definition id (form add|edit when claim definition is selected...)
let myClaim = new Map()

// store all uploaded files by leave application id
let myUploadedFiles = new Map()

// store requested leave max entitled days allowed
let maxClaimAllow = 0.0

// store if requested leave attachment is required or not
let attachmentRequired = 0

// when DOM is loaded
window.addEventListener('DOMContentLoaded', () => {
    // fetch uploaded files
    API.getUploadedFiles(connectedEmail).then(resp => {
        if (!resp.error)
            if (Object.keys(resp.data.Files).length > 0) myUploadedFiles = Helpers.populateUploadedFilesMap(resp.data.Files, myUploadedFiles)            
    })

    // fetch all needed informations  
    API.getAllInformationsMyClaims(connectedID, connectedEmail).then(resp => {
        console.log('resp');
        console.log(resp);
        // update variables
        allEmployees        = Common.updateEmployeeList(resp.AllEmployees, allEmployees)
        allClaimDefinition  = Helpers.populateClaimDefinitionMap(resp.AllClaimsDefinitions.Active, allClaimDefinition)
        if (resp.MyClaims != null) myClaim = Helpers.populateMyClaimMap(resp.MyClaims, myClaim)

        const mySeniority   = Number(resp.MySeniority)

        // insert rows master list table: myClaims
        Helpers.insertRows(resp.MyClaims)
        Helpers.makeEditable()

        // insert select option claim definition (form dropdown)
        Helpers.insertOptionsCD('claimDefinition', resp.AllClaimsDefinitions.Active)

        // when claim definition change (form)
        document.querySelector('#claimDefinition').addEventListener('change', (e) => {
            const claimDefinitionID = e.target.value,
                  docRequired = allClaimDefinition.get(claimDefinitionID).docRequired,
                  claimDefinition = allClaimDefinition.get(claimDefinitionID)
            let alreadyRequested = null;
            
            if (resp.MyClaims != null) alreadyRequested = myClaim.get(claimDefinitionID)
            // check if claim requires attachment
            if (docRequired == '1'){
                attachmentRequired = 1
                Common.showDivByID('myAttachmentDiv')
            }else{
                attachmentRequired = 0
                Common.hideDivByID('myAttachmentDiv')
            } 

            // display selected claim requirements
            Helpers.selectedClaimRequirements(claimDefinition, alreadyRequested, mySeniority)
        })

        // when attachments is clicked
        $('#myClaimTable').on('click', '.myAttachments', function (e) {
            const appID = e.currentTarget.dataset.id
            Helpers.populateAttachments(appID)
        })
    })


    // when form is submitted (save button)
    document.querySelector('#claimFormSubmit').addEventListener('click', () => {
        const checkRequired = Common.validateRequiredFields(myRIF)
        if (checkRequired == 0){
            checkApplication = Helpers.validateApplication()
            if (checkApplication == 0){
                const myData = Common.getForm('claimForm', connectedID)
                API.createClaim(myData).then(resp => {
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
        }
    })

    // close warning message
    const myWarningMessage = document.querySelector('#hideWarningMessage')
    myWarningMessage.addEventListener('click', () => {
        Common.hideDivByID('warningMessageDiv')
    })
    // clear form (when open form is clicked)
    document.querySelector('#openCreateMyClaim').addEventListener('click', () => {
        Common.clearForm('claimForm', myRIF)
        Common.hideDivByID('myAttachmentDiv')
        Common.hideDivByID('myAuthorizationDiv')
        document.querySelector('#uploadedFiles').value = ''
        attachmentRequired = 0
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
    document.querySelector('#deleteAllMyClaim').addEventListener('click', () => {
        const checked = Helpers.selectedClaim()

        if (typeof checked != 'undefined' && checked.length > 0) {
            Helpers.populateConfirmDelete(checked.length)
            myConfirm.show()
        }

    })

    // When confirm delete all is clicked
    const confirmedDelete = document.querySelector('#confirmDeleteSubmit')

    confirmedDelete.addEventListener('click', () => {
        const checked = Helpers.selectedClaim()

        API.softDeleteClaim(checked, connectedEmail).then(resp => {
            if (!resp.error) {
                myConfirm.hide()
                location.reload()
            }
        })

    })
  
    // make modals draggable
    Draggable.draggableModal('myClaim')
    Draggable.draggableModal('confirmDelete')
})
