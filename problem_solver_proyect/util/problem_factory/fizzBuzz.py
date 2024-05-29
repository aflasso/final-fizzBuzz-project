from util.problem_factory.i_problem_solver import IProblemSolver

class FizzBuzz(IProblemSolver):

    def solve_problem(self, data):
        result = []
        line = ""

        for element in data:
            line =  self.__fizz_buzz(int(element))
            result.append(str(element) + " " +line)
        return result
    

    def __fizz_buzz(self,number: int) -> str:      
        result = str(number)
        fizz_flag = False
      
        if number % 3 == 0:
            result = "Fizz"
            fizz_flag = True

        if number % 5 == 0:
            if fizz_flag:
                result += "Buzz"
                return result
            result = "Buzz"
            
        return result