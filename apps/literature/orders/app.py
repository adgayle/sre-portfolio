"""Flask application to track literature orders"""
import os
from flask import Flask
from flask_sqlalchemy import SQLAlchemy

app = Flask(__name__)
db = SQLAlchemy(app)

app.config["SECRET_KEY"] = os.getenv("SECRET_KEY")
app.config["SQLALCHEMY_DATABASE_URI"] = os.getenv(
    "SQLALCHEMY_DATABASE_URI", "sqlite:///orders.db"
)


class Publishers(db.Model):
    """Publisher defination"""

    id = db.Column(db.Integer, primary_key=True, index=True)
    name = db.Column(db.String(64), nullable=False)
    phone = db.Column(db.String(16), nullable=False, unique=True)
    email = db.Column(db.String(32), nullable=True)

    def __repr__(self) -> str:
        """Prints the publisher"""
        return f"Publisher <{self.id}->{self.name}>"


class LiteratureType(db.Model):
    """LiteratureType definition"""

    id = db.Column(db.Integer, primary_key=True, index=True)
    type = db.Column(db.String(16), nullable=False, unique=True)

    def __repr__(self) -> str:
        """Prints LiteratureType"""
        return f"LiteratureType <{self.id}->{self.type}"


class Literature(db.Model):
    """Literature definition"""

    id = db.Column(db.Integer, primary_key=True, index=True)
    name = db.Column(db.String(32), nullable=False, unique=True)
    type = db.relationship("LiteratureType", backref="type", lazy="dynamic")

    def __repr__(self) -> str:
        """Prints Literature"""
        return f"Literature <{self.id}->{self.name}"


@app.route("/")
def home() -> dict:
    """Returns a message from the flask app on GET request"""
    return {"message": "GET response from Flask app"}, 200


if __name__ == "__main__":
    app.run(debug=True)
