# STATEXPERT

# 💻 Tech Stack:
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

## DESCRIPTION
STATEXPERT is a command-line tool designed to process a list of numbers from a file and calculate statistical metrics such as the average, median, variance, and standard deviation. The program reads the numbers, processes them, and outputs the calculated statistics to the console.

## AUTHORS
- [hujaafar](https://github.com/hujaafar)

## Usage

### How to Run
1. git clone repoLink
2. cd into the directory
3. go run main.go filename

## Implementation Details

### Algorithm
The program reads integers from the provided file, ignoring any invalid entries. It then calculates the statistical metrics: the average is computed by dividing the sum of all numbers by their count; the median is determined as the middle value in a sorted list, or the average of the two middle values if the list length is even. Variance is calculated as the average of the squared differences from the mean, and the standard deviation is simply the square root of the variance. Finally, the program outputs these statistics, rounded for clarity, directly to the console.
