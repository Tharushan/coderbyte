/* Have the function LetterChanges(str) take the str parameter being passed and modify it using the following algorithm.
 * Replace every letter in the string with the letter following it in the alphabet (ie. c becomes d, z becomes a).
 * Then capitalize every vowel in this new string (a, e, i, o, u) and finally return this modified string.
 */

function LetterChanges(str) {
  let newString = ""
  for (let i = 0; i < str.length; i++){
    let newLetter;
    if (/[a-z]/i.test(str[i])) {
      if (str[i] === 'z') {
        newLetter = 'a';
      } else if (str[i] === 'Z') {
        newLetter = 'A';
      } else {
        newLetter = String.fromCharCode(str.charCodeAt(i) + 1);
      }
    } else {
      newLetter = str[i];
    }
    newLetter = (['a', 'e', 'i', 'o', 'u'].indexOf(newLetter) > -1) ? newLetter.toUpperCase() : newLetter;
    newString += newLetter;
  }
  return newString;
}

console.log(LetterChanges("hello world"));
console.log(LetterChanges("sentence"));
console.log(LetterChanges("this long cake@&"));
console.log(LetterChanges("abcdefghijklmnopqrstuvwxyz"));
