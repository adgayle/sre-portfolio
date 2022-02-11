"""Tests for app"""
from flask import Flask
from tracker.app import create_app


def test_create_app():
    """Test for create app"""
    app = create_app("development")
    assert type(app) is Flask
