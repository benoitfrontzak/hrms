// EmployeeRead_API handle all API call to be sent to the broker
const broker = 'http://localhost:8088/'

class EmployeeConfigTablesAPI{

  // fetch all employee's config tables
  async getEmployeeCT() {
    const url = broker + 'route/employee/configTables/get'
    const response = await fetch(url);
    const result = await response.json();
    return result;
  }

  // add new config table entry
  async addNewPI(stringified) {
    const url = broker + 'route/payrollItem/create'

    const body = {
      method: 'POST',
      body: stringified,
    }

    const response = await fetch(url, body)
    const result = await response.json()
    return result
  }

   // update config table entry
   async updateEntryPI(stringified) {
    const url = broker + 'route/payrollItem/update'

    const body = {
      method: 'POST',
      body: stringified,
    }

    const response = await fetch(url, body)
    const result = await response.json()
    return result
  }

  // add new config table entry
  async addNewConfigTableEntry(ctName, ctValue, email) {
    const url = broker + 'route/employee/configTable/create'

    const payload = {
      name : ctName,
      value : ctValue,
      email : email
    }

    const body = {
      method: 'POST',
      body: JSON.stringify(payload),
    }

    const response = await fetch(url, body)
    const result = await response.json()
    return result
  }

  // update config table entry
  async updateConfigTableEntry(ctName, ctValue, ctID, email) {
    const url = broker + 'route/employee/configTable/update'

    const payload = {
      name : ctName,
      value : ctValue,
      rowid : ctID,
      email : email
    }

    const body = {
      method: 'POST',
      body: JSON.stringify(payload),
    }

    const response = await fetch(url, body)
    const result = await response.json()
    return result
  }
  
  // delete config table entry
  async deleteConfigTableEntry(list, ctName, email) {
    const url = broker + 'route/employee/configTable/softDelete'

    const payload = {
      list : list,
      name : ctName,
      email : email
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