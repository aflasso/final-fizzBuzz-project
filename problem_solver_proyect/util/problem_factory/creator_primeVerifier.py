from util.problem_factory.i_problem_creator import IProblemCreator
from util.problem_factory.i_problem_solver import IProblemSolver
from util.problem_factory.primeClasifier import PrimeClassifier

class CreatorPrimeClasifier(IProblemCreator):

    def factory_method(self) -> IProblemSolver:
        return PrimeClassifier()