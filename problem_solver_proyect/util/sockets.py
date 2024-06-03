import base64
import json
import socket as sk
import traceback
import requests
from util.problem_factory.creator_fibonacci import CreatorFibonacci
from util.problem_factory.creator_fizzBuzz import CreatorFizzBuzz
from util import cryptography as cy

import util.log
import logging

class Socket():

    
    @staticmethod
    def start_server():
        server_socket = sk.socket(sk.AF_INET, sk.SOCK_STREAM)

        server_socket.bind(('localhost', 65432))

        server_socket.listen(1)

        print("Server listening on port 65432...")

        logging.info("Socket started")

        flag = True

        while flag:
           

            try:
                connection, client_address = server_socket.accept()

                logging.info("Connection established with", client_address)

                data = connection.recv(1024)

                logging.info(f"Data recived of client: {client_address}")

                print('Received data:', data.decode('utf-8'))

                print(f"Conexión establecida con {client_address}")
                
                    
                recived_json = json.loads(data.decode('utf-8'))
                
                if recived_json['Kill']:
                    try:

                        response_server = requests.get('http://localhost:5000/shutdown')

                        if response_server.status_code == 200:
                            connection.sendall("Shuting down everything".encode('utf-8'))
                            flag = False

                            logging.info("Data server shutdowned")
                            logging.info("Problem solver shutdowned")
                        else:
                            connection.sendall("error shuting down data server".encode('utf-8'))
                            logging.error("Error shutting down data server")
                    
                    except requests.exceptions.RequestException as e:
                        connection.sendall(f"Error connecting to data server".encode('utf-8'))
                        logging.error("Error connecting to data server")
                
                else:
                    numbers_server = None

                    try:
                        
                        session_key = cy.get_session_key()

                        encrypted_session_key = cy.get_encrpyted_session_key(session_key)

                        json_str = json.dumps(recived_json)

                        encrypted_json = cy.encrypt_with_session_key(json_str, session_key)

                        data_to_send = {
                            "encrypted_session_key": encrypted_session_key,
                            "encrypted_data": encrypted_json
                        }


                        response = requests.post('http://localhost:5000/numbers', json=data_to_send)

                        logging.info("Problem solver send data to data server")

                        if response.status_code == 200:

                            logging.info("Problem solver recived data of data server")

                            problem = None

                            response_data = response.json()
                            encrypted_session_key = response_data["encrypted_session_key"]
                            encrypted_data = response_data["encrypted_data"]

                            session_key = cy.decrypt_data(encrypted_session_key, cy.get_private_client_key())
                            print("Encrypted response from server:", response_data)

                            decrypted_response = cy.decrypt_with_session_key(encrypted_data, session_key)
                            print("Decrypted response from server:", decrypted_response)

                            json_decrypted_response = json.loads(decrypted_response)

                            problem = createProblem(recived_json["Problem"])

                            if problem == None:
                                connection.sendall("that is not a handled problem".encode('utf-8'))
                                logging.error("Error solving problem")

                            else:
                                result = problem.solve_problem(json_decrypted_response["Numbers"])
                                logging.info("Problem solved")

                                json_result = {"Result": result}
                                connection.sendall(json.dumps(json_result).encode('utf-8'))
                                logging.info("Solution send to client")

                        else:
                            connection.sendall(f"error getting the numbers: {response.text}".encode('utf-8'))
                            logging.error("Error getting the numbers - ", response.text)
                    
                    except requests.exceptions.RequestException as e:
                        connection.sendall(f"Error connecting to data server".encode('utf-8'))
                        logging.error("Error connecting to data server")

            except Exception as e:
                print(f"Error general: {e}")
                traceback.print_exc()
                logging.error("Error loading ssocket")

            
            finally:
                connection.close()

def createProblem(problem):

    if problem == "FizzBuzz":
        return CreatorFizzBuzz()
    
    if problem == "Fibonacci":
        return CreatorFibonacci()
    
    return None