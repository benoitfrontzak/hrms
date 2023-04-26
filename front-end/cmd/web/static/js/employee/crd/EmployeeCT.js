class EmployeeCT {
     // Populate config tables for nationality, residence country, race, religion, country, relationship
     populateConfigTables(ct){
        this.insertNationalityOptions('nationality', ct.Country)
        this.insertOptions('residence', ct.Country)
        this.insertOptions('race', ct.Race)
        this.insertOptions('religion', ct.Religion)
        this.insertOptions('country', ct.Country)

        // set Malaysia as default value
        $('#residence option[value=135]').attr('selected','selected');
        $('#country option[value=135]').attr('selected','selected');
    }
    
    // Insert to selectID one option per element of data
    insertNationalityOptions(id, data){
        const target = document.querySelector('#'+id)

        target.innerHTML = '<option selected hidden value="0"></option>'

        data.forEach(element => {
            let opt = document.createElement('option')
            opt.value = element.ID
            opt.innerHTML = element.Nationality
            target.appendChild(opt)
        })

    }

    // Insert to selectID one option per element of data
    insertOptions(id, data){
        const target = document.querySelector('#'+id)

        target.innerHTML = '<option selected hidden value=""></option>'

        data.forEach(element => {
            let opt = document.createElement('option')
            opt.value = element.ID
            opt.innerHTML = element.Name
            target.appendChild(opt)
        })

    }
    
}