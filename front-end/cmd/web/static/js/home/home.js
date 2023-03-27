const Helpers = new HomeHelpers(),
      API     = new HomeAPI()

// When DOM is loaded
window.addEventListener('DOMContentLoaded', () => {
    API.getEmployeeInfoByEmail(connectedEmail).then(resp => {
        console.log(resp);
    })
    
})