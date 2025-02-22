

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

const ignore = document.getElementById("message");
const post = document.getElementById("posts");
const right_side_bare = document.getElementById("categories");
const area_msg = document.getElementById("area-msg");

ignore.addEventListener("click", () => {
  area_msg.style.display = "flex";
  right_side_bare.style.display = "none";
  post.style.display = "none";
});

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
    messageDiv.className = "message sent";
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

// STATUS
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