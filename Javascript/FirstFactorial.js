/* Using the JavaScript language, have the function FirstFactorial(num) take the num parameter being passed and return the factorial of it (e.g. if num = 4, return (4 * 3 * 2 * 1)).
 * For the test cases, the range will be between 1 and 18 and the input will always be an integer.
 */

function FirstFactorial(num) {
  let sum = 1;
  for (; num > 1; num--) {
    sum *= num
  }
  return sum;
}

console.log(FirstFactorial(1));
console.log(FirstFactorial(3));
console.log(FirstFactorial(5));
console.log(FirstFactorial(6));
