import { HomePage } from "./component/homePage.js";
import { LoginPage } from "./component/loginPage.js"
import { NavBar } from "./component/navbar.js";
import { MiddleWear } from "./middleWear.js"
const SPAContainer = document.querySelector(".SPAContainer");
const headElement = document.querySelector('head')
const BodyElement = document.querySelector("body")

var Logged = false

function clearSPAContainer(){
    SPAContainer.innerHTML = ""
    console.log(`Cleared SPA Container Sucessfuly`);
}

function createScript(src, type = "text/javascript"){
    const temp = document.createElement("div")
    const mainScript = document.createElement('script')
    mainScript.src = `/assets/js/${src}.js`
    mainScript.type = type
    mainScript.id = src
    temp.append(mainScript) 
    return temp.firstChild
}

window.onload = async () => {
    const response = await MiddleWear()
    Logged = response.status
    const Page = Logged ? "home" : "login"
    console.log(`User Logged Statuse => ${Logged} --> Redirected to ${Page}`);
    LoadPage(Page)
}


export async function LoadPage(page = "home"){
    console.log(`Loading Page => ${page}`);

    // ChangeUrl(page)
    // clearSPAContainer()
    // SPAContainer.append(LoginPage())
    // const registerScript = createScript("register")
    // BodyElement.append(registerScript)
    clearSPAContainer()

    if (page == "home"){
        removeStyleElements("home")
        console.log("----=>",NavBar());
        SPAContainer.appendChild(HomePage())
        // SPAContainer.appendChild(NavBar())
        // BodyElement.appendChild(createScript("script","module"))
        // const HomePageHMTL = HomePage();
        // console.log(bodyContainer);
        
        // // bodyContainer.innerHTML = HomePageHMTL
        BodyElement.appendChild(createScript("profile"))
        BodyElement.appendChild(createScript("createPost"))
        BodyElement.appendChild(createScript("likes"))
        BodyElement.appendChild(createScript("script","module"))
        // BodyElement.appendChild(createScript("chat"))
    }else if (page == "login"){
        removeStyleElements("register")
        SPAContainer.appendChild(LoginPage())
        BodyElement.appendChild(createScript("register","module"))
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

function removeStyleElements(id){
    console.log(`Removing all Style Elements Except => ${id}`);
    const headStyles = document.head.querySelectorAll('link')
        headStyles.forEach(elem => {
            if (elem.id && elem.id != id){
                elem.remove()
            }
        })
}

