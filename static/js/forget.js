// <!-- sweet alerts -->

function showToast(icon, title) {
    const Toast = Swal.mixin({
        toast: true,
        position: 'top-end',
        showConfirmButton: false,
        timer: 3000,
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


// <!-- handle forget password -->

document.addEventListener('DOMContentLoaded', function () {
    const resetPasswordForm = document.getElementById('forgot-password-form');

    resetPasswordForm.addEventListener('submit', function (e) {
        e.preventDefault();

        // Get the email address from the form
        const email = document.getElementById('email').value;

        // Make an AJAX request to the /forget-password endpoint
        fetch('/forget-password', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ email }),
        })
            .then(response => response.json())
            .then(data => {
                if (data.success) {
                    showToast("success", data.message);

                    // Display success message
                    const successMessage = document.getElementById('successMessage');
                    successMessage.style.display = 'block';
                    successMessage.innerHTML = data.message;
                    setTimeout(() => {
                        window.location.href = "http://localhost:8080";
                    }, 3000);
                } else {
                    // Display an error message
                    showToast("error", data.message);
                }
            })
            .catch(error => {
                console.error('An error occurred:', error);
                showToast("error", error.message);
            });
    });
});


