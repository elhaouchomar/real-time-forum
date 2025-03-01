import { HomePage } from "./component/homePage.js";
import { LoginPage } from "./component/loginPage.js"
import { apiRequest } from "./apiRequest.js"
import { ROUTES } from "./routes/routes.js";
const SPAContainer = document.querySelector(".SPAContainer");
const headElement = document.querySelector('head')
const BodyElement = document.querySelector("body")
export const AVATAR_URL = 'https://ui-avatars.com/api/?name=';//${userName}
export const USERNAME = 'Mohamed TEST';//${userName}
var MAIN_URL = window.location.pathname.split("/")[1]
MAIN_URL =  MAIN_URL == "" ? "home" : MAIN_URL
console.log("XXXX", MAIN_URL);

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

window.onload = async () => {
    const response = await apiRequest("checker")
    Logged = response.status
    MAIN_URL = Logged ? MAIN_URL : "login"
    console.log(`User Logged Statuse => ${Logged} --> Redirected to ${MAIN_URL}`);
    LoadPage(MAIN_URL)
}


export async function LoadPage(page = "home"){
    console.log(`Loading Page => ${page}`);
    ChangeUrl(page)
    clearSPAContainer()

    if (page == "home"){
        removeStyleElements()
        removeScriptElements()
        SPAContainer.appendChild(HomePage())
        console.log("Append HomePage");
        ROUTES[page]["scripts"].forEach(elem => {
            BodyElement.appendChild(createScript(elem, "module"))
        })
        ROUTES[page]["styles"].forEach(elem => {
            headElement.appendChild(createStyle(elem, page))
        })
        // SPAContainer.appendChild(NavBar())
        // BodyElement.appendChild(createScript("script","module"))
        // const HomePageHMTL = HomePage();
        // console.log(bodyContainer);
        
        // // bodyContainer.innerHTML = HomePageHMTL
        // BodyElement.appendChild(createScript("createPost"))
        // BodyElement.appendChild(createScript("likes"))
        // BodyElement.appendChild(createScript("script","module"))
        // BodyElement.appendChild(createScript("chat"))
    } else if (page == "login") {
        removeStyleElements();
        removeScriptElements();
        SPAContainer.appendChild(LoginPage());
        console.log("Append LoginPage");
        ROUTES[page]["scripts"].forEach(name => {
            console.log("Adding JS of Login", page);
            BodyElement.appendChild(createScript(name, "module"));
        });
        ROUTES[page]["styles"].forEach(name => {
            console.log("Adding CSS of Login", page);
            headElement.appendChild(createStyle(name, page));
        });
    }
    // await fetchPost(`http://localhost:8080/${page}`).then(
    //     elem => {
    //         htmlPage.innerHTML = elem
    //         // const htmlPage = document.querySelector('html')
    //         const bodyElement = htmlPage.querySelector('body')
    //         console.log("BODY + ", bodyElement);
            
    //         console.log("test=:>");
            
    //         return
    //     }
    // )
   
        
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