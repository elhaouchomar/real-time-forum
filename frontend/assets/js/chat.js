let ws;

// function appendMessage(message) {
//   const messagesContainer = document.querySelector(".messages-list");
//   const messageElement = document.createElement("div");
//   messageElement.className = "message-card";
//   messageElement.innerHTML = `
//         <div class="message-header">
//             <span class="message-author">${message.author}</span>
//             <span class="message-time">${new Date(
//               message.timestamp
//             ).toLocaleString()}</span>
//         </div>
//         <div class="message-content">${message.content}</div>
//     `;
//   messagesContainer.appendChild(messageElement);
//   messagesContainer.scrollTop = messagesContainer.scrollHeight;
// }

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

/// STATUS
// function setUserStatus(userId, isOnline) {
//   const userElement = document.querySelector(`#user-${userId} .status`);
//   if (userElement) {
//     userElement.classList.toggle("online", isOnline);
//     userElement.classList.toggle("offline", !isOnline);
//   }
// }

// // Exemple d'utilisation
// setUserStatus("Sarah-Smith", true); // En ligne
// setUserStatus("Mike-Johnson", false); // Hors ligne

// websockets

function connectWebSocket() {
  ws = new WebSocket("ws://localhost:9090/ws"); // Ensure WebSocket server is running and accessible
  ws.onopen = () => {
    console.log("Connected to chat");
  };
  ws.onmessage = (event) => {
    const data = JSON.parse(event.data);
    console.log("----------------------", event.data);
    if (data.type === "users") {
      addFriend(data.data);
      return;
    }
    const messagesContainer = document.querySelector(".messages-area");
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
  };
  ws.onerror = (event) => {
    console.log(event);
  };
  ws.onclose = () => {
    console.log("Disconnected from chat");
    setTimeout(connectWebSocket, 5000); // Retry connection after 5 seconds
  };
}
connectWebSocket(); // Establish the WebSocket connection

function addFriend(frinds) {
  const friendsList = document.querySelector(".allfriends");
  friendsList.innerHTML = ""; // Clear existing friends
  frinds.forEach((friend) => {
    const friendElement = document.createElement("div");
    friendElement.className = "friend";
    const friendhtml = `
    <div class="friend-avatar">
      <img src="../../assets/images/profile.png" class="profile-img" alt="Sarah">
      <!-- <img src="/api/placeholder/50/50" class="profile-img" alt="Sarah"> -->
      <div class="status-dot online"></div>
    </div>
    <div class="friend-info">
      <div class="friend-name">${friend}</div>
      <!-- <div class="last-message">Hey, how are you?</div> -->
    </div>`;
    friendElement.innerHTML = friendhtml;

    friendsList.appendChild(friendElement);
    const show_user = document.getElementById("user-receiver");
    friendElement.addEventListener("click", () => {
      friends_list.style.display = "none";
      chat_box.style.display = "flex";
      show_user.innerText = friend;
    });
  });
}
