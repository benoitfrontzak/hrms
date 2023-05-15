const Common  = new MainHelpers(),
      Helpers = new HomeHelpers(),
      API     = new HomeAPI()

// initiate of google core chart is done within the header

// store all employees (active, inactive & deleted) 
let allEmployees = new Map()
allEmployees.set(0, 'not defined')

// When DOM is loaded
window.addEventListener('DOMContentLoaded', () => {
  // fetch & populate connected user information
  API.getEmployeeInfoByEmail(connectedEmail).then(resp => {
    Helpers.populateMyProfile(resp.data)
    // Helpers.populateMyProfileDates(resp.data.Employment.joinDate, resp.data.Employment.confirmDate)
  })

  // fetch & populate all employees
  API.getAllEmployees().then(resp => {
    allEmployees = Common.updateEmployeeList(resp.data, allEmployees)
    // fetch & populate current leaves (today & tomorrow)
    API.getLeaveByDate().then(resp => {
      Helpers.populateTodayAndTomorrowLeaves(resp.data)
    })
  })
  
  // fetch & populate connected user leaves (entitled)
  API.getEmployeeLeaveDetailsByID(connectedID).then(resp => {
    Helpers.populateMyLeaveDetails(resp.data)
    Helpers.populateMyLeaveDetailsProgress(resp.data)
  })
  
  // fetch & populate connected user claims (total)
  API.getEmployeeClaimByID(connectedID).then(resp => {
      Helpers.populateMyClaims(resp.data)        
      google.charts.setOnLoadCallback(Helpers.pieChartClaims(resp.data))
  })

  // fetch & populate connected user claims (details)
  API.getEmployeeClaimDetailsByID(connectedID).then(resp => {      
      Helpers.populateMyClaimDetails(resp.data)        
  })

  // scroll back to top
  var btn = $('#backToTop');
  $(window).on('scroll', function() {
      if ($(window).scrollTop() > 300) {
          btn.addClass('show');
      } else {
          btn.removeClass('show');
      }
  });
  btn.on('click', function(e) {
      e.preventDefault();
      $('html, body').animate({
          scrollTop: 0
      }, '300');
  });

  
})