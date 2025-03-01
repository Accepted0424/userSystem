document.getElementById("forgetForm").addEventListener("submit", async function (event) {
    event.preventDefault(); 

    const email = document.getElementById("email").value;
    const new_password = document.getElementById("password").value;
    const new_password_repeat = document.getElementById("password_repeat").value; 
    const message = document.getElementById("message");

    //验证两次密码是否一致
    if (new_password != new_password_repeat) {
        message.innerText = "两次密码输入不一致，请重新输入";
        message.style.color = "red";
        return;
    }

    const response = await fetch("http://localhost:8080/loginHandle", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, new_password })
    });

    const result = await response.json();
    message.innerText = result.message;
});

document.getElementById("back2login").addEventListener("click", function () {
    window.location.href = "/login"
})
