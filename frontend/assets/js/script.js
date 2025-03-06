import { apiRequest } from "./apiRequest.js";
import { CommentInputEventListenner, ExpandComments } from "./comments.js";
import { HandleLikes } from "./likes.js";
import { AVATAR_URL, BodyElement, ChangeUrl, ListnerMap, LoadPage } from "./spa.js";

// const sidebardLeft = document.querySelector(".sidebar-left");
const windowMedia = window.matchMedia("(min-width: 768px)");

// let errorr = "";

export async function fetchPosts(offset, type) {
  const UrlParams = new URLSearchParams(window.location.search);
  var type = UrlParams.get("type");
  const username = UrlParams.get("username");
  console.log("====> fetchPosts CALLED <======", username);

  type = type || "home";
  let category_name = UrlParams.get("category");
  console.log("Category name ", category_name);
  console.log("type name ", type, window.location.search);
  
  const postsContainer = document.querySelector(".main-feed");
  try {
    const response = await fetch(
      `/infinite-scroll?offset=${offset}&type=${type}${
        category_name ? `&category=${category_name}` : ""
      }${username ? `&username=${username}` : ""}`
    );
    const posts = await response.json();
    console.log("POST +>>>>>>>", posts);
    
    if (posts) {
      updateProfile(posts.profile);
      posts.posts.forEach((post) => {
        postsContainer.append(createPostCard(post));
      });
    }
    HandleLikes()
    readPost();
  } catch (error) {
    console.log(error);
  }
}

function updateProfile(profile) {
  const userName = profile.UserName
  const pImage = document.querySelector(".profileImage img");
  const pName = document.querySelector(".profileName");
  const pCounts = document.querySelector(".posts .postCounts");
  const cCounts = document.querySelector(".comments .postCounts");
  pImage.src =`${AVATAR_URL}${userName}`
  pName.textContent = userName
  pCounts.textContent = `${profile.ArticleCount} Articles`;
  cCounts.textContent = `${profile.CommentCount} Comments`;
}

function createPostCard(post) {
  const postCard = document.createElement("div");
  postCard.classList.add("post-card");
  postCard.append(
    createProfileLink(post.author_username),
    createPostDetails(post)
  );
  return postCard;
}

function createProfileLink(username) {
  const profileLink = document.createElement("a");
  profileLink.href = `/?type=profile&username=${username}`;
  const profileImage = document.createElement("div");
  profileImage.className = "ProfileImage tweet-img";
  profileImage.style.backgroundImage = `url('${AVATAR_URL}${username}')`;
  profileLink.appendChild(profileImage);
  return profileLink;
}

function createPostDetails(post) {
  const postDetails = document.createElement("div");
  postDetails.className = "post-details";
  postDetails.append(
    createRowTweet(post),
    createPostContent(post),
    createHashtag(post),
    createPostFooter(post)
  );
  return postDetails;
}

function createRowTweet(post) {
  const rowTweet = document.createElement("div");
  rowTweet.className = "row-tweet";
  const postHeader = document.createElement("div");
  postHeader.className = "post-header";
  const tweeterName = document.createElement("span");
  tweeterName.className = "tweeter-name post";
  tweeterName.id = post.post_id;
  tweeterName.innerHTML = `${post.post_title
    .replace(/</g, "&lt;")
    .replace(/>/g, "&gt;")}<br>
        <span class="tweeter-handle">@${post.author_username}</span>
        <span class="material-symbols-outlined" id="timer">schedule</span>
        <span class="post-time" data-time="${post.post_creation_time}"> ${
    post.post_creation_time
  }</span>`;
  postHeader.appendChild(tweeterName);
  rowTweet.append(postHeader);
  return rowTweet;
}

function createPostContent(post) {
  const postContent = document.createElement("div");
  const postParagraph = document.createElement("p");
  postContent.className = "post-content";
  postParagraph.innerText = post.post_content;
  postContent.appendChild(postParagraph);
  return postContent;
}

function createHashtag(post) {
  const hashtag = document.createElement("div");
  hashtag.className = "Hashtag";
  if (post.post_categories) {
    post.post_categories.forEach((category) => {
      const categoryLink = document.createElement("a");
      categoryLink.href = "/?type=category&&category=" + category;
      categoryLink.innerHTML = `<span>#${category}</span>`;
      hashtag.appendChild(categoryLink);
    });
  }
  return hashtag;
}

