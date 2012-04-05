// This package is more of redundant since copying the same material from the comments of the coding

/*

Usage: expr Expression

Expression has to be valid and match in parentheses. 
A slash has to be present to avoid pitfalls like inclusion of opening and closing braces
Uses the shunt yard Algorithm to convert infix to post fix and then calculate post fix expression.

*/

/*

Program to evaluate a expression as per the expr command in Unix
Authors: Juan Golphin, Soham Sadhu

References:
http://en.wikipedia.org/wiki/Shunting-yard_algorithm Shunting yard algorithm for conversion from infix to post fix expression.
http://en.wikipedia.org/wiki/Reverse_Polish_notation The post fix evaluation algorithm has been referenced from this place.

Please note the I have used BODMAS rule ( Brackets off, Division, Multiplication, Addition and Subtraction ) in respective order of precedence. So division has higher precedence than multiplication rather than having equal precedence. This results in some results differing from expr like the assignment example evaluates to -6 instead of -7 due to precedence rule. Also the modulus precedence is below division and multiplication but higher than addition and subtraction.

*/

/*
The packages imported and used
fmt for printing and stuff
os for getting the arguments from command line
strconv for converting and checking if the command line arguments are numbers or operators
strings package this was imported just for fun used one function of it Join to get a hang of slice and how go joins string array to get a string
*/

/*
The function precedence() is the nested switch case which just returns the precedence as a boolean value.
If operator 1 precedence is greater than operator 2 then return true else return false.
*/

package documentation
