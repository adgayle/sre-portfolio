"""Test orders"""
import requests
from orders import __version__, Publisher, Literature, db


def test_version():
    """Test the version"""
    assert __version__ == "0.1.0"


def test_orders():
    """Test order response to GET"""
    response = requests.get("http://localhost:5000/")
    assert response.status_code == 200, "Failed to get a 200 response"
    assert (
        response.json()["message"] == "GET response from Flask app"
    ), "Failed to retrieve correct message"
