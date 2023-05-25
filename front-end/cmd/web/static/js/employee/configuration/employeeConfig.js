const Common    = new MainHelpers(),
      Draggable = new DraggableModal(),
      Helpers   = new EmployeeConfigTablesHelpers(),
      API       = new EmployeeConfigTablesAPI()

// When DOM is loaded
window.addEventListener('DOMContentLoaded', () => {

     // fetch all employees & update DOM & set delete by icon event
     API.getEmployeeCT().then(resp => { 
        // populate config table data when sidebar element is clicked
        document.querySelectorAll('.myConfigT').forEach(element => {
            element.addEventListener('click', (e) => {
                const wanted = e.target.id,
                      data = resp.data[wanted]
                // show page
                Common.showDivByID('pageDiv')
                // make link active
                Helpers.makeLinkActive(wanted)
                // populate config table rows data.shift()
                Helpers.populateCT(data, e.target.innerHTML, wanted)

                // when one PI is clicked (edit name)
                document.querySelectorAll('.showSave').forEach(element => {             
                    element.addEventListener('click', (e) => {
                        const rowID = e.target.parentElement.dataset.id
                        Common.showDivByID('saveDivPI'+rowID)
                        Common.hideDivByID('deleteDivPI'+rowID)
                    })
                })

                // when one entry is clicked (edit name)
                document.querySelectorAll('.showValue').forEach(element => {             
                    element.addEventListener('click', (e) => {
                        const rowID = e.target.dataset.id
                        Helpers.editableRow(rowID)
                    })
                })

                // when save entry is clicked (save button)
                document.querySelectorAll('.saveValue').forEach(element => {             
                    element.addEventListener('click', (e) => {
                        const rowID = e.target.dataset.id,
                              ctName = document.querySelector('#ctName').value,
                              ctValue = document.querySelector('#editValue'+rowID).value
                        API.updateConfigTableEntry(ctName,ctValue,rowID,connectedEmail).then(resp => {
                            if (!resp.error) location.reload()
                        })
                    })
                })

                // when save PI entry is clicked (save button)
                document.querySelectorAll('.saveValuePI').forEach(element => {             
                    element.addEventListener('click', (e) => {
                        const rowID = e.target.dataset.id,
                              data = Helpers.getRowPI(rowID)
                        API.updateEntryPI(data).then(resp => {
                            if (!resp.error) location.reload()
                        })
                    })
                })

                // initiate tooltips
                const tooltipTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="tooltip"]'));
                tooltipTriggerList.map(function (tooltipTriggerEl) {
                    return new bootstrap.Tooltip(tooltipTriggerEl);
                })

            })
        })

    }) 

    // when add button is clicked (clear form)
    document.querySelector('#openModal').addEventListener('click', () => {
        Common.insertInputValue('', 'ctNewValue')
    })
    // when add PI button is clicked (clear form)
    document.querySelector('#openModalPI').addEventListener('click', () => {
        Helpers.cleanForm()
    })

    // when add new entry is clicked (submit)
    document.querySelector('#ctSubmit').addEventListener('click', () => {
        if (document.querySelector('#ctNewValue').value != '') {
            const myCT = document.querySelector('#ctName').value,
                  myValue = document.querySelector('#ctNewValue').value
            API.addNewConfigTableEntry(myCT, myValue, connectedEmail).then(resp => {
                if (!resp.error) location.reload()
            })
        }
    })

    // when add new PI entry is clicked (submit)
    document.querySelector('#piSubmit').addEventListener('click', () => {
        const errors = Helpers.validateForm()
        if (errors == 0){
            const data = Helpers.getForm()
            API.addNewPI(data).then(resp => {
                if (!resp.error) location.reload()
            })
        }        
    })

   // when delete entry is clicked (checkboxes)
   document.querySelector('#deleteCT').addEventListener('click', () => {
    const entryChecked = Helpers.selectedEntry(),
          ctName = document.querySelector('#ctName').value
    API.deleteConfigTableEntry(entryChecked, ctName, connectedEmail).then(resp => {
        if (!resp.error) location.reload()
    })
   })

   // make modals draggable
   Draggable.draggableModal('ctModal')
   Draggable.draggableModal('piModal')
   Draggable.draggableModal('confirmDelete')

})
