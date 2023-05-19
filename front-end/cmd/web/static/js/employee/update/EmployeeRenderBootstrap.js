class EmployeeRenderBootstrap {
    inputFloatingRequired(id, value, label){
        let myValue = ''
        if (typeof value != 'undefined') myValue = value

        return `<div class="form-floating">
                    <input type="text" class="form-control" id="${id}" name="${id}" value="${myValue}" />
                    <label for="${id}"><i class="bi-shield-fill-exclamation"></i> ${label}</label>
                </div>
                <div class="fw-bolder text-danger fst-italic smaller" id="${id}Error"></div>`
    }

    inputFloating(id, value, label){
        let myValue = ''
        if (typeof value != 'undefined') myValue = value

        return `<div class="form-floating">
                    <input type="text" class="form-control" id="${id}" name="${id}" value="${myValue}" />
                    <label for="${id}">${label}</label>
                </div>
                <div class="fw-bolder text-danger fst-italic smaller" id="${id}Error"></div>`
    }

    inputDateFloating(id, value, label, calculation){
        let myValue = ''
        if (typeof value != 'undefined') myValue = value

        return `<div class="form-floating">
                    <input type="date" class="form-control" id="${id}" name="${id}" value="${myValue}" />
                    <label for="${id}">${label} <span class="fw-bolder text-secondary">${calculation}</span></label>
                </div>
                <div class="fw-bolder text-danger fst-italic smaller" id="${id}Error"></div>`
    }

    radioButtons(id, data, value){
        let result =''
        for (let i = 0; i < data.length; i++) {
            const element = data[i];
            let checked = ''
            if (value == element.value) checked = 'checked' 
            result += `<div class="form-check">
                            <input class="form-check-input" type="radio" name="${id}" id="${id}${i}" value=${element.value} ${checked} />
                            <label class="form-check-label" for="${id}${i}">${element.label}</label>
                        </div>`
        }        
        return result
    }

    checkBox(id, value, label){
        let checked = ''
        if (value == 1) checked = 'checked'

        return `<div class="form-check">
                    <input class="form-check-input" type="checkbox" value="${value}" id="${id}" name="${id}" ${checked} />
                    <label class="form-check-label fw-lighter fst-italic smaller" for="${id}">${label}</label>
                </div>`
    }

    selectBoxFloatingRequired(id, data, label, value, nationality=false){
        let result = `<div class="form-floating">
                        <select class="form-select" id="${id}" name="${id}">`
        for (let i = 0; i < data.length; i++) {
            const element = data[i];
            let optDisplay =  element.Name
            if (nationality) optDisplay = element.Nationality
            let selected = ''
            if (value == element.ID) selected = 'selected' 
            result += `<option value="${element.ID}" ${selected}>${optDisplay}</option>`
        }
        result += `</select>
                   <label for="${id}"><i class="bi-shield-fill-exclamation"></i> ${label}</label>
                </div>
                <div class="fw-bolder text-danger fst-italic smaller" id="${id}Error"></div>`     
        return result
    }

    selectBoxFloating(id, data, label, value){
        let result = `<div class="form-floating">
                        <select class="form-select" id="${id}" name="${id}">`
        for (let i = 0; i < data.length; i++) {
            const element = data[i];
            let selected = ''
            if (value == element.ID) selected = 'selected' 
            result += `<option value="${element.ID}" ${selected}>${element.Name}</option>`
        }
        result += `</select>
                   <label for="${id}"><i class="bi-shield-fill-exclamation"></i> ${label}</label>
                </div>
                <div class="fw-bolder text-danger fst-italic smaller" id="${id}Error"></div>`     
        return result
    }

    textareaFloating(id, value, label){
        let myValue = ''
        if (typeof value != 'undefined') myValue = value

        return `<div class="form-floating">
                    <textarea class="form-control" name="${id}" id="${id}" style="height: 200px">${myValue}</textarea>
                    <label for="${id}">${label}</label>
                </div>
                <div class="fw-bolder text-danger fst-italic smaller" id="${id}Error"></div>`
    }

    uploadFeature(id, uploadedID, archivedID){
        return `<div class="row mt-1">
                    <div class="col"><i class="bi-upload largeIcon pointer" id="${id}" data-bs-toggle="tooltip" data-bs-placement="top" title="Upload files"></i></div>
                    <div class="col text-end">
                    <div class="d-flex justify-content-end">
                        <span data-bs-toggle="modal" data-bs-target="#${uploadedID}"><i class="bi-folder-check largeIcon pointer" data-bs-toggle="tooltip" data-bs-placement="top" title="Latest files"></i></span>
                        <span>&nbsp;</span>
                        <span data-bs-toggle="modal" data-bs-target="#${archivedID}"><i class="bi-folder-x largeIcon pointer" data-bs-toggle="tooltip" data-bs-placement="top" title="Archive files"></i></span>
                    </div>                    
                    </div>
                </div>`
    }
    
}