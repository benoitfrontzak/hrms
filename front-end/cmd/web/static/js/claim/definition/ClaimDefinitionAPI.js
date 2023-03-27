const broker = 'http://localhost:8088/'

class ClaimDefinitionAPI {

  // fetch all claim definition
  async getAllClaimDefinition() {
    const url = broker + 'route/claim/definition/get/all';
    const response = await fetch(url);
    const result = await response.json();
    return result;
  }

  // fetch all config tables
  async getClaimCT() {
    const url = broker + 'route/claim/configTables/get'
    const response = await fetch(url);
    const result = await response.json();
    return result;
  }

  // create new claim definition   
  async createClaimDefinition(stringifiedJSON) {
    const url = broker + 'route/claim/definition/create'
    const body = {
      method: 'POST',
      body: stringifiedJSON,
    }
    const response = await fetch(url, body)
    const result = await response.json()
    return result
  }

  // update claim definition   
  async updateClaimDefinition(stringifiedJSON) {
    const url = broker + 'route/claim/definition/update'
    const body = {
      method: 'POST',
      body: stringifiedJSON,
    }
    const response = await fetch(url, body)
    const result = await response.json()
    return result
  }

  // soft delete all claim definition
  async softDeleteClaimDefinition(claimDefinitionList, connectedUser) {
    const url = broker + 'route/claim/definition/softDelete'

    const payload = {
      list : claimDefinitionList,
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
