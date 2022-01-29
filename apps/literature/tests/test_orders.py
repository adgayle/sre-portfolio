"""Test orders"""
import requests
from orders import __version__


def test_version():
    assert __version__ == "0.1.0"


def test_orders():
    response = requests.get("http://127.0.0.1:5000")
    assert response.status_code == 200, "Failed to get a 200 response"
