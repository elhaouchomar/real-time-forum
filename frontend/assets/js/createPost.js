import { CreatePostPage } from "./component/createPostPage.js";
import { HandleLikes } from "./likes.js";
import { readPost } from "./script.js";
import { ListnerMap, AVATAR_URL, SPAContainer } from "./spa.js";

const ERROR_DISPLAY_TIME = 5000;

function showError(message) {
    const errorBox = document.querySelector(".ErrorMessage");
    errorBox.style.display = "flex";
    document.querySelector(".message").innerText = message;
    setTimeout(() => errorBox.style.display = "none", ERROR_DISPLAY_TIME);
}

function createPostTemplate(post) {
    return `
        <a href="/?type?profile&username=${post.UserName}">
            <div class="ProfileImage tweet-img" style="background-image: url('${AVATAR_URL}${post.UserName}')"></div>
        </a>
        <div class="post-details">
            <div class="row-tweet">
                <div class="post-header">
                    <span class="tweeter-name post" id="${post.ID}">
                        <span class="text-title"></span>
                        <br><span class="tweeter-handle">@${post.UserName} ${post.CreatedAt}.</span>
                    </span>
                </div>
            </div>
            <div class="post-content"><p></p></div>
            <span class="see-more">See More</span>
            <div class="Hashtag">
                ${post.Categories.map(category => `<a href=""><span>#${category}</span></a>`).join('')}
            </div>
            <div class="post-footer">
                <div class="react">
                    <div isPost="true" class="counters like" id="${post.ID}">
                        <i class="material-symbols-outlined popup-icon">thumb_up</i>
                        <span>${post.LikeCount}</span>
                    </div>
                    <div isPost="true" class="counters dislike" id="${post.ID}">
                        <i class="material-symbols-outlined popup-icon">thumb_down</i>
                        <span>${post.DislikeCount}</span>
                    </div>
                </div>
                <div class="comment post" id="${post.ID}">
                    <i class="material-symbols-outlined showCmnts">comment</i>
                    <span id="${post.ID}">0</span>
                </div>
            </div>
        </div>`;
}

export function createPostListner() {
    const CreatePostArea = document.querySelectorAll(".new-post-header");
    CreatePostArea.forEach((elem) => {
        const handler = () =>  createPost()
        if (ListnerMap.has(elem)){
            elem.removeEventListener('click', ListnerMap.get(elem))
        }
        elem.addEventListener("click", handler);
        ListnerMap.set(elem, handler)
    });
}

export async function createPost() {
    const CreatePostPageHtml = CreatePostPage()
    const firShild = SPAContainer.querySelector('div')
    SPAContainer.insertBefore(CreatePostPageHtml, firShild)
    const modal = document.querySelector(".postModal");
    const closeBtn = document.querySelector(".titleInput .close-post");
    const form = document.querySelector('.CreatePostContainer form');
  
    modal.style.display = "flex";
    document.querySelector(".titleInput input").focus();

    const closeModal = () => {
        SPAContainer.querySelector('#CreatePostModal').remove()
    };

    window.onclick = e => {
        if (e.target === modal) closeModal();
    };

    closeBtn.addEventListener('click', closeModal);

    form.addEventListener('submit', async (e) => {
        e.preventDefault();

        // Get checked categories and sanitize IDs
        const categories = Array.from(form.category)
            .filter(input => input.checked)
            .map(input => input.id);

        if (!categories.length) {
            showError("Please Select At Least One Category");
            return;
        }

        // Get category names safely
        const categoriesList = categories.map(id => {
            const element = document.querySelector(`[id="${id}"]`);
            return element ? element.value : '';
        }).filter(Boolean);

        const data = {
            title: form.title.value,
            content: form.content.value,
            categories,
            categoriesList
        };

        try {
            const response = await fetch('/createPost', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(data)
            });

            const post = await response.json();
            console.log("======>", post);
            
            if (response.status === 200) {
                updateUI(post, closeModal, form);
            } else {
                showError(post.data.message)
            }
        } catch (error) {
            showError("Failed to create post, try later...");
        }
    });
}

function updateUI(post, closeModal, form) {
    const postCard = document.createElement('div');
    postCard.classList.add('post-card', 'PostAdded');
    postCard.innerHTML = createPostTemplate(post);

    postCard.querySelector('.post-content p').innerText = post.Content;
    postCard.querySelector('.text-title').innerText = post.Title;

    const mainFeed = document.querySelector('.main-feed');
    mainFeed.insertBefore(postCard, mainFeed.children[1]);
    closeModal();
    form.reset();

    // Refresh event listeners
    [readPost, HandleLikes, createPostListner].forEach(fn => fn());
}