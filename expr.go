/*

Program to evaluate a expression as per the expr command in Unix
Authors: Juan Golphin, Soham Sadhu

References:
http://en.wikipedia.org/wiki/Shunting-yard_algorithm Shunting yard algorithm for conversion from infix to post fix expression.
http://en.wikipedia.org/wiki/Reverse_Polish_notation The post fix evaluation algorithm has been referenced from this place.

Please note the I have used BODMAS rule ( Brackets off, Division, Multiplication, Addition and Subtraction ) in respective order of precedence. So division has higher precedence than multiplication rather than having equal precedence. This results in some results differing from expr like the assignment example evaluates to -6 instead of -7 due to precedence rule. Also the modulus precedence is below division and multiplication but higher than addition and subtraction.

*/

// Declaring the main package
package main

/*
The packages imported and used
fmt for printing and stuff
os for getting the arguments from command line
strconv for converting and checking if the command line arguments are numbers or operators
strings package this was imported just for fun used one function of it Join to get a hang of slice and how go joins string array to get a string
*/
import (
	"fmt"
	"os"
	"strconv"
	//"strings"
)

/*
The variables declared below are global because of my intial laziness and hence a left over of bad design.
The below variables are used for evaluation of the postfix expression that I get in intermediate stage.
*/
var (
	stack_top int = 0
	numbers   [30]int
)

/*
Usage function created a separate one so that any error could be handled.
This is not exhaustive and could have been made better with more error cases thrown in.
*/
func usage(err string) {
	switch err {
	case "less_args":
		fmt.Print("usage: ./expr.out 'expression you want to evaluate' ")
	case "first_operator":
		fmt.Print("operator in the first position cannot be evaluated")
	case "incorrect_parentheses":
		fmt.Print("the number of parentheses do not match or the nesting is incorrect")
	}
}

/*
The below function precedence() is the nested switch case which just returns the precedence as a boolean value.
If operator 1 precedence is greater than operator 2 then return true else return false.
*/
func precedence(operator1, operator2 string) bool {
	switch operator1 {
	case "/":
		if operator2 == "*" || operator2 == "%" || operator2 == "+" || operator2 == "-" || operator2 == "(" {
			return true
		} else {
			return false
		}
	case "*":
		if operator2 == "%" || operator2 == "+" || operator2 == "-" || operator2 == "(" {
			return true
		} else {
			return false
		}
	case "%":
		if operator2 == "+" || operator2 == "-" || operator2 == "(" {
			return true
		} else {
			return false
		}
	case "+":
		if operator2 == "-" || operator2 == "(" {
			return true
		} else {
			return false
		}
	case "-":
		if operator2 == "(" {
			return true
		} else {
			return false
		}
	}
	// The below statement is required else go compiler shoots a error.
	return false
}

