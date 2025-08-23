export function LeftSideBar(){
    return `<section class="sidebar-left">
        <div class="nav-links margin-top">
            <a class="Links selected" href="/" id="home" >
                <i class="material-symbols-outlined">home</i>
                <span>Home</span>
            </a>
            <a  class="Links" id="message"  href="#">
                <i class="material-symbols-outlined">mail</i>
                <span>Messages</span>
            </a>
            <a  class="Links" href="/?type=liked" id="liked">
                <i class="material-symbols-outlined">favorite</i>
                <span>Liked Posts</span>
            </a>
            <a class="Links" href="/?type=profile" id="profile">
                <i class="material-symbols-outlined">person</i>
                <span>Profile</span>
            </a>
            <a class="Links" href="/?type=trending" id="trending">
                <i class="material-symbols-outlined">trending_up</i>
                <span>Trending</span>
            </a>
            <a href="/?type=recent" class="disabled">
                <i class="material-symbols-outlined disabled">update</i>
                <span>Recent Posts</span>
            </a>
            <button class="tweet-button new-post-header">
                <span class="material-symbols-outlined">
                edit_square
                </span>  New Post</button>
        </div>
    </section>`
}