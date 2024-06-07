from util.problem_factory.creator_primeVerifier import CreatorPrimeClasifier
from util.problem_factory.creator_fizzBuzz import CreatorFizzBuzz
from util.problem_factory.creator_fibonacci import CreatorFibonacci
from util.sockets import createProblem
import pytest
import random


@pytest.fixture
def fibonacci():
    return createProblem("Fibonacci")

@pytest.fixture
def fizzBuzz():
    return createProblem("FizzBuzz")

@pytest.fixture
def prime():
    return createProblem("Prime")


def test_facthoy_method(fizzBuzz, fibonacci):

    problem = createProblem("FizzBuzz")
    
    assert isinstance(problem, CreatorFizzBuzz)

    problem = createProblem("Fibonacci")
    assert isinstance(problem, CreatorFibonacci)

    problem = createProblem("Prime")
    assert isinstance(problem, CreatorPrimeClasifier)



def test_solve_problems(fizzBuzz, fibonacci, prime):

    data = [1,2,3,4,5,6,7,8,9,10]


    result = fizzBuzz.solve_problem(data)
    assert result == ["1 1", "2 2", "3 Fizz", "4 4", "5 Buzz", "6 Fizz", "7 7", "8 8", "9 Fizz", "10 Buzz"]

    result = fibonacci.solve_problem(data)
    assert result == ["1 True", "2 True", "3 True", "4 False", "5 True", "6 False", "7 False", "8 True", "9 False", "10 False"]

    result = prime.solve_problem(data)
    assert result == ["1 prime", "2 prime", "3 prime", "4 cuadratic semiprime" ,"5 prime", "6 semiprime", "7 prime", "8 8", "9 cuadratic semiprime", "10 semiprime"]

