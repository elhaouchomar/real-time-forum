
export function NavBar(){
    return `
    <div class="notif">
        <a  class="selected">
            <i  id="home" class="material-symbols-outlined">home</i>
        </a>
        <a href="/?type=liked">
            <i id="liked" class="material-symbols-outlined">favorite</i>
        </a>
        <a  id="message" class="selected">
            <i class="material-symbols-outlined">mail</i>
        </a>
        <a class="selected"  id="category">
        <i id="category" class="material-symbols-outlined">list</i>
        </a>
        <a class="selected" >
            <i id="profile" class="material-symbols-outlined">person</i>
            
        </a>
        <div class="dropMenu">
                <div class="Links" id="profile">Profile</div>
                <div class="dark">
                <li class="theme-toggle">
                    <input id="switch" type="checkbox">
                    <label for="switch">
                        Theme
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
            <div class="Links" id="logout">Logout</div>
        </div>
    </div> `
}