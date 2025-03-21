export function RightSideBar(){
    return  `
    <div class="sidebar-right " id="categories">
        <h1 class="h1-title sub-main ">
            <div class="switch-icon">
                <span class="material-symbols-outlined">
                    autorenew
                </span>
            </div>

            <div class="switch-buttons">
                <span class="sub">Categories</span>
                <span class="main">Messages</span>
            </div>
        </h1>
        
        <div class="MessageCard">
            <div class="friends-list-right">
                <div class="allfriends"></div>
            </div> 
        </div>
        <div class="Categories ">
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