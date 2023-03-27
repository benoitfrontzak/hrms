const broker = 'http://localhost:8088/'
class HomeAPI{
    // fetch all employee summary
    async getEmployeeInfoByEmail(email) {
        const url = broker + 'route/employee/get/email'

        const body = {
        method: 'POST',
        body: JSON.stringify(email),
        }

        const response = await fetch(url, body)
        const result = await response.json()
        return result
    }
    
}