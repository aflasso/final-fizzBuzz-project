import base64
import json
import socket as sk
import traceback
import requests
from util.problem_factory.creator_fibonacci import CreatorFibonacci
from util.problem_factory.creator_fizzBuzz import CreatorFizzBuzz
from util import cryptography as cy


class Socket():

    @staticmethod
    def start_server():
        server_socket = sk.socket(sk.AF_INET, sk.SOCK_STREAM)

        server_socket.bind(('localhost', 65432))

        server_socket.listen(1)

        print("Servidor escuchando en el puerto 65432...")

        flag = True

        while flag:
           

            try:
                connection, client_address = server_socket.accept()

                data = connection.recv(1024)

                print('Received data:', data.decode('utf-8'))

                print(f"Conexión establecida con {client_address}")
                
                    
                recived_json = json.loads(data.decode('utf-8'))
                
                if recived_json['Kill']:
                    try:

                        response_server = requests.get('http://localhost:5000/shutdown')

                        if response_server.status_code == 200:
                            connection.sendall("Shuting down everything".encode('utf-8'))
                            flag = False
                        else:
                            connection.sendall("error shuting down data server".encode('utf-8'))
                    
                    except requests.exceptions.RequestException as e:
                        connection.sendall(f"Error connecting to data server".encode('utf-8'))
                
                else:
                    numbers_server = None

                    try:
                        
                        key = cy.generate_random_key()

                        json_str = json.dumps(recived_json)

                        encrypted_json = cy.encrypt_data(json_str, key)

                        data_to_send = {
                            "key": base64.b64encode(key).decode('utf-8'),
                            "encrypted_data": encrypted_json
                        }

                        print(encrypted_json)

                        response = requests.post('http://localhost:5000/numbers', json=data_to_send)

                        if response.status_code == 200:

                            encrypted_response = response.json().get("encrypted_data")
                            print("Encrypted response from server:", encrypted_response)

                            decrypted_response = cy.decrypt_data(encrypted_response, key)
                            print("Decrypted response from server:", decrypted_response)

                            json_decrypted_response = json.loads(decrypted_response)

                            problem = createProblem(recived_json["Problem"])

                            if problem == None:
                                connection.sendall("that is not a handled problem".encode('utf-8'))

                            else:
                                result = problem.solve_problem(json_decrypted_response["Numbers"])

                                json_result = {"Result": result}
                                connection.sendall(json.dumps(json_result).encode('utf-8'))

                        else:
                            connection.sendall("error getting the numbers".encode('utf-8'))
                    
                    except requests.exceptions.RequestException as e:
                        connection.sendall(f"Error connecting to data server".encode('utf-8'))
                    
                
                print("No hay más datos de", client_address)

            except Exception as e:
                print(f"Error general: {e}")
                traceback.print_exc()
            
            finally:
                connection.close()

def createProblem(problem):

    if problem == "FizzBuzz":
        return CreatorFizzBuzz()
    
    if problem == "Fibonacci":
        return CreatorFibonacci()
    
    return None