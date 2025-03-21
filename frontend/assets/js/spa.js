import { HomePage } from "./component/homePage.js";
import { LoginPage } from "./component/loginPage.js"
import { apiRequest } from "./apiRequest.js"
import { ROUTES } from "./routes/routes.js";
import { NotifyButtons, fetchPosts, infiniteScroll, postControlList, readPost } from "./script.js";
import { registerFunctions } from "./register.js";
import {createPostListner} from "./createPost.js"
import { ErrorPage } from "./component/error.js";
import { connectWebSocket, initChat } from "./chat.js";
import { profileEffect } from "./profiles.js";


export const SPAContainer = document.querySelector(".SPAContainer");
const headElement = document.querySelector('head')
export const BodyElement = document.querySelector("body")
export const AVATAR_URL = 'https://ui-avatars.com/api/?name=';//${userName}
export let USRNAME  = ""
export let ID  = 0
export const ListnerMap = new WeakMap()
export let previousUrl = document.location.href;
export var Logged = false
var MAIN_URL = window.location.href.split("/")[3]

MAIN_URL =  MAIN_URL == "" ? "home" : MAIN_URL
window.onload = async () => {
    const response = await apiRequest("checker")
    if (response){
        console.log("Response of Checker" , response);
        Logged = response.status
        if(Logged){
            USRNAME =  response.data.UserName
            ID = response.data.ID
        }
        console.log("user Name ", response.data);
   }
   MAIN_URL = Logged ? MAIN_URL : "login"
   console.log(`User Logged Statuse => ${Logged} --> Redirected to ${MAIN_URL}`);
   ChangeUrl(MAIN_URL)
}


function clearSPAContainer(){
    SPAContainer.innerHTML = ""
    console.log(`Clear Main Container`);
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

    const response = await apiRequest("checker")
    if (!response || !response.status){
       page = "login"
    }else{
        USRNAME = response.data.UserName
        Logged = response.status
        ID = response.data.ID
    }

    if (!skip) ChangeUrl(page)
    clearSPAContainer()

    if (page == "home" || page == "category" ||
         page == "trending" || page == "profile" || page == "liked"){
        removeStyleElements()
        ROUTES["home"]["styles"].forEach(elem => {
            headElement.appendChild(createStyle(elem, page))
        })
        SPAContainer.appendChild(HomePage(page))
        console.log("Append HomePage");
        infiniteScroll("spa")
        fetchPosts(0, page, "LoadPage")
        postControlList()
        readPost()
        profileEffect()
        createPostListner()
        NotifyButtons()
        connectWebSocket()
        initChat()
    } else if (page == "login") {
        if (response.status){
            LoadPage("home")
            return
        }
        Logged = false
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


export function ChangeUrl(url, data = {}) {
    console.log("URL Changed to =>", url);
    history.pushState(data, "", url)
    window.dispatchEvent(new PopStateEvent("popstate"));

    // LoadPage("home")
}

window.addEventListener("popstate", async (event) => {
    previousUrl = document.location.href;
    console.log('Previous URL: ', previousUrl);
    const Url = new URL(document.location.href)
    const params = new URLSearchParams(Url.search)
    var type = params.get("type")
    if (!type){
        const tmp = Url.pathname.split("/")[1]
        type =  tmp == "" ? "home" : tmp
    }
    console.log("=====================================", type)

    LoadPage(type, null, null, true)
    console.log("=====================================")
    console.log("=====================================")
    console.log("=====================================")
    console.log("=====================================")
    
})