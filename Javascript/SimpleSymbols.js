/* Using the JavaScript language, have the function SimpleSymbols(str) take the str parameter being passed and determine if it is an acceptable sequence by either returning the string true or false.
 * The str parameter will be composed of + and = symbols with several letters between them (ie. ++d+===+c++==a) and for the string to be true each letter must be surrounded by a + symbol.
 * So the string to the left would be false. The string will not be empty and will have at least one letter.
 */

function SimpleSymbols(str) {
  for (let i = 0; i < str.length; i++) {
    if (/[^a-z]/i.test(str[i])) continue;
    if (str[i - 1] && str[i + 1] && '+' === str[i - 1] && '+'=== str[i + 1])
      continue;
    else
    return false;
  }
  return true;
}

console.log(SimpleSymbols("+d===+a+")); // false
console.log(SimpleSymbols("+a=")); // false
console.log(SimpleSymbols("2+a+a+")); // true
console.log(SimpleSymbols("==a+")); // false