/*
The below function is just a straight forward implementation of Edgser Dijkstra's Shunt-yard Algorithm to convert a expression from infix to postifx.
Note that in the function we are passing and returning array of strings rather than strings. Since we want tokens not individual characters.
*/
func shunt_yard(input_string [30]string) ([30]string, bool) {
	/*
	 Have to give the arrays a definite length. Was not adventurous to try [...] notation
	 Unlike other languages [] declaration and hoping for a dynamic declaration does not work. Since [] stands for slices.
	 flag variable initialised no error so true
	 A output_queue string array that will be cast to result string and sent back.
	 A operator_stack required in the shunt yard algorithm.
	 Counter or index for both the queue and stack.
	*/
	var (
		output_queue           [30]string
		operator_stack         [30]string
		output_queue_counter   int  = 0
		operator_stack_counter int  = 0
		flag                   bool = true
		//result                 string
	)

	/*
			 Do not know why gofmt command treats label with exception other wise hell bent of indenting.
			 This is the main loop that reads the tokens one by one and have labelled so can break off in case of invalid parentheses
		         condition set on line 165.
			*/
token_read_loop: for i := 0; i < len(input_string); i++ {

		/*
			The below if condition checks to see if the token is operator or number.
			If number then dealt in if loop; if operator then dealt in the else loop below.
			Notice use of blank variable (_) in line 132. This is since I am not interested in numerical value but only boolean value
			if the token is number or operator. If you are not gonna use a variable then declare same as blank else it will be a error.
			Remember you have to capture each of the return type else it will be a error.
		*/
		if _, err := strconv.Atoi(string(input_string[i])); err == nil {
			output_queue[output_queue_counter] = string(input_string[i])
			output_queue_counter++
		} else {

			// As per algorithm if "(" then just push it into operator stack
			if string(input_string[i]) == "(" {
				operator_stack[operator_stack_counter] = "("
				operator_stack_counter++
			}

			//If you encounter a closing parentheses then start poping off the operator stack and push onto output queue.
			if string(input_string[i]) == ")" {
				for operator_stack_counter > 0 {

					// As soon closing parentheses encountered break out of the immedidate for loop
					if string(operator_stack[operator_stack_counter-1]) == "(" {
						operator_stack_counter--
						break
					}

					// Pushing the operators onto output queue
					output_queue[output_queue_counter] = string(operator_stack[operator_stack_counter-1])
					output_queue_counter++
					operator_stack_counter--

				}

				/*
				 At this point if stack is empty ( no opening parentheses found ) then it is a error and break out of the token loop.
				 Please note we are still in the closing parentheses loop.
				*/
				if operator_stack_counter == 0 {
					output_queue[0], flag = "incorrect_parentheses", false
					break token_read_loop
				}
			}

			/*
			  The loop that inserts the operator token into stack if the top variable has lower precedence than the token.
			  Else it keeps poping off the stack the operator and putting onto the output queue till the end of stack or
			  till the token variable has higher precedence.
			*/
			// couldn't get an idea of how to break this long expression into two lines will have to see documentation again.
			if string(input_string[i]) == "/" || string(input_string[i]) == "%" || string(input_string[i]) == "*" || string(input_string[i]) == "+" || string(input_string[i]) == "-" {
				/*
				 GO supports short circuiting of the conditional loops. Advantage I could write the below condition for for loop
				 that makes it while.
				 Putting the stack counter condition previous to passing stack call avoids compiler raising error of array access
				 beyond bounds in case trying to insert first operator like stack counter being zero.
				*/
				for (operator_stack_counter > 0) && precedence(string(operator_stack[operator_stack_counter-1]), string(input_string[i])) {
					output_queue[output_queue_counter] = string(operator_stack[operator_stack_counter-1])
					operator_stack_counter--
				}
				operator_stack[operator_stack_counter] = string(input_string[i])
				operator_stack_counter++
			}
		}
	}

	// If the result from token_read_loop failure; then return value immediately.
	if flag == false {
		return output_queue, flag
	}

	/*
	  If there are still operators on the stack then pop them and push onto the result queue.
	  If you find a opening parentheses on the stack then it is an error.
	*/
	if operator_stack_counter > 0 {
		for operator_stack_counter > 0 {
			if string(operator_stack[operator_stack_counter-1]) == "(" {
				output_queue[0], flag = "incorrect_parentheses", false
				return output_queue, flag
			}
			output_queue[output_queue_counter] = string(operator_stack[operator_stack_counter-1])
			output_queue_counter++
			operator_stack_counter--
		}
	}

	// Below two lines make variables for use of strings.Join() function
	//separator := ""
	//temp := output_queue[0:len(output_queue)]

	// The Join function takes a slice not array and a separator that will be filled between array elements when joined together.
	//result = strings.Join(temp, separator)
	return output_queue, flag
}

/*
 Below is the push function used by the postfix evaluator.
 Since it works on global data (line# 35) no need to return any value that is modified by this function.
*/
func push(num int) {
	numbers[stack_top] = num
	stack_top++
}

// Function that calculates the postfix expression from the shunt yard function.
func calculate(str [30]string) {
	var result int // The result variable stores evaluation that will be returned.

	// If expression is just one single thing just return the same
	if len(str) == 1 {
		fmt.Println(str)
	} else {
		// Read the tokens from the post fix expression.
		for i := 0; i < len(str); i++ {
			// If the token is a integer then push it into the single stack.
			if stk_var, err := strconv.Atoi(string(str[i])); err == nil {
				push(stk_var)
			} else {
				// If the token is not a number then it is a operator.
				op_var := string(str[i])

				// Use the operator to calculate the top 2 numbers from the stack and then push the result back to stack.
				switch op_var {
				case "+":
					result = numbers[stack_top-2] + numbers[stack_top-1]
					stack_top -= 2
					push(result)
				case "*":
					result = numbers[stack_top-2] * numbers[stack_top-1]
					stack_top -= 2
					push(result)
				case "/":
					result = numbers[stack_top-2] / numbers[stack_top-1]
					stack_top -= 2
					push(result)
				case "%":
					result = numbers[stack_top-2] % numbers[stack_top-1]
					stack_top -= 2
					push(result)
				case "-":
					result = numbers[stack_top-2] - numbers[stack_top-1]
					stack_top -= 2
					push(result)
				}
			}
		}
		// Print the result.
		fmt.Println(result)
	}
}

func main() {
	// The result and the error variables.
	var (
		s [30]string
		p bool
	)
	// The usage line one of the cases.
	if len(os.Args) < 2 {
		usage("less_args")
		os.Exit(1)
	}

	// Taking in the input and stripping out the back slashes or the escape characters.
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] != "\\" {
			s[i-1] = os.Args[i]
		}
	}

	// If the first token is a operator then error print usage.
	if string(s[0]) == "/" || string(s[0]) == "%" || string(s[0]) == "*" || string(s[0]) == "+" || string(s[0]) == "-" || string(s[0]) == "%" {
		usage("first_operator")
		os.Exit(1)
	}

	// Get the postfix expression from shunt yard and if success then call calculate else print error message.
	if s, p = shunt_yard(s); p {
		calculate(s)
	} else {
		usage("incorrect_parentheses")
	}
}
