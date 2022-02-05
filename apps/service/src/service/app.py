"""Flask application"""
from flask import Flask
from .config import app_config


def create_app(environment: str) -> Flask:
    """Create flask app using environment configuration"""
    app = Flask(__name__)
    app.config.from_object(app_config[environment])

    @app.route("/", methods=["GET"])
    def index():
        """Root endpoint"""

        return "Congratulations! Your first endpoint is working"

    return app
