document.getElementById("registerForm").addEventListener("submit", async function (event) {
    event.preventDefault(); 

    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    const response = await fetch("http://localhost:8080/registerHandle", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password })
    });

    const result = await response.json();
    document.getElementById("message").innerText = result.message;
});

document.getElementById("back2login").addEventListener("click", function () {
    window.location.href = "/login"
})
