WELCOME



Yo Welcome!!! so its actually one of my first projects out of fellowship and the tool I’ve created is a text editor and completion tool that functions seamlessly, handling more edge cases than not.



DESCRIPTION:

Let me explain…. In the program written above, we take user inputs as arguments and edit based on attached commands. All other 6 functions handle different test/edge cases, the main function handles filing from opening input-file as one argument to writing to output-file as another argument with Processing functions as the 3rd argument in this program.

In the nitty grit of this program, the functions output modified cases and bases through annotated commands of the user input (type string), the strings also get modified in cases where quotes(single and double) and punctuations appear misplaced, also in instances where there’s a grammatical error due to A/An misplacement.



LOGIC AND IMPLEMENTATION:

starting from our main function (func main) that handles filing. First we open the input-file that takes user input, read then write to the output file. Having structured the func main this way, the user is able to input non-modified texts into the input.txt file and then run the program which automatically creates an output file containing the modified texts.

Followed by other functions namely:



Func handleconv (inside the “base.go” file) – the logic here states that after closing in the spaces within open bracket and the number base, separate the string into sub-strings then find any substring thats a number base (i.e: (hex) or (bin)) and then convert the previous substring to the expected output-base (i.e: (decimal or binary)). In this program I am looking to transform only hexadecimal and binary number base to decimal number base, followed by an append expression that deletes the command after transformation. Our final return is a string-joined output.



Func handlecases (inside the “casein.go” file) – firstly, I got the substrings of parameters by the strings.Fields function, then proceed to check if the substring at current index in our scope is same as the case-conversion commands(up, low, cap), if true then the previous sub-string should be effected according to case-conversion command, else if there is another sub-string attached to the command (I.e: cap, 2 or 3 or 4 etc...)

then we trim the closing bracket off the command and convert the sub-string at current index from string to an integer using the strconv package then take the integer as number of prior sub-strings to be effected by command..after effecting the whole command, command (up, num) is then deleted from output string by appending the sub-strings without the commands



func handleAtotheAn( inside the “atothean.go” file) – first, split the string into sub-strings and then loop through to check if substring at current index is “a” or “A”, we then open the scope to check the next character in our string, at this point we condition the program to check if there is a punctuation or quote after the substring “a” or “A”, there should be a skip and continue to check next character and if a vowel sound or “H” is found then the program changes previous substring (“a” or “A”) to “An”. The same logic applies if the sub-string before a consonant is “an” or “An”..Making for proper articulation and better readability. The final return is the joined on output.



Func punc(found in the “punch.go” file) -

for the proper arrangement and behavior of punctuations, ive declared variables of type string named: punch(for holding punctuation characters) alph(for holding alphabetical characters) and output(for holding output post-modification)

what happens next is we iterate through each character and condition output to add the current slice into modified string, only if the current character is a whitespace and the next character exists and contains punctuation. Followed by another if statement that only adds a space to the modified strings if current and next character is a punctuation followed by another character which is an alphabet and not a whitespace. Then add space and finally return output.



Func handlePunchnQ(found inside the “quoti.go” file) -

this function takes care of quotes, both single and double, arranging them by trimming all spaces between the words and quotes INSIDE a quoted word or sentence (left and right) and leaves a trailing whitespace to the left and right OUTSIDE the quoted word or sentence. The tool used here is a super-tool if I say so myself. The “regexp” function takes note of patterns and by using the mustcompile tool im able to match all patterns that involve extra spaces inside a quoted word or sentence , spaces between word and quote are then replaced with a tight fitting between word and quote by using the ReplaceAllString tool.

After these modifications. The string is then returned, modified without much splitting and any conversions whatsoever.

KNOWLEDGE ACQUIRED:
*STRING MANIPULATION
*EXPLICIT WORKING OF INDEX AND SLICES
*LOGICAL IMPLEMENTATION USING GOLANG
*ERROR HANDLING
*FILE OPENING, READING AND WRITING
*UPGRADE ON PSEUDOCODE TO CODING APPROACH
*SCOPING AND CONDITIONING

TARGETED WEAK AREAS:
*STRING MANIPULATION USING REGEXP
