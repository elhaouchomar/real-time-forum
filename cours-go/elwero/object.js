let myVar = "country";

let user = {
  // Properties
  theName: "Omar",
  theAge: 25,
  country: "USA",
  "country of": "Morocco",
  // Methodes
  sayHello: function () {
    return "hello";
  },
};

console.log(user.theName);
console.log(user["theAge"]);
console.log(user["country of"]);
console.log(user.country);
console.log(user.myVar);
console.log(user[myVar]);
console.log(user.sayHello());

let person = {
  // properties
  name: "omar",
  age: 25,
  skills: ["html", "css", "js"],
  available: false,
  address: {
    usa: "California",
    morroco: {
      one: "Taounate",
      two: "fes",
    },
  },
  // methodes
  checkav: function () {
    if (this.available === true) {
      return "free for work";
    } else {
      return "not free";
    }
  },
};

console.log(person.name);
console.log(person.age);
console.log(person.skills);
console.log(person.skills.join(" "));
console.log(person.skills[2]);
console.log(person.address.morroco.one);
console.log(person["address"].morroco.one);
console.log(person["address"].morroco["one"]);
console.log(person.checkav());

let some = new Object(); // it the same let some = {}
console.log(some);

some.age = 66;
some["country"] = "morocco";

console.log(some);

let copyobj = Object.create(person);

copyobj.age = 30;

console.log(copyobj.age);

let obj1 = {
  prop1: 1,
  meth1: function () {
    return this.prop1;
  },
};

let obj2 = {
  prop2: 2,
  meth2: function () {
    return this.prop2;
  },
};

let targetObject = {
  prop1: 100,
  prop3: 3,
};

let finalObject = Object.assign(targetObject, obj1, obj2);
finalObject.prop1 = 200;
finalObject.prop4 = 4;

console.log(finalObject);

let newobject = Object.assign({}, obj1, { prop5: 5, prop6: 6 });
console.log(newobject);
