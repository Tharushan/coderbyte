/* Have the function FirstReverse(str) take the str parameter being passed and return the string in reversed order.
 *   For example: if the input string is "Hello World and Coders" then your program should return the string sredoC dna dlroW olleH.
 */

function FirstReverse(str) {
  let reversed = "";
  for (let i = str.length - 1; i >= 0; i--) {
    reversed += str[i];
  }
  return reversed;
}

console.log(FirstReverse('Hello World'));
console.log(FirstReverse('123456789'));
