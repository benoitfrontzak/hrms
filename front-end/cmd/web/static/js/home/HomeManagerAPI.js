const broker = 'http://localhost:8088/'
class HomeAPI{

    // fetch connected employee's information
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
    
    // fetch connected employee's claims information (total)
    async getEmployeeClaimByID(eid) {
        const url = broker + 'route/myclaim/get/thisYear'

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
    
    // fetch connected employee's claims information (details)
    async getEmployeeClaimDetailsByID(eid) {
        const url = broker + 'route/myclaim/get/thisYearDetails'

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

    // fetch all leaves of today & tomorrow
    async getLeaveByDate() {
        const url = broker + 'route/leave/get/today'
        const response = await fetch(url)
        const result = await response.json()
        return result
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

    // fetch all approved leaves for connected user
    async getAllApprovedLeaves(eid){
        const url = broker + 'route/leave/get/manager';
        const payload = {
            id: eid
        }

        const body = {
        method: 'POST',
        body: JSON.stringify(payload),
        }
        const response = await fetch(url, body)
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
    
    // fetch all public holidays
    async getAllPublicHolidays(){
        const url = broker + 'route/publicHoliday/get/all';
        const response = await fetch(url)
        const result = await response.json();
        return result;
    }
}