// EmployeeRead_API handle all API call to be sent to the broker
const broker = 'http://localhost:8088/',
      fileServer = 'http://localhost/'

class EmployeeReadAPI{
  // fetch all employee's uploaded files
  async getUploadedFiles(email){
    const url = fileServer + 'api/v1/employee/getUploadedFiles/'+email
    const response = await fetch(url)
    const result = await response.json()
    return result
  }

  // create new employee   
  async createEmployee(stringifiedJSON) {
    const url = broker + 'route/employee/create'
    const body = {
      method: 'POST',
      body: stringifiedJSON,
    }
    const response = await fetch(url, body)
    const result = await response.json()
    return result
  }

  // soft delete all selected employee
  async softDeleteEmployeeByID(employeeList, connectedUser) {
    const url = broker + 'route/employee/softDelete'

    const payload = {
      list : employeeList,
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

  // fetch all employee's config tables
  async getEmployeeCT() {
    const url = broker + 'route/employee/configTables/get'
    const response = await fetch(url);
    const result = await response.json();
    return result;
  }

  // fetch all employees
  async getAllEmployees() {
    const url = broker + 'route/employee/get/all'
    const response = await fetch(url);
    const result = await response.json();
    return result;
  }
  
  // fetch all employee information by id
  async getEmployeeInfo(id) {
    const url = broker + 'route/employee/get/id'

    const body = {
      method: 'POST',
      body: JSON.stringify(id),
    }

    const response = await fetch(url, body)
    const result = await response.json()
    return result
  }

}