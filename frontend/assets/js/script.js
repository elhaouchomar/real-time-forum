import { apiRequest } from "./apiRequest.js";
import { CommentInputEventListenner, ExpandComments, PostButtonSwitcher } from "./comments.js";
import { HandleLikes } from "./likes.js";
import { AVATAR_URL, BodyElement, ChangeUrl, ListnerMap, LoadPage, USRNAME } from "./spa.js";

// const sidebardLeft = document.querySelector(".sidebar-left");
const windowMedia = window.matchMedia("(min-width: 768px)");

// let errorr = "";

export async function fetchPosts(offset, type, where) {
  console.log("This is called from ", where);
  
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
      `/infinite-scroll?offset=${offset}&type=${type}${category_name ? `&category=${category_name}` : ""
      }${username ? `&username=${username}` : ""}`
    );
    const posts = await response.json();
    console.log("POST +>>>>>>>", posts);

    if (posts) {
      updateProfile(posts.profile);
      updateCategoriesCount(posts.categories)
      if (posts.posts) {
        posts.posts.forEach((post) => {
          postsContainer.append(createPostCard(post));
        });
      }

    }
    HandleLikes()
    readPost();
  } catch (error) {
    console.log(error);
  }
}

function updateCategoriesCount(categoriesCount) {
  const categories = document.querySelectorAll(".Categories .trending-item span");
  categories.forEach((category) => {
    const categoryName = category.parentElement.querySelector(".item-category p").textContent.trim();
    if (categoriesCount[categoryName]) {
      category.textContent = `${categoriesCount[categoryName]} Posts`;
    }
  });
}

