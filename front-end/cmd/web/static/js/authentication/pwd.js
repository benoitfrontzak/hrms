const Helpers = new PwdHelpers(),
  API = new PwdAPI(),
  Common = new MainHelpers()

// set form's parameters (Required Input Fields...)
const myRIF = ['oldPassword', 'newPassword', 'confirmPassword']

// When DOM is loaded
window.addEventListener('DOMContentLoaded', () => {

  //get data from form
  let form = document.getElementById('updatePasswordForm')
  document.querySelector('#submitPwdBtn').addEventListener('click', () => {
    const error = Common.validateRequiredFields(myRIF)

    let dataStringified = Common.getForm("updatePasswordForm", connectedID)
    let data = JSON.parse(dataStringified);

    // validate data
    if (!error) {
      if (Helpers.validatePwd(data)) {
        // send data to server
        API.upPwd(connectedID, data).then(resp => {

          if (!resp.error) {
            window.location.href = '/logout'
          } else {
            alert('Password change failed')
          }
        })
      }
    }
  })
})

  //get currentEmailFrom hidden field



  