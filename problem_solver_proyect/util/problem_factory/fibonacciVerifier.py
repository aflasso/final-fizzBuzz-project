from util.problem_factory.i_problem_solver import IProblemSolver

class FibonacciVerifier(IProblemSolver):

    def solve_problem(self, data):
        result = []
        print("tu puta madre fibonacci")

        for number in data:

            line = f"{number} {self.is_fibonacci(number)}"

            result.append(line)

        return result

    def is_fibonacci(self, number):

        if number == 1 or number == 0:
            return True

        numer_1 = 0
        numer_2 = 1

        is_fibonacci = False

        actual_fibonacci_number = 0

        while actual_fibonacci_number <= number:

            if actual_fibonacci_number == number:
                is_fibonacci = True
            
            actual_fibonacci_number = numer_1 + numer_2

            numer_1 = numer_2

            numer_2 = actual_fibonacci_number
        
        return is_fibonacci