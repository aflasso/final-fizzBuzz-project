import base64
import os
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
        key = base64.b64decode(request_data["key"])
        encrypted_data = request_data["encrypted_data"]

        try:

            decrypted_data = cy.decrypt_data(encrypted_data, key)

            json_data = json.loads(decrypted_data)

            print(encrypted_data)
            print(json_data)
            
            desc_random = random.choice((0,1))
            random_numbers = None

            if desc_random:
                distribution = NormalRandomDistribution()
            else:
                distribution = UniformRandomDistribution()

            random_numbers = distribution.get_numbers(json_data["MinNumber"], json_data["MaxNumber"], json_data["CantData"])

            response_json = {"Numbers": random_numbers}

            response_data = json.dumps(response_json)

            encrypted_response_data = cy.encrypt_data(response_data, key)

            return jsonify({"encrypted_data":encrypted_response_data})
        
        except Exception as e:
            return str(e), 400

    return app

def shutdown_server():

    sys.exit()