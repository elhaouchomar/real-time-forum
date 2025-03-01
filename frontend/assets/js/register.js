import { LoadPage } from "./spa.js";

const container = document.getElementById('container');
const registerBtnToggle = document.getElementById('registerBtnToggle');
const loginBtn = document.getElementById('loginBtn');
let pass = document.getElementById("pass");
let confirmPass = document.getElementById("confirmPass");
let user = document.getElementById("user");
let email = document.getElementById("email");
let registerBtn = document.getElementById("registerBtn");
let usernameMessage = document.getElementById("usernameMessage");
let emailMessage = document.getElementById("emailMessage");
let confirmPassMessage = document.getElementById("confirmPassMessage");

loginBtn.addEventListener('click', () => {
    container.classList.remove("active");
});

registerBtnToggle.addEventListener('click', () => {
    container.classList.add("active");
});

const RegForm = document.getElementById("RegisterForm")
const LogForm = document.getElementById("LoginForm")

RegForm.addEventListener("submit", function(event){
    event.preventDefault()
})

LogForm.addEventListener("submit", function(event){
    event.preventDefault()
    console.log("Submit Login Form");
    
    var email = LogForm.querySelector("input").value
    var pass = LogForm.querySelector(".password-container input").value
    
    fetch("/login", {
        method: "POST",
        headers: {
            "Content-Type" : "application/json"
        },
        body: JSON.stringify({
            "email": email,
            "password": pass,
        })

    }).then(response => {
            if (!response.ok){
                throw new Error("Error Connecting to login route")
            }
            return response.json()
        })
        .then(data =>  {
            console.log(data);
            if (data.status){
                console.log("inside the Json Response");
                LoadPage("home")
            }
        }).catch(err => {
            console.log(err);
        })
})


function checkPassword() {
    const patterns = {
        lower: /(?=.*[a-z])/,
        upper: /(?=.*[A-Z])/,
        number: /(?=.*\d)/,
        symbol: /(?=.*[\W_])/,
        length: /^.{8,}$/
    };

    let Divs = {
        lower: document.querySelector('.lower'),
        upper: document.querySelector('.upper'),
        number: document.querySelector('.number'),
        symbol: document.querySelector('.symbol'),
        length: document.querySelector('.length'),
        passChecker: document.querySelector('.passChecker'),
    };

    Divs.passChecker.style.display = pass.value ? "block" : "none";

    for (const [k, v] of Object.entries(patterns)) {
        const icon = Divs[k].querySelector('.material-icons');
        icon.innerHTML = v.test(pass.value) ? '&#xe86c;' : '&#xe5c9;';
        icon.style.color = v.test(pass.value) ? "green" : "red";
        Divs[k].style.color = v.test(pass.value) ? "green" : "red";
    }
}

function validateForm() {
    let usernameValid = /^[\w]+$/.test(user.value);
    let emailValid = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/.test(email.value);
    let confirmPasswordValid = pass.value === confirmPass.value;
    let passwordValid = pass.value.match(/(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[\W_]).{8,}/);

    if (!usernameValid && user.value !== "") {
        usernameMessage.style.display = "block";
    } else {
        usernameMessage.style.display = "none";
    }

    if (!emailValid && email.value !== "") {
        emailMessage.style.display = "block";
    } else {
        emailMessage.style.display = "none";
    }

    if (!confirmPasswordValid && confirmPass.value !== "") {
        confirmPassMessage.style.display = "block";
    } else {
        confirmPassMessage.style.display = "none";
    }

    registerBtn.disabled = !(usernameValid && emailValid && confirmPasswordValid && passwordValid);
}

pass.addEventListener("input", () => {
    checkPassword();
    validateForm();
});

user.addEventListener("input", validateForm);
email.addEventListener("input", validateForm);

confirmPass.addEventListener("input", validateForm);


// shoz password
document.querySelectorAll('.toggle-password').forEach(toggle => {
    toggle.addEventListener('click', function () {
        const input = this.previousElementSibling;
        const type = input.getAttribute('type') === 'password' ? 'text' : 'password';
        input.setAttribute('type', type);
        this.textContent = type === 'password' ? 'visibility_off' : 'visibility';
    });
});


function switchForm(type) {
    const loginForm = document.querySelector('.sign-in');
    const registerForm = document.querySelector('.sign-up');
    const buttons = document.querySelectorAll('.mobile-switch button');

    // Remove active class from all forms
    [loginForm, registerForm].forEach(form => {
        form.classList.remove('active-form');
    });

    // Add active class to selected form
    if (type === 'login') {
        loginForm.classList.add('active-form');
    } else {
        registerForm.classList.add('active-form');
    }

    // Update button states
    buttons.forEach(button => {
        button.classList.remove('active');
        if (button.textContent.toLowerCase() === type) {
            button.classList.add('active');
        }
    });
}
// Initialize mobile view
window.addEventListener('load', () => {
    if (window.innerWidth <= 768) {
        switchForm('login');
    }
});