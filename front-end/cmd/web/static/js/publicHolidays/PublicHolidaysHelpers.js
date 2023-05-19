class MyPublicHolidaysHelpers {
    // convert timestamp (remove timestamp)
    formatDate(t){
        let b
        // get only the 10 first characters of the string
        const d = t.substring(0,10)
        // the zero value of a date is 0001-01-01
        return (d == '0001-01-01') ? b = '' : b = d
    }

    // insert to datatable one row per element of data
    insertRows(data) {
        const target = document.querySelector('#publicHolidayBody')
        target.innerHTML = ''
        
        const myColumns = [
            {title: "Date"},
            {title: "Name"},
            {title: "Description"},
            {title: "Created At"},
            {title: "Created By"},
            {title: "Updated At"},
            {title: "Updated By"},
            {title: "Action"}
        ]
        
        if (data == null){
            DT.initiateMyTable('publicHolidayTable', myColumns)
            $('#publicHolidayTable').DataTable().clear().draw()
            return
        }

        data.forEach(element => { 
            let row = document.createElement('tr')
            row.id = 'myPH' + element.id
            row.innerHTML = `<td class="row-data pointer">${this.formatDate(element.date)}</td>
                             <td class="row-data pointer">${element.name}</td>
                             <td class="row-data pointer">${element.description}</td>
                             <td class="row-data pointer">${this.formatDate(element.createdAt)}</td>
                             <td class="row-data pointer">${allEmployees.get(Number(element.createdBy))}</td>
                             <td class="row-data pointer">${this.formatDate(element.updatedAt)}</td>
                             <td class="row-data pointer">${allEmployees.get(Number(element.updatedBy))}</td>
                             <td>
                                <div class="row">
                                    <div class="col text-end">
                                        <div class="form-check">
                                            <label class="form-check-label fw-lighter fst-italic smaller" for="softDelete"><i class="bi-trash2-fill largeIcon pointer deletePH"></i></label>
                                            <input class="form-check-input deleteCheckboxes"  type="checkbox" value="${element.id}" name="softDelete">                                    
                                        </div> 
                                    </div>
                                </div>                                                 
                             </td>`
            
                             
            target.appendChild(row)
        })

        DT.initiateMyTable('publicHolidayTable', myColumns)
        this.createRowClickEvent()
    }

    // make datatable rows clickable but the last column 
    // open employee to new tab for edit 
    createRowClickEvent(){
        $('#publicHolidayTable tbody').on('click', 'td:not(:last-child)', function() {
          const tr = $(this).closest('tr'),
          rowID = tr[0].id
                console.log(rowID);
        }) 
      }

    // returns list of selected leave id to be deleted
    selectedPH(){
        const checked = document.querySelectorAll('input[name=softDelete]:checked')
        if (checked.length == 0) document.querySelector('#deleteWarningMessageDiv').style.display = 'block'
        if (checked.length > 0){
            let allChecked = []
            checked.forEach(element => {
                allChecked.push(element.value)
            })
            
            return allChecked
        }        
    }

    // populate confirm detele message
    populateConfirmDelete(nb){
        let msg
        (nb == 1)? msg = 'Do you really want to delete this Public Holiday?' : msg = `Do you really want to delete ${nb} Public Holidays?`

        const myBody = document.querySelector('#confirmDeleteBody')
        myBody.innerHTML = msg
    }

    // clean form when open
    cleanForm(){
        // clean/hide errors
        Common.hideDivByID('warningMessageDiv')
        Common.insertHTML('', 'datePublicHolidayError')
        Common.insertHTML('', 'namePublicHolidayError')
        // reset input values
        Common.insertInputValue('','datePublicHoliday')
        Common.insertInputValue('','namePublicHoliday')
        Common.insertInputValue('','description')
    }

    // clean form from CSV when open
    cleanFormCSV(){
        Common.insertHTML('', 'csvFileError')
        Common.insertInputValue('','csvFile')
    }

    // sort array
    sortArray(arr){
        arr.sort(function(a,b) {
            a = a.split('/').reverse().join('')
            b = b.split('/').reverse().join('')
            return a.localeCompare(b)
            // return a > b ? 1 : a < b ? -1 : 0; // <-- alternative   
        })
        return arr
    }

    // populate all public holidays array (and sort it)
    populatePublicHolidays(data, ph){
        // get ph dates
       data.forEach(element => {
           ph.push(this.formatDate(element.date))
       });

       // sort dates
       return this.sortArray(ph)
    }

    // generate warning message
    populateWarningMessage(notValid, myMsg){
        let msg = myMsg + ` <ul>`

        notValid.forEach(field => {
            msg += `<li>${Common.replaceCamelCase(field)}</li>`
        })

        msg += `</ul>`

        return msg
    }

    // display main warning message
    displayWarningMessage(notValid, myMsg){
        const myWarningDivID  = 'warningMessageDiv',
              myWarningBodyID = 'warningMessageBody'

        if (!notValid?.length) {
            Common.hideDivByID(myWarningDivID)
        }else{
            Common.showDivByID(myWarningDivID)
            Common.insertHTML(this.populateWarningMessage(notValid, myMsg), myWarningBodyID)
        }
      
    }
    
    // validate form (required fields & uniq date)
    validateForm(myRequired){
        // store not valid fields
        let notValid = []

        // check if new public holiday doesn't already exists
        const myPH = document.querySelector('#datePublicHoliday')
        if (ph.includes(myPH.value)){
            Common.insertHTML('This date is already set', 'datePublicHolidayError')
            notValid.push('datePublicHolidayError')
        }

        // display main warning message
        this.displayWarningMessage(notValid, 'This Public Holiday already exists')

        return notValid.length
    }

    getForm(connectedID){
        const myjson = {}

        myjson['date']        = document.querySelector('#datePublicHoliday').value
        myjson['name']        = document.querySelector('#namePublicHoliday').value
        myjson['description'] = document.querySelector('#description').value
        myjson['createdBy']   = connectedID

        return JSON.stringify(myjson, function replacer(key, value) { return value})
    }

    // check if date format is valid
    isDateFormatValid(dateString) {
        const regex = /^\d{4}\/\d{2}\/\d{2}$/;
        return regex.test(dateString);
    }

    // CSV parsing 
    parseCSV(csvData) {
        const lines = csvData.split('\n');
        const headers = lines[0].split(',');
      
        const data = [];
        for (let i = 1; i < lines.length; i++) {
          const line = lines[i].replace(/\r/g, ''); // Remove \r characters
          const values = line.split(',');
          const entry = {};
          for (let j = 0; j < headers.length; j++) {
            if (typeof values[j] !== 'undefined' && values[j] !== '') {
              entry[headers[j]] = values[j];
            }
          }
      
          data.push(entry);
        }
      
        return data;
    }
    
    // fetch uploaded CSV
    validateFormCSV(){
        // store not valid fields
        let notValid = []

        // check if new public holiday doesn't already exists
        const myCSV = document.querySelector('#csvFile')
        if (myCSV.value == ''){
            Common.insertHTML('You must select a file', 'csvFileError')
            notValid.push('csvFileError')
        }

        return notValid.length
    }

    // returns only valid dates (well formatted & not already existing)
    validateParsedData(parsedData){
        // store not valid fields
        let valid = []
        parsedData.forEach(element => {
   
            // check if element is not empty
            if(Object.keys(element).length != 0){
                // check if date format is valid
                if (this.isDateFormatValid(element.Date)){
                    // clean potential / in the date and replace it by -
                    const myDate = element.Date.replace(/\//g, "-")
                    // check if date is empty or already exist
                    if ( (!ph.includes(myDate)) && (myDate != '') ){
                        let entry = { }
                        for (const key in element){
                            entry[key] = element[key]
                        }
                        valid.push(entry)
                    }
                }
            }
            
                
        })
        return valid
    }

    removeCarriageReturn(obj) {
        if (typeof obj === 'object' && obj !== null) {
          const result = Array.isArray(obj) ? [] : {};
          for (const key in obj) {
            if (obj.hasOwnProperty(key)) {
              const newKey = key.replace(/\r/g, ''); // Remove \r characters from key
              result[newKey] = this.removeCarriageReturn(obj[key]);
            }
          }
          return result;
        } else if (typeof obj === 'string') {
          return obj.replace(/\r/g, ''); // Remove \r characters from string value
        }
        return obj;
    }

    getFormCSV(validPH, connectedID){
        const myjson = {}

        myjson['PublicHolidays'] = validPH
        myjson['CreatedBy']      = connectedID

        return JSON.stringify(this.removeCarriageReturn(myjson))
    }

    
}