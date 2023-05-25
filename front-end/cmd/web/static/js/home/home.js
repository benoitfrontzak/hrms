const Common  = new MainHelpers(),
      Helpers = new HomeHelpers(),
      API     = new HomeAPI()

// initiate of google core chart is done within the header

// store all employees (active, inactive & deleted) 
let allEmployees = new Map()
allEmployees.set(0, 'not defined')

// When DOM is loaded
window.addEventListener('DOMContentLoaded', () => {
  const homeLink = document.querySelector('#homeMenuLink')
  homeLink.className = 'nav-link active border-bottom'
  localStorage.setItem("activeLink", "home")
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
    // fetch Public Holidays
    API.getAllPublicHolidays().then(resp => {
      const holidays = resp.data
      // if user is not admin
      if (connectedRole != 2){
        API.getAllApprovedLeaves(connectedID).then(resp => {
          let leaves = []
          if (resp.data != null) {
            if (resp.data.Approved != null) {
              resp.data.Approved.forEach(leave => {
                leave.details.forEach(element => {
                  oneLeave = {
                    date: Helpers.formatDate(element.requestedDate),
                    name: allEmployees.get(Number(leave.employeeid)),
                    description: `${leave.leaveDefinitionName} ${Helpers.myBooleanIsHalf(element.isHalf)}`
                  }
                  leaves.push(oneLeave)
                });
              });
            }
          }          
  
          $('#calendar').fullCalendar({
            header: {
              left: 'prev',
              center: 'title',
              right: 'next'
            },
            defaultView: 'month',
            // Limit the number of events per day
            eventLimit: true,
            eventLimitClick: 'popover',
            views: {
              month: {
                eventLimit: 5 // Set the maximum number of events to show per day
              }
            },
            events: [
              ...holidays.map(function(holiday) {
                return {
                  title: holiday.name,
                  start: Helpers.formatDate(holiday.date),
                  className: 'myHarmonyTurquoise', // Holiday color
                  extendedProps: {
                    description: holiday.description,
                    createdAt: holiday.createdAt,
                    createdBy: holiday.createdBy
                  }
                };
              }),
              ...leaves.map(function(event) {
                return {
                  title: event.name,
                  start: event.date,
                  color: event.color,
                  extendedProps: {
                    description: event.description,
                    createdAt: event.createdAt,
                    createdBy: event.createdBy
                  }
                };
              })
            ],
            eventRender: function(event, element) {
              const description = event.extendedProps.description;
        
              element.on('mouseenter', function() {
                const tooltipContent = `<div class="tooltip-title">${event.title}</div>
                                        <div class="tooltip-description">${description}</div>`;
        
                $(this).tooltip({
                  title: tooltipContent,
                  html: true,
                  container: 'body',
                  placement: 'top',
                  trigger: 'manual'
                });
        
                $(this).tooltip('show');
              });
        
              element.on('mouseleave', function() {
                $(this).tooltip('hide');
              });
              if (event.className.includes('myHarmonyTurquoise')) {
                const divElement = element.find('.fc-content'),
                      spanElement = element.find('.fc-title');                    
                element.attr('style', 'background-color: #35DCCC !important;border:none !important;') // Custom background color for the span element
                divElement.attr('style', 'background-color: #35DCCC !important;border:none !important;') // Custom background color for the span element
                spanElement.attr('style', 'background-color: #35DCCC !important;border:none !important;'); // Custom background color for the span element
              }
            }
          });
   
        })
      }else{
        API.getAllApprovedLeaves().then(resp => {
          let leaves = []
          if (resp.data != null) {
            if (resp.data.Approved != null) {
              resp.data.Approved.forEach(leave => {
                leave.details.forEach(element => {
                  oneLeave = {
                    date: Helpers.formatDate(element.requestedDate),
                    name: allEmployees.get(Number(leave.employeeid)),
                    description: `${leave.leaveDefinitionName} ${Helpers.myBooleanIsHalf(element.isHalf)}`
                  }
                  leaves.push(oneLeave)
                });
              });

            }
          }
  
          $('#calendar').fullCalendar({
            header: {
              left: 'prev',
              center: 'title',
              right: 'next'
            },
            defaultView: 'month',
            // Limit the number of events per day
            eventLimit: true,
            eventLimitClick: 'popover',
            views: {
              month: {
                eventLimit: 5 // Set the maximum number of events to show per day
              }
            },
            events: [
              ...holidays.map(function(holiday) {
                return {
                  title: holiday.name,
                  start: Helpers.formatDate(holiday.date),
                  className: 'myHarmonyTurquoise', // Holiday color
                  extendedProps: {
                    description: holiday.description,
                    createdAt: holiday.createdAt,
                    createdBy: holiday.createdBy
                  }
                };
              }),
              ...leaves.map(function(event) {
                return {
                  title: event.name,
                  start: event.date,
                  color: event.color,
                  extendedProps: {
                    description: event.description,
                    createdAt: event.createdAt,
                    createdBy: event.createdBy
                  }
                };
              })
            ],
            eventRender: function(event, element) {
              const description = event.extendedProps.description;
        
              element.on('mouseenter', function() {
                const tooltipContent = `<div class="tooltip-title">${event.title}</div>
                                        <div class="tooltip-description">${description}</div>`;
        
                $(this).tooltip({
                  title: tooltipContent,
                  html: true,
                  container: 'body',
                  placement: 'top',
                  trigger: 'manual'
                });
        
                $(this).tooltip('show');
              });
        
              element.on('mouseleave', function() {
                $(this).tooltip('hide');
              });
              if (event.className.includes('myHarmonyTurquoise')) {
                const divElement = element.find('.fc-content'),
                      spanElement = element.find('.fc-title');                    
                element.attr('style', 'background-color: #35DCCC !important;border:none !important;') // Custom background color for the span element
                divElement.attr('style', 'background-color: #35DCCC !important;border:none !important;') // Custom background color for the span element
                spanElement.attr('style', 'background-color: #35DCCC !important;border:none !important;'); // Custom background color for the span element
              }
            }
          });

        })
      }      
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