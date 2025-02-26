export function LeftSideBar(){
    return `<section class="sidebar-left">
        <div class="nav-links margin-top">
            <a class="Links seelcted" href="/">
                <i class="material-symbols-outlined">home</i>
                <span>Home</span>
            </a>
            <a  class="Links" id="message" class="selected" >
                <i class="material-symbols-outlined">mail</i>
                <span>Messages</span>
            </a>
            <a  class="Links" href="/?type=liked">
                <i class="material-symbols-outlined">favorite</i>
                <span>Liked Posts</span>
            </a>
            <a class="Links" href="/?type=profile">
                <i class="material-symbols-outlined">person</i>
                <span>Profile</span>
            </a>
            <a class="Links" href="/?type=trending" class="selected">
                <i class="material-symbols-outlined">trending_up</i>
                <span>Trending</span>
            </a>
            <a href="/?type=recent" class="disabled">
                <i class="material-symbols-outlined disabled">update</i>
                <span>Recent Posts</span>
            </a>
            <!-- TODO For Omar => this part should only work on mobiles, and hide the other one on header --> 
            <div class="theme-toggle-mobile">
                <ul>
                    <li class="theme-toggle theme-toggle-mobile">
                        <input id="switch" type="checkbox">
                        <label for="switch">
                            <div class="toggle">
                                <span class="material-symbols-outlined dark_mode">
                                    dark_mode
                                </span>
                                <span class="material-symbols-outlined light_mode">
                                    light_mode
                                </span>
                            </div>
                            <span class="theme">Theme</span>
                        </label>
                    </li>
                </ul>
            </div>
            <button class="tweet-button new-post-header">
                <span class="material-symbols-outlined">
                edit_square
                </span>  New Post</button>
        </div>
    </section>`
}