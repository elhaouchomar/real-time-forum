import { mapClass, mapArray } from "./map.js";
import { diedhero, MapHero, heart, nbrOfKilled } from "./bomberman.js";
import { EnemyGenerator } from "./enemy.js";

export let numbreofenemy = 5
// When the block is destroyed and becomes gress, The Hero go below the gress
// Maybe because we draw the map in index.js and bomberman.js
let mapSence = document.getElementById("map")

export const tileImage = new Image
export const blockImage = new Image
export const greenBlockImage = new Image
export const enemymoveemage = new Image
export const boombimage = new Image
export const exploImg = new Image
const heartimg = new Image

tileImage.src = `./assets/tile.png`
blockImage.src = `./assets/block.png`
greenBlockImage.src = `./assets/greenBlock.png`
enemymoveemage.src = `./assets/enemy.png`
boombimage.src = `./assets/bomb.png`
exploImg.src = `./assets/3.png`

heartimg.src = `./assets/heart.png`

export const directions = {
    up: "ArrowUp",
    down: "ArrowDown",
    left: "ArrowLeft",
    right: "ArrowRight",
    destroy: "destroy",
};

export const keys = {
    ArrowUp: directions.up,
    ArrowDown: directions.down,
    ArrowLeft: directions.left,
    ArrowRight: directions.right
};

export const heroConfig = {
    tileSize: 32,
    initialGridX: 1,
    initialGridY: 1,
    speed: 2
};


export let enemyCoords = {
    1: {
        x: -1,
        y: -1,
        w: 0,
        z: 0
    },
    2: {
        x: -1,
        y: -1,
        w: 0,
        z: 0
    },
    3: {
        x: -1,
        y: -1,
        w: 0,
        z: 0
    },
    4: {
        x: -1,
        y: -1,
        w: 0,
        z: 0
    },
    5: {
        x: -1,
        y: -1,
        w: 0,
        z: 0
    }
}

export function SetEnemyCoords(x, y, e) {

    enemyCoords[e].x = x
    enemyCoords[e].y = y

}


function khamazat() {
    let enemies = document.querySelectorAll(".enemy")

    for (const e of enemies) {
        e.remove()
    }
}


let game = document.getElementById("game")
export let settingScreen = document.getElementById("settingScreen")
export let gameSetting = document.getElementById("gameSetting")
settingScreen.classList.add("animate");

game.insertBefore(settingScreen, game.firstChild)

export let instructions = document.getElementById("instructions")
export let title = document.getElementById("title")
// let gameAudio = document.getElementById("gameAudio")
export let startGame = true
export let gamePaused = false


// let audio = new Audio("./assets/playGame.mp3")
let Enemies
let hero
let gameTime = 200
let id
let heartLeft = 3

export function startGameLoop() {
    let CountPerFrame = 0
    let numOfHeart = document.getElementById("numOfHeart")

    numOfHeart.innerText = heartLeft
    let Time = document.getElementById("Time")
    Time.innerText = "Time" + " " + gameTime

    const gameLoop = () => {
        if (!gamePaused) {

            hero.moveHero();
            CountPerFrame += 16.7
            if (CountPerFrame >= 1000) {
                gameTime--
                Time.innerText = "Time" + " " + gameTime
                CountPerFrame = 0
            }
            hero.render();
            if (gameTime <= 0) {
                alert("GAME OVER")
                cancelAnimationFrame(id)
                gameTime = 1
                window.location.href = 'index.html'
            }

            for (const e of Enemies) {
                e.move()
            }

            if (diedhero) {
                settingScreen.style.opacity = ".1"
                gameSetting.style.opacity = ".5"
                settingScreen.classList.remove("animate");
                if (numOfHeart == 0) {
                    title.innerHTML = "Game Over"

                } else {
                    title.innerHTML = "You are Dead"
                    instructions.innerHTML = "Click D to continue"
                    cancelAnimationFrame(id)
                }
            }
            id = window.requestAnimationFrame(gameLoop);
        }
    };
    gameLoop();
}

window.addEventListener('keydown', (e) => {
    if ((e.code == "KeyR" && gamePaused)) {
        window.location.href = 'index.html'
    }
    if (e.code == "KeyP" || ((diedhero) && (e.code == "KeyD"))) {


        if ((diedhero) && (e.code == "KeyD")) {
            startGame = true
            cancelAnimationFrame(id)
            gameTime = 200
            khamazat()

            if (heartLeft == 1) {
                alert("GAME OVER")
                window.location.href = 'index.html'
            }

            heartLeft -= 1
            numOfHeart.innerText = heartLeft
        }

        if (startGame) {

            settingScreen.classList.remove("animate");
            settingScreen.style.opacity = "0"
            gameSetting.style.opacity = "0"
            startGame = false

            hero = new MapHero(mapSence);
            hero.initializeControls()
            
            let E = new EnemyGenerator(mapSence, numbreofenemy - nbrOfKilled)
            Enemies = E.init()
            console.log("enemyCoords Afetr Creation", enemyCoords);

            startGameLoop()

            // audio.play()
            // audio.pause()

            // diedhero=false
            return
        }

        if (!gamePaused && !startGame && !diedhero) {

           // audio.pause()
            gamePaused = true
            settingScreen.style.opacity = "0.1"
            gameSetting.style.opacity = "1"
            settingScreen.classList.add("animate");
            title.innerHTML = "Game is Paused"
            instructions.innerHTML = "Press C to Continue or R to Restart"
        }
    }
    if (e.code == "KeyC") {
        if (gamePaused) {
            settingScreen.style.opacity = "0"
            gameSetting.style.opacity = "0"
            settingScreen.classList.remove("animate");
            // audio.play()
            gamePaused = false
            startGameLoop()
        }
    }
});

let level1 = new mapClass()
level1.drawMap(mapArray, .2)
