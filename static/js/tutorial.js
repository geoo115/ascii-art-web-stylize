// Define your function first
function positionPopup(element) {
  var rect = element.getBoundingClientRect();
  var popup = document.getElementById('tutorial-modal');
}

document.addEventListener("DOMContentLoaded", function() {
  document.getElementById('tutorial-modal').style.display = "block";


document.getElementsByClassName('close-btn')[0].onclick = function() {
  document.getElementById('tutorial-modal').style.display = "none";
}

document.getElementById('next-step').onclick = function() {
  var currentStep = parseInt(this.getAttribute('data-step'), 10);
  var nextStep = currentStep + 1;
  
  if (nextStep === 2) {
      document.getElementById('tutorial-step').textContent = 'Step 2:Type in any words to generate ascii-art.';
  } else if (nextStep === 3) {
      document.getElementById('tutorial-step').textContent = 'Step 3: Click the "Download" button to download your ascii art.';
      this.textContent = 'Finish';
  } else if (nextStep === 4) {
      document.getElementById('tutorial-modal').style.display = "none";
  }
  
  this.setAttribute('data-step', nextStep);
}

window.onclick = function(event) {
  if (event.target == document.getElementById('tutorial-modal')) {
      document.getElementById('tutorial-modal').style.display = "none";
  }
}
  // Tutorial button event
  document.getElementById('tutorial-btn').onclick = function() {
    positionPopup(document.getElementById('banner-select'));
    document.getElementById('tutorial-modal').style.display = "block";
    document.getElementById('tutorial-step').textContent = 'Step 1: Select a banner from the dropdown.';
    document.getElementById('next-step').setAttribute('data-step', 1);
  }
});