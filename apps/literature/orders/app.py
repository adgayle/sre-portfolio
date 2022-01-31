"""Flask application to track literature orders"""
from crypt import methods
import os
from flask import Flask, request, jsonify
from flask_sqlalchemy import SQLAlchemy
from flask_marshmallow import Marshmallow

app = Flask(__name__)
app.config["SECRET_KEY"] = os.getenv("SECRET_KEY")
app.config["SQLALCHEMY_DATABASE_URI"] = os.getenv(
    "SQLALCHEMY_DATABASE_URI", "sqlite:///orders.db"
)
app.config["SQLALCHEMY_TRACK_MODIFICATIONS"] = False
db = SQLAlchemy(app)
ma = Marshmallow(app)


class Publisher(db.Model):
    """Publisher definition"""

    id = db.Column(db.Integer, primary_key=True, index=True)
    name = db.Column(db.String(64), nullable=False)
    phone = db.Column(db.String(16), nullable=True)
    email = db.Column(db.String(32), nullable=True)

    def __init__(self, name, phone, email):
        self.name = name
        self.phone = phone
        self.email = email


class PublisherSchema(ma.Schema):
    class Meta:
        fields = ("id", "name", "phone", "email")


class Literature(db.Model):
    """Literature definition"""

    id = db.Column(db.Integer, primary_key=True, index=True)
    name = db.Column(db.String(32), nullable=False, unique=True)

    def __repr__(self) -> str:
        """Prints Literature"""
        return f"Literature <{self.id}->{self.name}"


class LiteratureSchema(ma.Schema):
    class Meta:
        fields = ("id", "name")


publisher_schema = PublisherSchema()
literature_schema = LiteratureSchema()


@app.route("/publisher", methods=["POST"])
def add_publisher() -> dict:
    """Create a publisher"""
    name = request.json["name"]
    phone = request.json["phone"]
    email = request.json["email"]

    new_publisher = Publisher(name, phone, email)
    db.session.add(new_publisher)
    db.session.commit()

    return publisher_schema.jsonify(new_publisher)


@app.route("/publisher", methods=["GET"])
def get_publishers():
    """Get publishers"""
    all_publishers = Publisher.query.all()
    return publisher_schema.jsonify(all_publishers)


@app.route("/publisher/<id>", methods=["GET"])
def get_publisher(id: int):
    """Get a publisher by id"""
    publisher = Publisher.query.get(id)
    return publisher_schema.jsonify(publisher)


if __name__ == "__main__":
    app.run(debug=True)
