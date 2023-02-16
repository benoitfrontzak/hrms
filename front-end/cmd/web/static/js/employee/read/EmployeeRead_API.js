class EmployeeRead_API{
    async createEmployee(stringifiedJSON) {
        const url = 'http://localhost:8088/createEmployee';

        const body = {
          method: 'POST',
          body: stringifiedJSON,
        }

        const response = await fetch(url, body)
        const result = await response.json()
        return result
    }
}