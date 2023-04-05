// EmployeeRead_API handle all API call to be sent to the broker & the file server
const broker = 'http://localhost:8088/',
      fileServer = 'http://localhost:8484/'

class EmployeeUpdateAPI{

  // fetch all employee's uploaded files
  async getUploadedFiles(email){
    const url = fileServer + 'api/v1/employee/getUploadedFiles/'+email
    const response = await fetch(url)
    const result = await response.json()
    return result
  }

  // fetch all employees
  async getAllEmployees() {
    const url = broker + 'route/employee/get/all'
    const response = await fetch(url);
    const result = await response.json();
    return result;
  }
  
  // update employee   
  async updateEmployee(stringifiedJSON) {
    const url = broker + 'route/employee/update'
    const body = {
      method: 'POST',
      body: stringifiedJSON,
    }
    const response = await fetch(url, body)
    const result = await response.json()
    return result
  }

  // fetch all employee list
  async getActiveEmployees() {
    const url = broker + 'route/employee/get/active'
    const response = await fetch(url);
    const result = await response.json();
    return result;
  }

  // fetch all employee's config tables
  async getEmployeeCT() {
    const url = broker + 'route/employee/configTables/get'
    const response = await fetch(url);
    const result = await response.json();
    return result;
  }

  // fetch all employee information
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