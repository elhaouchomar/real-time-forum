import { AVATAR_URL, USRNAME } from "../spa.js"
import { Header } from "./header.js"
import { LeftSideBar } from "./leftSideBar.js"
import { Messages } from "./messages.js"
import { NavBar } from "./navbar.js"
import { RightSideBar } from "./rightSideBar.js"

export function HomePage(page) {
    let profileStatics = ""
    if (page == "profile"){
      profileStatics = `
      <div class="ProfileCard">
            <div class="profileStatics">
                <span class="analytics profileName">analytics</span>
                <div class="posts">
                    <span class="material-symbols-outlined">
                        article
                    </span>
                        
                        <span class="postCounts"><a href="/login">Login</a></span>
                </div>
                <div class="comments">
                    <span class="material-symbols-outlined">
                        comment
                    </span>
                    <span class="postCounts"><a href="/register">Register</a></span>
                </div>
            </div>
        </div>
        `
    }
    let createPost = ""
    const url = new URL(window.location.href);
    const params = new URLSearchParams(url.search)
    const profileName = params.get("username") || USRNAME
    if (USRNAME == profileName){
      createPost = `
      <div class="new-tweet">
            <div
              class="ProfileImage tweet-img no-border"
              style="background-image: url('${AVATAR_URL}${USRNAME}')"
            ></div>
            <div class="new-post-header">
              <div class="textarea">What's happening?</div>
            </div>
        </div>`
    }
    const HeaderFile = Header()
    const RightSideBarFile = RightSideBar()
    const LeftsideBarFile = LeftSideBar()
    const MessagesFile = Messages()
    const tmp = document.createElement('div')
    const NavBarFile = NavBar()
    tmp.innerHTML = `
    
    <div class="ParentContainer">
        ${HeaderFile}
        ${LeftsideBarFile}
        ${RightSideBarFile}
        ${MessagesFile}
        ${NavBarFile}
        <div class="main-flex" id="posts">
        <div class="main-feed">
          <!-- Create New Post -->
          ${createPost}
          <!-- End Of Create New Post -->
          ${profileStatics}
        </div>
      </div>
    </div>
    `
    return tmp.firstElementChild
}