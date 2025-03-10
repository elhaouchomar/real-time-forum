import { AVATAR_URL, USERNAME } from "../spa.js";

export function Header(){
    return    `<nav>
        <!-- Logo -->
        <a href="/" class="Logo">
            <span class="Name">Edu<span>Talks</span>
        </a>
        <div class="RightHeaderSection">
            <div class="main-account">
                <div class="dropdown">
                    <div class="ProfileImage tweet-img"
                        style="background-image: url('${AVATAR_URL}${USERNAME}')">
                    </div>
                    <div class="content">
                        <!-- Control Posts -->
                        <ul>
                            <a class="Links" id="profile" href="/type=profile?username=${USERNAME}">
                                <li><span class="material-symbols-outlined">account_circle</span>Profile</li>
                            </a>
                            <!-- Darck && light mode -->
                            <hr>
                            <a class="Links" id="logout" href="/logout">
                                <li><span class="material-symbols-outlined">logout</span>Logout</li>
                            </a>
                        </ul>
                    </div>
                </div>
            </div>
            <div class="dark">
            <li class="theme-toggle">
                <input id="switch" type="checkbox">
                <label for="switch">
                    <div class="toggle">
                        <span class="material-symbols-outlined dark_mode">
                            dark_mode
                        </span>
                        <span class="material-symbols-outlined light_mode">
                            light_mode
                        </span>
                    </div>
                </label>
            </li>
        </div>
        </div>
    </nav>`
}