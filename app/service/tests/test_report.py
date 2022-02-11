"""Tests for report"""
from faker import Faker
from tracker.models.report import Report


MAX_HOURS = 200
MAX_PLACEMENTS = 1000
MAX_VIDEO_SHOWINGS = 500
MAX_RETURN_VISITS = 100
MAX_BIBLE_STUDIES = 50
ENVIRONMENT = "development"


def test_new_report() -> None:
    """Create a new Report and test that it is defined correctly"""
    fake = Faker()
    data = {
        "hours": fake.random_int(0, MAX_HOURS),
        "placements": fake.random_int(0, MAX_PLACEMENTS),
        "video_showings": fake.random_int(0, MAX_VIDEO_SHOWINGS),
        "return_visits": fake.random_int(0, MAX_RETURN_VISITS),
        "studies": fake.random_int(0, MAX_BIBLE_STUDIES),
        "comments": fake.sentence(nb_words=10),
    }
    report = Report(data)
    assert report.hours == data.get("hours"), "Incorrect hours returned"
    assert report.placements == data.get("placements"), "Incorrect placements returned"
    assert report.video_showings == data.get(
        "video_showings"
    ), "Incorrect video showings returned"
    assert report.return_visits == data.get(
        "return_visits"
    ), "Incorrect return visits returned"
    assert report.studies == data.get("studies"), "Incorrect studies returned"
    assert report.comments == data.get("comments"), "Incorrect comments returned"
