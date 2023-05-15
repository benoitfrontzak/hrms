const broker     = 'http://localhost:8088/',
      fileServer = 'http://localhost/'

class EmployeeClaimAPI {
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

  // fetch enmployee email by employee id
  async findEmployeeEmail(eid) {
    const url = broker + 'route/employee/find/email'

    const payload = {
        id: eid
    }

    const body = {
    method: 'POST',
    body: JSON.stringify(payload),
    }

    const response = await fetch(url, body)
    const result = await response.json()
    return result
  }

  // fetch connected employee's claims information (total)
  async getEmployeeClaimByID(eid) {
    const url = broker + 'route/myclaim/get/thisYear'

    const payload = {
        id: eid
    }

    const body = {
    method: 'POST',
    body: JSON.stringify(payload),
    }

    const response = await fetch(url, body)
    const result = await response.json()
    return result
  }

  // fetch all my claim
  async getEmployeeClaims(connectedID, connectedEmail) {
    const url = broker + 'route/myclaim/get/all';
    const payload = {
      id : connectedID,
      email : connectedEmail
    }

    const body = {
      method: 'POST',
      body: JSON.stringify(payload),
    }
    const response = await fetch(url, body)
    const result = await response.json();
    return result;
  }
}
