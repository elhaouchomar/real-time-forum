let ws;

// Choose send messages
const ignore = document.getElementById("message");
const post = document.getElementById("posts");
const right_side_bare = document.getElementById("categories");
const area_msg = document.getElementById("area-msg");
const notif = document.querySelector(".notif");

document.addEventListener("click", (event) => {
  const messageElement = event.target.closest("#message");
  if (messageElement) {
    notif.style.display = "none";
    area_msg.style.display = "flex";
    right_side_bare.style.display = "none";
    post.style.display = "none";
  }
});

// Change friend
const friends_list = document.querySelector(".friends-list");
const chat_box = document.querySelector(".chat-box");
if (window.innerWidth <= 780) {
  const friend = document.querySelector(".friend");
  const back = document.querySelector(".back");
  const close_message = document.querySelector(".close-message");

  if (friend) {
    friend.addEventListener("click", () => {
      friends_list.style.display = "none";
      chat_box.style.display = "flex";
    });
  }

  if (back) {
    back.addEventListener("click", () => {
      friends_list.style.display = "block";
      chat_box.style.display = "none";
    });
  }

  if (close_message) {
    close_message.addEventListener("click", () => {
      area_msg.style.display = "none";
      post.style.display = "flex";
      notif.style.display = "flex";
    });
  }
}

// SCROLL TO BOTTOM
const messageInput = document.getElementById("messageInput");
const sendButton = document.getElementById("sendButton");
const messagesArea = document.getElementById("messages");
const r = document.getElementById("user-receiver");

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
    const messageDiv = document.createElement("div");
    messageDiv.className = "messages sent";
    messageDiv.innerHTML = `
    <div>${message}</div>
<div class="message-time">${time}</div>
`;
    messagesArea.appendChild(messageDiv);
    messageInput.value = "";
    messagesArea.scrollTop = messagesArea.scrollHeight;
  }
}

sendButton.addEventListener("click", sendMessage);
messageInput.addEventListener("keypress", (e) => {
  if (e.key === "Enter") {
    sendMessage();
  }
});

function connectWebSocket() {
  ws = new WebSocket("ws://localhost:9090/ws");

  ws.onopen = () => {
    console.log("Connected to chat");
  };

  ws.onmessage = (event) => {
    const data = JSON.parse(event.data);
    console.log(event.data); // {"type":"users_list","usernames":["basta","omarel"],"user_statuses":{"basta":"offline","omarel":"online"},"user_ids":{"basta":2,"omarel":1}}

    if (data.type === "users_list") {
      console.log(data.usernames, data.user_ids, data.user_statuses);

      addFriend(data.usernames, data.user_ids, data.user_statuses);
      return;
    }

    if (data.type === "message") {
      const messagesContainer = document.getElementById("messages");
      const messageElement = document.createElement("div");
      messageElement.className = "received";
      messageElement.innerHTML = `
              <div class="message-header" style="display: flex; justify-content: space-between">
                  <span class="message-author">${data.username}</span>
                  <span class="message-time">${data.timestamp}</span>
              </div>
              <div class="message-content">${data.content}</div>
          `;
      messagesContainer.appendChild(messageElement);
      messagesContainer.scrollTop = messagesContainer.scrollHeight;
    }
  };
  ws.onerror = (event) => console.log(event);
  ws.onclose = () => {
    console.log("Disconnected from chat");
    setTimeout(connectWebSocket, 5000);
  };
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

    // Set initial status
    const statusElement = friendElement.querySelector(`#user-${userId}`);
    statusElement.classList.toggle("online", status === "online");
    statusElement.classList.toggle("offline", status === "offline");

    friendElement.addEventListener("click", () => {
      const show_user = document.getElementById("user-receiver");
      friends_list.style.display = window.innerWidth <= 780 ? "none" : "block";
      chat_box.style.display = "flex";
      show_user.innerText = friend;
    });

    friendsList.appendChild(friendElement);
  });
}
connectWebSocket();
