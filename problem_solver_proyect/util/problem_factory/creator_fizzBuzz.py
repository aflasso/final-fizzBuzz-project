from util.problem_factory.i_problem_creator import IProblemCreator
from util.problem_factory.i_problem_solver import IProblemSolver
from util.problem_factory.fizzBuzz import FizzBuzz

class CreatorFizzBuzz(IProblemCreator):

    def factory_method(self) -> IProblemSolver:
        return FizzBuzz()
        