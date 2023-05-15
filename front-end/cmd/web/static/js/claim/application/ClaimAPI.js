const broker = 'http://localhost:8088/'
const fileServer = 'http://localhost/'

class ClaimAPI {
  // fetch all employee's uploaded files
  async getUploadedFiles(email){
    const url = fileServer + 'api/v1/claim/getUploadedFiles/'+email,
          response = await fetch(url),
          result = await response.json()

    return result
  }

  // fetch all employees
  async getAllEmployees() {
    const url = broker + 'route/employee/get/all'
    const response = await fetch(url);
    const result = await response.json();
    return result;
  }
  
  // fetch all claim 
  async getAllClaims() {
    const url = broker + 'route/claim/get/all';
    const response = await fetch(url)
    const result = await response.json();
    return result;
  }

  // fetch employee by id
  async getEmployeeEmailByID(id) {
    const url = broker + 'route/employee/find/email'
    
    const payload = {
      id : id
    }
    const body = {
      method: 'POST',
      body: JSON.stringify(payload),
    }

    const response = await fetch(url, body)
    const result = await response.json()
    return result
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

  // reject claims   
  async rejectClaims(claimList, reason, email, id) {
    const url = broker + 'route/claim/reject'

    const payload = {
      list : claimList,
      reason : reason,
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
