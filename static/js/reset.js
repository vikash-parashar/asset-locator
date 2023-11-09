
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

document.addEventListener('DOMContentLoaded', function () {
    const passwordResetForm = document.getElementById('password-reset-form');

    passwordResetForm.addEventListener('submit', function (e) {
        e.preventDefault();

        // Get the new password and confirm password from the form
        const newPassword = document.getElementById('new-password').value;
        const confirmNewPassword = document.getElementById('confirm-password').value;

        // Check if passwords match
        if (newPassword !== confirmNewPassword) {
            showToast("error", "Password Mismatch");
            return;
        }

        // Prepare data for submission
        const data = {
            new_password: newPassword,
        };

        // Get the token from the query parameter in the URL
        const urlParams = new URLSearchParams(window.location.search);
        const token = urlParams.get('token');

        if (token) {
            // Now you can use the 'token' variable in your fetch request
            fetch(`/reset-password/?token=${token}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(data),
            })
                .then(response => response.json())
                .then(data => {
                    if (data.success) {
                        showToast("success", "Your Password Is Changed Now");
                        // Optionally, you can redirect the user to a success page or perform other actions.
                    } else {
                        showToast("error", "Something Wrong , Try Again Later !");
                    }
                })
                .catch(error => {
                    console.error('An error occurred:', error);
                    showToast("error", "Something Wrong , Try Again Later !");
                });
        } else {
            console.error("token is not present in URL")
            // Handle the case where the token is not found in the URL
            showToast("error", "Something Wrong , Try Again Later !");
        }

    });
});