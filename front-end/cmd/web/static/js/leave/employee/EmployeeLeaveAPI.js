const broker = 'http://localhost:8088/',
      fileServer = 'http://localhost/'

class EmployeeLeaveAPI {
  // fetch all employee's uploaded files
  async getUploadedFiles(email){
    const url = fileServer + 'api/v1/leave/getUploadedFiles/'+email,
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

  // fetch connected employee's leave details
  async getEmployeeLeaveDetailsByID(eid) {
    const url = broker + 'route/myleave/get/today'

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
  
  // fetch all required information at once
  async getEmployeeLeaves(eid, email){
    const url = broker + 'route/myleave/get/all';
    const payload = {
      id : eid,
      email : email
    }

    const body = {
      method: 'POST',
      body: JSON.stringify(payload),
    }
    const response = await fetch(url, body)
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

  // fetch connected employee's leave details
  async updateEmployeeLeaves(stringifyData) {
    const url = broker + 'route/leave/update/credits'

    const body = {
    method: 'POST',
    body: stringifyData,
    }

    const response = await fetch(url, body)
    const result = await response.json()
    return result
  }
}
