import { HandleLikes } from "./likes.js";
import { AVATAR_URL, ListnerMap, LoadPage } from "./spa.js";

function toggleCollapse(elem, comments) {
    elem.classList.toggle("collapse");
    comments.forEach(second_elem => {
        if (second_elem != elem)
            second_elem.classList.add("collapse");
    });
}
export function ExpandComments(flag) {
    // Expand Comment and read Content...
    let comments = document.querySelectorAll(".commentData")
    comments.forEach(elem => {
    if (ListnerMap.has(elem)){
        elem.removeEventListener('click', ListnerMap.get(elem))
    }
    const toggleCollapseWrapper = () => toggleCollapse(elem, comments);
    elem.addEventListener('click', toggleCollapseWrapper);
    ListnerMap.set(elem, toggleCollapseWrapper);
    })
}

function CommentErrorMsg(msg) {
    const commentError = document.querySelector('.CommentErrorMessage');
    commentError.style.display = "block"
    commentError.innerText = msg;
    setTimeout(() => {
        commentError.style.display = "none"
        commentError.innerText = "";
    }, 5000);
}

// Remove duplicate 500 status check
async function handleCommentEvent(e) {
    if (e.type === 'click' || (e.type === 'keypress' && e.key === 'Enter')) {
        e.preventDefault();
        const commentValue = e.target.closest('.commentInput').querySelector('input');

        const comment = commentValue.value;
        if (comment.trim() === '' || comment.length === 0) {
            return;
        }

        const postID = commentValue.id;
        const response = await fetch('/CreateComment', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                postID,
                comment
            })
        });

        if (response.status === 401) {
            LoadPage("login")
            return;
        }
        if (response.status === 429) {
            CommentErrorMsg(`Slow down! Good comments take time—quality over speed! try again after 1 minute 😊`);
            return;
        }
        if (response.status === 500) {
            CommentErrorMsg("Oops! It looks like you've already posted this comment. Please try something new!");
            return;
        }
       
        const data = await response.json();
        commentValue.value = '';
        if (data["status"] == "ok") {
            const commentContainer = document.querySelector('.Comments');
            const commentCard = document.createElement('div');
            commentCard.classList.add('commentCard');
            commentCard.classList.add('CommentAdded');
            const userName =  "Mohamed Tawil"
            commentCard.innerHTML = `
                <div class="commentAuthorImage">
                    <img src="${AVATAR_URL}${userName}" alt="">
                </div>
                <div class="commentAuthor">
                    <div class="commentAuthorInfo">
                        <span class="commentAuthorName">
                            @${data["UserName"]}
                            <span class="commentTime">
                                ${data["CreatedAt"]}
                            </span>
                        </span>
                        <div class="commentReaction DisableUserSelect">
                            <span isPost="false" class="like" id="${data["CommentID"]}">
                                <i class="material-symbols-outlined">
                                    thumb_up
                                </i>
                                <span>0</span>
                            </span>
                            <span isPost="false" class="dislike" id="${data["CommentID"]}">
                                <i class="material-symbols-outlined">thumb_down</i>
                                <span>0</span>
                            </span>
                        </div>
                    </div>
                    <div class="commentInfo">
                        <p class="commentData collapse"></p>
                    </div>
                </div>
            `;
            commentCard.querySelector('.commentData').innerText = data["Content"];
            commentContainer.prepend(commentCard);
            document.querySelector('.commentCount').textContent = data.CommentCount
            document.querySelector(`.comment.post span[id="${postID}"]`).textContent = data.CommentCount
            // call new listeners
            HandleLikes();
            // call new Listners
            ExpandComments();
        }
    }
}



export function CommentInputEventListenner() {
    const send_comment = document.querySelector('.send-comment');
    const commentInput = document.querySelector('.commentInput input');
    if (ListnerMap.has(send_comment)) {
        send_comment.removeEventListener('click', ListnerMap.get(send_comment));
    }
    if (ListnerMap.has(commentInput)) {
        commentInput.removeEventListener('keypress', ListnerMap.get(commentInput));
    }
    
    const handleCommentEventWrapper = (e) => handleCommentEvent(e);
    ListnerMap.set(commentInput, handleCommentEventWrapper);
    ListnerMap.set(send_comment, handleCommentEventWrapper);

    commentInput.addEventListener('keypress', handleCommentEventWrapper);
    send_comment.addEventListener('click', handleCommentEventWrapper);
}

function DisplayComments() {

    const commentSection = document.querySelector('.postComments');
    const postSection = document.querySelector('.ProfileAndPost');
    commentSection.style.display = 'none';
    postSection.style.display = 'flex';
}


function PostButtonSwitcher(){
    const postButton = document.querySelector('.PostButton');
    if (eventListenerMapx.has(postButton)) {
        postButton.removeEventListener('click', eventListenerMapx.get(postButton));
    }

    const handleDisplayComments = () => DisplayComments();
    eventListenerMapx.set(postButton, handleDisplayComments);
    postButton.addEventListener('click', handleDisplayComments);
}
