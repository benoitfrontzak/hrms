// EmployeeRead_API handle all API call to be sent to the broker & the file server
const broker = 'http://localhost:8088/',
      fileServer = 'http://localhost/'

class EmployeeUpdateAPI{

  // fetch all employee's uploaded files
  async getUploadedFiles(email){
    const url = fileServer + 'api/v1/employee/getUploadedFiles/'+email,
          response = await fetch(url),
          result = await response.json()

    return result
  }

  // fetch all requires information for employee update
  async getAllRequiresInfo(employeeID, connectedID, connectedEmail) {
    const url = broker + 'route/employee/get/id'

    const payload = {
      id : employeeID,
      connectedUser : connectedEmail,
      userID : connectedID
    }

    const body = {
      method: 'POST',
      body: JSON.stringify(payload),
    }

    const response = await fetch(url, body)
    const result = await response.json()

    return result
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
  
  // add payroll item   
  async addPayrollItem(stringifiedJSON) {
    const url = broker + 'route/employee/payrollItem/create'

    const body = {
      method: 'POST',
      body: stringifiedJSON,
    }

    const response = await fetch(url, body)
    const result = await response.json()

    return result
  }

  // soft delete all selected employee
  async softDeletePayrollItem(piList, connectedUser) {
    const url = broker + 'route/employee/payrollItem/softDelete'

    const payload = {
      list : piList,
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