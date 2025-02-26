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
                        style="background-image: url('https://api.multiavatar.com/{{.UserName}}.svg')">
                    </div>
                    <div class="content">
                        <!-- Control Posts -->
                        <ul>
                            <a href="/?type=profile&username={{.UserName}}">
                                <li><span class="material-symbols-outlined">account_circle</span>Profile</li>
                            </a>
                            <!-- Darck && light mode -->
                            <hr>
                            <a href="/logout">
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