function updateProfile(profile) {
  const userName = profile.UserName
  const pImage = document.querySelector(".profileImage img");
  const pName = document.querySelector(".profileName");
  const pCounts = document.querySelector(".posts .postCounts");
  const cCounts = document.querySelector(".comments .postCounts");
  pImage.src = `${AVATAR_URL}${userName}`
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
  console.log("Post Details:");
  console.log(post);

  tweeterName.innerHTML = `${post.post_title
    .replace(/</g, "&lt;")
    .replace(/>/g, "&gt;")}<br>
        <span class="tweeter-handle">@${post.author_username}</span>
        <span class="material-symbols-outlined" id="timer">schedule</span>
        <span class="post-time" data-time="${post.CreatedAtString}"> ${post.CreatedAtString
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
  likeCounter.className = `counters like ${post.view && post.view === "1" ? "FILL" : ""
    }`;
  likeCounter.id = post.post_id;
  likeCounter.innerHTML = `<i class="material-symbols-outlined popup-icon" id="${post.ID}">thumb_up</i><span id="${post.post_id}">${post.like_count}</span>`;
  return likeCounter;
}

function createDislikeCounter(post) {
  const dislikeCounter = document.createElement("div");
  dislikeCounter.setAttribute("isPost", "true");
  dislikeCounter.className = `counters dislike ${post.view && post.view === "0" ? "FILL" : ""
    }`;
  dislikeCounter.id = post.post_id;
  dislikeCounter.innerHTML = `<i class="material-symbols-outlined popup-icon" id="${post.ID}">thumb_down</i><span id="${post.post_id}">${post.dislike_count}</span>`;
  return dislikeCounter;
}

function selectedItem(id) {
  console.log("Element Selected", id);
  
  const Links = document.querySelectorAll(".Links")
  Links.forEach(elem => {
    if (elem.parentElement.classList.contains("nav-links")) {
    console.log("", id);

      if (elem.id != id) {
        elem.classList.remove("selected")
      } else {
        elem.classList.add("selected")
      }
    }
  })

}
export function infiniteScroll(where) {
  console.log("infiniteScroll This is called from ", where);
  const themeToggle = document.querySelectorAll("#switch");

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
        await fetchPosts(offset, type, "infiniteScroll");

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
    elem.addEventListener("click", async (event) => {
      event.preventDefault();
      const response = await apiRequest("checker")
      if (!response || !response.status) {
        ChangeUrl("login")
        return
      }
      if (elem.id == "logout") {
        const respons = await apiRequest("logout")
        console.log(respons);
        if (respons.status) {
          console.log("Clicked on logout icon");
          LoadPage("login")
          return
        }
      }

      if (elem.id == "message") {
        document.querySelector("#area-msg").hidden = false
        setTimeout(selectedItem(elem.id), 2000)
        console.log("aciba");
        return
      }
      if (elem.id == "profile") {
        ChangeUrl(`?type=profile&username=${USRNAME}`)
        return
      }

      const LinkHref = elem.getAttribute("href")
      const params = new URLSearchParams(LinkHref.split("?")[1])
      const type = params.get("type") || "home"

      ChangeUrl(LinkHref)
      LoadPage(type, null, null, true)
      setTimeout(selectedItem(elem.id), 2000)

    })
  })

}

async function fetchPost(url) {
  try {
    const response = await fetch(url);
    if (response.status === 401) return LoadPage("login")
    if (response.status === 400) return LoadPage("error", 400, "Bad request")
    if (response.status === 500) return LoadPage("error", 500, "Internal server error")
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
    if (ListnerMap.has(elem)) {
      elem.removeEventListener('click', ListnerMap.get(elem))
    }
    elem.addEventListener("click", handler);
    ListnerMap.set(elem, handler)
  });
}

async function loadPostContent(elem) {
  console.log("====> loadPostContent CALLED <======");
  const response = await apiRequest("checker")
  if (!response || !response.status) {
    ChangeUrl("login")
    return
  }
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

const body = document.body;

function toggleDarkMode(isDark) {
  body.classList.toggle("dark-mode", isDark);
  localStorage.setItem("darkMode", isDark);
  body.classList.add("theme-transitioning");
  body.classList.remove("theme-transitioning");
}

const darkModeStored = localStorage.getItem("darkMode") === "true";
toggleDarkMode(darkModeStored);


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
      const dropDown = document.querySelector(".dropMenu")
      if (dropDown && event.target != dropDown) {
        dropDown.style.display = "none"
      }
    };

    drop.addEventListener('click', handleClick);
    document.addEventListener('click', handleClickOutside);

    ListnerMap.set(drop, handleClick);
    ListnerMap.set(document, handleClickOutside);
  })
}

async function handleClickNotify(ele) {
  ele.preventDefault()
  const response = await apiRequest("checker")
  if (!response || !response.status) {
    ChangeUrl("login")
    return
  }
  const postContainer = document.getElementById("posts");
  const messagesContainer = document.getElementById("area-msg");
  const sidebarRight = document.querySelector(".sidebar-right")
  if (ele.target.id == "home") {
    ChangeUrl("home")
  } else if (ele.target.id == "liked") {
    ChangeUrl("?type=liked")
  } else if (ele.target.id == "profile") {
    const dropMenu = document.querySelector(".dropMenu")
    if (dropMenu) {
      dropMenu.style.display = "block"
    }

  } else if (ele.target.id == "category") {
    sidebarRight.style.display = "flex"
    messagesContainer.style.display = "none"
    postContainer.style.display = "none"
  }
  console.log(ele.target, "|Im /here>D|");

}

export function NotifyButtons() {

  const notifyButtons = document.querySelectorAll(".notif a")
  console.log("inside notify", notifyButtons);
  notifyButtons.forEach(button => {
    if (ListnerMap.has(button)) {
      button.removeEventListener('click', ListnerMap.get(button))
    }
    button.addEventListener('click', handleClickNotify)
    ListnerMap.set(button, handleClickNotify)
  })

}

function handleMediaChange(event) {
  const commentSection = document.querySelector(".postComments");
  const postSection = document.querySelector(".ProfileAndPost");
  const friendsList = document.querySelector(".friends-list");
  if (event.matches) {
    friendsList.style.display = "block"
    if (postSection) {
      postSection.style.display = "flex"
      commentSection.style.display = "flex"
    }

  } else {
    commentSection.style.display = "none"
  }
}
windowMedia.addEventListener('change', handleMediaChange)