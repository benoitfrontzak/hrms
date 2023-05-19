const homeLink      = document.querySelector('#homeMenuLink'),
      employeeLink  = document.querySelector('#employeeMenuLink'),
      claimLink     = document.querySelector('#claimDropdownMenuLink'),
      myClaimLinks  = document.querySelectorAll('.claimMenuLink'),      // class submenu dropdown
      leaveLink     = document.querySelector('#leaveDropdownMenuLink'),
      myLeaveLinks  = document.querySelectorAll('.leaveMenuLink'),      // class submenu dropdown
      phLink        = document.querySelector('#phMenuLink'),            // public holidays
      userLink      = document.querySelector('#userDropdownMenuLink'),
      myUserLinks   = document.querySelectorAll('.userMenuLink')        // class submenu dropdown

// Set localStorage active link to home by default
if (localStorage.getItem("activeLink") === null) {
    localStorage.setItem("activeLink", "home")
}

// Fetch active link from localStorage
let activeLink = localStorage.getItem("activeLink")

// when home is clicked
homeLink.addEventListener('click', () => {
    homeLink.className = 'nav-link active border-bottom'
    localStorage.setItem("activeLink", "home")
})
// when home is clicked
employeeLink.addEventListener('click', () => {
    employeeLink.className = 'nav-link active border-bottom'
    localStorage.setItem("activeLink", "employee")
})
// when claim is clicked
claimLink.addEventListener('click', () => {
    claimLink.className = 'nav-link dropdown-toggle active border-bottom'
    localStorage.setItem("activeLink", "claim")
})
// when submenu claim is clicked
myClaimLinks.forEach(claim => {
    claim.addEventListener('click', () => {
        claimLink.className = 'nav-link dropdown-toggle active border-bottom'
        localStorage.setItem("activeLink", "claim")
    })
})
// when leave is clicked
leaveLink.addEventListener('click', () => {
    leaveLink.className = 'nav-link dropdown-toggle active border-bottom'
    localStorage.setItem("activeLink", "leave")
})
// when submenu leave is clicked
myLeaveLinks.forEach(leave => {
    leave.addEventListener('click', () => {
        leaveLink.className = 'nav-link dropdown-toggle active border-bottom'
        localStorage.setItem("activeLink", "leave")
    })
})
// when public holidays is clicked
phLink.addEventListener('click', () => {
    phLink.className = 'nav-link active border-bottom'
    localStorage.setItem("activeLink", "ph")
})
// when user is clicked
userLink.addEventListener('click', () => {
    userLink.className = 'nav-link dropdown-toggle active border-bottom'
    localStorage.setItem("activeLink", "user")
})
// when submenu user is clicked
myUserLinks.forEach(user => {
    user.addEventListener('click', () => {
        userLink.className = 'nav-link dropdown-toggle active border-bottom'
        localStorage.setItem("activeLink", "user")
    })
})

// set active link up to localStorage value
switch (activeLink) {
    case 'home':
        homeLink.className = 'nav-link active border-bottom'
        break;                
    case 'employee':
        employeeLink.className = 'nav-link active border-bottom'
        break;
    case 'claim':
        claimLink.className = 'nav-link dropdown-toggle active border-bottom'
        break;
    case 'leave':
        leaveLink.className = 'nav-link dropdown-toggle active border-bottom'
        break;
    case 'user':
        userLink.className = 'nav-link dropdown-toggle active border-bottom'
        break;
}