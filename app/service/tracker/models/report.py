"""Reports model"""
from datetime import datetime
from marshmallow import fields, Schema
from . import db


class Report(db.Model):
    """Report model class"""

    __table_name__ = "reports"

    id = db.Column(db.Integer, primary_key=True)
    hours = db.Column(db.Integer, nullable=False)
    placements = db.Column(db.Integer, nullable=True)
    video_showings = db.Column(db.Integer, nullable=True)
    return_visits = db.Column(db.Integer, nullable=True)
    studies = db.Column(db.Integer, nullable=True)
    comments = db.Column(db.String(256), nullable=True)
    publisher_id = db.Column(db.Integer, db.ForeignKey("publisher.id"), nullable=False)
    created_at = db.Column(db.DateTime)
    modified_at = db.Column(db.DateTime)

    def __init__(self, data: dict) -> None:
        """Report constructor"""
        self.hours = data.get("hours")
        self.placements = data.get("placements")
        self.video_showings = data.get("video_showings")
        self.return_visits = data.get("return_visits")
        self.studies = data.get("studies")
        self.comments = data.get("comments")
        self.publisher_id = data.get("publisher_id")
        self.created_at = datetime.utcnow()
        self.modified_at = datetime.utcnow()

    def add(self) -> None:
        """Add a report"""
        db.session.add(self)
        db.commit()

    def update(self, data: dict) -> None:
        """Update a report"""
        for key, value in data.items():
            setattr(self, key, value)
        self.modified_at = datetime.utcnow()
        db.session.commit()

    def delete(self) -> None:
        """Delete a report"""
        db.session.delete(self)
        db.commit()

    @staticmethod
    def get_report(id) -> dict:
        """Return the single publisher with id"""
        return Report.query.get(id)

    def __repr__(self) -> str:
        """Return the string representation of a report"""
        return f"<id {self.id}>"


class ReportSchema(Schema):
    """Report schema definition"""

    id = fields.Int(dump_only=True)
    hours = fields.Int(required=True)
    placements = fields.Int(required=True)
    video_showings = fields.Int(required=False)
    return_visits = fields.Int(required=False)
    studies = fields.Int(required=False)
    comments = fields.Str(required=False)
    publisher_id = fields.Int(required=True)
    created_at = fields.DateTime(dump_only=True)
    modified_at = fields.DateTime(dump_only=True)
