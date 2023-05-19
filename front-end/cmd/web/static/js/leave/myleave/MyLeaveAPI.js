const broker = 'http://localhost:8088/',
      fileServer = 'http://localhost/'

class MyLeaveAPI {
  // fetch all employee's uploaded files
  async getUploadedFiles(email){
    const url = fileServer + 'api/v1/leave/getUploadedFiles/'+email,
          response = await fetch(url),
          result = await response.json()

    return result
  }

  // fetch all public holidays
  async getAllPublicHolidays(){
    const url = broker + 'route/publicHoliday/get/all';
    const response = await fetch(url)
    const result = await response.json();
    return result;
  }

  // fetch all required information at once
  async getAllInformationsMyLeave(connectedID, connectedEmail){
    const url = broker + 'route/myleave/get/all';
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

  // create new leave definition   
  async createLeave(stringifiedJSON) {
    const url = broker + 'route/myleave/create'
    const body = {
      method: 'POST',
      body: stringifiedJSON,
    }
    const response = await fetch(url, body)
    const result = await response.json()
    return result
  }

  // update leave definition   
  async updateLeave(stringifiedJSON) {
    const url = broker + 'route/myleave/update'
    const body = {
      method: 'POST',
      body: stringifiedJSON,
    }
    const response = await fetch(url, body)
    const result = await response.json()
    return result
  }

  // soft delete all leave definition
  async softDeleteLeave(leaveList, connectedUser, userID) {
    const url = broker + 'route/myleave/softDelete'

    const payload = {
      list : leaveList,
      connectedUser : connectedUser,
      userID : userID
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
