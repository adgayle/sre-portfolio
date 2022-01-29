"""Flask application to track literature orders"""
from flask import Flask

app = Flask(__name__)


@app.route("/")
def home() -> dict:
    """Returns a message from the flask app on GET request"""
    return {"message": "GET response from Flask app"}, 200


if __name__ == "__main__":
    app.run(debug=True)
