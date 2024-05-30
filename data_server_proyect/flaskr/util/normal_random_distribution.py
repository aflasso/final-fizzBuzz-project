import numpy as np
from flaskr.util.i_number_solver import INumberRandomSolver

class NormalRandomDistribution(INumberRandomSolver):

    def get_numbers(self, min, max, cant):
        
        numbers = np.random.normal(min, max, cant)
        int_numbers = numbers.astype(int)
        return int_numbers.tolist()