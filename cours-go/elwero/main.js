// window.onload = function () {
//   document.querySelector("h1").style.color = "blue";
// };

// document.write("<h1>helooo</h1>");
// // window.alert("Hello")

// console.table(["omar", "ahmed", "med"]);

// console.log("Hello from %cJs file", "color:red; font-size:20px");

let swappingCases = "elWERo";
let invertedNumbers = [1, -10, 15, 100, -30];
let ignoreBooleans = "Elz123er4o";

let sw = swappingCases
  .split("")
  .map((el) => (el === el.toLowerCase() ? el.toUpperCase() : el.toLowerCase()))
  .join("");

console.log(sw);

let inv = invertedNumbers.map(function (el) {
  return -el;
});

console.log(inv);

let ign = ignoreBooleans
  .split("")
  .map((el) => (isNaN(parseInt(el)) ? el : ""))
  .join("");

console.log(ign);

// filter

let friends = ["Ahmed", "Sameh", "Sayed", "Asmaa", "Amgad", "Israa"];

let numbers = [11, 20, 2, 5, 17, 10];

let filterFriends = friends.filter(function (el) {
  return el.includes("A");
});

console.log(filterFriends);

let sentence = "I Love Foood Code Too Playing Much";

let smallWords =  sentence.split(" ").filter(function(el) {
  return el.length <= 4
}).join(" ")

console.log(smallWords);

let mix = "A13BS2ZX"

let ell = mix.split("").filter(function (el) {
  return !isNaN(parseInt(el))
})

console.log(ell);


let choose = ell.map(function (el) {
  return el * el
})

console.log(choose);


let nums = [10, 20, 15, 30]

let add = nums.reduce(function (acc, current, index, arr) {
  console.log(`Accumulator => ${acc}`);
  console.log(`Current Element => ${current}`);
  console.log(`Current Element Index => ${index}`);
  console.log(`Array ${arr}`);
  console.log(`/////////////////////////////`);
  return acc + current
})

let numss = [10, 20, 15, 30]

let addd = numss.reduce(function (acc, current, index, arr) {
  console.log(`Accumulator => ${acc}`);
  console.log(`Current Element => ${current}`);
  console.log(`Current Element Index => ${index}`);
  console.log(`Array ${arr}`);
  console.log(`/////////////////////////////`);
  return acc + current
}, 5)

console.log(add);
console.log(addd);


let removeChar = ["E", "@", "@", "L", "Z", "@", "@", "E", "R", "@", "O"]


let finalString = removeChar.filter((el) => {
  return !el.includes("@")
}).reduce((acc, curr) => {
  return `${acc}${curr}`
})

console.log(finalString);

let allLis = document.querySelectorAll("ul li")
allLis.forEach(function (el){
  el.onclick = function () {
    // this.style.display = "none" 
    allLis.forEach(function (el) {
      el.classList.remove("active")
    })
    this.classList.add("active")
    
  }
})

 