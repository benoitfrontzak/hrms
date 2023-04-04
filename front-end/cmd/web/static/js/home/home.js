const Common  = new MainHelpers(),
      Helpers = new HomeHelpers(),
      API     = new HomeAPI()

// initiate of google core chart is done within the header

// store all employees (active, inactive & deleted) 
let allEmployees = new Map()
allEmployees.set(0, 'not defined')

// When DOM is loaded
window.addEventListener('DOMContentLoaded', () => {
  
  // fetch & populate all employees
  API.getAllEmployees().then(resp => {
    console.log('get all employee');
    console.log(resp);
    allEmployees = Common.updateEmployeeList(resp.data, allEmployees)
    // fetch & populate current leaves (today & tomorrow)
    API.getLeaveByDate().then(resp => {
      console.log('leave by date');
      console.log(resp);
      Helpers.populateTodayAndTomorrowLeaves(resp.data)
    })
  })

  // fetch & populate connected user information
  API.getEmployeeInfoByEmail(connectedEmail).then(resp => {
      console.log('employee by email');
      console.log(resp);
      Helpers.populateMyProfile(resp.data)
      Helpers.populateMyProfileDates(resp.data.Employment.joinDate, resp.data.Employment.confirmDate)
  })
  
  // fetch & populate connected user claims (total)
  API.getEmployeeClaimByID(connectedID).then(resp => {
      console.log('claim by connected user id');
      console.log(resp);
      Helpers.populateMyClaims(resp.data)        
      google.charts.setOnLoadCallback(Helpers.pieChartClaims(resp.data))
  })

  // fetch & populate connected user claims (details)
  API.getEmployeeClaimDetailsByID(connectedID).then(resp => {
      console.log('claim details by connected user id');
      console.log(resp);
       
      Helpers.populateMyClaimDetails(resp.data)        
  })

  // fetch & populate connected user leaves (entitled)
  API.getEmployeeLeaveDetailsByID(connectedID).then(resp => {
      console.log('leave details');
      console.log(resp);
      Helpers.populateMyLeaveDetails(resp.data)        
      google.charts.setOnLoadCallback(Helpers.chartLeaves(resp.data))
  })


})