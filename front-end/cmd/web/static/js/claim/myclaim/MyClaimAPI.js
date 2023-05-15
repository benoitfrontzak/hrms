const broker = 'http://localhost:8088/'
const fileServer = 'http://localhost/'

class MyClaimAPI {
  // fetch all employee's uploaded files
  async getUploadedFiles(email){
    const url = fileServer + 'api/v1/claim/getUploadedFiles/'+email,
          response = await fetch(url),
          result = await response.json()

    return result
  }
  // fetch all my claim
  async getAllInformationsMyClaims(connectedID, connectedEmail) {
    const url = broker + 'route/myclaim/get/all';
    const payload = {
      id : connectedID,
      email : connectedEmail
    }

    const body = {
      method: 'POST',
      body: JSON.stringify(payload),
    }
    const response = await fetch(url, body)
    const result = await response.json();
    return result;
  }

  // create new claim definition   
  async createClaim(stringifiedJSON) {
    const url = broker + 'route/myclaim/create'
    const body = {
      method: 'POST',
      body: stringifiedJSON,
    }
    const response = await fetch(url, body)
    const result = await response.json()
    return result
  }

  // update claim definition   
  async updateClaim(stringifiedJSON) {
    const url = broker + 'route/myclaim/update'
    const body = {
      method: 'POST',
      body: stringifiedJSON,
    }
    const response = await fetch(url, body)
    const result = await response.json()
    return result
  }

  // soft delete all claim definition
  async softDeleteClaim(claimList, connectedUser) {
    const url = broker + 'route/myclaim/softDelete'

    const payload = {
      list : claimList,
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
