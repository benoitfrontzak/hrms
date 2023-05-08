const broker = 'http://localhost:8088/'

class LeaveDefinitionAPI {
  // fetch all employees
  async getAllEmployees() {
    const url = broker + 'route/employee/get/all'
    const response = await fetch(url);
    const result = await response.json();
    return result;
  }
  
  // fetch all leave definition
  async getAllLeaveDefinition() {
    const url = broker + 'route/leave/definition/get/all';
    const response = await fetch(url);
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

  // create new claim definition   
  async createLeaveDefinition(stringifiedJSON) {
    const url = broker + 'route/leave/definition/create'
    const body = {
      method: 'POST',
      body: stringifiedJSON,
    }
    const response = await fetch(url, body)
    const result = await response.json()
    return result
  }

  // update leave definition   
  async updateLeaveDefinition(stringifiedJSON) {
    const url = broker + 'route/leave/definition/update'
    const body = {
      method: 'POST',
      body: stringifiedJSON,
    }
    const response = await fetch(url, body)
    const result = await response.json()
    return result
  }

  // soft delete all leave definition
  async softDeleteLeaveDefinition(leaveDefinitionList, connectedUser) {
    const url = broker + 'route/leave/definition/softDelete'

    const payload = {
      list : leaveDefinitionList,
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
