export function Messages(){
return `<div id="area-msg" hidden>
            <div class="chat-container">
                <div class="chat-box">
                    <div class="chat-header">
                        <span class="material-symbols-outlined back">
                            arrow_back
                        </span>
                        <div class="friend-avatar">
                            <img src="../../assets/images/profile.png" class="profile-img">
                            <div class="status"></div>
                        </div>
                        <span id="user-receiver"></span>
                    </div>
                    <div id="chat">
                        <div class="messages-area" id="messages"></div>
<<<<<<< Updated upstream
=======
                        <div style="background-image: url('https://img.freepik.com/premium-photo/blackandwhite-doodle-pattern-with-various-objects-icons_1228868-16146.jpg?w=1380');
                            position: absolute;
                            width: 100%;
                            height: 100%;
                            top: 0;
                            left: 0;opacity: 0.1;"  >
                        </div>
>>>>>>> Stashed changes
                    </div>
                    <div class="chat-footer">
                            <input type="text" class="msg-input" id="messageInput" placeholder="Type a message...">
                            <button class="send-btn" id="sendButton">
                                <i class="material-symbols-outlined">send</i>
                            </button>
                    </div>
                </div>

                <div class="friends-list">
                    <div class="friends-header">
                        <h2>Messages</h2>
                        <span class="material-symbols-outlined close-message">
                            close
                        </span>
                    </div>
                    <div class="allfriends"></div>
                </div>
            </div>
        </div>
`
}