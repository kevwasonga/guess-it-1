Here’s a README outlining the objectives, instructions, and testing process for the project:

---

# Guess It - Number Range Prediction Program

## Objectives

In this project, the goal is to use your math skills to predict a range for the next number in a sequence based on a given input. This exercise helps you implement statistical and probability calculations to make informed guesses about the next value.

## Instructions

You are tasked with building a program that, given a sequence of numbers as standard input, predicts a range for the next number in the sequence.

### How it works:
- The program will receive a sequence of numbers, each on a new line. 
- The x-axis represents the position of each number in the sequence (0, 1, 2, ...).
- The y-axis represents the value of the numbers (e.g., 189, 113, 121, etc.).
  
Your task is to print a range (lower and upper limits) for the next number in the sequence. The range should be separated by a space, and the program should output the current number followed by the predicted range for the next number.

### Example

```bash
$ ./your_program
189  --> the current input number
120 200    --> the predicted range for the next number (113 in this case)
113  --> the next input number
160 230    --> the predicted range for the next number (121 in this case)
121  --> the next input number
110 140    --> the predicted range for the next number (114 in this case)
...
```

- Some ranges might be too large or incorrect. You can choose to define a specific range for better accuracy. 
- The purpose of the program is to use your knowledge of statistics, probabilities, and patterns to predict the next number’s range.

## Testing

Your program will be tested extensively for performance and accuracy. Here’s how testing will work:

- **Scoring**: Your score is based on how well you predict the range. The smaller and more accurate your range is, the better your score.
- **Range size**: If the range is too large, your score will decrease. Aim to find a balance between range size and accuracy.
  
### Steps to Run the Test

1. **Create a folder named `student`**:
   - Place the necessary files for your program inside this folder.

2. **Create a shell script named `script.sh`**:
   - This script will contain the commands to run your program.
   - Ensure the script runs from the root directory of the tester.
   
   Example `script.sh` for a program written in JavaScript:
   ```sh
   #!/bin/sh
   # We assume that we are on the root folder, so we have to enter the
   # student folder in order to run the `solution.js` file
   node ./student/solution.js
   ```

3. **Languages**:
   - You can write your program in **Golang**, **JavaScript**, **Rust**, or **Python**.
   
4. **Tester**:
   - Download the provided tester (zip file), and follow the instructions to place the `student` folder in the correct location.
   - Data sets (Data 1, Data 2, and Data 3) will be used to test your program.
   
### Important Note:
Make sure to run the test yourself using the tester provided in the zip file to confirm that everything is working properly.

## Learning Outcomes

This project will help you improve in:
- **Statistical and Probability Calculations**.
- **Scripting and Automation**.
  
For more details, visit the repository:
[Guess It Project](https://learn.zone01kisumu.ke/git/kevwasonga/guess-it-1)

---

This README explains the key objectives and process clearly, giving an overview of what needs to be done and how to run tests on the program.