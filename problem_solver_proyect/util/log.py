import logging

logging.basicConfig(
        filename="problem_solver_proyect/files/problemSolver.log",
        filemode='a',
        format='%(asctime)s - %(name)s - %(levelname)s - %(message)s',
        datefmt='%Y-%m-%d %H:%M:%S',
        level = logging.DEBUG
    )

