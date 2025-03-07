let ws;
let userID = 0;
let isLoading = false;
let messageOffset = 0;
const MESSAGES_PER_PAGE = 10;

// DOM Elements
const ignore = document.getElementById("message");
const post = document.getElementById("posts");
const right_side_bare = document.getElementById("categories");
const area_msg = document.getElementById("area-msg");
const notif = document.querySelector(".notif");
const friends_list = document.querySelector(".friends-list");
const chat_box = document.querySelector(".chat-box");
const messageInput = document.getElementById("messageInput");
const sendButton = document.querySelector(".send-btn");
const messagesArea = document.getElementById("messages");
const r = document.getElementById("user-receiver");
const chat = document.getElementById("chat");
const friend_avatar = document.querySelector(".friend-avatar");

// Event Listeners
document.addEventListener("click", handleDocumentClick);
sendButton.addEventListener("click", sendMessage);
messageInput.addEventListener("keypress", handleKeyPress);

if (window.innerWidth <= 780) {
  const friend = document.querySelector(".friend");
  const back = document.querySelector(".back");
  const close_message = document.querySelector(".close-message");

  if (friend) friend.addEventListener("click", showChatBox);
  if (back) back.addEventListener("click", showFriendsList);
  if (close_message) close_message.addEventListener("click", closeMessageArea);
}

// Functions
function handleDocumentClick(event) {
  const messageElement = event.target.closest("#message");
  if (messageElement) {
    notif.style.display = "none";
    area_msg.style.display = "flex";
    right_side_bare.style.display = "none";
    post.style.display = "none";
  }
}

function handleKeyPress(e) {
  if (e.key === "Enter") {
    sendMessage();
  }
}

function showChatBox() {
  friends_list.style.display = "none";
  chat_box.style.display = "flex";
}

function showFriendsList() {
  friends_list.style.display = "block";
  chat_box.style.display = "none";
}

function closeMessageArea() {
  area_msg.style.display = "none";
  post.style.display = "flex";
  notif.style.display = "flex";
}

function sendMessage() {
  const message = messageInput.value.trim();
  if (message) {
    const time = new Date().toLocaleTimeString([], {
      hour: "2-digit",
      minute: "2-digit",
    });
    if (ws) {
      ws.send(
        JSON.stringify({
          type: "message",
          username: r.innerText,
          content: message,
          timestamp: new Date().toISOString(),
        })
      );
    }
    const messageDiv = createMessageElement(message, time, "sent");
    messagesArea.appendChild(messageDiv);
    messageInput.value = "";
    messagesArea.scrollTop = messagesArea.scrollHeight;
  }
}

function createMessageElement(content, time, type) {
  const messageDiv = document.createElement("div");
  messageDiv.className = `messages ${type}`;
  messageDiv.innerHTML = `
    <div class="message-bubble">
      <div class="message-content">${content}</div>
      <div class="message-time">${time}</div>
    </div>
  `;
  return messageDiv;
}

function connectWebSocket() {
  ws = new WebSocket("ws://localhost:9090/ws");

  ws.onopen = () => console.log("Connected to chat");
  ws.onmessage = handleWebSocketMessage;
  ws.onerror = (event) => console.log(event);
  ws.onclose = () => {
    console.log("Disconnected from chat");
    setTimeout(connectWebSocket, 5000);
  };
}
let button = document.getElementById("sendButton");
// let checkStatus = document.querySelector(".status");
let currentUserId = null;

function setCurrentUser(userId) {
  currentUserId = userId;
}

function handleWebSocketMessage(event) {
  const data = JSON.parse(event.data);
  console.log("WebSocket received:", data);

  if (data.type === "users_list") {
    console.log("Updating friends list", data);
    addFriend(
      data.usernames,
      data.user_ids,
      data.user_statuses,
      data.last_messages,
      data.last_times,
      data.unread_counts
    );
    return;
  }

  if (data.type === "message") {
    console.log("Message received:", data);
    const time = new Date().toLocaleTimeString([], {
      hour: "2-digit",
      minute: "2-digit",
    });

    let checkStatus = document.querySelector(".status");
    console.log("Status element:", checkStatus);

    if (!checkStatus || !checkStatus.id) {
      console.log("checkStatus غير موجود، كاين مشكل فـ DOM!");
      alert(`${data.username} sent you a message!`);
      return;
    }

    let activeUserId = null;
    if (checkStatus.id.includes(`${currentUserId}`)) {
      alert(`${data.username} sent you a message!`);
      return;
    } else if (checkStatus.id.startsWith("user-status-")) {
      activeUserId = parseInt(checkStatus.id.replace("user-status-", ""));
    } else {
      const matches = checkStatus.id.match(/\d+/);
      activeUserId = matches ? parseInt(matches[0]) : NaN;
    }

    console.log("Active User ID:", activeUserId, "Sender ID:", data.sender_id);

    if (!isNaN(activeUserId) && activeUserId === data.sender_id) {
      const messagesArea = document.querySelector(".messages-area");
      if (!messagesArea) {
        console.error("messagesArea is not defined!");
        return;
      }
      const messageElement = createMessageElement(data.content, time, "received");
      messagesArea.appendChild(messageElement);
      messagesArea.scrollTop = messagesArea.scrollHeight;
    } else {
      console.log("Message is from another user");
      alert(`${data.username} sent you a message!`);
    }
  }
}


