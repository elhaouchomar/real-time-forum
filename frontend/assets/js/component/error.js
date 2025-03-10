export function ErrorPage(code, msg){
    const tmp = document.createElement('div')
    tmp.innerHTML =  `<div class="ErrorContainer">
        <img src="/assets/images/Errors.svg" alt="">
        <h1>${code}</h1>
        <h2>${msg}</h2>
        <br>
        <br>
        <a style="cursor: pointer; padding:10px; background-color:white; color: black; font-size:20px" onclick="navigation.back()">return</a>
    </div>`
    return tmp.firstChild
}