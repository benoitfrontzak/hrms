class PasswordHelpers{

    // generate warning message
    populateWarningMessage(notValid){
        let msg
        (notValid.length == 1) ? msg = `Sorry there is an error with this field <ul>` : 
                                 msg = `Sorry there is still few errors with those fields <ul>`

        notValid.forEach(field => {
            msg += `<li>${Common.replaceCamelCase(field)}</li>`
        })

        msg += `</ul>`

        return msg
    }

    // display main warning message
    displayWarningMessage(notValid){
        const myWarningDivID  = 'warningMessageDiv',
              myWarningBodyID = 'warningMessageBody'

        if (!notValid?.length) {
            Common.hideDivByID(myWarningDivID)
        }else{
            Common.showDivByID(myWarningDivID)
            Common.insertHTML(this.populateWarningMessage(notValid), myWarningBodyID)
        }
      
    }

    // validate form
    validateForm(){
        let valid = true,
            notValid = []

        const oldP = document.querySelector('#oldPassword'),
              newP = document.querySelector('#newPassword'),
              confirmP = document.querySelector('#confirmPassword')

        if (newP.value != confirmP.value) {
            Common.insertHTML('Not matching with confirm password', newPasswordError.id)
            Common.insertHTML('Not matching with new password', confirmPasswordError.id)
            valid = false
            notValid.push('newPassword', 'confirmPassword')
        }
       
        if (newP.value == oldP.value) {
            Common.insertHTML('New password must be different from old password', newPasswordError.id)
            valid = false
            notValid.push('newPassword')
        }

        // display main warning message
        this.displayWarningMessage(notValid)

        // check password matching policy
        const newPassordMatch = this.newPasswordValid(newP.value)
        if (!newPassordMatch) valid = false   

        return valid
    }

    // new passord policy rules
    newPasswordValid(password){
        let valid = true,
            notValid = ['newPassword']
        
        // password must at least 8 characters
        if (password.length < 8) {
            Common.insertHTML('New password must be at least 8 characters', newPasswordError.id)
            valid = false
            this.displayWarningMessage(notValid)
            return valid
        }

        // password must contain at least one letter
        if (password.search(/[a-z]/i) < 0) {
            Common.insertHTML('New password must contain at least one letter', newPasswordError.id)
            valid = false
            this.displayWarningMessage(notValid)
            return valid
        }

        // password must contain at least one uppercase letter
        if (password.search(/[A-Z]/i) < 0) {
            Common.insertHTML('New password must contain at least one uppercase letter', newPasswordError.id)
            valid = false
            this.displayWarningMessage(notValid)
            return valid
        }

        // password must contain at least one digit
        if (password.search(/[0-9]/) < 0) {
            Common.insertHTML('New password must contain at least one digit', newPasswordError.id)
            valid = false
            this.displayWarningMessage(notValid)
            return valid
        }

        // password must contain at least one special character
        const specialC = /[ `!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?~]/
        if (!specialC.test(password)) {
            Common.insertHTML('New password must contain at least one special character', newPasswordError.id)
            valid = false
            this.displayWarningMessage(notValid)
            return valid
        }
        
        return valid
    }

    oldPasswordNotValid(){
        const notValid = ['oldPassword']
        Common.insertHTML('Not matching with the DB password', oldPasswordError.id)
        this.displayWarningMessage(notValid)
    }
    
}