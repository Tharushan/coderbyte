/* Have the function TimeConvert(num) take the num parameter being passed and return the number of hours and minutes the parameter converts to (ie. if num = 63 then the output should be 1:3).
 * Separate the number of hours and minutes with a colon.
 */

function TimeConvert(num) {
  let hours = parseInt(num / 60);
  let minutes = num - 60 * hours;
  return `${hours}:${minutes}`;
}

console.log(TimeConvert(120)) // 2:0
console.log(TimeConvert(58)) // 0:58
console.log(TimeConvert(63)) // 1:3
