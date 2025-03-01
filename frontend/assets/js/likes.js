// Store handler references
const likeHandlers = new WeakMap();
const dislikeHandlers = new WeakMap();

async function handleReaction(btn, type) {
    const postId = btn.id;
    const postOrComment = btn.getAttribute('isPost') === "true";
    
    try {
        const res = await fetch('/PostReaction', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ 
                postId,
                type,
                post: postOrComment
            })
        });

        if (res.status === 401) return popUp();
        if (res.status !== 200) return;

        const data = await res.json();
        const otherBtn = type === "like" ? 
            btn.nextElementSibling : 
            btn.previousElementSibling;
        
        if (data.added) {
            btn.classList.add("FILL");
            otherBtn?.classList.remove("FILL");
        } else {
            btn.classList.remove("FILL");
        }
        
        if (otherBtn) {
            otherBtn.querySelector('span').innerText = type === "like" ? data.dislikes : data.likes;
        }
        btn.querySelector('span').innerText = type === "like" ? data.likes : data.dislikes;
    } catch (error) {
        console.error(`${type} operation failed:`, error);
    }
}

// Usage:
const LikePostAndComments = btn => handleReaction(btn, "like");
const DisLikePostAndComments = btn => handleReaction(btn, "dislike");

export function HandleLikes() {
    const likeBtns = document.querySelectorAll('.like');
    const dislikeBtns = document.querySelectorAll('.dislike');

    likeBtns.forEach(btn => {
        // Remove old handler if exists
        if (likeHandlers.has(btn)) {
            btn.removeEventListener('click', likeHandlers.get(btn));
        }
        
        // Create and store new handler
        const handler = () => LikePostAndComments(btn);
        likeHandlers.set(btn, handler);
        btn.addEventListener('click', handler);
    });

    dislikeBtns.forEach(btn => {
        // Remove old handler if exists
        if (dislikeHandlers.has(btn)) {
            btn.removeEventListener('click', dislikeHandlers.get(btn));
        }
        
        // Create and store new handler
        const handler = () => DisLikePostAndComments(btn);
        dislikeHandlers.set(btn, handler);
        btn.addEventListener('click', handler);
    });
}