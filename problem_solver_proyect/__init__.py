from util.sockets import Socket
from util.problem_factory.i_problem_creator import IProblemCreator
from util.problem_factory.creator_fizzBuzz import CreatorFizzBuzz
from util.problem_factory.creator_fibonacci import CreatorFibonacci


def client(creator: IProblemCreator):

    numbers = [1,2,3,4,5,6,7,8,9,10]

    result = creator.solve_problem(numbers)

    return result

if __name__ == "__main__":

    Socket.start_server()