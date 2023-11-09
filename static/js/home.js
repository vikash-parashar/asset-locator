// <!-- offcanvas script -->

$(document).ready(function () {
    // Initialize the offcanvas component
    var offcanvas = new bootstrap.Offcanvas(document.getElementById('offcanvasExample'));

    // Track the offcanvas state
    var offcanvasOpen = false;

    // Function to toggle the offcanvas state
    function toggleOffcanvas() {
        offcanvasOpen = !offcanvasOpen;
        if (offcanvasOpen) {
            offcanvas.show();
        } else {
            offcanvas.hide();
        }
    }

    // Event listener for mouseenter on the left side of the screen
    $(document).on('mouseenter', function (event) {
        if (event.clientX < 50) { // Adjust the value as needed
            toggleOffcanvas();
        }
    });

    // Event listener for mouseleave on the offcanvas element
    $('#offcanvasExample').on('mouseleave', function () {
        toggleOffcanvas();
    });
});




// <!-- sweet alerts -->

function showToast(icon, title) {
    const Toast = Swal.mixin({
        toast: true,
        position: 'top-end',
        showConfirmButton: false,
        timer: 1000,
        timerProgressBar: false,
        didOpen: (toast) => {
            toast.addEventListener('mouseenter', Swal.stopTimer);
            toast.addEventListener('mouseleave', Swal.resumeTimer);
        }
    });

    Toast.fire({
        icon: icon,
        title: title
    });
}
// showToast('success', 'Signed in successfully');



// <!-- get current user details -->


function capitalize(str) {
    return str.slice(0).toUpperCase();
}


// Function to fetch and display user details when the page loads
function fetchCurrentUserDetails() {
    fetch('/api/v1/get-current-user', {
        method: 'GET',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        }
    })
        .then(response => {
            console.log(response)
            if (!response.ok) {
                console.error("Network response was not ok");
                showToast("error", "Try Again !")
            }
            return response.json();
        })
        .then(data => {
            document.getElementById('firstNamePlaceholder').innerHTML = capitalize(data.user.first_name);
            // document.getElementById('lastNamePlaceholder').innerHTML = capitalize(data.user.last_name);
            document.getElementById('phonePlaceholder').innerHTML = capitalize(data.user.phone);
            // document.getElementById('emailPlaceholder').innerHTML = capitalize(data.user.email);
            document.getElementById('rolePlaceholder').innerHTML = capitalize(data.user.role);
        })

        .catch(error => {
            showToast("error", "Something went wrong !")
            console.error('Error fetching user details:', error);
        });
}

// Add an event listener to fetch user details when the page loads
document.addEventListener('DOMContentLoaded', fetchCurrentUserDetails);


// <!-- handle logout -->

// Function to handle logout
function handleLogout() {
    fetch("/logout", {
        method: "POST",
    })
        .then(response => {
            if (!response.ok) {
                showToast("error", "Try Again !")
                console.error("Network response was not ok");
            }
            return response.json();
        })
        .then(data => {
            if (data.success) {
                showToast("success", "Logout Success");
                window.location.href = "/"; // Redirect to the root path
            } else {
                showToast("error", "Try Again !");
            }
        })
        .catch(error => {
            console.error("An error occurred while logging out:", error);
            showToast(3, "An error occurred while logging out.");
        });
}

// Attach the handleLogout function to the logout button's click event
const logoutButton = document.getElementById("logout-button");
logoutButton.addEventListener("click", handleLogout);
