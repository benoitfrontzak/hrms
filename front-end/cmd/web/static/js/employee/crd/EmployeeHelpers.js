class EmployeeReadHelpers{

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

    // Insert to datatable one row per element of data
    insertRows(data, sPage){

        $('#employeeSummary').DataTable().destroy()
        const target = document.querySelector('#employeeSummaryBody')
        target.innerHTML = ''
        if (data != null){
            data.forEach(element => {
                const updateURL = sPage + element.ID
                let opt = document.createElement('tr')
                opt.id = element.ID
                opt.innerHTML = `<td><a href="${updateURL}" class="link-dark myLink">${element.Code}</a></td>
                                <td><a href="${updateURL}" class="link-dark myLink">${element.Fullname}</a></td>
                                <td><a href="${updateURL}" class="link-dark myLink">${element.Email}</a></td>
                                <td><a href="${updateURL}" class="link-dark myLink">${element.Mobile}</a></td>
                                <td><a href="${updateURL}" class="link-dark myLink">${this.birthdate(element.Birthdate)}</a></td>
                                <td><a href="${updateURL}" class="link-dark myLink">${element.Race}</a></td>
                                <td><a href="${updateURL}" class="link-dark myLink">${this.gender(element.Gender)}</a></td>
                                <td>
                                    <div class="form-check">
                                        <input class="form-check-input deleteCheckboxes"  type="checkbox" value="${element.ID}" name="softDelete">
                                        <label class="form-check-label fw-lighter fst-italic smaller" for="softDelete"><i class="bi-trash2-fill largeIcon pointer deleteEmployee"></i></label>
                                    </div>                                
                                </td>`
                target.appendChild(opt)
            })
             // trigger datatable
            $('#employeeSummary').DataTable()
            // set pointer mouse on mouseover
            $('#employeeSummary'+' tr').css('cursor', 'pointer')
        }else{
            $('#employeeSummary').DataTable()
            $('#employeeSummary').DataTable().clear().draw()
        }
    }
    // convert gender_id
    gender(id){
        let g
        return (id == 1) ? g = 'M <i class="bi-gender-male"></i>' : g = 'F <i class="bi-gender-female"></i>'
    }
    // convert birthdate (remove timestamp)
    birthdate(t){
        let b
        // get only the 10 first characters of the string
        const d = t.substring(0,10)
        // the zero value of a date is 0001-01-01
        return (d == '0001-01-01') ? b = '' : b = d
    }

    // trigger datatable and row click event
    triggerDT(){
        // trigger datatable
        const table = $('#employeeSummary').DataTable()
        // set pointer mouse on mouseover
        $('#employeeSummary'+' tr').css('cursor', 'pointer')
    }

    // returns list of selected employee id to be deleted
    selectedEmployee(myDeleteWarning){
        const checked = document.querySelectorAll('input[name=softDelete]:checked')
        if (checked.length == 0) document.querySelector('#'+myDeleteWarning).style.display = 'block'
        if (checked.length > 0){
            let allChecked = []
            checked.forEach(element => {
                allChecked.push(element.value)
            })
            return allChecked
        }        
    }


    // populate confirm detele message
    populateConfirmDelete(id, nb){
        let msg
        (nb == 1)? msg = 'Do you really want to delete this employee?' : msg = `Do you really want to delete ${nb} employees?`

        const myBody = document.querySelector('#'+id)
        myBody.innerHTML = msg
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

    getForm(formID){
        const form = document.querySelector(`#${formID}`),
              data = new FormData(form)
        const myjson        = {}

        myjson['EmployeeCode']        = data.get('employeeCode')
        myjson['Fullname']            = data.get('fullName')
        myjson['Nickname']            = data.get('nickName')
        myjson['IcNumber']            = data.get('icNumber')
        myjson['PassportNumber']      = data.get('passportNumber')
        myjson['PassportExpiryAt']    = (data.get('passportExpiryAt') == '') ? '0001-01-01' : data.get('passportExpiryAt')
        myjson['Birthdate']           = (data.get('birthdate') == '') ? '0001-01-01' : data.get('birthdate')
        myjson['Nationality']         = data.get('nationality')
        myjson['Residence']           = data.get('residence')
        myjson['Maritalstatus']       = (data.get('maritalstatus') == '0') ? '0' : '1'
        myjson['Gender']              = (data.get('gender') == '0') ? '0' : '1'
        myjson['Race']                = (data.get('race') == '') ? '0' : data.get('race')
        myjson['Religion']            = (data.get('religion') == '') ? '0' : data.get('religion')
        myjson['Streetaddr1']         = data.get('streetaddr1')
        myjson['Streetaddr2']         = data.get('streetaddr2')
        myjson['City']                = data.get('city')
        myjson['Zip']                 = data.get('zip')
        myjson['State']               = data.get('state')
        myjson['Country']             = data.get('country')
        myjson['PrimaryPhone']        = data.get('primaryPhone')
        myjson['SecondaryPhone']      = data.get('secondaryPhone')
        myjson['PrimaryEmail']        = data.get('primaryEmail')
        myjson['SecondaryEmail']      = data.get('secondaryEmail')
        myjson['IsForeigner']         = (data.get('isForeigner') == null) ? '0' : '1'
        myjson['ImmigrationNumber']   = data.get('immigrationNumber')
        myjson['IsDisabled']          = (data.get('isDisabled') == null) ? '0' : '1'
        myjson['IsActive']            = (data.get('isActive') == null) ? '0' : '1'
        myjson['Role']                = data.get('role')
        myjson['Eclaim']              = (data.get('eclaim') == null) ? '0' : '1'
        myjson['Eleave']              = (data.get('eleave') == null) ? '0' : '1'
        myjson['ConnectedUser']       = connectedEmail
        myjson['CreatedBy']           = connectedID
        myjson['UpdatedBy']           = connectedID
        
        return JSON.stringify(myjson, function replacer(key, value) { return value})
    }

}