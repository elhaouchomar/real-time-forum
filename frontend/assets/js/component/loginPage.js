export function LoginPage(){
    const temp = document.createElement('div')
    temp.innerHTML =  `
     <div class="container" id="container">
        <div class="form-container  sign-up" id="sign-up-form">
            <form id="RegisterForm">
                <h1>Create Account</h1>
                <div style="justify-content: space-between; display: flex; gap: 10px;">
                    <input type="text" name="first-name" placeholder="First Name" required>
                    <!-- <div id="usernameMessage" class="message"><i class="material-icons check-uncheck"
                            style="font-size: 10px;">&#xe5c9;</i> Username must contain only letters and numbers.</div> -->
                    <input type="text" name="last-name" placeholder="Last Name" required>

                </div>
                <div style="justify-content: space-between; display: flex; gap: 30px; width: 100%;" >
                    <select name="gender" id="gender" style="display: flex; justify-content: space-between;">
                        <option value="">Gender</option>
                        <option value="male">Male</option>
                        <option value="female">Female</option>
                    </select>         
                    <input type="number" id="" name="age" placeholder="Age" required>      
                </div>
                <input type="text" id="user" name="username" placeholder="Username" required>
                <div id="usernameMessage" class="message"><i class="material-icons check-uncheck"
                        style="font-size: 10px;">&#xe5c9;</i> Username must contain only letters and numbers.</div>

                <input type="email" id="email" name="email" placeholder="Email" required>
                <div id="emailMessage" class="message"><i class="material-icons check-uncheck"
                        style="font-size: 10px;">&#xe5c9;</i> Invalid email format.</div>
                <div class="password-container">
                    <input type="password" name="password" id="pass" placeholder="Password" required>
                    <span class="material-symbols-outlined toggle-password">visibility_off</span>
                </div>
                <div class="passChecker">
                    <div class="number"><i class="material-icons check-uncheck" style="font-size: 10px;">&#xe5c9;</i> A
                        number</div>
                    <div class="lower"><i class="material-icons check-uncheck" style="font-size: 10px;">&#xe5c9;</i> A
                        lowercase letter</div>
                    <div class="symbol"><i class="material-icons check-uncheck" style="font-size: 10px;">&#xe5c9;</i> A
                        symbol ex(@\']...)</div>
                    <div class="length"><i class="material-icons check-uncheck" style="font-size: 10px;">&#xe5c9;</i>
                        Minimum 8 characters</div>
                    <div class="upper"><i class="material-icons check-uncheck" style="font-size: 10px;">&#xe5c9;</i> A
                        capital (uppercase) letter</div>
                </div>
                <div class="password-container">
                    <input id="confirmPass" type="password" name="password" placeholder="Password" required>
                    <span class="material-symbols-outlined toggle-password">visibility_off</span>
                </div>
                <div id="confirmPassMessage" class="message"><i class="material-icons check-uncheck"
                        style="font-size: 10px;">&#xe5c9;</i> Confirmation password does not match.</div>
                <button type="submit" id="registerBtn" disabled>Register</button>
                <p>You have an account? <a href="#" onclick="switchForm('login')">Sign in</a></p>
                <div class="ErrorMessage">
                    <i class="material-icons" style="font-size: 14px;">&#xe000;</i>
                    <span class="Content">
                    </span>
                </div>
            </form>
        </div>

        <div class="form-container active-form sign-in" id="login-form">
            <form id="LoginForm">
                <h1>Sign In</h1>
                <input type="text" name="email" placeholder="Email" required>
                <div class="password-container">
                    <input type="password" name="password" placeholder="Password" required>
                    <span class="material-symbols-outlined toggle-password">visibility_off</span>
                </div>
                <p>Don't have an account? <a href="#" onclick="switchForm('register')">Sign Up</a></p>
                <button type="submit">Sign In</button>
                <div class="ErrorMessage">
                    <i class="material-icons" style="font-size: 14px;">&#xe000;</i>
                    <span class="Content">
                    </span>
                </div>
            </form>
        </div>

        <div class="toggle-container">
            <div class="toggle">
                <div class="toggle-panel toggle-left">
                    <h1>Welcome Back!</h1>
                    <p>Enter your personal details to use all of site features</p>
                    <button class="hidden" id="loginBtn">Sign In</button>
                </div>
                <div class="toggle-panel toggle-right">
                    <h1>Hello, Friend!</h1>
                    <p>Register with your personal details to use all of site features</p>
                    <button class="hidden" id="registerBtnToggle">Sign Up</button>
                </div>
            </div>
        </div>
    </div>
    `
    return temp.firstElementChild
}