const Common    = new MainHelpers(),
      DT        = new DataTableFeatures(),
      Draggable = new DraggableModal(),
      Helpers   = new ClaimDefinitionHelpers(),
      API       = new ClaimDefinitionAPI()

// set form's parameters (Required Input Fields...)
const myRIF = ['name', 'description', 'category']

let rowDetailsNumber = 1

// store all employees (active, inactive & deleted) 
let allEmployees = new Map()
allEmployees.set(0, 'not defined')

// when DOM is loaded
window.addEventListener('DOMContentLoaded', () => {

    // fetch all employee information
    API.getAllEmployees().then(resp => {
        allEmployees = Common.updateEmployeeList(resp.data, allEmployees)
        // fetch all claim's definition & update DOM 
        API.getAllClaimDefinition().then(resp => {
            // display active claim definition
            Helpers.insertRows(resp.data.Active)
            Helpers.makeEditable()
        })
    })

    // fetch claim category & update DOM (form dropdown)
    API.getClaimCT().then(resp => {
        Helpers.insertOptions('category', resp.data.Category)
    })

    // when open form (add button) is clicked (clear warning message & field's values)
    document.querySelector('#openCreateClaimDefinition').addEventListener('click', () => {
        Helpers.removeRowFromDivsByClass('detailRemovable')
        Helpers.insertDefaultDetailsRow()
        Common.clearForm('createClaimDefinitionForm', myRIF)
        Common.insertInputValue('create', 'formAction')
    })

    // when form is submitted (save button)
    document.querySelector('#createClaimDefinitionSubmit').addEventListener('click', () => {
        const error = Common.validateRequiredFields(myRIF),
            formAction = document.querySelector('#formAction').value

        // create form   
        if (error == '0' && formAction == 'create') {
            myData = Common.getForm('createClaimDefinitionForm', connectedID)
            myData = Helpers.extractDetailsFromForm(myData)
            API.createClaimDefinition(myData).then(resp => {
                if (!resp.error) location.reload()
            })
        }

        // edit form
        if (error == '0' && formAction == 'edit') {
            myData = Common.getForm('createClaimDefinitionForm', connectedID)
            myData = Helpers.extractDetailsFromForm(myData)

            const allRows = document.querySelectorAll(".existingDetailsRow");
            const deleteArray = [];
            const updateArray = [];

            allRows.forEach(function (row) {
                const rowId = row.id;
                const senValue = row.querySelector(".existingSeniority").value;
                const limValue = row.querySelector(".existingLimitation").value;
                const rowData = {
                    rowid: rowId,
                    seniority: senValue,
                    limitation: limValue
                };

                const checkbox = row.querySelector(".toggle-checkbox");
                if (checkbox.checked) {
                    deleteArray.push(rowId); //push row.id instead since backed takes slice of string
                } else {
                    updateArray.push(rowData);
                }
            });
            const newRows = document.querySelectorAll(".newRowDetails");
            newRows.forEach(function (row2) {
                const rowId = row2.id;
                const senValue = row2.querySelector(".newSeniority").value;
                const limValue = row2.querySelector(".newLimitation").value;
                const rowData = {
                    rowid: rowId,
                    seniority: senValue,
                    limitation: limValue
                };
            });

            myData = Helpers.addUpdateDetailsAndDeleterowIDToForm(myData, updateArray, deleteArray);

            API.updateClaimDefinition(myData).then(resp => {
                if (!resp.error) location.reload()
            })
        }
    })

    // close warning message
    document.querySelector('#hideWarningMessage').addEventListener('click', () => {
        Common.hideDivByID('warningMessageDiv')
    })

    // initiate delete confirm modal
    const myConfirm = new bootstrap.Modal(document.getElementById('confirmDelete'), {
        keyboard: false
    })

    // cleaned checked checkboxes when modal is close
    document.getElementById('confirmDelete').addEventListener('hidden.bs.modal', function (event) {
        const myDeleteCheckboxes = document.querySelectorAll('.deleteCheckboxes')

        myDeleteCheckboxes.forEach(element => {
            element.checked = false
        })

    })

    // when delete all is clicked
    const myDeleteAll = document.querySelector('#deleteAllClaimDefinition')

    myDeleteAll.addEventListener('click', () => {
        const checked = Helpers.selectedClaimDefinition()

        if (typeof checked != 'undefined' && checked.length > 0) {
            Helpers.populateConfirmDelete(checked.length)
            myConfirm.show()
        }

    })

    // when confirm delete all is clicked
    const confirmedDelete = document.querySelector('#confirmDeleteSubmit')

    confirmedDelete.addEventListener('click', () => {
        const checked = Helpers.selectedClaimDefinition()

        API.softDeleteClaimDefinition(checked, connectedEmail).then(resp => {
            if (!resp.error) {
                myConfirm.hide()
                location.reload()
            }
        })

    })
    // when add details row is clicked
    document.querySelector('#addDetails').addEventListener('click', () => {
        Helpers.insertDetailsRow()
    })

     // make modals draggable
     Draggable.draggableModal('createClaimDefinition')
     Draggable.draggableModal('confirmDelete')

})