function addFriend(friends, userIds, userStatuses, lastMessages, lastTimes, unreadCounts) {
  const friendsList = document.querySelector(".allfriends");
  if (!friendsList) {
    console.error("friendsList is not found in the DOM!");
    return;
  }

  friendsList.innerHTML = "";

  friends.forEach((friend) => {
    const userId = userIds[friend];
    const status = userStatuses[friend] || "offline";
    const lastMessage = lastMessages[friend] || "No messages yet";
    const unreadCount = unreadCounts[friend] || 0;
    const lastTime = lastTimes[friend] ? 
      new Date(lastTimes[friend]).toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" }) : "—";

    const friendElement = document.createElement("div");
    friendElement.className = "friend";
    if (unreadCount > 0) {
      friendElement.classList.add("has-unread");
    }

    const img = new Image();
    img.src = `https://api.multiavatar.com/${friend}.svg`;
    img.onerror = function () {
      img.src = `https://ui-avatars.com/api/?name=${friend}&background=random`;
    };

    friendElement.innerHTML = `
      <div class="friend-avatar">
          <img src="${img.src}" class="profile-img" alt="${friend}">
          <div class="status ${status}" id="user-${userId}"></div>
      </div>
      <div class="friend-info">
          <div>
              <div class="friend-name">${friend}</div>
              <div class="last-message">${lastMessage.length > 30 ? lastMessage.substring(0, 30) + '...' : lastMessage}</div>
          </div>
          <div class="time-notif">
              <div class="last-time">${lastTime}</div>
              ${unreadCount > 0 ? `<div class="notification">${unreadCount}</div>` : ''}
          </div>
      </div>`;

    friendElement.addEventListener("click", () => {
      console.log("Friend clicked:", friend);

      const messagesArea = document.querySelector(".messages-area");
      const show_user = document.querySelector("#show-user");
      const friend_avatar = document.querySelector("#friend-avatar");
      const button = document.querySelector("#chat-button");

      if (!messagesArea || !show_user || !friend_avatar || !button) {
        console.error("One or more DOM elements are missing! Check `.messages-area`, `#show-user`, `#friend-avatar`, `#chat-button`.");
        return;
      }

      messagesArea.innerHTML = "";
      show_user.innerHTML = friend;
      friend_avatar.innerHTML = `
        <div class="friend-avatar">
          <img src="../../assets/images/profile.png" class="profile-img" alt="${friend}">
          <div class="status ${status}" id="user-status-${userId}"></div>
        </div>`;
      button.id = `user-status-${userId}`;
      currentUserId = userId;

      fetchChatHistory(userId, messageOffset);

      if (messagesArea && userID == 0) {
        userID = userId;
        messagesArea.addEventListener("scroll", async () => {
          if (messagesArea.scrollTop === 0 && !isLoading && userID !== 0) {
            isLoading = true;
            await fetchChatHistory(userID, messageOffset);
            isLoading = false;
          }
        });
      }

      if (unreadCount > 0) {
        markMessagesAsRead(userId);
        friendElement.classList.remove("has-unread");
        const notifElement = friendElement.querySelector(".notification");
        if (notifElement) notifElement.classList.add("hidden");
      }
    });

    friendsList.appendChild(friendElement);
  });
}


function markMessagesAsRead(senderId) {
  if (!currentUserId) {
    return;
  }

  fetch('/api/mark-read', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
      sender_id: senderId,
      receiver_id: currentUserId
    })
  })
    .then(response => {
      if (response.ok) {
        console.log(`Marked messages from ${senderId} as read`);
      } else {
        console.error('Failed to mark messages as read');
      }
    })
    .catch(error => {
      console.error('Error marking messages as read:', error);
    });
}


function createMessageElement(content, time, type) {
  const messageElement = document.createElement("div");
  messageElement.className = `message ${type}`;
  messageElement.innerHTML = `
      <div class="message-content">${content}</div>
      <div class="message-time">${time}</div>
  `;
  return messageElement;
}



async function fetchChatHistory(userId, offset = 0) {
  await fetch(`/api/chat/history?user_id=${userId}&offset=${offset}`)
    .then((response) => {
      if (!response.ok) {
        throw new Error(`Error fetching chat history: ${response.statusText}`);
      }
      return response.json();
    })
    .then((data) => {

      messageOffset += 10;
      console.log(messageOffset);
      data.messages.forEach((mess) => {
        const messageDiv = displayMessage(mess, userId);
        messagesArea.prepend(messageDiv);
      });
      messagesArea.scrollTop =
        messageOffset <= 10 ? messagesArea.scrollHeight : 60;
    })
    .catch((error) => {
      console.error(error);
      return [];
    });
}

function displayMessage(message, currentUserId) {
  const time = new Date(message.timestamp).toLocaleTimeString([], {
    hour: "2-digit",
    minute: "2-digit",
  });
  const isSent = parseInt(message.sender_id) !== parseInt(currentUserId);
  return createMessageElement(
    message.content,
    time,
    isSent ? "sent" : "received"
  );
}

connectWebSocket();
