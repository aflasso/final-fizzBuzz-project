import base64
import os
import signal
from flask import Flask, request, jsonify
import sys
import random
from flaskr.util.normal_random_distribution import NormalRandomDistribution
from flaskr.util.uniform_random_distribution import UniformRandomDistribution
import flaskr.util.cryptography as cy
import json


def create_app(test_config=None):
    # create and configure the app
    app = Flask(__name__, instance_relative_config=True)
    app.config.from_mapping(
        SECRET_KEY='dev',
        DATABASE=os.path.join(app.instance_path, 'flaskr.sqlite'),
    )

    if test_config is None:
        # load the instance config, if it exists, when not testing
        app.config.from_pyfile('config.py', silent=True)
    else:
        # load the test config if passed in
        app.config.from_mapping(test_config)

    # ensure the instance folder exists
    try:
        os.makedirs(app.instance_path)
    except OSError:
        pass

    # a simple page that says hello
    @app.route('/shutdown', methods=("GET",))
    def shutdown():
        shutdown_server()
        return "Server shutting down..."
    
    @app.route("/numbers", methods=("POST",))
    def numbers():
        request_data = request.json
        encrypted_session_key = request_data["encrypted_session_key"]
        encrypted_data = request_data["encrypted_data"]

        try:

            session_key = cy.decrypt_data(encrypted_session_key, cy.get_private_server_key())

            decrypted_data = cy.decrypt_with_session_key(encrypted_data, session_key)

            json_data = json.loads(decrypted_data)

            print(encrypted_data)
            print(json_data)
            
            desc_random = random.choice((0,1))
            random_numbers = None

            if desc_random:
                distribution = NormalRandomDistribution()
            else:
                distribution = UniformRandomDistribution()
            
            if json_data["TestMode"]:
                random_numbers = [1,2,3,4,5,6,7,8,9,10]
            
            else:
                random_numbers = distribution.get_numbers(json_data["MinNumber"], json_data["MaxNumber"], json_data["CantData"])

            encrypted_session_key = cy.get_encrpyted_session_key(session_key)

            
            response_json = {"Numbers": random_numbers}

            response_data = json.dumps(response_json)

            encrypted_response_data = cy.encrypt_with_session_key(response_data, session_key)

            return jsonify({
                            "encrypted_session_key": encrypted_session_key,
                            "encrypted_data": encrypted_response_data
                        })

            # return jsonify({"Numbers": random_numbers})
        
        except Exception as e:
            return str(e), 400

    return app

def shutdown_server():
    pid = os.getpid()
    os.kill(pid, signal.SIGINT)