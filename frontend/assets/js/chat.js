import * as spa from "./spa.js";

// Chat Vars
export let ws;
let userID = 0;
let isLoading = false;
let messageOffset = 0;
const MESSAGES_PER_PAGE = 10;
// DOM Elements
const ignore = document.getElementById("message");


function initializeDOMElements() {
  window.ignore = document.getElementById("message");
  window.post = document.getElementById("posts");
  window.right_side_bare = document.getElementById("categories");
  window.area_msg = document.getElementById("area-msg");
  window.notif = document.querySelector(".notif");
  window.friends_list = document.querySelector(".friends-list");
  window.chat_box = document.querySelector(".chat-box");
  window.messageInput = document.getElementById("messageInput");
  window.sendButton = document.querySelector(".send-btn");
  window.messagesArea =
    document.getElementById("messages") ||
    document.getElementById("messages-area");
  window.r = document.getElementById("user-receiver");
  window.chat = document.getElementById("chat");
  window.friend_avatar =
    document.querySelector(".friend-profile") ||
    document.querySelector(".friend-avatar");
}

export function initChat() {
  initializeDOMElements()
  const chat = document.getElementById("chat");
  const messageInput = document.getElementById("messageInput");
  const sendButton = document.getElementById("sendButton");
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
}

function handleDocumentClick(event) {
  const messageElement = event.target.closest("#message");
  const post = document.getElementById("posts");
  const right_side_bare = document.getElementById("categories");
  const area_msg = document.getElementById("area-msg");

  if (messageElement) {
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
  if (window.friends_list && window.chat_box) {
    window.friends_list.style.display = "none";
    window.chat_box.style.display = "flex";
  }
}

function showFriendsList() {
  if (window.friends_list && window.chat_box) {
    window.friends_list.style.display = "block";
    window.chat_box.style.display = "none";
  }
}

function closeMessageArea() {
  if (window.area_msg && window.post && window.notif) {
    window.area_msg.style.display = "none";
    window.post.style.display = "flex";
    window.notif.style.display = "flex";
  }
}

// WebSocket Functions
export function connectWebSocket() {
  try {
    ws = new WebSocket("ws://localhost:9090/ws");

    ws.onopen = () => console.log("Connected to chat server");
    ws.onmessage = handleWebSocketMessage;
    ws.onerror = (event) => console.error("WebSocket error:", event);
    ws.onclose = () => {
      console.log("Disconnected from chat server, attempting to reconnect...");
      setTimeout(connectWebSocket, 5000);
    };
  } catch (error) {
    console.error("WebSocket connection error:", error);
    setTimeout(connectWebSocket, 5000);
  }
}


function handleWebSocketMessage(event) {
  try {
    const data = JSON.parse(event.data);

    if (data.type === "users_list") {
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
      handleIncomingMessage(data);
    }
  } catch (error) {
    console.error("Error processing WebSocket message:", error);
  }
}

function handleIncomingMessage(data) {
  const time = new Date(data.timestamp || new Date()).toLocaleTimeString([], {
    hour: "2-digit",
    minute: "2-digit",
    hour12: false,
  });
  let activeUserId = getActiveChatUserId();

  console.log("Active User ID:", activeUserId, "Sender ID:", data.sender_id);
  console.log(data.sender_id, activeUserId);
  console.log("omar : ", data);
  
  if (!activeUserId || activeUserId != data.sender_id) {
    // TODO ba mohamed if you want show me your creativity 😉😉
    // create Notification her
    // alert(`${data.username} send you message!`)
    return
  }
  if (activeUserId === data.sender_id) {
    // Hadi 4at5daaam mnin n7aydo input without select 
    markMessagesAsRead(activeUserId)
    const messagesContainer =
      window.messagesArea ||
      document.getElementById("messages") ||
      document.getElementById("messages-area");
    if (messagesContainer) {
      const messageElement = createMessageElement(
        data.content,
        time,
        "received",
        data.username
      );
      messagesContainer.appendChild(messageElement);
      messagesContainer.scrollTop = messagesContainer.scrollHeight;
    } else {
      console.error("Messages container not found!");
    }
  } else {
    console.log("Message is from another user");
  }
}

function getActiveChatUserId() {
  const statusElement = document.querySelector(".status[id^='user-status-']");
  if (statusElement && statusElement.id) {
    const matches = statusElement.id.match(/user-status-(\d+)/);
    if (matches && matches[1]) {
      return parseInt(matches[1]);
    }
  }
  const alternativeStatus = document.querySelector(".status[id^='user-']");
  if (alternativeStatus && alternativeStatus.id) {
    const matches = alternativeStatus.id.match(/user-(\d+)/);
    if (matches && matches[1]) {
      return parseInt(matches[1]);
    }
  }
  return userID || currentUserId;
}

// Message Functions
function sendMessage() {
  const message = messageInput.value.trim();
  const messagesArea = document.getElementById("messages");
  const r = document.getElementById("user-receiver");

  if (message) {
    const time = new Date().toLocaleTimeString([], {
      hour: "2-digit",
      minute: "2-digit",
      hour12: false,
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
    const messageDiv = createMessageElement(message, time, "sent", "You");
    messagesArea.appendChild(messageDiv);
    messageInput.value = "";
    messagesArea.scrollTop = messagesArea.scrollHeight;
  }
}

function createMessageElement(content, time, type, username) {
  const messageDiv = document.createElement("div");
  messageDiv.className = `messages ${type}`;
  
  // Add a line break if content is more than 30 characters
  const formattedContent = content.length > 30 ? content.replace(/(.{30})/g, "$1<br>") : content;

  messageDiv.innerHTML = `
    <div class="message-bubble">
      <div class="message-content">${formattedContent}</div>
      <div class="message-time">${time}</div>
    </div>
    <div class="message-author">${username}</div>
  `;
  return messageDiv;
}
export function displayMessage(message, currentUserId) {
  const time = new Date(message.timestamp).toLocaleTimeString([], {
    hour: "2-digit",
    minute: "2-digit",
    hour12: false,
  });
  const isSent = parseInt(message.sender_id) !== parseInt(currentUserId);
  console.log("Is Sent  ;", isSent);
  
  return createMessageElement(message.content, time, isSent ? "sent" : "received", message.username);
}

// Friend List Functions
function addFriend(
  friends,
  userIds,
  userStatuses,
  lastMessages,
  lastTimes,
  unreadCounts
) {
  const friendsList = document.querySelector(".allfriends");
  const messagesArea = document.getElementById("messages");
  const friends_list = document.querySelector(".friends-list");
  const chat_box = document.querySelector(".chat-box");

  if (!friendsList) {
    console.error("friendsList is not found in the DOM!");
    return;
  }

  friendsList.innerHTML = "";
  if (friends){

    friends.forEach((friend) => {
      const userId = userIds[friend];
      const status = userStatuses[friend] || "offline";
    const lastMessage = lastMessages[friend] || "No messages yet";
    const unreadCount = unreadCounts[friend] || 0;
    const lastTime = lastTimes[friend]
      ? new Date(lastTimes[friend]).toLocaleTimeString([], {
          hour: "2-digit",
          minute: "2-digit",
          hour12: false,
        })
      : "—";

      const friendElement = document.createElement("div");
      friendElement.className = "friend";
      if (unreadCount > 0) {
      friendElement.classList.add("has-unread");
    }

    friendElement.innerHTML = `
        <div class="friend-avatar">
            <img src="../../assets/images/profile.png" class="profile-img" alt="${friend}">
            <div class="status ${status}" id="user-${userId}"></div>
        </div>
        <div class="friend-info">
            <div>
              <div class="friend-name">${friend}</div>
                <div class="last-message">${
                lastMessage.length > 30
                  ? lastMessage.substring(0, 15) + "..."
                  : lastMessage
              }</div>
          </div>
          <div class="time-notif">
              <div class="last-time">${lastTime}</div>
              <div class="notification ${
                unreadCount === 0 ? "hidden" : ""
              }">${unreadCount}</div>
          </div>
      </div>`;

      friendElement.addEventListener("click", () =>
        handleFriendClick(friend, userId, status, unreadCount)
      );
      friendsList.appendChild(friendElement);

      // const statusElement = friendElement.querySelector(`#user-${userId}`);
      // statusElement.classList.toggle("online", status === "online");
      // statusElement.classList.toggle("offline", status === "offline");

      // friendElement.addEventListener("click", () => {
      //   messagesArea.innerHTML = "";
      //   const show_user = document.getElementById("user-receiver");
      //   friends_list.style.display = window.innerWidth <= 780 ? "none" : "block";
      //   chat_box.style.display = "flex";
      //   show_user.innerText = friend;
      //   messageOffset = 0;
      //   fetchChatHistory(userId, messageOffset);
      //   if (messagesArea && userID == 0) {
      //     userID = userId;
      //     messagesArea.addEventListener("scroll", async () => {
      //       if (messagesArea.scrollTop === 0 && !isLoading && userID !== 0) {
      //         isLoading = true;
      //         await fetchChatHistory(userID, messageOffset);
      //         isLoading = false;
      //       }
      //     });
      //   }
      // });

      // friendsList.appendChild(friendElement);
    });
  }
  
}

function handleFriendClick(friend, userId, status, unreadCount) {
  console.log(userId);

  console.log("Friend clicked:", friend, "ID:", userId);
  if (unreadCount > 0 ) {
    markMessagesAsRead(userId)
  }
  if (window.messagesArea) {
    window.messagesArea.innerHTML = "";
  } else {
    console.error("Messages area not found!");
    return;
  }
  if (window.r) window.r.innerText = friend;
  if (window.friends_list)
    window.friends_list.style.display =
      window.innerWidth <= 780 ? "none" : "block";
  if (window.chat_box) window.chat_box.style.display = "flex";

  if (window.friend_avatar) {
    window.friend_avatar.innerHTML = `
      <div class="friend-avatar">
        <img src="../../assets/images/profile.png" class="profile-img" alt="${friend}">
        <div class="status ${status}" id="user-status-${userId}"></div>
      </div>`;
  }
  messageOffset = 0;
  fetchChatHistory(userId, messageOffset);
  setupScrollListener(userId);
}

function setupScrollListener(userId) {
  if (!window.messagesArea) return;

  const oldElement = window.messagesArea;
  const newElement = oldElement.cloneNode(true);
  oldElement.parentNode.replaceChild(newElement, oldElement);
  window.messagesArea = newElement;

  window.messagesArea.addEventListener("scroll", async () => {
    if (window.messagesArea.scrollTop === 0 && !isLoading) {
      isLoading = true;
      await fetchChatHistory(userId, messageOffset);
      isLoading = false;
    }
  });
}

// Chat History Functions
async function fetchChatHistory(userId, offset = 0) {
  try {
    const response = await fetch(
      `/api/chat/history?user_id=${userId}&offset=${offset}`
    );

    if (!response.ok) {
      throw new Error(`Error fetching chat history: ${response.statusText}`);
    }
    const data = await response.json();
    if (!window.messagesArea) {
      console.error("Messages area not found!");
      return;
    }
    if (!data.messages || data.messages.length === 0) {
      if (offset === 0) {
        const emptyMessage = document.createElement("div");
        emptyMessage.className = "no-messages";
        emptyMessage.textContent = "No messages yet. Start a conversation!";
        window.messagesArea.appendChild(emptyMessage);
      }
      return;
    }
    console.log(data);
    
    messageOffset += data.messages.length;
    console.log("Updated messageOffset:", messageOffset);
    const fragment = document.createDocumentFragment();
    
    data.messages.forEach((message) => {
      if (!document.querySelector(`[data-message-id="${message.id}"]`)) {
        const messageDiv = displayMessage(message, userId);
        messageDiv.setAttribute("data-message-id", message.id);
        fragment.prepend(messageDiv);
      }
    });
    window.messagesArea.prepend(fragment);
    if (offset === 0) {
      window.messagesArea.scrollTop = window.messagesArea.scrollHeight;
    } else {
      window.messagesArea.scrollTop = 60;
    }
    return data.messages.length;
  } catch (error) {
    console.error("Error fetching chat history:", error);
    return 0;
  }
}
function markMessagesAsRead(receiverId) {
  console.log("Sending request to mark messages as read:", {
    receiver_id: receiverId,
  });

  fetch("/api/mark-read", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      receiver_id: receiverId,
    }),
  })
    .then((response) => {
      if (!response.ok) {
        throw new Error("error in markmessage read");
      }
      // Update the friend list notification
      const friendElement = document.querySelector(
        `.friend:has([id="user-${receiverId}"])`
      );
      if (friendElement) {
        friendElement.classList.remove("has-unread");
        const notifElement = friendElement.querySelector(".notification");
        if (notifElement) {
          notifElement.classList.add("hidden");
          notifElement.textContent = "0";
        }
      }
    })
    .catch((error) => {
      console.error(error);
    });
}