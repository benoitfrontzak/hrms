class EmployeeRender{
    // UPDATE EMPLOYEE SKELETON PAGE
    updateEmployeePage(employeeInfo, uploadedFiles){
      const employeeID = employeeInfo.Employee.id
        return `
        <form name="updateEmployeeForm${employeeID}" id="updateEmployeeForm${employeeID}" method="POST">
          <div class="zoomOut mt-3 p-3 borderRadiusTop borderRadiusBottom shadow">
              
              <!-- Title -->
              <div class="row w-100 mt-3 align-items-center">
                  <div class="col-2"><div class="text-center text-secondary" style="font-size: 64px"><i class="bi-person-circle"></i></div></div>
                  <div class="col-10">
                    <div class="row align-items-center">
                      <div class="col-10"><h1 class="text-secondary fw-bolder ms-3">${employeeInfo.Employee.fullname}</h1></div>
                      <div class="col-2 text-end"><a href="#" class="btn btn-secondary myButton updateEmployeeSubmit" data-id="${employeeID}" data-bs-toggle="tooltip" title="Save Employee"><i class="bi bi-save2-fill largeIcon"></i></a></div>
                    </div>
                  </div>
              </div>
              <!-- Title -->

              <!-- Main -->
              <div class="row w-100 mt-3">
                  <div class="col-2 text-white">
                      <!-- Vertical menu -->
                      <div class="row border-start border-top border-end border-danger ms-3 p-3 myTint3BG shadowOnly borderRadiusTop"> <div class="col pointer personal">Personal</div> </div>
                      <div class="row border-start border-end border-danger ms-3 p-3 myTint3BG shadowOnly"> <div class="col pointer spouse">Spouse</div> </div>
                      <div class="row border-start border-end border-danger ms-3 p-3 myTint3BG shadowOnly"> <div class="col pointer employment">Employment</div> </div>
                      <div class="row border-start border-bottom border-end border-danger ms-3 p-3 myTint3BG shadowOnly borderRadiusBottom"> <div class="col pointer statutory">Statutory</div> </div>
                      <!-- Vertical menu -->
                  </div>
                  <div class="col-10">
                      <!-- Content -->
                      <div class="show">${Render.personalPage(employeeInfo.Employee, employeeID)}</div>
                      <div class="hide">${Render.spousePage(employeeInfo.Spouse, employeeID)}</div>
                      <div class="hide">${Render.employmentPage(employeeInfo.Employment, employeeID)}</div>
                      <div class="hide">${Render.statutoryPage(employeeInfo.Statutory, employeeID)}</div>
                      <!-- Content -->
                  </div>
              </div>
              <!-- Main -->

          </div>
        </form>
        `
    }
    
    // PERSONAL SKELETON PAGE
    personalPage(data, employeeID){
        return `
        <div class="row pointer ms-3 border p-3 myIdentity"> <div class="col-10">Identity</div> <div class="col-2 text-end disabled"><i class="bi-caret-down"></i></div> </div>
        <div class="row pointer ms-3 border p-3 myIdentityContent hide"> <div class="col">${Render.identityCard(data, employeeID)}</div>  </div>

        <div class="row pointer ms-3 border p-3 myContact"> <div class="col-10">Contact</div> <div class="col-2 text-end disabled"><i class="bi-caret-down"></i></div> </div>
        <div class="row pointer ms-3 border p-3 myContactContent hide"> <div class="col">${Render.contactCard(data, employeeID)}</div> </div>

        <div class="row pointer ms-3 border p-3 myBank"> <div class="col-10">Bank</div> <div class="col-2 text-end disabled"><i class="bi-caret-down"></i></div> </div>
        <div class="row pointer ms-3 border p-3 myBankContent hide"> <div class="col">${Render.bankCard(data, employeeID)}</div> </div>

        <div class="row pointer ms-3 border p-3 myEmergencyContact"> <div class="col-10">Emergency Contact</div> <div class="col-2 text-end disabled"><i class="bi-caret-down"></i></div> </div>
        <div class="row pointer ms-3 border p-3 myEmergencyContactContent hide"> <div class="col">${Render.emergencyContactCard(data, employeeID)}</div> </div>

        <div class="row pointer ms-3 border p-3 myOtherInformation"> <div class="col-10">Other Information</div> <div class="col-2 text-end disabled"><i class="bi-caret-down"></i></div> </div>
        <div class="row pointer ms-3 border p-3 myOtherInformationContent hide"> <div class="col">${Render.otherInformationCard(data, employeeID)}</div> </div>`
    }
    /*
    PERSONAL HELPERS
    */ 
    // identity card
    identityCard(data, employeeID){
      const gender             = [{label:'Male', value:'1'}, {label:'Female', value:'0'}],
            maritalstatus      = [{label:'Single', value:'0'}, {label:'Married', value:'1'}],
            role               = [{label:'Admin', value:'2'}, {label:'Manager', value:'3'}, {label:'User', value:'4'}],
            passportExpiryAt   = Helpers.convertToDate(data.passportExpiryAt),
            passportExpiryDays = Helpers.daysDifferenceNow(passportExpiryAt),
            birthdate          = Helpers.formatDate(data.birthdate),
            age                = Helpers.calculateAge(birthdate)
      
      return `
        <div class="row my-3">
          <div class="col-sm-8">${RenderBootstrap.inputFloatingRequired('fullname'+employeeID, data.fullname, 'Full Name')}</div>
          <div class="col-sm-4">${RenderBootstrap.inputFloatingRequired('nickName'+employeeID, data.nickname, 'Nick Name')}</div>
        </div>
        <div class="space"></div>
        <div class="row my-3">
          <div class="col-sm-4">${RenderBootstrap.inputFloatingRequired('employeeCode'+employeeID, data.employeeCode, 'Employee Code')}</div>
          <div class="col-sm-4"></div>
          <div class="col-sm-4">
            <div class="row">
              <div class="col">${RenderBootstrap.radioButtons('gender'+employeeID, gender, data.gender)}</div> 
              <div class="col">${RenderBootstrap.radioButtons('maritalstatus'+employeeID, maritalstatus, data.maritalstatus)}</div> 
            </div>
          </div>
        </div>
        <div class="space"></div>
        <hr />
        <div class="row my-3">
          <div class="col-sm-4">
            ${RenderBootstrap.inputFloating('icNumber'+employeeID, data.icNumber, 'IC Number')}
            ${RenderBootstrap.uploadFeature('uploadICButton'+employeeID, 'uploadedIC', 'archivedIC')}
          </div>
          <div class="col-sm-4">
            ${RenderBootstrap.inputFloating('passportNumber'+employeeID, data.passportNumber, 'Passport Number')}
            ${RenderBootstrap.uploadFeature('uploadPassportButton'+employeeID, 'uploadedPassport', 'archivedPassport')}
          </div>
          <div class="col-sm-4">
            ${RenderBootstrap.inputDateFloating('passportExpiryAt'+employeeID, passportExpiryAt, 'Passport Expiry Date', passportExpiryDays)}
          </div>
        </div>
        <div class="space"></div>
        <hr />
        <div class="row my-3">
          <div class="col-sm-4">${RenderBootstrap.selectBoxFloatingRequired('nationality'+employeeID, employeeCT.Country, 'Nationality', data.nationality, true)}</div>
          <div class="col-sm-4">${RenderBootstrap.selectBoxFloatingRequired('residence'+employeeID, employeeCT.Country, 'Residence', data.residence)}</div>
          <div class="col-sm-4">${RenderBootstrap.inputFloating('immigrationNumber'+employeeID, data.immigrationNumber, 'Immigration Number')}</div>
        </div> 
        <div class="space"></div>
        <hr />
        <div class="row my-3">
          <div class="col-sm-4">${RenderBootstrap.inputDateFloating('birthdate'+employeeID, birthdate, 'Date Of Birth', age)}</div>
          <div class="col-sm-4">${RenderBootstrap.selectBoxFloating('race'+employeeID, employeeCT.Race, 'Race', data.race)}</div>
          <div class="col-sm-4">${RenderBootstrap.selectBoxFloating('religion'+employeeID, employeeCT.Religion, 'Religion', data.religion)}</div>
        </div>  
        <div class="space"></div>
        <hr />
        <div class="row mb-3">
          <div class="col-sm-4">
            ${RenderBootstrap.checkBox('isActive'+employeeID, data.isActive, 'Is Active')}
            ${RenderBootstrap.checkBox('isForeigner'+employeeID, data.isForeigner, 'Is Foreigner')}
            ${RenderBootstrap.checkBox('isDisabled'+employeeID, data.isDisabled, 'Is Disabled Person')}
          </div>
          <div class="col-sm-4"></div>
          <div class="col-sm-4">Role ${RenderBootstrap.radioButtons('role'+employeeID, role, data.role)}</div>
        </div>`
    }
    // contact card
    contactCard(data, employeeID){
        return `
        <div class="row mb-3">
          <div class="col-sm-4">${RenderBootstrap.inputFloatingRequired('streetAddressLine1'+employeeID, data.streetaddr1, 'Street Address Line 1')}</div>
          <div class="col-sm-4">${RenderBootstrap.inputFloatingRequired('streetAddressLine2'+employeeID, data.streetaddr2, 'Street Address Line 2')}</div>
          <div class="col-sm-4">${RenderBootstrap.inputFloatingRequired('zip'+employeeID, data.zip, 'Zip Code')}</div>      
        </div>
        <div class="row mb-3">
          <div class="col-sm-4">${RenderBootstrap.inputFloatingRequired('city'+employeeID, data.city, 'City')}</div>
          <div class="col-sm-4">${RenderBootstrap.inputFloatingRequired('state'+employeeID, data.state, 'State')}</div>
          <div class="col-sm-4">${RenderBootstrap.selectBoxFloatingRequired('country'+employeeID, employeeCT.Country, 'Country', data.country)}</div>
        </div>
        <div class="space"></div>
        <hr />
        <div class="row mb-3">
          <div class="col-sm-4">${RenderBootstrap.inputFloatingRequired('primaryPhone'+employeeID, data.primaryPhone, 'Primary Phone')}</div>
          <div class="col-sm-4">${RenderBootstrap.inputFloating('secondaryPhone'+employeeID, data.secondaryPhone, 'Secondary Phone')}</div>
          <div class="col-sm-4"></div>
        </div>
        <div class="row mb-3">
          <div class="col-sm-4">${RenderBootstrap.inputFloatingRequired('primaryEmail'+employeeID, data.primaryEmail, 'Primary Email')}</div>
          <div class="col-sm-4">${RenderBootstrap.inputFloating('secondaryEmail'+employeeID, data.secondaryEmail, 'Secondary Email')}</div>
          <div class="col-sm-4"></div>
        </div>`
    }
    // contact card
    bankCard(data, employeeID){
        return `
        <div class="row mb-3">
          <div class="col-sm-4">${RenderBootstrap.inputFloating('bankName'+employeeID, data.bankName, 'Bank')}</div>
          <div class="col-sm-4">${RenderBootstrap.inputFloating('accountNumber'+employeeID, data.accountNumber, 'Account Number')}</div>
          <div class="col-sm-4">${RenderBootstrap.inputFloating('accountName'+employeeID, data.accountName, 'Beneficiary Name')}</div>
        </div>`
    }
    // contact card
    emergencyContactCard(data, employeeID){
        return `
        <div class="row mb-3">
          <div class="col-sm-4">${RenderBootstrap.inputFloating('fullnameEC'+employeeID, data.fullnameEC, 'Full Name')}</div>
          <div class="col-sm-4">${RenderBootstrap.inputFloating('mobileEC'+employeeID, data.mobileEC, 'Mobile')}</div>
          <div class="col-sm-4">${RenderBootstrap.selectBoxFloatingRequired('relationshipEC'+employeeID, employeeCT.Relationship, 'Relationship', data.relationshipEC)}</div>
        </div>`
    }
    // contact card
    otherInformationCard(data, employeeID){
        return `
        <div class="row mb-3">
          <div class="col-sm-4">${RenderBootstrap.textareaFloating('education'+employeeID, data.education, 'Education')}</div>
          <div class="col-sm-4">${RenderBootstrap.textareaFloating('experience'+employeeID, data.experience, 'Experience')}</div>
          <div class="col-sm-4">${RenderBootstrap.textareaFloating('notes'+employeeID, data.notes, 'Notes')}</div>
        </div>`
    }

    // SPOUSE SKELETON PAGE
    spousePage(data, employeeID){
        return `<div class="row pointer ms-3 border p-3 spouseIdentity"> <div class="col-10">Identity</div> <div class="col-2 text-end disabled"><i class="bi-caret-down"></i></div> </div>
                <div class="row pointer ms-3 border p-3 spouseIdentityContent hide"> <div class="col">${Render.spouseIdentityCard(data, employeeID)}</div>  </div>

                <div class="row pointer ms-3 border p-3 spouseWorking"> <div class="col-10">Working</div> <div class="col-2 text-end disabled"><i class="bi-caret-down"></i></div> </div>
                <div class="row pointer ms-3 border p-3 spouseWorkingContent hide"> <div class="col">${Render.spouseWorkingCard(data, employeeID)}</div> </div>

                <div class="row pointer ms-3 border p-3 spouseContact"> <div class="col-10">Contact</div> <div class="col-2 text-end disabled"><i class="bi-caret-down"></i></div> </div>
                <div class="row pointer ms-3 border p-3 spouseContactContent hide"> <div class="col">${Render.spouseContactCard(data, employeeID)}</div> </div>`
    }
    /*
    SPOUSE HELPERS
    */
    spouseIdentityCard(data, employeeID){
      return `spouse identity ${employeeID}`
    }
    spouseWorkingCard(data, employeeID){
      return `spouse working ${employeeID}`
    }
    spouseContactCard(data, employeeID){
      return `spouse contact ${employeeID}`
    }

    // EMPLOYEMENT SKELETON PAGE
    employmentPage(data, employeeID){
        return `<div class="row pointer ms-3 border p-3 myPayroll"> <div class="col-10">Payroll</div> <div class="col-2 text-end disabled"><i class="bi-caret-down"></i></div> </div>
                <div class="row pointer ms-3 border p-3 myPayrollContent hide"> <div class="col">${Render.myPayrollCard(data, employeeID)}</div>  </div>

                <div class="row pointer ms-3 border p-3 myEmployment"> <div class="col-10">Employment</div> <div class="col-2 text-end disabled"><i class="bi-caret-down"></i></div> </div>
                <div class="row pointer ms-3 border p-3 myEmploymentContent hide"> <div class="col">${Render.myEmploymentCard(data, employeeID)}</div> </div>`
    }
    /*
    EMPLOYEMENT HELPERS
    */
    myPayrollCard(data, employeeID){
      return `payroll ${employeeID}`
    }
    myEmploymentCard(data, employeeID){
      return `employment ${employeeID}`
    }

    // STATUTORY SKELETON PAGE
    statutoryPage(data, employeeID){
        return `<div class="row pointer ms-3 border p-3 myEPF"> <div class="col-10">EPF</div> <div class="col-2 text-end disabled"><i class="bi-caret-down"></i></div> </div>
                <div class="row pointer ms-3 border p-3 myEPFContent hide"> <div class="col">${Render.myEPFCard(data, employeeID)}</div>  </div>

                <div class="row pointer ms-3 border p-3 mySOCSO"> <div class="col-10">SOCSO</div> <div class="col-2 text-end disabled"><i class="bi-caret-down"></i></div> </div>
                <div class="row pointer ms-3 border p-3 mySOCSOContent hide"> <div class="col">${Render.mySOCSOCard(data, employeeID)}</div> </div>

                <div class="row pointer ms-3 border p-3 myIncomeTax"> <div class="col-10">Income Tax</div> <div class="col-2 text-end disabled"><i class="bi-caret-down"></i></div> </div>
                <div class="row pointer ms-3 border p-3 myIncomeTaxContent hide"> <div class="col">${Render.myIncomeTaxCard(data, employeeID)}</div> </div>
                
                <div class="row pointer ms-3 border p-3 myOthers"> <div class="col-10">Others</div> <div class="col-2 text-end disabled"><i class="bi-caret-down"></i></div> </div>
                <div class="row pointer ms-3 border p-3 myOthersContent hide"> <div class="col">${Render.myOthersCard(data, employeeID)}</div> </div>`
    }
    /*
    STATUTORY HELPERS
    */
    myEPFCard(data, employeeID){
      return `EPF ${employeeID}`
    }
    mySOCSOCard(data, employeeID){
      return `SOCSO ${employeeID}`
    }
    myIncomeTaxCard(data, employeeID){
      return `Income Tax ${employeeID}`
    }
    myOthersCard(data, employeeID){
      return `others ${employeeID}`
    }

}