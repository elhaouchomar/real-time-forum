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
const sendButton = document.getElementById("sendButton");
const messagesArea = document.getElementById("messages");
const r = document.getElementById("user-receiver");
const chat = document.getElementById("chat");

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

function handleWebSocketMessage(event) {
  const data = JSON.parse(event.data);
  if (data.type === "users_list") {
    addFriend(data.usernames, data.user_ids, data.user_statuses);
    return;
  }
  if (data.type === "message") {
    const time = new Date().toLocaleTimeString([], {
      hour: "2-digit",
      minute: "2-digit",
    });
    const messageElement = createMessageElement(data.content, time, "received");
    messagesArea.appendChild(messageElement);
    messagesArea.scrollTop = messagesArea.scrollHeight;
  }
}

function addFriend(friends, userIds, userStatuses) {
  const friendsList = document.querySelector(".allfriends");
  friendsList.innerHTML = "";

  friends.forEach((friend) => {
    const userId = userIds[friend];
    const status = userStatuses[friend];

    const friendElement = document.createElement("div");
    friendElement.className = "friend";
    friendElement.innerHTML = `
      <div class="friend-avatar">
        <img src="../../assets/images/profile.png" class="profile-img" alt="${friend}">
        <div class="status ${status}" id="user-${userId}"></div>
      </div>
      <div class="friend-info">
        <div class="friend-name">${friend}</div>
      </div>`;

    const statusElement = friendElement.querySelector(`#user-${userId}`);
    statusElement.classList.toggle("online", status === "online");
    statusElement.classList.toggle("offline", status === "offline");

    friendElement.addEventListener("click", () => {
      messagesArea.innerHTML = "";
      const show_user = document.getElementById("user-receiver");
      friends_list.style.display = window.innerWidth <= 780 ? "none" : "block";
      chat_box.style.display = "flex";
      show_user.innerText = friend;
      messageOffset = 0;
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
    });

    friendsList.appendChild(friendElement);
  });
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
      data.messages.forEach((mess) => {
        const messageDiv = displayMessage(mess, userId);
        messagesArea.prepend(messageDiv);
      });
      messagesArea.scrollTop = messageOffset <= 10 ? messagesArea.scrollHeight : 60;
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
  return createMessageElement(message.content, time, isSent ? "sent" : "received");
}

connectWebSocket();
