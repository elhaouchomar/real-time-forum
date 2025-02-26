package handlers

import (
	"fmt"
	"forum/structs"
	"log"
	"net/http"

	"forum/database"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	structError["StatuCode"] = http.StatusNotFound
	// 	structError["MessageError"] = "page not found"
	// 	ErrorPage(w, "error.html", structError)
	// 	return
	// }

	w.Write([]byte(`
	<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8" />
			<meta http-equiv="X-UA-Compatible" content="IE=edge" />
			<meta name="viewport" content="width=device-width, initial-scale=1.0" />
			<link href="/assets/style/google-icons/google-icons.css" rel="stylesheet" />
			<link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
			<link id="home" rel="stylesheet" href="/assets/style/root.css" />
			<link id="home" rel="stylesheet" href="/assets/style/create_post.css" />
			<link id="home" rel="stylesheet" href="/assets/style/header.css" />
			<link id="home" rel="stylesheet" href="/assets/style/left_sidebar.css" />
			<link id="message" rel="stylesheet" href="/assets/style/message.css" />
			<link id="home" rel="stylesheet" href="/assets/style/popstyle.css" />
			<link id="home" rel="stylesheet" href="/assets/style/post.css" />
			<link id="register" rel="stylesheet" href="/assets/style/register.css" />
			<link id="home" rel="stylesheet" href="/assets/style/right_sidebar.css" />
			<link id="home" rel="stylesheet" href="/assets/style/style.css" />
			<link id="home" rel="stylesheet" href="/assets/style/ussely_by_js.css" />
			<link rel="icon" href="/assets/images/logo.svg" type="image/svg+xml" />

			<title>{{/*.Settings.Title*/}}</title>
			<style>
				.SPAContainer {
					display: flex;
					justify-content: center;
					align-items: center;
					height: 100vh;
					width: 100vw;
				}
			</style>	
		</head>

		<body>
			<div class="SPAContainer">
				
			</div>			
			<script>
				function removeCreatePostListner() {
					const CreatePostArea = document.querySelectorAll(".new-post-header");
					CreatePostArea.forEach((elem) => {
					elem.removeEventListener("click", createPostListner);
					});
				}
				function createPostListner() {
					const CreatePostArea = document.querySelectorAll(".new-post-header");
					CreatePostArea.forEach((elem) => {
					elem.addEventListener("click", () => {
						createPost();
					});
					});
				}
				<!-- createPostListner(); --->
			</script>
			<script src="/assets/js/spa.js" type="module" id="spa.js"></script>
		</body>
		</html>

	`))
	return
	template := getHtmlTemplate()

	userId, err := CheckAuthentication(w, r)

	if err != nil {
		return
	}

	var profile structs.Profile
	if userId != 0 {
		profile, err = database.GetUserProfile(DB, userId)
		if err != nil {
			log.Fatal(err)
		}
	}

	categories, err := database.GetCategoriesWithPostCount(DB)

	if err != nil {
		structError["StatuCode"] = http.StatusInternalServerError
		structError["MessageError"] = "error getting categories from database " + err.Error()
		ErrorPage(w, "error.html", structError)
		return
	}
	if r.FormValue("type") != "" {
		profile.CurrentPage = r.FormValue("type")
		profile.Category = r.FormValue("category")
	} else {
		profile.CurrentPage = "home"
	}
	categor := structs.Categories(categories)

	err = template.ExecuteTemplate(w, "index.html", struct {
		Posts      []structs.Post
		Profile    structs.Profile
		Categories structs.Categories
	}{nil, profile, categor})
	if err != nil {
		fmt.Println("error executing template", err)
	}
}
