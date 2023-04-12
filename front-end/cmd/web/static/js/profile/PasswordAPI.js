const broker = 'http://localhost:8088/'
class PasswordAPI{

    // Send new password to auth service
    async updateMyPassword(data) {
        const url = broker + 'route/authentication/update/password'

        const body = {
        method: 'POST',
        body: data,
        }

        const response = await fetch(url, body)
        const result = await response.json()
        return result
    }
}