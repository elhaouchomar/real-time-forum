
function removeCreatePostListner() {
    const CreatePostArea = document.querySelectorAll(".new-post-header")
    CreatePostArea.forEach(elem => {
        elem.removeEventListener('click',createPostListner)
    })
}
function createPostListner() {
    const CreatePostArea = document.querySelectorAll(".new-post-header")
    CreatePostArea.forEach(elem => {
        elem.addEventListener('click', () => {
            createPost.createPost();
        })
    })
}
createPostListner()