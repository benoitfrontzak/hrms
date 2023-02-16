const Common  = new Main_Helpers(),
      Helpers = new EmployeeRead_Helpers(),
      API     = new EmployeeRead_API()

// Define timeout before closing info message (in milliseconds)
const myTimeout = 10000 

// When DOM is loaded
window.addEventListener('DOMContentLoaded', () => {
    // set Required Input Fields
    const myRIF = [ 'firstName', 'middleName', 'familyName', 'employeeCode', 
                    'streetAddressLine1','streetAddressLine2', 'zip', 'city', 'state', 'country', 
                    'nationality', 'residence',
                    'primaryPhone', 'primaryEmail'],
          myForm = 'createEmployeeForm'
          
    // when openCreateEmployee is clicked
    const myModal = document.querySelector('#openCreateEmployee')
    myModal.addEventListener('click', () =>{
        Common.clearForm(myForm, myRIF)
    })

    // When form is submitted
    const mySubmit = document.querySelector('#employeeAddSubmit')
    mySubmit.addEventListener('click', () => {
        const error = Common.validateRequiredFields(myRIF)
        if (error == '0'){
            myData = Common.getForm(myForm)
              API.createEmployee(myData).then(resp => {
                console.log(resp)
              })
            
        }
    })

    // Set parameters to enable show|hide card function
    const showCard = true,
          hideCard = false,
          myCards = ['identity', 'contact', 'emergency'];
    // Enable show|hide function for each card
    myCards.forEach(card => {
        document.querySelector(`#${card}Hide`).addEventListener('click', () => { Common.displayCardByID(card, hideCard) })
        document.querySelector(`#${card}Show`).addEventListener('click', () => { Common.displayCardByID(card, showCard) })
    })

    // Close warning message
    const myWarningMessage = document.querySelector('#hideWarningMessage')
    myWarningMessage.addEventListener('click', () => {
        Common.hideDivByID('warningMessageDiv')
    })

})
