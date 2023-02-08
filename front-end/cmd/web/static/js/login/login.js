const API = new Login_API()

function setCookie(name,value,days) {
    var expires = "";
    if (days) {
        var date = new Date();
        date.setTime(date.getTime() + (days*24*60*60*1000));
        expires = "; expires=" + date.toUTCString();
    }
    document.cookie = name + "=" + (value || "")  + expires + "; path=/";
}

window.addEventListener('DOMContentLoaded', (event) => {

    const myBtn = document.querySelector('#submit_button'),
          myEmail = document.querySelector('#email_input'),
          myPassword = document.querySelector('#password_input')
    
    //When submit button is clicked 
    myBtn.addEventListener('click', () => {

        API.verifyLogin(myEmail.value, myPassword.value).then(resp => {
            console.log('resp')
            console.log(resp)

            if(!resp.error){
                console.log('Successfully logged')
                // setCookie('jwt',resp.token,1)
                // const headers = new Headers();
                // headers.append("Authorization", "Bearer: " + resp.token);
                // window.location.href = "/home"
                // window.location.replace("/home")
            }else{
                console.log('Something went wrong')
                // window.location.replace("http://localhost:8080")
            }
        })
  
    })

})