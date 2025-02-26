import { Header } from "./header.js"
import { LeftSideBar } from "./leftSideBar.js"
import { Messages } from "./messages.js"
import { RightSideBar } from "./rightSideBar.js"

export function HomePage() {
    const HeaderFile = Header()
    const RightSideBarFile = RightSideBar()
    const LeftsideBarFile = LeftSideBar()
    const MessagesFile = Messages()
    const tmp = document.createElement('div')
    const userName = "Mohamed Tawil"
    tmp.innerHTML = `
    
    <!--TODO PostContainer to be Enabled <div class="postContainer closed"></div> -->
    <div class="ParentContainer">
      ${HeaderFile}
      <!-- <div class="nav-mobile" id="right-side">
        <a href="/#posts">
          <div>For you</div>
        </a>
        <a href="#categories">
          <div>Categories</div>
        </a>
      </div> -->
        ${LeftsideBarFile}
        ${RightSideBarFile}
        ${MessagesFile}
      <div class="main-flex" id="posts">
        <div class="main-feed">
          <!-- Create New Post -->
          <div class="new-tweet">
            <!-- TODO User image -->
            <div
              class="ProfileImage tweet-img no-border"
              style="background-image: url('https://ui-avatars.com/api/?name=${userName}')"
            ></div>
            <div class="new-post-header">
              <div class="textarea">What's happening?</div>
            </div>
          </div>
          <!-- End Of Create New Post -->
        </div>
      </div>
    </div>
    `
    return tmp.firstElementChild
}