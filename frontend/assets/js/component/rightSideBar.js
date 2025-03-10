export function RightSideBar(){
    return  `
    <div class="sidebar-right" id="categories">
        <h1 class="h1-title sub-main reverse">
            <div class="switch-icon rotate">
                <span class="material-symbols-outlined">
                    autorenew
                </span>
            </div>

            <div class="switch-buttons">
                <span class="main">Profile</span>
                <span class="sub">Categories</span>
            </div>
        </h1>
        
        <div class="ProfileCard">
            <div class="profileImage">
                <img src="/assets/images/profile.png" alt="">
            </div>
            <div class="profileName">
                Please Login first
            </div>
            <div class="profileStatics">
                <span class="analytics">analytics</span>
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
        <!-- <h1 class="h1-title Categories">Categories</h1> -->
        <div class="Categories display">
            <hr>
            <a class="Links" href="/?type=category&amp;category=Business">
                <div class="trending-item">
                    <div class="item-category">
                        <p>Business</p>
                    </div>
                    <span>0 Posts</span>
                </div>
            </a>

            <hr>
            <a class="Links" href="/?type=category&amp;category=Entertainment">

                <div class="trending-item">
                    <div class="item-category">
                        <p>Entertainment</p>
                    </div>
                    <span>0 Posts</span>
                </div>
            </a>

            <hr>
            <a class="Links" href="/?type=category&amp;category=General">

                <div class="trending-item">
                    <div class="item-category">
                        <p>General</p>
                    </div>
                    <span>0 Posts</span>
                </div>
            </a>

            <hr>
            <a class="Links" href="/?type=category&amp;category=Health">

                <div class="trending-item">
                    <div class="item-category">
                        <p>Health</p>
                    </div>
                    <span>0 Posts</span>
                </div>
            </a>

            <hr>
            <a class="Links" href="/?type=category&amp;category=Sports">

                <div class="trending-item">
                    <div class="item-category">
                        <p>Sports</p>
                    </div>
                    <span>0 Posts</span>
                </div>
            </a>

            <hr>
            <a class="Links" href="/?type=category&amp;category=Technology">

                <div class="trending-item">
                    <div class="item-category">
                        <p>Technology</p>
                    </div>
                    <span>0 Posts</span>
                </div>
            </a>

        </div>
    </div>`
}