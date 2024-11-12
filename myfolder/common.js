// Retrieve the hidden div element
const hiddenDiv = document.getElementById('hiddenDiv');

// Check if the hidden div is visible
if (hiddenDiv.style.display !== 'none') {
    // Update the navigation element to be selected
    const navigationElement = document.getElementById('navigationElement');
    navigationElement.classList.add('selected');
}