import { HomePage } from "./component/homePage.js";
import { LoginPage } from "./component/loginPage.js"
import { apiRequest } from "./apiRequest.js"
import { ROUTES } from "./routes/routes.js";
import { fetchPosts, infiniteScroll, postControlList, readPost } from "./script.js";
import { registerFunctions } from "./register.js";
import {createPostListner} from "./createPost.js"
import { ErrorPage } from "./component/error.js";

export const SPAContainer = document.querySelector(".SPAContainer");
const headElement = document.querySelector('head')
export const BodyElement = document.querySelector("body")
export const AVATAR_URL = 'https://ui-avatars.com/api/?name=';//${userName}
export const USERNAME = null
export const ListnerMap = new WeakMap()

var MAIN_URL = window.location.pathname.split("/")[1]
MAIN_URL =  MAIN_URL == "" ? "home" : MAIN_URL
console.log("XXXX", MAIN_URL);


window.onload = async () => {
    const response = await apiRequest("checker")
    Logged = response.status
    MAIN_URL = Logged ? MAIN_URL : "login"
    console.log(`User Logged Statuse => ${Logged} --> Redirected to ${MAIN_URL}`);
    LoadPage(MAIN_URL)
}

var Logged = false

function clearSPAContainer(){
    SPAContainer.innerHTML = ""
    console.log(`Clear Main Container`);
}

function createScript(src, type = "text/javascript"){
    console.log(`Create JS Script File = ${src} - Type = ${type}`);
    const temp = document.createElement("div")
    const mainScript = document.createElement('script')
    mainScript.src = `/assets/js/${src}.js`
    mainScript.type = type
    mainScript.id = src
    temp.append(mainScript) 
    return temp.firstChild
}

function createStyle(src, page){
    console.log(`Create Css Style File = ${src} - page = ${page}`);
    const temp = document.createElement("div")
    const styleElement = document.createElement('link')
    styleElement.rel = "stylesheet"
    styleElement.href = `/assets/style/${src}.css`
    styleElement.id = `${src}-${page}`
    temp.append(styleElement) 
    return temp.firstChild
}


export async function LoadPage(page = "home", code, msg, skip = false){
    console.log(`Loading Page => ${page}`);
    if (!skip) ChangeUrl(page)
    clearSPAContainer()

    if (page == "home" || page == "category" ||
         page == "trending" || page == "profile" || page == "liked"){
        removeStyleElements()
        SPAContainer.appendChild(HomePage())
        console.log("Append HomePage");
        infiniteScroll()
        fetchPosts(0, page)
        postControlList()
        readPost()
        createPostListner()
        ROUTES["home"]["styles"].forEach(elem => {
            headElement.appendChild(createStyle(elem, page))
        })
    } else if (page == "login") {
        removeStyleElements();
        SPAContainer.appendChild(LoginPage());
        console.log("Append LoginPage");
        registerFunctions()
        ROUTES[page]["styles"].forEach(name => {
            console.log("Adding CSS of Login", page);
            headElement.appendChild(createStyle(name, page));
        });
    }else if (page == "error"){
        SPAContainer.appendChild(ErrorPage(code, msg));
    }else{
        removeStyleElements();
        SPAContainer.appendChild(ErrorPage(404, "Page Not Found"));
        headElement.appendChild(createStyle(ROUTES["error"]["styles"][0], "error"));
    }
}

function removeStyleElements(){
    console.log(`Removing all Style Elements`);
    const headStyles = document.head.querySelectorAll('link')
        headStyles.forEach(elem => {
            if (elem.id){
                elem.remove()
            }
        })
}
function removeScriptElements(){
    console.log(`Removing all ScriptJs Elements`);
    const JsFiles = document.querySelectorAll('script')
        JsFiles.forEach(elem => {
            if (elem.id){
                elem.remove()
            }
        })
}


export function ChangeUrl(url, data = {}) {
    history.pushState(data, "", url)
    console.log("URL Changed to =>", url);
    // LoadPage("home")
}

navigation.addEventListener("navigate", (event) => {
    const Url = new URL(event.destination.url)
    const params = new URLSearchParams(Url.search)
    var type = params.get("type")
    var category = params.get("category")
    if (!type){
        type = Url.pathname.split("/")[1]
    }
    LoadPage(type, null, null, true)
    console.log("=====================================")
    console.log("=====================================")
    console.log("  Any changes ",event.destination.url, event.pushState )
    console.log("=====================================")
    console.log("=====================================")

})