class EmployeeAddHelpers{
    // hide | show div by ID
    hideDivByID(id){ document.querySelector('#'+id).style.display = 'none' }
    showDivByID(id){ document.querySelector('#'+id).style.display = 'block' }

    // hide card by ID
    hideCardByID(bodyID, iconHide, iconShow){
        this.hideDivByID(bodyID)    // hide card's body
        this.hideDivByID(iconHide)  // remove icon 'hide'
        this.showDivByID(iconShow)  // display icon 'show'
    }
    // hide card by ID
    showCardByID(bodyID, iconHide, iconShow){
        this.showDivByID(bodyID)    // display card's body
        this.showDivByID(iconHide)  // display icon 'hide'
        this.hideDivByID(iconShow)  // remove icon 'show'
    }

    // show (or hide) card by ID
    // the icons show & hide must follow the naming convention: id = cardnameShow & id = cardnameHide
    // the card's body must follow the naming convention : id = cardnameBodyCard
    displayCardByID(card, display){
        const showIcon = card + 'Show',
              hideIcon = card + 'Hide',
              cardBody = card + 'BodyCard';
        (display === true) ? this.showCardByID(cardBody, hideIcon, showIcon) : this.hideCardByID(cardBody, hideIcon, showIcon);
    }

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

    // valid personal info
    validatePersonalInfo(){
        const myIC = document.querySelector('#icNumber'),
              myPassport = document.querySelector('#passportNumber'),
              myPassportExpiry = document.querySelector('#passportExpiryAt')
        
        if (myIC.value == '') {
            console.log('inside condition');
            document.querySelector('#icNumberError').innerHTML = 'This field is required'
        }
        if (myPassport.value == '') {
            document.querySelector('#passportNumberError').innerHTML = 'This field is required'
        }
        if (myPassportExpiry.value == '') {
            document.querySelector('#passportExpiryAtError').innerHTML = 'This field is required'
        }
    }
}