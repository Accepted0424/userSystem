document.getElementById("send").addEventListener("click", async function name() {
    const email = document.getElementById("email").value;
    const message = document.getElementById("message")
    //发送验证码
    const response = await fetch("http://localhost:8080/emailvertify", {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({email})
    });
    
    const result = await response.json();
    message.innerText = result.message;
});

document.getElementById("registerForm").addEventListener("submit", async function (event) {
    event.preventDefault(); 

    const email = document.getElementById("email").value;
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;
    const password_repeat = document.getElementById("password_repeat").value;
    const code = document.getElementById("verifyCode").value;
    const message = document.getElementById("message")

    if (password !== password_repeat) {
        message.innerText = "两次密码输入不一致，请重新输入";
        message.style.color = "red";
        return;
    }

    const response = await fetch("http://localhost:8080/registerHandle", {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({ username, email, code, password })
    });

    const result = await response.json();
    message.innerText = result.message;
});

document.getElementById("back2login").addEventListener("click", function () {
    window.location.href = "/login"
})
