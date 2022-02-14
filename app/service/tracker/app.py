"""Flask application"""
import os
from flask import Flask, Blueprint
from .config import app_config, Dev
from .models import bcrypt, db
from .views.publisher_view import publisher_api


def create_app(environment: str) -> Flask:
    """Create flask app using environment configuration"""
    app = Flask(__name__)
    app.config.from_object(Dev)
    app.config["SQLALCHEMY_TRACK_MODIFICATIONS"] = False
    db_file_path = os.getenv("APP_DB_FILE", "/tmp/tracker.db")
    app.config["SQLALCHEMY_DATABASE_URI"] = "sqlite:///" + db_file_path

    bcrypt.init_app(app)
    db.init_app(app)
    db.create_all

    app.register_blueprint(publisher_api, url_prefix="/api/v1/publisher")

    @app.route("/", methods=["GET"])
    def home():
        """Home endpoint"""
        return "Made it to home"

    return app
