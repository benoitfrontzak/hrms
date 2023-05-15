class EmployeeClaimHelpers {
    // populate option list with active employee
    populateEmployeeList(data){
        const target = document.querySelector('#datalistOptions')
        target.innerHTML = ''
        data.forEach(element => {
            const row = document.createElement('option')
            row.id = element.ID
            row.value = element.Fullname
            target.appendChild(row)
        })
    }

    // update connected user claims (total)
    populateMyClaims(data){
        let total
        total = Number(data.Pending) + Number(data.Rejected) + Number(data.Approved)
        document.querySelector('#claimsTotal').innerHTML = total
        document.querySelector('#claimsTotalPending').innerHTML = data.Pending
        document.querySelector('#claimsTotalRejected').innerHTML = data.Rejected
        document.querySelector('#claimsTotalApproved').innerHTML = data.Approved      
         
    }

     // draw claims Google pie
     pieChartClaims(myData) {
        const data = google.visualization.arrayToDataTable([
        ['Claim', 'Total'],
        ['Pending', myData.Pending],
        ['Approved', myData.Approved],
        ['Rejected', myData.Rejected]
        ])

        const options = {
        // title: 'My Claims ' + new Date().getFullYear(),
        is3D: true,
        legend: 'none',
        pieSliceText: 'percentage',
        backgroundColor: 'transparent',
        chartArea: {'width': '80%', 'height': '80%'},
        colors: ['#3587DC', '#35DCCC', '#515151']
        }

        const chart = new google.visualization.PieChart(document.getElementById('pieChartClaims'))
        chart.draw(data, options)

        
    }

    // populate my claim map (already requested amount by claim id)
    populateMyClaimMap(data, myClaim){
        data.forEach(element => {
            let previousAmount = 0.0;
            if (typeof myClaim.get(element.claimDefinitionID) != 'undefined') previousAmount = Number(myClaim.get(element.claimDefinitionID))
                 myClaim.set(element.claimDefinitionID,Number(element.amount)+previousAmount)
        })
        return myClaim
    }

    // populate all uploaded files map
    populateUploadedFilesMap(data, uploadedFiles){
        for (const [key, value] of Object.entries(data)) {
            uploadedFiles.set(key, value)
        }
        return uploadedFiles
    }

    // create attachments icon
    createAttachments(appID){
        return (myUploadedFiles.has(appID)) ? '<i class="bi-check-square pointer"></i>' : '<i class="bi-file-excel-fill not-allowed"></i>'
    }

    // populate attachments modal & open it
    populateAttachments(appID, email){
        if (myUploadedFiles.has(appID)){
            const myFiles = myUploadedFiles.get(appID),
                  target = document.querySelector('#uploadedFilesModalBody'),
                  myModal = new bootstrap.Modal(document.getElementById('uploadedFilesModal'), { 
                    backdrop: 'static',
                    keyboard: false 
                  })
            target.innerHTML = ''
            myFiles.forEach(element => {
                const oneAttachment = document.createElement('div')
                oneAttachment.classList = 'd-flex justify-content-center mb-3'
                oneAttachment.innerHTML = `<div>
                                                <div><img class="myPicture" src="/upload/${email}/claims/${appID}/${element}" /></div>
                                                <div class="text-end"><a href="/upload/${email}/claims/${appID}/${element}" class="myLink" download="${element}"><i class="bi-download pointer"></i> download</a></div>
                                           </div>`
                target.appendChild(oneAttachment)
            })
            myModal.show()
        }
    }

    // insert to datatable one row per element of data
    insertRows(data) {
        const target = document.querySelector('#claimEmployeeBody')
        target.innerHTML = ''

        const myColumns = [
            {title: "Allowance"},
            {title: "Description"},
            {title: "Category"},
            {title: "Amount"},
            {title: "Status"},
            {title: "Approved At"},
            {title: "Approved By"},
            {title: "Approved Amount"},
            {title: "Reason"},
            {title: "Attachment"},
            {title: "Created By"},
            {title: "Created At"},
            {title: "Updated By"},
            {title: "Updated At"}
        ]
        // can initiate DT only one time
        if ( $.fn.dataTable.isDataTable( '#claimEmployeeTable' ) ) {
            $('#claimEmployeeTable').DataTable().destroy()
        }
        
        // when no data
        if (data == null){
          DT.initiateMyTable('claimEmployeeTable', myColumns)
          $('#claimEmployeeTable').DataTable().clear().draw()
          return
        }

        data.forEach(element => {
            // since omitempty has been set golang side...
            let approvedBy, approvedAmount;
            (typeof element.approvedAmount == 'undefined') ? approvedAmount = 0 : approvedAmount = element.approvedAmount;
            (typeof element.approvedBy == 'undefined') ? approvedBy = 0 : approvedBy = Number(element.approvedBy);

            let attachments = this.createAttachments(element.rowid)

            let row = document.createElement('tr')
            row.id = 'myClaim' + element.rowid
            row.innerHTML = `<td class="row-data" data-id="${element.claimDefinitionID}">${element.claimDefinition}</td>
                             <td class="row-data">${element.description}</td>
                             <td class="row-data">${element.category}</td>
                             <td class="row-data">${element.amount}</td>
                             <td class="row-data" data-id="${element.statusID}">${element.status}</td>
                             <td class="row-data">${this.formatDate(element.approvedAt)}</td>
                             <td class="row-data">${allEmployees.get(approvedBy)}</td>
                             <td class="row-data">${approvedAmount}</td>
                             <td class="row-data">${element.rejectedReason}</td>
                             <td class="row-data myAttachments text-center" data-id=${element.rowid}>${attachments}</td>
                             <td class="row-data">${allEmployees.get(Number(element.createdBy))}</td>
                             <td class="row-data">${this.formatDate(element.createdAt)}</td>
                             <td class="row-data">${allEmployees.get(Number(element.updatedBy))}</td>
                             <td class="row-data">${this.formatDate(element.updatedAt)}</td>`
         
            target.appendChild(row)
        })
        
        // initiate datatable & show/hide columns
        DT.initiateMyTable('claimEmployeeTable', myColumns)
    }

    // convert birthdate (remove timestamp)
    formatDate(t) {
        let b
        // get only the 10 first characters of the string
        const d = t.substring(0, 10)
        // the zero value of a date is 0001-01-01
        return (d == '0001-01-01') ? b = '' : b = d
    }

}