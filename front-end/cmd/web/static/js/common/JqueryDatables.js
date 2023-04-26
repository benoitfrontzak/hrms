class DataTableFeatures {
  initiateMyTable(myTableId, mycolumns){
    const table = $('#'+myTableId).DataTable({
      columns: mycolumns,
      scrollX: true,
      scrollCollapse: true,
      fixedColumns:   {
          left: 1,
          right: 1
      },        
      colReorder: true,
      stateSave:  true,
      lengthMenu: [
          [10, 25, 50, -1],
          ['10 rows', '25 rows', '50 rows', 'Show all']
      ],
      dom: '<"row"<"col"B><"col"fr>>t<"row"<"col"l><"col"p>>',
      buttons: [
                {extend:'excel',
                 exportOptions: {columns: ':visible' }
                }, 
                {extend:'pdf',
                 exportOptions: {columns: ':visible' },
                 customize: function(doc) {
                  for (var row = 0; row < doc.content[1].table.headerRows; row++) {
                    var header = doc.content[1].table.body[row];
                    for (var col = 0; col < header.length; col++) {
                      header[col].fillColor = 'red';
                    }
                  }
                }   
                }
              ],
      "bInfo" : false,
      "initComplete": () => {
          $('#processing').remove()
      }
    })

    // show/hide columns feature
    const defaultConf = mycolumns.map( obj => {
      return obj.title.replace(/ /g,'');
    })
    this.showHideColumnsDT(table, defaultConf, myTableId+'Columns', 'columnMenu')

    return table

  }

  showHideColumnsDT(table, defaultConf, storageName, classInputName){
    const confExist = localStorage.getItem(storageName)
    const allCheckboxes = document.querySelectorAll('.' + classInputName)
    const removedColumns = this.loadColumnsDT(table, defaultConf, storageName)
    const availableColumns = defaultConf.filter(n => !removedColumns.includes(n))

    //Update columns to hide/display when checkboxes are unchecked/checked
    for (const [key, myInput] of Object.entries(allCheckboxes)) {
      myInput.addEventListener('click', () => {
        if (myInput.checked === true) {
          table.column(key).visible(true);
          availableColumns.push(myInput.name)
          localStorage.setItem(storageName, JSON.stringify(availableColumns))
        } else {
          table.column(key).visible(false);
          availableColumns.splice(availableColumns.indexOf(myInput.name), 1)
          localStorage.setItem(storageName, JSON.stringify(availableColumns))
        }
      });
    }
  }
  loadColumnsDT(table, defaultConf, storageName){
    const confExist = localStorage.getItem(storageName)
    let removedColumns
    //If conf exist: hide columns not stored in localStorage
    //If conf doesn't exist : set default columns to show in localStorage
    if (confExist) {
      removedColumns = defaultConf.filter(n => !confExist.includes(n))
      for (const [keyRemoved, removedColumn] of Object.entries(removedColumns)) {
        if (typeof(document.getElementById(removedColumn)) !== 'undefined' && document.getElementById(removedColumn) !== null){
          document.getElementById(removedColumn).checked = false
          table.column(defaultConf.indexOf(removedColumn)).visible(false);
        }
      }
    } else {
      localStorage.setItem(storageName, JSON.stringify(defaultConf))
      removedColumns = []
    }
    return removedColumns
  }
  showHideColumns(defaultConf, storageName, classInputName){
    const confExist = localStorage.getItem(storageName)
    const allCheckboxes = document.querySelectorAll('.' + classInputName)
    const removedColumns = this.loadColumns(defaultConf, storageName)
    const availableColumns = defaultConf.filter(n => !removedColumns.includes(n))

    //Update columns to hide/display when checkboxes are unchecked/checked
    for (const [key, myInput] of Object.entries(allCheckboxes)) {
      myInput.addEventListener('click', () => {
        if (myInput.checked === true) {
          $('.' + myInput.name).toggle()
          availableColumns.push(myInput.name)
          localStorage.setItem(storageName, JSON.stringify(availableColumns))
        } else {
          $('.' + myInput.name).hide();
          availableColumns.splice(availableColumns.indexOf(myInput.name), 1)
          localStorage.setItem(storageName, JSON.stringify(availableColumns))
        }
      });
    }
  }
  loadColumns(defaultConf, storageName, classInputName){
    const confExist = localStorage.getItem(storageName)
    let removedColumns
    //If conf exist: hide columns not stored in localStorage
    //If conf doesn't exist : set default columns to show in localStorage
    if (confExist) {
      removedColumns = defaultConf.filter(n => !confExist.includes(n))
      for (const [keyRemoved, removedColumn] of Object.entries(removedColumns)) {
        if (typeof(document.getElementById(removedColumn)) != 'undefined' && document.getElementById(removedColumn) != null){
          document.getElementById(removedColumn).checked = false
          $('.' + removedColumn).hide()
        }
      }
    } else {
      localStorage.setItem(storageName, JSON.stringify(defaultConf))
      removedColumns = []
    }
    return removedColumns
  }
}