function createPostFooter(post) {
  const postFooter = document.createElement("div");
  postFooter.className = "post-footer";
  const react = document.createElement("div");
  react.className = "react";
  react.id = post.ID;
  react.append(createLikeCounter(post), createDislikeCounter(post));
  const comment = document.createElement("div");
  comment.className = "comment post";
  comment.id = post.post_id;
  comment.innerHTML = `<i class="material-symbols-outlined showCmnts">comment</i><span id="${post.post_id}">${post.comment_count}</span>`;
  postFooter.append(react, comment);
  return postFooter;
}

function createLikeCounter(post) {
  const likeCounter = document.createElement("div");
  likeCounter.setAttribute("isPost", "true");
  likeCounter.className = `counters like ${
    post.view && post.view === "1" ? "FILL" : ""
  }`;
  likeCounter.id = post.post_id;
  likeCounter.innerHTML = `<i class="material-symbols-outlined popup-icon" id="${post.ID}">thumb_up</i><span id="${post.post_id}">${post.like_count}</span>`;
  return likeCounter;
}

function createDislikeCounter(post) {
  const dislikeCounter = document.createElement("div");
  dislikeCounter.setAttribute("isPost", "true");
  dislikeCounter.className = `counters dislike ${
    post.view && post.view === "0" ? "FILL" : ""
  }`;
  dislikeCounter.id = post.post_id;
  dislikeCounter.innerHTML = `<i class="material-symbols-outlined popup-icon" id="${post.ID}">thumb_down</i><span id="${post.post_id}">${post.dislike_count}</span>`;
  return dislikeCounter;
}

export function infiniteScroll() {
    console.log("====> infiniteScroll CALLED <======");

    var UrlParams = new URLSearchParams(window.location.search);
    const type = UrlParams.get("type");
    
    let offset = 10;
    let timeout = null;
    window.addEventListener("scroll", () => {
      clearTimeout(timeout);
      timeout = setTimeout(async () => {
        const { scrollTop, scrollHeight, clientHeight } =
          document.documentElement;
        if (scrollTop + clientHeight >= scrollHeight - 5) {
          await fetchPosts(offset, type);
        }
      }, 1000);
    });

    themeToggle.forEach((elem) => {
      elem.checked = darkModeStored;
      elem.addEventListener("change", () => {
        toggleDarkMode(elem.checked);
      });
    });

    const Links = document.querySelectorAll(".Links")
    Links.forEach(elem => {
        elem.addEventListener("click", async (event)=>{
            event.preventDefault();
            if (elem.id == "logout") {
              const respons = await apiRequest("logout")
              console.log(respons);
              if (respons.status){
                console.log("Clicked on logout icon");
                LoadPage("login")
                return
              }
            }
            if (elem.id == "message"){
              elem.hidden = false
              return
            }
            const LinkHref = elem.getAttribute("href")
            const params = new URLSearchParams(LinkHref.split("?")[1])
            const type = params.get("type") || "home"
            
            ChangeUrl(LinkHref)
            LoadPage(type, null, null, true)            
            
            //     <a class="CategoriesLinks" href="/?type=category&category={{$key}}">
            //     <!-- TODO // selected-category /// Class For Specific Selected once -->
            //     <div class="trending-item">
            //         <div class="item-category">
            //             <p>{{$key}}</p>
            //         </div>
            //         <span>{{$value}} Posts</span>
            //     </div>
            // </a>
        })
    })

}

async function fetchPost(url) {
  try {
    const response = await fetch(url);
    if (response.status === 401) return LoadPage("login")
    if (response.status != 200) {
      LoadPage("error", 404, "Page Not Found")
      return false;
    }
    return await response.text();
  } catch (error) {
    errorr = error;
  }
}


export function readPost() {
  console.log("====> readPost CALLED <======");
  
  document.querySelectorAll(".post").forEach((elem) => {
    const handler = () => loadPostContent(elem)
    if (ListnerMap.has(elem)){
      elem.removeEventListener('click', ListnerMap.get(elem))
    }
    elem.addEventListener("click", handler);
    ListnerMap.set(elem, handler)
  });
}

