"""Flask configuration parameters"""
import os


class Dev(object):
    """Development environment"""

    DEBUG = True
    TESTING = True
    JWT_SECRET_KEY = os.getenv("JWT_SECRET_KEY")
    SQLALCHEMY_DATABASE_URI = os.getenv("DATABASE_URI")


app_config = {"development": Dev}
