class EmployeeRead_Helpers{

    // Calculate approximated age
    calculateAge(birthdate){
        const today     = new Date(),
              thisYear  = today.getFullYear(),
              birthYear = birthdate.split('-')[0],
              age       = thisYear - birthYear
        return age
    }
    
    // Populate age
    populateAge(years){
        const age = document.querySelector('#age')
        age.value = years
    }

}