const myObj = { Mark: 33, Sveta: 31 };
const myObjIterable = Object.entries(myObj);
const myMap = new Map(myObjIterable);

for (const entry of myMap) {
  console.log(entry);
}

for (const property in myObj) {
  console.log(property);
}

let i = 0;
while (i < 10) {
  i++;
  console.log(i);
}
