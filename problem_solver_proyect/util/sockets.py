import json
import socket as sk
import traceback
import requests

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

                try:
                    print(f"Conexión establecida con {client_address}")

                    data = connection.recv(1024)

                    while data:
                        
                        recived_json = json.loads(data.decode('utf-8'))
                        
                        if recived_json['Kill']:

                            response_server = requests.get('http://localhost:5000/shutdown')

                            response_socket = {"message": "zzz"}
                            connection.sendall(json.dumps(response_socket).encode('utf-8'))
                            flag = False
                        
                        else:
                            numbers_server = requests.post('http://localhost:5000/numbers', json=recived_json)

                        data = connection.recv(1024)
                        
                        print("No hay más datos de", client_address)
                            
                finally:
                    connection.close()


            except Exception as e:
                print(f"Error general: {e}")
                traceback.print_exc()
