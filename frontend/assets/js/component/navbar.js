export function NavBar(){
    const tmp = document.createElement('div')
    tmp.innerHTML = `
    <div class="notif">
        <a href="/"  class="selected" >
            <i class="material-symbols-outlined">home</i>
        </a>
        <a  href="#"  href="/?type=liked">
            <i class="material-symbols-outlined">favorite</i>
        </a>
        <a  id="message" class="selected" >
            <i class="material-symbols-outlined">mail</i>
        </a>
        <a href="/" class="selected">
        <i class="material-symbols-outlined">list</i>
        </a>
        <!-- <a href="/" class="disabled">
        <i class="material-symbols-outlined disabled">notifications</i>
        </a> -->
        <a href="/"  class="selected" >
            <i class="material-symbols-outlined">person</i>
        </a>
    </div> `
    return tmp.firstChild
}