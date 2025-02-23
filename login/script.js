document.getElementById("loginForm").addEventListener("submit", async function (event) {
    event.preventDefault(); 

    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    const response = await fetch("http://localhost:8080/loginHandle", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password })
    });

    const result = await response.json();
    document.getElementById("message").innerText = result.message;
});

document.getElementById("back2register").addEventListener("click", function () {
    window.location.href = "/register"
})

document.getElementById("forget").addEventListener("click", function() {
    window.location.href = "/forget"
})
