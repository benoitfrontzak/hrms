class PwdHelpers{
    // update connected user profile information
    validatePwd(data){
        if (data.newPassword != data.confirmPassword) {
            alert('Passwords do not match')
            return false
        }   
        else if (data.newPassword == data.oldPassword) {
            alert('New password must be different from old password')
            return false
        }
        // regex to check for at least one number, one lowercase, one uppercase, one special character and at least 8 characters
        // else if (!data.newPassword.match(/^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[a-zA-Z]).{8,}$/gm)) {
        //     alert('Password must contain at least one number, one lowercase, one uppercase, one special character')
        //     return false
        // }
        else{
            return true
        }
    }

}