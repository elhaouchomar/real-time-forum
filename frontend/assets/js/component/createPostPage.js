export function CreatePostPage(){
    return `
    <div class="postModal">
        <div class="CreatePostContainer">
            <form>
                <div class="titleInput">
                    <input type="text" name="title" placeholder="Enter a Title for your post" required maxlength="60"
                        aria-label="Post Title">
                    <span class="material-symbols-outlined close-post DisableUserSelect">
                        close
                    </span>
                </div>
                <div class="postInput">
                    <textarea name="content" placeholder="Share your thoughts, insights, or stories..." required
                        maxlength="1000" aria-label="Post Content"></textarea>
                </div>
                <div class="CategoriesSelector DisableUserSelect">
                    <input type="checkbox" name="category" id="1" value="General">
                    <label for="1">General</label>
                    <input type="checkbox" name="category" id="2" value="Entertainment">
                    <label for="2">Entertainment</label>
                    <input type="checkbox" name="category" id="3" value="Health">
                    <label for="3">Health</label>
                    <input type="checkbox" name="category" id="4" value="Business">
                    <label for="4">Business</label>
                    <input type="checkbox" name="category" id="5" value="Sports">
                    <label for="5">Sports</label>
                    <input type="checkbox" name="category" id="6" value="Technology">
                    <label for="6">Technology</label>
                </div>
                <input type="submit" value="Publish Post" class="submitButton">
                <div class="ErrorMessage">
                    <div class="MsgContainer">
                        <span class="material-symbols-outlined">error</span>
                        <span class="message">Please fill out all the fields</span>
                    </div>
                </div>
            </form>
        </div>  
    </div>
`
}