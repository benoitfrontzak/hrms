class GoogleDT{

    drawTableAddEmployee(places) {
        // Prepare data for google DT addEmployee
        const dtAddEmployee = [
            ['Region','Subregion','Country','City','Place','Department','Add']
        ]
        places.forEach(place => {
            const departments = place.Departments
            departments.forEach(department => {
                const oneRow = []
                oneRow.push(`${place.Region}`)
                oneRow.push(`${place.Subregion}`)
                oneRow.push(`${place.Country}`)
                oneRow.push(`${place.City}`)
                oneRow.push(`${place.Name}`)
                oneRow.push(`${department}`)
                oneRow.push(`<i class="bi-plus-square-fill">`)
                dtAddEmployee.push(oneRow)
            })
        })
        // Load data
        const dataAddEmployee = google.visualization.arrayToDataTable(dtAddEmployee)  
        // Specify HTML target  
        const table = new google.visualization.Table(document.getElementById('table_add_employee'))        
        // Resize action column by adding iconCol css class
        // Define eventListener on icon click (redirect link)
        google.visualization.events.addListener(table, 'ready', readyHandler)
        function readyHandler(){
            const headerCells = table.VS.children[0]
            const addColHeader = headerCells.children[0].children[6]
            addColHeader.classList.add("iconCol")
            const tbody = table.VS.children[1],
                  rowNumber = tbody.children.length         
            for (let index = 0; index < rowNumber; index++) {
                const row = tbody.children[index],
                      region = row.children[0].innerText,
                      subregion = row.children[1].innerText,
                      country = row.children[2].innerText,
                      city = row.children[3].innerText,
                      place = row.children[4].innerText,
                      dept = row.children[5].innerText,
                      add = row.children[6]
                add.addEventListener('click', () => {
                    window.location.href = `/employee/add/${region}/${subregion}/${country}/${city}/${place}/${dept}`
                })
            }            
        } 
        // Draw DT   
        table.draw(dataAddEmployee, {showRowNumber: false, width: '100%', height: '100%',allowHtml:true})        
    }

    drawTableWorldMapEmployees(places) {
        // Prepare data for google DT worldMap
        const dtWorldMap = [
            ['Place','Department','Employee ID','Job','Type','Edit']
        ]
        places.forEach(place => {
            const employees = place.Employees
            for (const [eid, employee] of Object.entries(employees)) {
            const jobs = employee.Job
            jobs.forEach(job => {
                if (job.EndDate == ""){
                const oneRow = []
                oneRow.push(`${place.Name}`)
                oneRow.push(`${job.Department}`)
                oneRow.push(`${eid}`)
                oneRow.push(`${job.Name}`)
                oneRow.push(`${job.WorkingType}`)
                oneRow.push(`<i class="bi-pencil-square">`)
                dtWorldMap.push(oneRow)
                }
            })
            }
        })
        // Load data
        const dataWorldMap = google.visualization.arrayToDataTable(dtWorldMap)
        // Specify HTML target
        const table = new google.visualization.Table(document.getElementById('table_world_employees'))      
        
        // Resize action column by adding iconCol css class
        // Define eventListener on icon click (redirect link)
        google.visualization.events.addListener(table, 'ready', readyHandler)
        function readyHandler(){
            const headerCells = table.VS.children[0]         
            const editColHeader = headerCells.children[0].children[5]
            editColHeader.classList.add("iconCol")
            const tbody = table.VS.children[1],
                  rowNumber = tbody.children.length        
            for (let index = 0; index < rowNumber; index++) {
                const row = tbody.children[index],
                      place = row.children[0].innerText,
                      department = row.children[1].innerText,
                      employeeID = row.children[2].innerText,
                    //   job = row.children[3].innerText,
                    //   type = row.children[4],
                      edit = row.children[5]
                edit.addEventListener('click', () => {
                    window.location.href = `/employee/edit/${place}/${employeeID}`
                })
            }            
        }
        // Draw DT   
        table.draw(dataWorldMap, {showRowNumber: false, width: '100%', height: '100%',allowHtml:true})        
    }

    drawRegionsMap(locationEmployees, countryCode="", chartDivID) {
        // Prepare data for google geoChart worldMap
        const geoChartWorldMap = [
            ['Location','Employees']
        ]
        for (const [location, employeesNumber] of Object.entries(locationEmployees)) {
            const oneRow = []
            oneRow.push(`${location}`)
            oneRow.push(`${employeesNumber}`)
            geoChartWorldMap.push(oneRow)
        }
        const data = google.visualization.arrayToDataTable(geoChartWorldMap)

        // Prepare options for countryMap
        let options = {}
        if (countryCode != ""){
            options = {
                region: countryCode,
                displayMode: 'markers',
                resolution: "provinces"
            }
        }
        
        const chart = new google.visualization.GeoChart(document.getElementById(chartDivID))

        // google.visualization.events.addListener(chart, 'select', selectClick)
        // function selectClick(){
        //     if (chart.getSelection().length > 0) {
        //         const node = data.getValue(chart.getSelection()[0].row, 0)
        //         window.location.href = `${window.location.pathname}/${node}`
        //     }
        // }
        chart.draw(data, options)      
    }
}
