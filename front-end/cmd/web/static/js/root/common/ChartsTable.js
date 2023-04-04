class ChartsTable{
    // This function takes the data and HTML target as required parameters
    // There is an option to create a readyHandler:
    // - To resize specific columns for icons
    // - To create redirect click event on each icon (or specify noLink)
    // - To create show/hide columns edit (must have link option set to noLink)
    // 
    // Examples:
    // 
    // - Create single charts table
    //                                  dataDT = [
    //                                   [header1, header2, ...] 
    //                                   [row1.col1, row1.col2, ...] 
    //                                   [row2.col1, row2.col2, ...] 
    //                                  ]
    //  drawTable(dataDT, idTargetHTML)
    // 
    // - Create charts table with 2 redirect columns (so 2 columns to resize to icon)
    //                                  dataDT = [
    //                                   [header1, header2, ...] 
    //                                   [row1.col1, row1.col2, ..., row1.col8] 
    //                                   [row2.col1, row2.col2, ..., row2.col8] 
    //                                  ]
    //                                  editableCol = []                                            (used to show/hide form)
    //                                  resizeCol = [6,7]                                           (two last columns of table: col7 & col8)
    //                                  prefixURL = ['/somewhere/details', '/somewhere/delete']     (prefix URL of each column)
    //                                  EITHER
    //                                  wantedCol = [0,2,4]                                         (prefixURL/col1/col3/col5 of each column)
    //                                  or
    //                                  hiddenRowID = 7                                             (prefixURL/rowID of each column)
    //                                  the hidden input must be just after the icon
    //  
    //  drawTable(dataDT, idTargetHTML, editableCol, resizeCol, prefixURL, wantedCol, hiddenRowID) 
    // 
    //  when two last columns are clicked it redirect to  /somewhere/details/rowID or /somewhere/delete/rowID
    //  drawTable(dataDT, idTargetHTML, [], [6,7], ['/somewhere/details', '/somewhere/delete'], [], 7) 
    //  
    //  when two last columns are clicked it redirect to  /somewhere/details/col1/col3/col5 or /somewhere/delete/col1/col3/col5
    //  drawTable(dataDT, idTargetHTML, [], [6,7], ['/somewhere/details', '/somewhere/delete'], [0,2,4], null) 
    //  
    // - Create charts table with 1 redirect column and 1 show/hide form (so 2 columns to resize to icon)
    //  drawTable(dataDT, idTargetHTML, editableCol, resizeCol, prefixURL, wantedCol, hiddenRowID) 
    //  drawTable(dataDT, idTargetHTML, [4,5], [6,7], ['noLink', '/somewhere/delete'], [], 7) 
    //  we have column 5 & 6 editable when we click on column 7
    //  we have redirect link with rowID when we click on column 8
    // 
    drawTable(dataDT, idTargetHTML, editableCol=null, resizeCol=null, prefixURL=null, wantedCol=null, hiddenRowID=null){
        const data = google.visualization.arrayToDataTable(dataDT)
        const table = new google.visualization.Table(document.getElementById(idTargetHTML))

        // If ready handler wanted (prefixURL, wantedCol, hiddenInfo may be null)
        // but we supposed we always got columns to resize when event is requested
        if (resizeCol != null){
            google.visualization.events.addListener(table, 'ready', readyHandler)
            function readyHandler(){
                const headerCells = table.VS.children[0]

                // Resize columns wanted
                for (let index = 0; index < resizeCol.length; index++) {
                    let colIndex = resizeCol[index],
                        col = headerCells.children[0].children[colIndex]
                    col.classList.add("iconCol")
                }

                // Split table row in columns to collect information wanted
                const tbody = table.VS.children[1],
                      rowNumber = tbody.children.length         
                for (let index = 0; index < rowNumber; index++) {
                    const row = tbody.children[index]

                    // Define redirect path (part after prefix)
                    let redirectPath = '/'

                    // Concatenate information wanted based on visible columns
                    if (wantedCol.length > 0) {
                        for (let i = 0; i < wantedCol.length; i++) {
                            const wcIndex = wantedCol[i], // index of column wanted
                                  wcData = row.children[wcIndex].innerText // data of column wanted
                            redirectPath += wcData + '/'
                        }
                    }
                    
                    // Concatenate information wanted based on hidden rowID
                    // We assume we hide a hidden input after the icon
                    if (hiddenRowID != null){
                        const element = row.children[hiddenRowID],
                              rowID = element.children[1].value
                        redirectPath += rowID
                    }

                    // Create listener
                    for (let n = 0; n < resizeCol.length; n++) {
                        const btnIndex = resizeCol[n],
                              btn = row.children[btnIndex],
                              prefix = prefixURL[n],
                              url = prefix + redirectPath,
                              btnClasses = btn.children[0].classList,
                              btnFirstClass = btnClasses[Object.keys(btnClasses)[0]]
                        
                        // Check if row must not be editable (case headquarters)
                        if (btnFirstClass !='disabledElements'){
                            btn.addEventListener('click', () =>{
                                // Check if button is a link
                                if (prefix == 'noLink'){
                                    // Show edit form columns
                                    if (editableCol.length > 0) {
                                        for (let x = 0; x < editableCol.length; x++) {
                                            const e = editableCol[x] // index column to show form
                                            let vDisplay = row.children[e].children[0], // view column display
                                                fDisplay = row.children[e].children[1]; // form column display
                                            
                                            if (vDisplay.style.display == 'block'){
                                                vDisplay.style.display = 'none'
                                                fDisplay.style.display = 'block'
                                            }else{
                                                vDisplay.style.display = 'block'
                                                fDisplay.style.display = 'none'
                                            }
                                
                                        }
                                    }
                                }else{
                                    // redirect to URL
                                    if (confirm('Are you sure ?')) window.location.href = url
                                }
    
                            })
                        }
                                                
                    }
                  
                }            
            }
        }

        

        // If no event wanted we simply draw the table
        table.draw(data, {showRowNumber: false, width: '100%', height: '100%', allowHtml:true})

        
    }

    drawDashboard(dataDT, idTargetHTML, idFilterHTML, editableCol=null, resizeCol=null, prefixURL=null, wantedCol=null, hiddenRowID=null){
        const data = google.visualization.arrayToDataTable(dataDT)
              
        // create a list of columns for the dashboard
        var columns = [{
            // this column aggregates all of the data into one column
            // for use with the string filter
            type: 'string',
            calc: function (dt, row) {
                for (var i = 0, vals = [], cols = dt.getNumberOfColumns(); i < cols; i++) {
                    vals.push(dt.getFormattedValue(row, i));
                }
                return vals.join('\n');
            }
        }];
        for (var i = 0, cols = data.getNumberOfColumns(); i < cols; i++) {
            columns.push(i);
        }
        
        // Define a slider control for the 'Donuts eaten' column
        var filter = new google.visualization.ControlWrapper({
            controlType: 'StringFilter',
            containerId: idFilterHTML,
            options: {
                filterColumnIndex: 0,
                matchType: 'any',
                caseSensitive: false,
                ui: {
                    label: 'Search:',
                    cssClass: 'inputSearch'
                }
            },
            view: {
                columns: columns
            }
        });

        // Define a table chart
        var table = new google.visualization.ChartWrapper({
            chartType: 'Table',
            containerId: idTargetHTML,
            options: {
                showRowNumber: false,
                page: 'enable',
                pageSize: 20,
                pagingSymbols: {
                    prev: '<i class="bi-arrow-left-circle"></i>',
                    next: '<i class="bi-arrow-right-circle"></i>'
                },
                pagingButtonsConfiguration: 'auto',
                width: '100%', 
                height: '100%',
                allowHtml:true,
            },
            view: {
                columns: [1,2,3,4]
            }
        });
        
        // If ready handler wanted (prefixURL, wantedCol, hiddenInfo may be null)
        // but we supposed we always got columns to resize when event is requested
        if (resizeCol != null){
            google.visualization.events.addListener(table, 'ready', readyHandler)
            function readyHandler(){
                const headerCells = table.HS.children[0].children[0].children[0].children[0].children[0]

                // Resize columns wanted
                for (let index = 0; index < resizeCol.length; index++) {
                    let colIndex = resizeCol[index],
                        col = headerCells.children[colIndex]
                    col.classList.add("iconCol")
                }

                // Split table row in columns to collect information wanted
                const tbody = table.HS.children[0].children[0].children[0].children[1],
                      rowNumber = tbody.children.length  
     
                for (let index = 0; index < rowNumber; index++) {
                    const row = tbody.children[index]

                    // Define redirect path (part after prefix)
                    let redirectPath = '/'

                    // Concatenate information wanted based on visible columns
                    if (wantedCol.length > 0) {
                        for (let i = 0; i < wantedCol.length; i++) {
                            const wcIndex = wantedCol[i], // index of column wanted
                                  wcData = row.children[wcIndex].innerText // data of column wanted
                            redirectPath += wcData + '/'
                        }
                    }
                    
                    // Concatenate information wanted based on hidden rowID
                    // We assume we hide a hidden input after the icon
                    if (hiddenRowID != null){
                        const element = row.children[hiddenRowID],
                              rowID = element.children[1].value
                        redirectPath += rowID
                    }

                    // Create listener
                    for (let n = 0; n < resizeCol.length; n++) {
                        const btnIndex = resizeCol[n],
                              btn = row.children[btnIndex],
                              prefix = prefixURL[n],
                              url = prefix + redirectPath,
                              btnClasses = btn.children[0].classList,
                              btnFirstClass = btnClasses[Object.keys(btnClasses)[0]]
                        
                        // Check if row must not be editable (case headquarters)
                        if (btnFirstClass !='disabledElements'){
                            btn.addEventListener('click', () =>{
                                // Check if button is a link
                                if (prefix == 'noLink'){
                                    // Show edit form columns
                                    if (editableCol.length > 0) {
                                        for (let x = 0; x < editableCol.length; x++) {
                                            const e = editableCol[x] // index column to show form
                                            let vDisplay = row.children[e].children[0], // view column display
                                                fDisplay = row.children[e].children[1]; // form column display
                                            
                                            if (vDisplay.style.display == 'block'){
                                                vDisplay.style.display = 'none'
                                                fDisplay.style.display = 'block'
                                            }else{
                                                vDisplay.style.display = 'block'
                                                fDisplay.style.display = 'none'
                                            }
                                
                                        }
                                    }
                                }else{
                                    // redirect to URL
                                    if (confirm('Are you sure ?')) window.location.href = url
                                }
    
                            })
                        }
                                                
                    }
                  
                }            
            }
        }
        
        var dashboard = new google.visualization.Dashboard(document.getElementById('dashboard'));
        dashboard.bind([filter], [table]);
        dashboard.draw(data);
    }

}
