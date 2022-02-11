"""Flask application"""
from flask import Flask
from .config import app_config
from .models import bcrypt, db


def create_app(environment: str) -> Flask:
    """Create flask app using environment configuration"""
    app = Flask(__name__)
    app.config.from_object(app_config[environment])
    app.config["SQLALCHEMY_TRACK_MODIFICATIONS"] = False

    bcrypt.init_app(app)
    db.init_app(app)

    return app
