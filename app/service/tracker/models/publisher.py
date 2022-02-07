"""Publisher model"""
from marshmallow import fields, Schema
from datetime import datetime
from . import bcrypt, db
from .report import ReportSchema


class Publisher(db.Model):
    """Publisher model class"""

    __table_name__ = "publishers"

    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(128), nullable=False)
    email = db.Column(db.String(128), nullable=False, unique=True)
    phone = db.Column(db.String(16), nullable=False)
    password = db.Column(db.String(128), nullable=False)
    created_at = db.Column(db.DateTime)
    modified_at = db.Column(db.DateTime)
    reports = db.relationship("Report", backref="publishers", lazy=True)

    def __init__(self, data: dict) -> None:
        """Publisher constructor"""
        self.name = data.get("name")
        self.email = data.get("email")
        self.phone = data.get("phone")
        self.password = self.__generate_hash(data.get("password"))
        self.created_at = datetime.utcnow()
        self.modified_at = datetime.utcnow()

    def add(self) -> None:
        """Add a Publisher"""
        db.session.add(self)
        db.session.commit()

    def update(self, data: dict) -> None:
        """Update a Publisher"""
        for key, value in data.items():
            if key == "password":
                self.password = self.__generate_hash(value)
            else:
                setattr(self, key, value)
        self.modified_at = datetime.utcnow()
        db.session.commit()

    def delete(self) -> None:
        """Delete a Publisher"""
        db.session.delete(self)
        db.session.commit()

    @staticmethod
    def get_publisher(id: int) -> dict:
        """Return the single publisher with id"""
        return Publisher.query.get(id)

    @staticmethod
    def get_publisher_by_email(email: str) -> dict:
        """Return the single publisher with email"""
        return Publisher.query.get(email)

    @staticmethod
    def get_all_publishers() -> dict:
        """Return all publishers"""
        return Publisher.query.all()

    def __repr__(self) -> str:
        """Return a string representation of publisher"""
        return f"<id {self.id}>"

    def __generate_hash(self, password: str) -> bytes:
        """Generate a password hash"""
        return bcrypt.generate_password_hash(password, rounds=12).decode("utf-8")

    def check_password(self, password: str) -> bool:
        """Check if the password provided matches the stored"""
        return bcrypt.check_password_hash(self.password, password)


class PublisherSchema(Schema):
    """Publisher schema definition"""

    id = fields.Int(dump_only=True)
    name = fields.Str(required=True)
    email = fields.Str(required=True)
    phone = fields.Str(required=False)
    password = fields.Str(required=True)
    created_at = fields.DateTime(dump_only=True)
    modified_at = fields.DateTime(dump_only=True)
    reports = fields.Nested(ReportSchema, many=True)
