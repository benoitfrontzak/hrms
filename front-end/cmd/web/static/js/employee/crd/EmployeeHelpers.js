class EmployeeHelpers{ 

    // cleaned checked checkboxes when modal is close
    clearSelectedEmployee(){
        document.getElementById('confirmDelete').addEventListener('hidden.bs.modal', function () {
            const myDeleteCheckboxes = document.querySelectorAll('.deleteCheckboxes')
    
            myDeleteCheckboxes.forEach(element => {
                element.checked = false
            })
    
        })
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

    // create data source event listener
    dataSourceListener(active, inactive, deleted){
        const dSource = ['active', 'inactive', 'deleted']
        dSource.forEach(element => {
            // create capitalize function
            const capitalize = element => (element && element[0].toUpperCase() + element.slice(1)) || ""
 
            document.querySelector('#'+element+'Btn').addEventListener('click', () => {
                $('#employeeSummary').DataTable().destroy()
                let ds
                switch (element) {
                    case 'active':
                        ds = active
                        break;
                    case 'inactive':
                        ds = inactive
                        break;
                    case 'deleted':
                        ds = deleted
                        break;                    
                } 
                Employee.insertRows(ds)
                document.querySelector('#employeeTitle').innerHTML = capitalize(element) + ' Employees'
            })
        });
    }
}