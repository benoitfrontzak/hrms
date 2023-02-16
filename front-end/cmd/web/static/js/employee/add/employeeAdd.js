const helpers = new EmployeeAddHelpers(),
      API     = new EmployeeAdd_API()

// Define timeout before closing info message (in milliseconds)
const myTimeout = 10000 

// When DOM is loaded
window.addEventListener('DOMContentLoaded', () => {  
    
    // When form is submitted
    const mySubmit = document.querySelector('#employeeAddSubmit')
    mySubmit.addEventListener('click', () => {
        helpers.validatePersonalInfo()
    })

    // Set parameters to enable show|hide card function
    const showCard = true,
          hideCard = false,
          myCards = ['identity', 'contact', 'emergency', 'spouseIdentity','spouseContact', 'payrollInfo', 'employmentInfo', 'epf', 'socso', 'incomeTax', 'others', 'payroll']
    // Enable show|hide function for each card
    myCards.forEach(card => {
        document.querySelector('#'+card+'Hide').addEventListener('click', () => { helpers.displayCardByID(card, hideCard) })
        document.querySelector('#'+card+'Show').addEventListener('click', () => { helpers.displayCardByID(card, showCard) })
    })

    // Calculate age
    document.querySelector('#birthdate').addEventListener('change', () => {
        const age = helpers.calculateAge(birthdate.value)
        helpers.populateAge(age)
    })

    // Auto close info message after myTimeout
    setTimeout( () => {
        helpers.hideDivByID('infoMessageDiv')
    }, myTimeout);

    // Close warning message
    const myWarningMessage = document.querySelector('#hideWarningMessage')
    myWarningMessage.addEventListener('click', () => {
        helpers.hideDivByID('warningMessageDiv')
    })

})
