const broker = 'http://localhost:8088/'

class EmployeeLeaveAPI {

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
