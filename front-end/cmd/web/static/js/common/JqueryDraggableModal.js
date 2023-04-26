class DraggableModal{
    draggableModal(modalId) {
        if (!($('.modal.in').length)) {
          $('.modal-dialog').css({
            top: 0,
            left: 0
          })
        }
        $('#' + modalId).modal({
          backdrop: 'static',
          keyboard: false,
          show: false
        })
    
        $('.modal-dialog').draggable({
          handle: ".modal-header"
        })
      }
}