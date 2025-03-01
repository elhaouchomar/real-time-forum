console.log("Sdsdscxxxc");
console.log("Sdsdscxxxc");
console.log("Sdsdscxxxc");
console.log("Sdsdscxxxc");
console.log("Sdsdscxxxc");
console.log("Sdsdscxxxc");
console.log("Sdsdscxxxc");
console.log("Sdsdscxxxc");
console.log("Sdsdscxxxc");
console.log("Sdsdscxxxc");
console.log("Sdsdscxxxc");
console.log("Sdsdscxxxc");
console.log("Sdsdscxxxc");
console.log("Sdsdscxxxc");

//###////////////////////  Controling The MainSidebar in Profile //////////

const switchIcons = document.querySelector(".sub-main")
const ProfileCard = document.querySelector(".ProfileCard")
const Categories = document.querySelector(".Categories")
const rotateIcon = document.querySelector(".switch-icon span")
const Parametrs = new URLSearchParams(window.location.search);

const Profile =  Parametrs.get('type') == "profile" 

if (switchIcons.classList.contains("reverse") && !Profile){ // change index to HomePage Path 
    ProfileCard.classList.remove("display")
    Categories.classList.add("display")
}else{
    switchIcons.classList.toggle("reverse")
    ProfileCard.classList.add("display")
    Categories.classList.remove("display")
}

switchIcons.addEventListener('click', () => {
    switchIcons.classList.toggle("reverse")
    rotateIcon.classList.toggle("rotate")
    ProfileCard.classList.toggle("display")
    Categories.classList.toggle("display")
})

