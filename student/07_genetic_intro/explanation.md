# EXPLANATION.md

## Task 1 — Nine Go Labs

This section documents the implementation and analysis of nine Go programming exercises. Following the worksheet methodology of "predict, implement, run, and explain," each lab addresses a specific computational or logical concept.

### 01 — Data Types
The first exercise focuses on foundational data type selection within the context of a campus festival registration system. Age was classified as an integer (int), mass as a floating-point number (float64), registration status as a boolean (bool), and student ID as a string. The student ID must be a string because it is an identifier rather than a mathematical operand, and using an int would strip the leading zeros. The implementation produced the expected output of 19 | 62.5 | true | 00123, confirming that the chosen types accurately represent the data and preserve semantic meaning.

### 02 — Café Sequence
This exercise highlights the significance of execution order in arithmetic operations. The scenario involves purchasing two coffees with a discount applied. The subtotal must be computed prior to the discount to avoid inaccurate calculations. The program successfully generated "Subtotal: 36000" and "Total: 31000." The primary assumption is that the discount cannot exceed the subtotal; otherwise, the total would become negative, which would require additional validation logic.

### 03 — Factor Checker
The third lab addresses input validation prior to mathematical operations. The objective is to determine if x divides y without a remainder. The implementation uses the condition x <= 0 || y <= 0 to validate input before executing the modulo operation y % x == 0. This prevents division-by-zero errors. The program correctly handled normal, boundary, and invalid test cases, such as 5 20, 20 5, and 0 20.

### 04 — Population
The population dashboard serves as a simplified mathematical model for demographic changes. The formula final = initial + births + arrivals - deaths - departures was implemented using the provided values. The expected output of Population: 1011 was achieved. A negative result would imply a flawed model or anomalous data, as population counts realistically cannot fall below zero. This illustrates that models are inherently simplifications of reality and require appropriate assumptions.

### 05 — Number Sequence
This exercise requires generating a number sequence where each new number is the sum of the previous two, mimicking a composer's pattern. For the input 5 and 3, the program computed the third, fourth, and fifth values as 8, 11, and 19, respectively. This sequential generation differs from parallel execution, as each calculation depends on the state of the preceding numbers. The implementation confirms that order and state management are crucial in sequence generation.

### 06 — Gravity
The gravity exercise demonstrates the distinction between mass and weight, emphasizing the importance of units and floating-point precision. Mass remains constant, while weight depends on gravitational force. Using the rules W = massKg × 9.8 and W = massKg × 9.8 × 0.38, the program output Earth: 686.0 N and Mars: 260.7 N. The use of float64 is essential here to maintain the fractional precision required for accurate physical measurements.

### 07 — Genetic Intro
This lab introduces genetic algorithm operations on four-bit strings. Fitness is calculated by counting the number of one-bits. Crossover involves mixing the first two bits of Parent A with the remaining bits of Parent B, resulting in the child 1010. A mutation is then applied to change the second bit, resulting in 1110 and increasing the fitness from 2 to 3. This exercise demonstrates that mutation does not always improve fitness; it is a stochastic operation that can introduce both beneficial and detrimental changes.

### 08 — Cube Root
The cube-root check tests precise arithmetic by verifying x*x*x == y using multiplication rather than floating-point approximation. The program correctly returned true for 5 125 and -3 -27, and false for 12 12. Negative numbers can have negative cubes because (-3)³ equals -27. The primary limitation is that extremely large integers may overflow the int type, which would require a larger data type or arbitrary-precision arithmetic.

### 09 — Bonus GA
The final exercise implements a population-based genetic algorithm to evolve four-bit strings toward the optimal solution of 1111. The algorithm calculates fitness, selects parents using tournament selection, performs crossover, and applies mutation based on a defined probability. It terminates when the fitness reaches chromosomeSize or when the maximum generation limit is reached. The algorithm does not guarantee finding 1111 within the budget, as random exploration can make candidates worse and the search space may not be fully traversed.

## Task 2 — Fill in the blank

### A — Git and GitHub
1. fork
2. clone
3. switch
4. status, diff
5. add
6. push
7. main
8. origin

### B — Go types, constants and operators
9. 3, int
10. 1, remainder
11. 3.5, implicitly
12. rune, byte
13. string, zeros
14. const, compile
15. false, true, false
16. not
17. assigns, compares
18. 14, 20
19. false, approximately
20. "A", Itoa

### C — Complete the program
21. const, float64, >
22. 66000 true
23. If the conversion is removed, the expression price * qty becomes float64 * int, which causes a compile error because Go does not convert types implicitly.