class MainHelpers{
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

    // insert HTML to divID
    insertHTML(message, divID){ document.querySelector('#'+divID).innerHTML = `${message}` }

    // insert value to input
    insertInputValue(value, inputID){document.querySelector('#'+inputID).value = `${value}`}
    
    // check radio button
    checkRadio(value, name){
        const myRadio = document.querySelectorAll(`input[name="${name}"]`)
        myRadio.forEach(element => {
            if (element.value == value) element.checked = true            
        })
    }

    // check checkbox
    checkCheckbox(value, checkboxID){
        const myCheckbox = document.querySelector(`#${checkboxID}`)
        if (value == 1) myCheckbox.checked = true
    }

    // select option selectBox
    selectBoxOptionSelected(value, selectBoxID){
        const mySelectBox = document.querySelector(`#${selectBoxID}`)
        for (let opt of mySelectBox.options) {
            if (opt.value == '') opt.remove()
            if (opt.value == value) {
                opt.selected = true
            }
        }
    }

    // replace camelCase to camel case
    replaceCamelCase(str){
        let result
        (Boolean(str.match(/[A-Z]/))) ? result = str.replace(/([a-z])([A-Z])/g, '$1 $2').toLowerCase() : result = str
        return result
    }

    // generate warning message
    populateWarningMessage(notValid){
        let msg = `Please fill-up all those fields <ul>`

        notValid.forEach(field => {
            msg += `<li>${this.replaceCamelCase(field)}</li>`
        })

        msg += `</ul>`

        return msg
    }

    // display main warning message
    displayWarningMessage(notValid){
        const myWarningDivID  = 'warningMessageDiv',
              myWarningBodyID = 'warningMessageBody'

        if (!notValid?.length) {
            this.hideDivByID(myWarningDivID)
        }else{
            this.showDivByID(myWarningDivID)
            this.insertHTML(this.populateWarningMessage(notValid), myWarningBodyID)
        }
      
    }
    
    // validate given required fields array
    validateRequiredFields(myRequired){
        // store not valid fields
        let notValid = []

        // loop over required fields and check wether it's empty or not
        myRequired.forEach(field => {
            let myField = document.querySelector(`#${field}`)
            if (myField.value == '') {
                this.insertHTML('This field is required', `${field}Error`)
                notValid.push(field)
            }else{
                this.insertHTML('', `${field}Error`)
            }
        })

        // display main warning message
        this.displayWarningMessage(notValid)

        return notValid.length
    }

    // construct a set of key/value pairs representing form fields and their values
    // returns JSON form data stringify
    getForm(formID, userID){
        const form = document.querySelector(`#${formID}`),
              data = new FormData(form)

        data.append('createdAt', userID)
        data.append('updatedAt', userID)
        
        const formJSON = JSON.stringify(Object.fromEntries(data))

        return formJSON
    }

    // clear form: reset form value & clear error message
    clearForm(formID, myRequired){
        // reset form fields
        document.querySelector("#"+formID).reset()

        // loop over required fields and remove error message
        myRequired.forEach(field => {
            this.insertHTML('', `${field}Error`)
        })

        let notValid = []
        this.displayWarningMessage(notValid)
    }

    // updateEmployeeList
    updateEmployeeList(data, allEmployees){
        // update allEmployee active
        if (data.Active !== null){
            data.Active.forEach(element => {
                allEmployees.set(element.ID, element.Fullname)
            })
        } 
        // update allEmployee inactive
        if (data.Inactive !== null) {
            data.Inactive.forEach(element => {
                allEmployees.set(element.ID, element.Fullname)
            })
        } 
        // update allEmployee deleted
        if (data.Deleted !== null) {
            data.Deleted.forEach(element => {
                allEmployees.set(element.ID, element.Fullname)
            })
        } 

        return allEmployees
    }
}