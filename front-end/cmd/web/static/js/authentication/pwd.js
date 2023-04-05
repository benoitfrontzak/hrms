const Helpers = new PwdHelpers(),
  API = new PwdAPI(),
  Common = new MainHelpers()

// When DOM is loaded
window.addEventListener('DOMContentLoaded', () => {

  //get data from form
  let form = document.getElementById('updatePasswordForm')
  form.addEventListener('submit', (e) => {
    e.preventDefault()
    let dataStringified = Common.getForm("updatePasswordForm",connectedID)
    let data = JSON.parse(dataStringified);

    // validate data
    if (Helpers.validatePwd(data)) {
      // send data to server
      API.upPwd(connectedID, data).then(resp => {

        if (! resp.error) {
          window.location.href = '/logout'
        } else {
          alert('Password change failed')
        }
      })
    }
  })
})

  //get currentEmailFrom hidden field

