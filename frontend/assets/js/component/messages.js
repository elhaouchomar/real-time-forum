export function Messages(){
    return `
        <div id="area-msg" hidden>
            <div class="chat-container">
                <div class="chat-box">
                <div class="chat-header">
                    <span class="material-symbols-outlined back">
                    arrow_back
                    </span>
                    <img src="../../assets/images/profile.png" class="profile-img" alt="Profile">
                    <span>Amin habchi</span>
                </div>
                
                <div class="messages-area" id="messages">
                    <div class="messages received">
                    <div>Hello! How are you?</div>
                    <div class="message-time">12:30 PM</div>
                    </div>
                    <div class="messages sent">
                    <div>I'm good, thanks! What about you?</div>
                    <div class="message-time">12:32 PM</div>
                    </div>
                    <div class="chat-footer">
                    <input type="text" class="msg-input" id="messageInput" placeholder="Type a message...">
                    <button class="send-btn" id="sendButton"><i class="material-symbols-outlined">send</i></button>
                    </div>
                </div>
                </div>
                <div class="friends-list">
                <div class="friends-header">
                    <h2>Messages</h2>
                    <span class="material-symbols-outlined close-message">
                    close
                    </span>
                </div>
                <div class="friend">
                    <div class="friend-avatar">
                    <img src="../../assets/images/profile.png" class="profile-img" alt="Sarah">
                    <!-- <img src="/api/placeholder/50/50" class="profile-img" alt="Sarah"> -->
                    <div class="status-dot online"></div>
                    </div>
                    <div class="friend-info">
                    <div class="friend-name">Mohamed Tawil</div>
                    <div class="last-message">Hey, how are you?</div>
                    </div>
                </div>
                <div class="friend">
                    <div class="friend-avatar">
                    <!-- <img src="/api/placeholder/50/50" class="profile-img" alt="Mike"> -->
                    <img src="../../assets/images/profile.png" class="profile-img" alt="Mike">
                    <div class="status-dot offline"></div>
                    </div>
                    <div class="friend-info">
                    <div class="friend-name">Youssef Basta</div>
                    <div class="last-message">See you later!</div>
                    </div>
                </div>
                <div class="friend">
                    <div class="friend-avatar">
                    <!-- <img src="/api/placeholder/50/50" class="profile-img" alt="Mike"> -->
                    <img src="../../assets/images/profile.png" class="profile-img" alt="Mike">
                    <div class="status-dot offline"></div>
                    </div>
                    <div class="friend-info">
                    <div class="friend-name">Hassan Ouhamou</div>
                    <div class="last-message">See you later!</div>
                    </div>
                </div>
                </div>
            </div>
        </div>`
}