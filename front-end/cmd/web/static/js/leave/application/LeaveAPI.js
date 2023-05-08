const broker = 'http://localhost:8088/',
      fileServer = 'http://localhost/'

class LeaveAPI {

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

  // fetch all leaves 
  async getAllLeaves() {
    const url = broker + 'route/leave/get/all';
    const response = await fetch(url)
    const result = await response.json();
    return result;
  }

  // fetch all config tables
  async getLeaveCT() {
    const url = broker + 'route/leave/configTables/get'
    const response = await fetch(url);
    const result = await response.json();
    return result;
  }

   // approve leaves   
   async approveLeaves(leaveList, email, id) {
    const url = broker + 'route/leave/approve'

    const payload = {
      list : leaveList,
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

  // reject leaves   
  async rejectLeaves(leaveList, reason, email, id) {
    const url = broker + 'route/leave/reject'

    const payload = {
      list : leaveList,
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
