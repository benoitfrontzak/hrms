const Common    = new MainHelpers(),
      DT        = new DataTableFeatures(),
      Draggable = new DraggableModal(),
      Helpers   = new MyPublicHolidaysHelpers(),
      API       = new MyPublicHolidaysAPI()

// store all employee by id
let allEmployees = new Map()
allEmployees.set(0, 'not defined')

// store public holidays 
let ph = []

// when DOM is loaded
window.addEventListener('DOMContentLoaded', () => {
    // fetch all employees  
    API.getAllEmployees().then(resp => {
        allEmployees = Common.updateEmployeeList(resp.data, allEmployees)
        API.getAllPublicHolidays().then(resp => {
            Helpers.insertRows(resp.data)
            Helpers.populatePublicHolidays(resp.data, ph)
        })
    })
    
    // when form is submited
    document.querySelector('#phSubmit').addEventListener('click', () => {
        const requiredErrors = Common.validateRequiredFields(['datePublicHoliday', 'namePublicHoliday'])
        // validate required fields
        if(requiredErrors == 0){
            const dateError = Helpers.validateForm()
            // validate new public holiday date doesn't already exists
            if (dateError == 0) {
                const myData = Helpers.getForm(connectedID)
                // add new public holiday and reload page
                API.createPH(myData).then(resp => {
                    if (!resp.error) location.reload()
                })
            }
        }
    })

    // when load from CSV is submitted
    document.getElementById('csvForm').addEventListener('submit', function (e) {
        e.preventDefault() // Prevent the default form submission
    
        const requiredError = Helpers.validateFormCSV()
        if (requiredError == 0){
            // Get the uploaded file
            const file = document.getElementById('csvFile').files[0]
            
            // Create a FileReader
            const reader = new FileReader()
        
            // Set the onload event handler
            reader.onload = function (event) {
                const csvData = event.target.result // CSV data as a string
            
                // Parse the CSV data (using a library or custom parsing logic)
                const parsedData = Helpers.parseCSV(csvData),
                      validPH = Helpers.validateParsedData(parsedData),
                      myCSV = Helpers.getFormCSV(validPH, connectedID)

                API.createFromCSV(myCSV).then(resp => {
                    if (!resp.error) location.reload()
                })
            }
        
            // Read the file as text
            reader.readAsText(file)
        }
        
    })
    
    // clean form CSV when open
    document.querySelector('#addCSV').addEventListener('click', () => {
        Helpers.cleanFormCSV()
    })

    // clean form when open
    document.querySelector('#openCreatePublicHoliday').addEventListener('click', () => {
        Helpers.cleanForm()
    })

    // close warning message (row)
    document.querySelector('#hideRowWarningMessage').addEventListener('click', () => {
        Common.hideDivByID('rowWarningMessageDiv')
    })

    // close warning message (form)
    document.querySelector('#hideWarningMessage').addEventListener('click', () => {
        Common.hideDivByID('warningMessageDiv')
    })

    // close warning message (delete)
    document.querySelector('#hidedeleteWarningMessage').addEventListener('click', () => {
        Common.hideDivByID('deleteWarningMessageDiv')
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
    document.querySelector('#deleteAllPublicHoliday').addEventListener('click', () => { 
        const checked = Helpers.selectedPH()

        if (typeof checked != 'undefined' && checked.length > 0){
            Helpers.populateConfirmDelete(checked.length)
            myConfirm.show()
        }

    })

    // When confirm delete all is clicked
    const confirmedDelete = document.querySelector('#confirmDeleteSubmit')

    confirmedDelete.addEventListener('click', () => {
        const checked = Helpers.selectedPH()
        
        API.softDeletePH(checked, connectedEmail, connectedID).then(resp => {
            if (!resp.error){
                myConfirm.hide()
                location.reload()
            }
        })

    })

    // make modals draggable 
    Draggable.draggableModal('confirmDelete')
    Draggable.draggableModal('createPublicHoliday')

})