async function loadPostContent(elem) {
  console.log("====> loadPostContent CALLED <======");
  
  
    const html = await fetchPost(`/post/${elem.id}`);
    if (!html) return;
    console.log("Post content :", html);
    
    const postContent = document.createElement("div");
    postContent.classList.add("postContainer")
    postContent.innerHTML = html;
    BodyElement.appendChild(postContent)
    BodyElement.classList.add("stop-scrolling");
    CommentInputEventListenner()
    ExpandComments()

    const handleClick = (event) => {
      if (
      event.target == postContent ||
      event.target.classList.contains("close-post")
      ) {
      ExpandComments(false);
      postContent.innerHTML = "";
      postContent.classList.add("closed");
      document.body.classList.remove("stop-scrolling");
      if (document.getElementById("ScriptInjected"))
        document.getElementById("ScriptInjected").remove();
      }
    };

    if (ListnerMap.has(document)) {
      document.removeEventListener("click", ListnerMap.get(document));
    }

    document.addEventListener("click", handleClick);
    ListnerMap.set(document, handleClick);
    postContent.classList.remove("closed");
    ListenOncommentButtom();
    HandleLikes();
}

function DisplayPost() {
  const commentSection = document.querySelector(".postComments");
  const postSection = document.querySelector(".ProfileAndPost");
  if (!windowMedia.matches) {
    commentSection.style.display = "flex";
    postSection.style.display = "none";
  }
  PostButtonSwitcher();
}

function ListenOncommentButtom() {
  const commentButton = document.querySelector(".CommentButton");
  if (ListnerMap.has(commentButton)) {
    commentButton.removeEventListener("click", ListnerMap.get(commentButton));
  } 
  commentButton.addEventListener("click", DisplayPost);
  ListnerMap.set(commentButton)
}

const themeToggle = document.querySelectorAll("#switch");
const body = document.body;

function toggleDarkMode(isDark) {
  body.classList.toggle("dark-mode", isDark);
  localStorage.setItem("darkMode", isDark);
  body.classList.add("theme-transitioning");
  body.classList.remove("theme-transitioning");
}

const darkModeStored = localStorage.getItem("darkMode") === "true";
toggleDarkMode(darkModeStored);



function timeAgo(date) {
  const seconds = Math.floor((new Date() - new Date(date)) / 1000);
  if (seconds < 60) return "just now";
  if (seconds < 3600) return `${Math.floor(seconds / 60)}m ago`;
  if (seconds < 86400) return `${Math.floor(seconds / 3600)}h ago`;
  if (seconds < 604800) return `${Math.floor(seconds / 86400)}d ago`;
  if (seconds < 2592000) return `${Math.floor(seconds / 604800)}w ago`;
  if (seconds < 31536000) return `${Math.floor(seconds / 2592000)}mo ago`;
  return `${Math.floor(seconds / 31536000)}y ago`;
}

function updateAllTimes() {
  const timeElements = document.querySelectorAll(
    ".post-time, .commentTime, .postDate"
  );
  timeElements.forEach((el) => {
    if (el.dataset.time) {
      el.textContent = timeAgo(el.dataset.time);
    }
  });
}

const observer = new MutationObserver(() => {
  updateAllTimes();
});

document.addEventListener("DOMContentLoaded", () => {
  updateAllTimes();
  observer.observe(document.body, {
    childList: true,
    subtree: true,
  });
  setInterval(updateAllTimes, 50000);
});

//////////////// Start Listning dropDown List For Posts ////////////
export function postControlList() {
  console.log("====> postControlList CALLED <======");

  const dropdown = document.querySelectorAll('.dropdown i, .dropdown .ProfileImage')
  dropdown.forEach(drop => {
      
      if (ListnerMap.has(drop)) {
          drop.removeEventListener('click', ListnerMap.get(drop));
          document.removeEventListener('click', ListnerMap.get(drop));
      }

      let contentSibling = drop.nextElementSibling;
      const handleClick = () => {
          contentSibling.classList.toggle("show");
      };
      const handleClickOutside = (event) => {
          if (!contentSibling.contains(event.target) && !drop.contains(event.target) && contentSibling.classList.contains("show")) {
          contentSibling.classList.remove('show');
          }
      };

      drop.addEventListener('click', handleClick);
      document.addEventListener('click', handleClickOutside);

      ListnerMap.set(drop, handleClick);
      ListnerMap.set(document, handleClickOutside);
  })
}
