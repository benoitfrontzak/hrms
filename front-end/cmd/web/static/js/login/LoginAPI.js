class Login_API {
    async verifyLogin(myEmail, myPassword) {
        const url = 'http://localhost:8081/authenticate';
        // const url = 'http://localhost:8888/login';
        // const url = 'http://broker-service/login';

        // const headers = new Headers();
        // headers.append("Content-Type", "application/json");
        // headers.append("Access-Control-Allow-Origin", "http://localhost:8080");

        // const payload = {
        //   action: "auth",
        //   auth: {
        //       email: myEmail,
        //       password: myPassword,
        //   }
        // }
        const payload = {
          email: myEmail,
          password: myPassword
        }

        const body = {
          method: 'POST',
          body: JSON.stringify(payload),
          // headers: headers,
        }

        const response = await fetch(url, body)
        const result = await response.json()
        return result
    }
    
}