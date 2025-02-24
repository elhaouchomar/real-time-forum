const UrlParams = new URLSearchParams(window.location.search);
const sidebardLeft = document.querySelector(".sidebar-left");
const windowMedia = window.matchMedia("(min-width: 768px)");
const type = UrlParams.get("type");
const username = UrlParams.get("username");
let errorr = "";

async function fetchPosts(offset, type) {
  type = type || "home";
  let category_name = UrlParams.get("category");
  const postsContainer = document.querySelector(".main-feed");
  try {
    const response = await fetch(
      `/infinite-scroll?offset=${offset}&type=${type}${
        category_name ? `&category=${category_name}` : ""
      }${username ? `&username=${username}` : ""}`
    );
    const posts = await response.json();
    if (posts) {
      updateProfile(posts.profile);
      posts.posts.forEach((post) => {
        postsContainer.append(createPostCard(post));
      });
    }
    handleLikes();
    removeReadPostListener();
    readPost();
  } catch (error) {
    console.log(error);
  }
}

function updateProfile(profile) {
  const pImage = document.querySelector(".profileImage img");
  const pName = document.querySelector(".profileName");
  const pCounts = document.querySelector(".posts .postCounts");
  const cCounts = document.querySelector(".comments .postCounts");
  pImage.src = profile.UserName
    ? `https://api.multiavatar.com/${profile.UserName}.svg`
    : "/assets/images/profile.png";
  pName.textContent = profile.UserName
    ? profile.UserName
    : "Please Login First";
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
  profileImage.style.backgroundImage = `url('https://api.multiavatar.com/${username}.svg')`;
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
  comment.innerHTML = `<i class="material-symbols-outlined showCmnts">comment</i><span>${post.comment_count}</span>`;
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

function infiniteScroll() {
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
}

async function fetchPost(url) {
  try {
    const response = await fetch(url);
    if (response.status != 200) {
      window.location.href = "/error?code=404&message=Page Not Found";
      return false;
    }
    return await response.text();
  } catch (error) {
    errorr = error;
  }
}

function removeReadPostListener() {
  document.querySelectorAll(".post").forEach((elem) => {
    elem.removeEventListener("click", loadPostContent(elem));
  });
}

function readPost() {
  document.querySelectorAll(".post").forEach((elem) => {
    elem.addEventListener("click", loadPostContent(elem));
  });
}

function loadPostContent(elem) {
  return async () => {
    const html = await fetchPost(`/post/${elem.id}`);
    if (!html) return;
    const postContent = document.querySelector(".postContainer");
    postContent.innerHTML = html;
    if (!document.getElementById("ScriptInjected")) {
      const script = document.createElement("script");
      script.id = "ScriptInjected";
      script.src = "/assets/js/comments.js";
      document.body.appendChild(script);
    }
    document.body.classList.add("stop-scrolling");
    document.addEventListener("click", (event) => {
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
    });
    postContent.classList.remove("closed");
    ListenOncommentButtom(false);
    ListenOncommentButtom(true);
    handleLikes();
  };
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

function ListenOncommentButtom(add) {
  const commentButton = document.querySelector(".CommentButton");
  if (add) {
    commentButton.addEventListener("click", DisplayPost);
  } else {
    commentButton.removeEventListener("click", DisplayPost);
  }
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

themeToggle.forEach((elem) => {
  elem.checked = darkModeStored;
  elem.addEventListener("change", () => {
    toggleDarkMode(elem.checked);
  });
});

function handleNavMobileClick() {
  document.querySelectorAll(".nav-mobile a div").forEach(function (div) {
    div.addEventListener("click", function () {
      document.querySelectorAll(".nav-mobile a div").forEach(function (item) {
        item.classList.remove("clicked");
      });
      this.classList.add("clicked");
    });
  });
}

window.addEventListener("load", () => {
  if (!window.location.hash) {
    window.location.hash = "#posts";
  }
});

window.addEventListener("hashchange", () => {
  const hash = window.location.hash;
  document.querySelectorAll("#posts, #categories").forEach((section) => {
    section.style.display = hash === "#posts" ? "block" : "none";
  });
});

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
  setInterval(updateAllTimes, 60000);
});

infiniteScroll();
fetchPosts(0, type);
readPost();
handleNavMobileClick();

// static/js/chat.js
// websockets
let ws;

function connectWebSocket() {
    ws = new WebSocket("ws://localhost:8080/ws"); // Correct WebSocket URL with quotes
    ws.onopen = () => {
        console.log("Connected to chat");
    };
    ws.onmessage = (event) => {
        console.log("----------------------", event.data);
        const messagesContainer = document.querySelector(".messages-list");
        const messageElement = document.createElement("div");
        messageElement.className = "message-card";
        messageElement.innerHTML = `
            <div class="message-header" style="display: flex; justify-content: space-between">
                <span class="message-author">Unknown Author</span>
                <span class="message-time">${new Date().toLocaleTimeString()}</span>
            </div>
            <div class="message-content">${event.data}</div>
        `;
        messagesContainer.appendChild(messageElement);
        messagesContainer.scrollTop = messagesContainer.scrollHeight;
    };
    ws.onerror = (event) => {
        console.log(event);
    };
}

connectWebSocket(); // Establish the WebSocket connection

const send = document.getElementById("send");
send.addEventListener("click", () => {
    let message = document.getElementById("msg").value;
    if (message === "") {
        return;
    }
    ws.send(message);
    document.getElementById("msg").value = "";
    const messagesContainer = document.querySelector(".messages-list");
    const messageElement = document.createElement("div");
    messageElement.className = "message-card";
    messageElement.innerHTML = `
        <div class="message-header" style="display: flex; justify-content: space-between">
            <span class="message-author">You</span>
            <span class="message-time">${new Date().toLocaleTimeString()}</span>
        </div>
        <div class="message-content">${message}</div>
    `;
    messagesContainer.appendChild(messageElement);
    messagesContainer.scrollTop = messagesContainer.scrollHeight;
});