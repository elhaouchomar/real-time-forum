

function appendMessage(message) {
  const messagesContainer = document.querySelector(".messages-list");
  const messageElement = document.createElement("div");
  messageElement.className = "message-card";
  messageElement.innerHTML = `
        <div class="message-header">
            <span class="message-author">${message.author}</span>
            <span class="message-time">${new Date(
              message.timestamp
            ).toLocaleString()}</span>
        </div>
        <div class="message-content">${message.content}</div>
    `;
  messagesContainer.appendChild(messageElement);
  messagesContainer.scrollTop = messagesContainer.scrollHeight;
}


// Add Message
document.addEventListener("DOMContentLoaded", () => {
  connectWebSocket();
  document
    .getElementById("messageForm")
    .addEventListener("submit", sendMessage);
});


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
  if (window.innerWidth <= 780) {
  const friends_list = document.querySelector(".friends-list");
  const chat_box = document.querySelector(".chat-box");
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

function sendMessage() {
  const message = messageInput.value.trim();
  if (message) {
    const time = new Date().toLocaleTimeString([], {
      hour: "2-digit",
      minute: "2-digit",
    });
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