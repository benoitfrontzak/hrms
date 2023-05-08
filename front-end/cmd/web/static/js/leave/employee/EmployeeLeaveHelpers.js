class EmployeeLeaveHelpers {
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

    // update connected user leaves (details)
    populateMyLeaveDetails(data){
        // populate main card with leave details
        const target = document.querySelector('#leaveDetailsBody')
        target.innerHTML = ''

        data.forEach(element => {
            const balance = Number(this.cleanDecimal(element.entitled)) + Number(element.credits) - Number(element.taken)
            let row = document.createElement('div')
            row.classList = 'row justify-content-md-center'
            row.innerHTML = `<div class="col-4 text-start"><input type="text" class="form-control form-control-sm transparentInput" value="${element.leaveDefinitionCode} - ${element.leaveDefinitionName}" disabled /></div>
                             <div class="col-2 text-start"><input type="text" class="form-control form-control-sm transparentInput" value="${this.cleanDecimal(element.entitled)}" disabled /></div>
                             <div class="col-2 text-start"><input type="number" step="0.01" class="form-control form-control-sm transparentInput allCredits" data-rowid="${element.rowid}" value="${element.credits}" /></div>
                             <div class="col-2 text-start"><input type="text" class="form-control form-control-sm transparentInput" value="${element.taken}" disabled /></div>
                             <div class="col-2 text-start"><input type="text" class="form-control form-control-sm transparentInput" value="${balance}" disabled /></div>`

            target.appendChild(row)
        })
    }

     // update connected user leaves (progress details)
     populateMyLeaveDetailsProgress(data){
        // populate main card with leave details
        const target = document.querySelector('#chartLeaves')
        target.innerHTML = ''

        data.forEach(element => {
            const balance = Number(this.cleanDecimal(element.entitled)) + Number(element.credits) - Number(element.taken),
                  max = Number(this.cleanDecimal(element.entitled)) + Number(element.credits)
     
            let percentage
            (element.taken != 0)? percentage = Math.round(Number(element.taken)*100/max) : percentage = 0

            let row = document.createElement('div')
            row.classList = 'mb-3 p-3 border borderRadiusTop borderRadiusBottom'
            row.innerHTML = `<div class="row mb-3">
                                <div class="col">${element.leaveDefinitionCode} - ${element.leaveDefinitionName}</div>
                                <div class="col text-end">${element.taken}/${max} days</div>
                            </div>
                            <div class="row">
                                <div class="col">
                                    <div class="progress">
                                        <div class="progress-bar myHarmonyBlue" role="progressbar" style="width: ${percentage}%" aria-valuenow="${percentage}" aria-valuemin="0" aria-valuemax="${max}">${percentage}%</div>
                                    </div>
                                </div>
                             </div>`

            target.appendChild(row)
        }) 
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
    
    // get form data
    getForm(connectedID, connectedEmail){
        const allCredits = []
        // fetch all credits of employee
        const rows = document.querySelectorAll('.allCredits')
        rows.forEach(credits => {
            const oneCredits = {
                rowid:credits.dataset.rowid,
                credits:credits.value,
                connectedUser:connectedEmail,
                userID:connectedID
            }
            allCredits.push(oneCredits)
        })
        
        return JSON.stringify(allCredits, function replacer(key, value) { return value})
    }

}