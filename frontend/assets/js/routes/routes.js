import { createPostListner } from "../createPost.js";
import { fetchPosts, infiniteScroll, postControlList, readPost } from "../script.js";

export const ROUTES = {
    "home": {
        "scripts": ["spa", "profile", "likes", "script", "chat"],
        "styles": ["root", "create_post", "header", "left_sidebar",
                    "message", "popstyle", "post", "right_sidebar", "style", "ussely_by_js"],
    },
    "login": {
        "scripts": ["spa", "register"],
        "styles": ["register"],

    },
    "error":{
        "styles": ["error"],
    }
}
//     <link rel="stylesheet" href="/assets/style/root.css" />
{/* <link rel="stylesheet" href="/assets/style/create_post.css" />
<link rel="stylesheet" href="/assets/style/header.css" />
<link rel="stylesheet" href="/assets/style/left_sidebar.css" />
<link rel="stylesheet" href="/assets/style/message.css" />
<link rel="stylesheet" href="/assets/style/popstyle.css" />
<link rel="stylesheet" href="/assets/style/post.css" />
<link rel="stylesheet" href="/assets/style/register.css" />
<link rel="stylesheet" href="/assets/style/right_sidebar.css" />
<link rel="stylesheet" href="/assets/style/style.css" />
<link rel="stylesheet" href="/assets/style/ussely_by_js.css" /> */}