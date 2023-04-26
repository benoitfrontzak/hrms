const broker = 'http://localhost:8088/'

class MyLeaveAPI {

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
  // // fetch all employees
  // async getAllEmployees() {
  //   const url = broker + 'route/employee/get/all'
  //   const response = await fetch(url);
  //   const result = await response.json();
  //   return result;
  // }

  // // fetch all my leave
  // async getAllMyLeave(connectedID, connectedEmail) {
  //   const url = broker + 'route/myleave/get/all';
  //   const payload = {
  //     id : connectedID,
  //     email : connectedEmail
  //   }

  //   const body = {
  //     method: 'POST',
  //     body: JSON.stringify(payload),
  //   }
  //   const response = await fetch(url, body)
  //   const result = await response.json();
  //   return result;
  // }

  // // fetch all leave definition
  // async getAllleaveDefinition() {
  //   const url = broker + 'route/leave/definition/get/all';
  //   const response = await fetch(url);
  //   const result = await response.json();
  //   return result;
  // }

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
  async softDeleteLeave(leaveList, connectedUser) {
    const url = broker + 'route/myleave/softDelete'

    const payload = {
      list : leaveList,
      connectedUser : connectedUser
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
