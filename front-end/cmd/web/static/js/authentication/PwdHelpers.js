class PwdHelpers{
    // update connected user profile information
    validatePwd(data){
        if (data.newPassword != data.confirmPassword) {
            Common.insertHTML('These two are not the same please check', newPasswordError.id);
            Common.insertHTML('These two are not the same please check', confirmPasswordError.id);
            return false
        }   
        else if (data.newPassword == data.oldPassword) {
            Common.insertHTML('New password must be different from old password', newPasswordError.id);
            return false
        }
        // // Maybe uncomment this during production ..... Commented for demo purposes.
        // // regex to check for at least one number, one lowercase, one uppercase, one special character and at least 8 characters
        // else if (!data.newPassword.match(/^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[a-zA-Z]).{8,}$/gm)) {
        //     Common.insertHTML('Password must contain at least one number, one lowercase, one uppercase, one special character', newPasswordError.id);
        //     return false
        // }
        else{
            return true
        }
    }

}