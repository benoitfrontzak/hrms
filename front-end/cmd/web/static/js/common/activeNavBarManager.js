const homeLink          = document.querySelector('#homeMenuLink'),
      leaveLink         = document.querySelector('#leaveDropdownMenuLink'),
      myLeaveLinks      = document.querySelectorAll('.leaveMenuLink'),      // class submenu dropdown
      userLink          = document.querySelector('#userDropdownMenuLink'),
      myUserLinks       = document.querySelectorAll('.userMenuLink')        // class submenu dropdown

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
        homeLink.className = 'nav-link active border-bottom border-danger'
        break;
    case 'leave':
        leaveLink.className = 'nav-link dropdown-toggle active border-bottom border-danger'
        break;
    case 'user':
        userLink.className = 'nav-link dropdown-toggle active border-bottom border-danger'
        break;
}