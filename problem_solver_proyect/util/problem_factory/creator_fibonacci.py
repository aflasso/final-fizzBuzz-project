from util.problem_factory.i_problem_creator import IProblemCreator
from util.problem_factory.i_problem_solver import IProblemSolver
from util.problem_factory.fibonacciVerifier import FibonacciVerifier

class CreatorFibonacci(IProblemCreator):

    def factory_method(self) -> IProblemSolver:
        return FibonacciVerifier()