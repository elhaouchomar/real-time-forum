
function toggleCollapse(elem, comments) {
    elem.classList.toggle("collapse");
    comments.forEach(second_elem => {
        if (second_elem != elem)
            second_elem.classList.add("collapse");
    });
}
function ExpandComments(flag) {
    // Expand Comment and read Content...
    let comments = document.querySelectorAll(".commentData")
    comments.forEach(elem => {
        if (flag){
            elem.addEventListener('click', ()=> toggleCollapse(elem, comments))
       
        }else{
            elem.removeEventListener('click',  ()=> toggleCollapse(elem, comments))
        }
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
    const commentError = document.querySelector('.CommentErrorMessage');

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
            popUp();
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
        if (response.status == 500) {
            commentError.style.display = "block"
            commentError.innerText = "Oops! It looks like you've already posted this comment. Please try something new!";
            setTimeout(() => {
                commentError.style.display = "none"
                commentError.innerText = "";
            }, 5000);
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
                    <img src="https://ui-avatars.com/api/?name=${userName}" alt="">
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
            // call new listeners
            handleLikes();

            // remove old Listners :
            ExpandComments(false)
            // call new Listners
            ExpandComments(true);
        }
    }
}



function CommentInputEventListenner() {
    const send_comment = document.querySelector('.send-comment');
    const commentInput = document.querySelector('.commentInput input');
    if (eventListenerMap.has(send_comment)) {
        send_comment.removeEventListener('click', eventListenerMap.get(send_comment));
    }
    if (eventListenerMap.has(commentInput)) {
        commentInput.removeEventListener('keypress', eventListenerMap.get(commentInput));
    }
    
    const handleCommentEventWrapper = (e) => handleCommentEvent(e);
    eventListenerMap.set(commentInput, handleCommentEventWrapper);
    eventListenerMap.set(send_comment, handleCommentEventWrapper);

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
CommentInputEventListenner()
ExpandComments(true)

