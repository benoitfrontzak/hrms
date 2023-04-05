class HomeHelpers{
    // update connected user profile information
    populateMyProfile(data){
        const target = document.querySelector('#myProfile')
        target.innerHTML = ''
        
        const myCredentials = this.getCredentials(data.Employee.icNumber, data.Employee.passportNumber)

        let row = document.createElement('div')
        row.innerHTML = `<h4>${data.Employee.fullname}</h4>
                         <div>${myCredentials}</div>
                         <div>${data.Employment.jobTitle}</div>`

        target.appendChild(row) 
    }

    // update connected user profile information (join | confirm date)
    populateMyProfileDates(joinD, confirmD){
        const target = document.querySelector('#myProfileDates')
        target.innerHTML = ''

        const jDate = this.formatDate(joinD),
              cDate = this.formatDate(confirmD),
              jDays = this.daysDifferenceNow(jDate),
              cDays = this.daysDifferenceNow(cDate)

        let row = document.createElement('div')

        row.innerHTML = `<div class="row">
                            <div class="col-2 smaller">Join:</div>
                            <div class="col-6 smaller">${jDate} (${jDays} days)</div>
                            <div class="col-4"></div>
                         </div>
                         <div class="row">
                            <div class="col-2 smaller">Confirm:</div>
                            <div class="col-6 smaller">${cDate} (${cDays} days)</div>
                            <div class="col-4"></div>
                         </div>`

        target.appendChild(row) 
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

    // update connected user claims (details)
    populateMyClaimDetails(data){
        // populate main card with total claims (default 0)
        document.querySelector('#myClaimsPending').innerHTML = 0
        document.querySelector('#myClaimsRequested').innerHTML = 0
        // populate main card with total claims (with data if exists)
        if (data.Pending != null)
            document.querySelector('#myClaimsPending').innerHTML = data.Pending.length
        if (data.Rejected != null || data.Approved != null)
            document.querySelector('#myClaimsRequested').innerHTML = data.Rejected.length + data.Approved.length
        
        // populate modal pending
        if (data.Pending != null){
            const targetPending = document.querySelector('#myPendingRequestBody')
            targetPending.innerHTML = ''

            data.Pending.forEach(element => {
                let row = document.createElement('div')
                row.classList = 'd-flex flex-row justify-content-between'
                row.innerHTML = `<div>${element.claimDefinition}</div>
                                <div>${element.description}</div>
                                <div>${element.amount}</div>
                                <div>${this.formatDate(element.createdAt)}</div>`

                targetPending.appendChild(row)
            });
        }
        // populate modal requested
        const targetRequested = document.querySelector('#requestedMyClaimsBody')
        targetRequested.innerHTML = ''

        if (data.Approved != null){
            data.Approved.forEach(element => {
                let row = document.createElement('div')
                row.classList = 'row'
                row.innerHTML = `<div class="col-2">${element.claimDefinition}</div>
                                <div class="col-3">${element.description}</div>
                                <div class="col">${element.status}</div>
                                <div class="col">${element.amount}</div>
                                <div class="col">${this.convertUndefined(element.approvedAmount)}</div>
                                <div class="col-3">${element.approvedReason}</div>
                                <div class="col">${this.formatDate(element.createdAt)}</div>`

                targetRequested.appendChild(row)
            });
        }

        if (data.Rejected != null){
            data.Rejected.forEach(element => {
                let row = document.createElement('div')
                row.classList = 'row'
                row.innerHTML = `<div class="col-2">${element.claimDefinition}</div>
                                <div class="col-3">${element.description}</div>
                                <div class="col">${element.status}</div>
                                <div class="col">${element.amount}</div>
                                <div class="col">${this.convertUndefined(element.approvedAmount)}</div>
                                <div class="col-3">${element.approvedReason}</div>
                                <div class="col">${this.formatDate(element.createdAt)}</div>`

                targetRequested.appendChild(row)
            });
        }
    }

    // update today & tomorrow leaves
    populateTodayAndTomorrowLeaves(data){
        // on leave today
        this.populateCurrentLeave(data.Today, 'onLeaveTodayBody', 'leaveToday')

        // on leave tomorrow
        this.populateCurrentLeave(data.Tomorrow, 'onLeaveTomorrowBody', 'leaveTomorrow')
        
    }

    // update connected user leaves (today & tomorrow)
    populateCurrentLeave(data, targetID, cardID){  
        const myCard =  document.querySelector('#'+cardID),
              target = document.querySelector('#'+targetID)

        // if data exists 
        if (data != null) {

            // populate main card with total leaves
            myCard.innerHTML = data.length

            // populate modal on leave today | tomorrow           
            target.innerHTML = ''

            data.forEach(element => {
                let row = document.createElement('div')
                row.classList = 'row'
                row.innerHTML = `<div class="col">${allEmployees.get(Number(element.EmployeeID))}</div>
                                <div class="col">${element.Code} - ${element.Description}</div>
                                <div class="col">${element.Reason}</div>
                                <div class="col">${this.myBooleanIcons(element.IsHalf)}</div>`

                target.appendChild(row)
            })

        }else{
            myCard.innerHTML = 0
            target.innerHTML = 'no data'
        }
    }

    // update connected user leaves (details)
    populateMyLeaveDetails(data){
        // populate main card with leave details
        const target = document.querySelector('#leaveDetailsBody')
        target.innerHTML = ''

        data.forEach(element => {
            const balance = Number(this.cleanDecimal(element.entitled)) - Number(element.taken)
            let row = document.createElement('div')
            row.classList = 'row'
            row.innerHTML = `<div class="col text-start">${element.leaveDefinitionCode} - ${element.leaveDefinitionName}</div>
                             <div class="col text-start">${this.cleanDecimal(element.entitled)}</div>
                             <div class="col text-start">${element.taken}</div>
                             <div class="col text-start">${balance}</div>`

            target.appendChild(row)
        })
    }

    // customize boolean (0|1) with icons
    myBooleanIcons(value){
        return (value == 1) ? '<i class="bi-check2-square"></i> true' : '<i class="bi-x-square"></i> false'
    }

    // convert undefined value to 0
    convertUndefined(myVal){
        let zeroVal
        (typeof myVal == 'undefined') ? zeroVal = 0 : zeroVal = myVal
        return zeroVal
    }

    // return ic or passport
    getCredentials(ic, passport){
        let credentials = 'Not defined'
        if (ic != '' && ic != null) credentials = ic
        if (passport != '' && passport != null) credentials = passport
        return credentials
    }

    // convert timestamp to date
    formatDate(t){
        let b
        // get only the 10 first characters of the string
        const d = t.substring(0,10)
        // the zero value of a date is 0001-01-01
        return (d == '0001-01-01') ? b = '' : b = d
    }

    // calculate numbers of days between date and now (round up)
    daysDifferenceNow(date){
        // To set two dates to two variables
        const now = new Date(),
              myDate = new Date(date);
        
        let timeDiff, daysDiff
        if (now < myDate){
            timeDiff = myDate.getTime() - now.getTime()
            daysDiff = timeDiff / (1000 * 3600 * 24)
        } else {
            timeDiff = now.getTime() - myDate.getTime()
            daysDiff = timeDiff / (1000 * 3600 * 24)
        }
        
        return Math.round(daysDiff)
    }

    // clean decimal value to follow rule: only .5 is allowed
    cleanDecimal(v){
        let cleaned = v
        // check if string contains decimal value (ie got a .)
        if (v.includes(".")){
            const n = Number(v),
                  int = Number(v.split('.')[0]),
                  fract = n%1
                  
            let cFract = 0
            if (fract.toFixed(4) >= 0.5) cFract = 0.5
            cleaned = int + cFract
        }
        return cleaned.toString()
    }
    // build leaves chart data
    buildLeavesChartData(data){
        const myData = [['Leave',   'Entitled', 'Taken', 'Balance']]
        
        data.forEach(element => {
          const leave = element.leaveDefinitionCode+' - '+element.leaveDefinitionName,
                entitled = Number(this.cleanDecimal(element.entitled)),
                taken = Number(element.taken),
                balance = entitled - taken,
                myAr = [leave, entitled, taken, balance]

          myData.push(myAr)
        })

        return myData
    }
    // draw leave Google chart
    chartLeaves(myData){
        const gData = this.buildLeavesChartData(myData)
        const data = google.visualization.arrayToDataTable(gData)

        const options = {
          title : 'Employee\'s leave',
          vAxis: {title: 'Days'},
          hAxis: {title: 'Leave'},
          seriesType: 'bars',
          colors: ['#3544DC', '#DC0015', '#45DC35']
        }

        const chart = new google.visualization.ComboChart(document.getElementById('chartLeaves'))
        chart.draw(data, options)

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
        title: 'My Claims ' + new Date().getFullYear(),
        is3D: true,
        pieSliceText: 'value-and-percentage',
        colors: ['#3544DC', '#45DC35', '#DC0015']
        }

        const chart = new google.visualization.PieChart(document.getElementById('pieChartClaims'))
        chart.draw(data, options)
    }

}