//###////////////////////  Controling The MainSidebar in Profile //////////
export function profileEffect(){
    const switchIcons = document.querySelector(".sub-main")
    const MessageCard = document.querySelector(".MessageCard")
    const Categories = document.querySelector(".Categories")
    const rotateIcon = document.querySelector(".switch-icon span")
    if (switchIcons.classList.contains("reverse")){ // change index to HomePage Path 
        MessageCard.classList.remove("display")
        Categories.classList.add("display")
    }else{
        switchIcons.classList.toggle("reverse")
        MessageCard.classList.add("display")
        Categories.classList.remove("display")
    }
    switchIcons.addEventListener('click', () => {
        switchIcons.classList.toggle("reverse")
        rotateIcon.classList.toggle("rotate")
        MessageCard.classList.toggle("display")
        Categories.classList.toggle("display")
    })
}