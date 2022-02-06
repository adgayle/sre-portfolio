"""Reports model"""
from datetime import datetime
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
    created_at = db.Column(db.DateTime)
    modified_at = db.Column(db.DateTime)

    def __init__(self, data: dict) -> None:
        """Report constructor"""
        self.hours = data.hours
        self.placements = data.placements
        self.video_showings = data.video_showings
        self.return_visits = data.return_visits
        self.studies = data.studies
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
