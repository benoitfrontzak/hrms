const broker = 'http://localhost:8088/'
class PwdAPI{

    // Send new password to auth service
    async upPwd(connectedID,data) {
        const url = broker + 'route/authentication/update/pwd/eid'
        
        const payload = {
            id: parseInt(connectedID),
            oldPassword: data.oldPassword,
            newPassword: data.newPassword,
        }

        const body = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(payload),
        }

        const response = await fetch(url, body)
        const result = await response.json()
        return result
    }
}