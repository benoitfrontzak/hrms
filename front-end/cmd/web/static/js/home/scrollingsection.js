// scroll to section
  // set up an array of section ids
  var sectionIds = ["section1", "section2", "section3"];

  // set up an index to keep track of the current section
  var currentSectionIndex = 0;

  // add an event listener for the mouse wheel event
  window.addEventListener("wheel", function(event) {
    // get the direction of the mouse wheel
    var delta = Math.sign(event.deltaY);

    // if the direction is negative (scrolling up)
    if (delta < 0) {
      // move to the previous section if we're not at the beginning
      if (currentSectionIndex > 0) {
        currentSectionIndex--;
      }
    }
    // if the direction is positive (scrolling down)
    else {
      // move to the next section if we're not at the end
      if (currentSectionIndex < sectionIds.length - 1) {
        currentSectionIndex++;
      }
    }

    // get the HTML element of the current section
    var currentSection = document.getElementById(sectionIds[currentSectionIndex]);

    // scroll to the current section using the scrollIntoView() method
    currentSection.scrollIntoView({ behavior: 'smooth' });

    // prevent the default scrolling behavior of the mouse wheel event
    event.preventDefault();
  });