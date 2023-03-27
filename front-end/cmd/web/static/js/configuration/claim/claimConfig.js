const Common  = new MainHelpers(),
      Helpers = new ClaimConfigTablesHelpers(),
      API     = new ClaimConfigTablesAPI()
         


// When DOM is loaded
window.addEventListener('DOMContentLoaded', () => {
     // fetch all employees & update DOM & set delete by icon event
     API.getClaimCT().then(resp => { 
        // populate list of config table to view data
        Helpers.populateViewData(resp.data)
        // populate config table data (when selected from view data list) 
        const myCT = document.querySelectorAll('.myConfigT')
        myCT.forEach(element => {
            element.addEventListener('click', (e) => {
                const wanted = e.target.innerHTML
                // populate config table rows
                Helpers.populateCT(resp.data[wanted], wanted)
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
                              ctName = document.querySelector('#configTableTitle').innerHTML,
                              ctValue = document.querySelector('#editValue'+rowID).value
                        console.log(ctName, ctValue, rowID);
                        API.updateConfigTableEntry(ctName,ctValue,rowID,connectedEmail).then(resp => {
                            if (!resp.error) location.reload()
                        })
                    })
                })
            })
        })        
    }) 

    // when add button is clicked (clear form)
    document.querySelector('#openModal').addEventListener('click', () => {
        Common.insertInputValue('', 'ctNewValue')
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

   // when delete entry is clicked (checkboxes)
   document.querySelector('#deleteCT').addEventListener('click', () => {
    const entryChecked = Helpers.selectedEntry(),
          ctName = document.querySelector('#configTableTitle').innerHTML
    API.deleteConfigTableEntry(entryChecked, ctName, connectedEmail).then(resp => {
        if (!resp.error) location.reload()
    })
   })

})
