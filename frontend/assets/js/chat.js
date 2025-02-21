

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

document.addEventListener("DOMContentLoaded", () => {
  connectWebSocket();
  document
    .getElementById("messageForm")
    .addEventListener("submit", sendMessage);
});
