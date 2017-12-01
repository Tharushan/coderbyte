/*  Using the JavaScript language, have the function LongestWord(sen) take the sen parameter being passed and return the largest word in the string.
 *  If there are two or more words that are the same length, return the first word from the string with that length.
 *  Ignore punctuation and assume sen will not be empty.
 */

function LongestWord(sen) {
  let longest = "";
  let arr = sen.replace(/[^\w]/g , ' ').split(' ');
  arr.forEach(word => {
    if (word.length > longest.length)
      longest = word;
  })
  return longest;
}

console.log(LongestWord("a beautiful sentence^&!"));
console.log(LongestWord("letter after letter!!"));
console.log(LongestWord("a confusing /:sentence:/[ this is not!!!!!!!~"));
