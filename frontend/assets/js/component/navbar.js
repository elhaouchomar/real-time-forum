export function NavBar(){
    return `
    <div class="notif">
        <a href="/"  class="selected">
            <i  id="home" class="material-symbols-outlined">home</i>
        </a>
        <a  href="#"  href="/?type=liked">
            <i id="liked" class="material-symbols-outlined">favorite</i>
        </a>
        <a  id="message" class="selected">
            <i class="material-symbols-outlined">mail</i>
        </a>
        <a href="/" class="selected"  id="category">
        <i id="category" class="material-symbols-outlined">list</i>
        </a>
        <a href="/"  class="selected" >
            <i id="profile" class="material-symbols-outlined">person</i>
        </a>
    </div> `
}