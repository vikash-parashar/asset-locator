
function notify(type, msg) {
    notie.alert({
        type: type,
        text: msg,
        position: 'top'
    })
}

document.addEventListener("DOMContentLoaded", function () {
    const registerForm = document.getElementById("register-form");

    registerForm.addEventListener("submit", async function (event) {
        event.preventDefault();
        try {
            const firstName = document.getElementById("first_name").value;
            const lastName = document.getElementById("last_name").value;
            const phone = document.getElementById("phone").value;
            const email = document.getElementById("register-email").value;
            const password = document.getElementById("register-password").value;

            const response = await fetch("/signup", {
                method: "POST",
                body: JSON.stringify({
                    first_name: firstName,
                    last_name: lastName,
                    phone: phone,
                    email: email,
                    password: password,
                }),
                headers: {
                    "Content-Type": "application/json",
                },
            });

            if (!response.ok) {
                const data = await response.json();
                console.error("Signup error:", data);
                notify("warning", data.message || "Something went wrong. Please try again.");
            } else {
                notify("success", "Registration successful. You can now log in.");
                setTimeout(() => {
                    window.location.href = "http://localhost:8080";
                }, 2000);
            }
        } catch (error) {
            console.error("Signup error:", error);
            notify("error", "An error occurred. Please try again later.");
        }
    });
});

document.addEventListener("DOMContentLoaded", function () {
    const loginForm = document.getElementById("login-form");

    loginForm.addEventListener("submit", async function (event) {
        event.preventDefault();
        try {
            const email = document.getElementById("login-email").value;
            const password = document.getElementById("login-password").value;

            const formData = new FormData(); // Create a new FormData object
            formData.append("email", email);
            formData.append("password", password);

            const response = await fetch("/login", {
                method: "POST",
                body: formData, // Send the form data
            });

            if (!response.ok) {
                const data = await response.json();
                console.error("Login error:", data);
                notify("warning", data.message || "Login failed. Please check your credentials.");
            } else {
                notify("success", "Login successful. Redirecting...");
                setTimeout(() => {
                    window.location.href = "http://localhost:8080/api/v1/homepage";
                }, 2000);
            }
        } catch (error) {
            console.error("Login error:", error);
            notify("error", "An error occurred. Please try again later.");
        }
    });
});

