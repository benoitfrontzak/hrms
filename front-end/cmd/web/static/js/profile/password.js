const Helpers = new PasswordHelpers(),
  API = new PasswordAPI(),
  Common = new MainHelpers()

// set form's parameters (Required Input Fields...)
const myRIF = ['oldPassword', 'newPassword', 'confirmPassword']

// When DOM is loaded
window.addEventListener('DOMContentLoaded', () => {

  //get data from form
  document.querySelector('#updatePasswordBtn').addEventListener('click', () => {
    const error = Common.validateRequiredFields(myRIF)         

    // check first that all required fields are filled
    if (!error) {
        // check if form is valid (password matching rules...)
        valid = Helpers.validateForm()
        if (valid){
            const data = Common.getForm('updatePasswordForm', connectedID)
            API.updateMyPassword(data).then(resp => {
                if (!resp.error) window.location.href = '/logout'
                if((resp.error) && (resp.message == 'old password is not matching')) Helpers.oldPasswordNotValid()          
            })
        }       
    }
  })

   // close warning message
   const myWarningMessage = document.querySelector('#hideWarningMessage')
   myWarningMessage.addEventListener('click', () => {
       Common.hideDivByID('warningMessageDiv')
   })
   
})




  