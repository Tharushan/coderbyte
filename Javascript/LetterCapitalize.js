/* Have the function LetterCapitalize(str) take the str parameter being passed and capitalize the first letter of each word.
 * Words will be separated by only one space.
 */

function LetterCapitalize(str) {
  return str.split(' ').map(word => word.charAt(0).toUpperCase() + word.slice(1)).join(' ');
}

console.log(LetterCapitalize('hello world'));
console.log(LetterCapitalize('my car is red.'));
