const broker = 'http://localhost:8088/'

class MyPublicHolidaysAPI {

  // fetch all public holidays
  async getAllPublicHolidays(){
    const url = broker + 'route/publicHoliday/get/all';
    const response = await fetch(url)
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

  // create new Public Holiday   
  async createPH(stringifiedJSON) {
    const url = broker + 'route/publicHoliday/create',
          body = {
            method: 'POST',
            body: stringifiedJSON,
          },
          response = await fetch(url, body),
          result = await response.json()

    return result
  }

  // create new Public Holidays from CSV
  async createFromCSV(stringifiedJSON) {
    const url = broker + 'route/publicHoliday/create/csv',
          body = {
            method: 'POST',
            body: stringifiedJSON,
          },
          response = await fetch(url, body),
          result = await response.json()
          
    return result
  }

  // soft delete all Public Holidays
  async softDeletePH(phList, connectedUser, userID) {
    const url = broker + 'route/publicHoliday/softDelete'

    const payload = {
      list : phList,
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