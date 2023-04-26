class EmployeeRead{

    // Insert to datatable one row per element of data
    insertRows(data){
        const target = document.querySelector('#employeeSummaryBody')
        target.innerHTML = ''
  
        const myColumns = [
          {title: "Code"},
          {title: "Fullname"},
          {title: "Email"},
          {title: "Mobile"},
          {title: "Birthdate"},
          {title: "Race"},
          {title: "Gender"},
          {title: "Created At"},
          {title: "Created By"},
          {title: "Updated At"},
          {title: "Updated By"},
          {title: "Action"}
        ]
  
        // when no data
        if (data == null){
          DT.initiateMyTable('employeeSummary', myColumns)
          $('#employeeSummary').DataTable().clear().draw()
          return
        }
  
        // insert rows
        data.forEach(element => {
          let row = document.createElement('tr')
          row.classList.add('newTab')
          row.id = element.ID
          row.innerHTML = `<td class="pointer"><a href="#" class="link-dark myLink">${element.Code}</a></td>
                          <td class="pointer"><a href="#" class="link-dark myLink">${element.Fullname}</a></td>
                          <td class="pointer"><a href="#" class="link-dark myLink">${element.Email}</a></td>
                          <td class="pointer"><a href="#" class="link-dark myLink">${element.Mobile}</a></td>
                          <td class="pointer"><a href="#" class="link-dark myLink">${this.removeTimestamp(element.Birthdate)}</a></td>
                          <td class="pointer"><a href="#" class="link-dark myLink">${element.Race}</a></td>
                          <td class="pointer"><a href="#" class="link-dark myLink">${this.gender(element.Gender)}</a></td>
                          <td class="pointer"><a href="#" class="link-dark myLink">${this.removeTimestamp(element.CreatedAt)}</a></td>
                          <td class="pointer"><a href="#" class="link-dark myLink">${allEmployees.get(Number(element.CreatedBy))}</a></td>
                          <td class="pointer"><a href="#" class="link-dark myLink">${this.removeTimestamp(element.UpdatedAt)}</a></td>
                          <td class="pointer"><a href="#" class="link-dark myLink">${allEmployees.get(Number(element.UpdatedBy))}</a></td>
                          <td>
                              <div class="form-check">
                                  <input class="form-check-input deleteCheckboxes"  type="checkbox" value="${element.ID}" name="softDelete">
                                  <label class="form-check-label fw-lighter fst-italic smaller" for="softDelete"><i class="bi-trash2-fill largeIcon pointer deleteEmployee"></i></label>
                              </div>                                
                          </td>`
          target.appendChild(row)
        })
  
        // initiate datatable & show/hide columns
        DT.initiateMyTable('employeeSummary', myColumns)
  
        // when one row is clicked, create new tab with selected employee information
        this.createRowClickEvent()  
  
      }

    /*
     HELPERS
    */ 

    // convert gender_id
    gender(id){
        let g
        return (id == 1) ? g = 'M <i class="bi-gender-male"></i>' : g = 'F <i class="bi-gender-female"></i>'
    }

    // remove timestamp from date yyyy-mm-dd
    removeTimestamp(t){
      let myDate
      // get only the 10 first characters of the string
      const d = t.substring(0,10)

      // the zero value of a date is 0001-01-01
      return (d == '0001-01-01') ? myDate = '' : myDate = d
    }

    // make datatable rows clickable but the last column 
    // open employee to new tab for edit 
    createRowClickEvent(){
      $('#employeeSummary tbody').on('click', 'td:not(:last-child)', function() {
        const tr = $(this).closest('tr'),
              employeeID = tr[0].id
              window.open('/employee/update/'+employeeID, '_blank')
      }) 
    }
}