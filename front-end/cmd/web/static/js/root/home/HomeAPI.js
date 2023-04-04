const broker = 'http://localhost:8088/'

class HomeAPI {

  // fetch all employees
  async getAllEmployees() {
    const url = broker + 'route/employee/get/all'
    const response = await fetch(url);
    const result = await response.json();
    return result;
  }
  
   // approve claims   
   async approveClaims(claimList, amount, email, id) {
    const url = broker + 'route/claim/approve'

    const payload = {
      list : claimList,
      amount : amount,
      connectedUser : email,
      userID : id
    }
    const body = {
      method: 'POST',
      body: JSON.stringify(payload),
    }
    const response = await fetch(url, body)
    const result = await response.json()
    return result
  }

}
