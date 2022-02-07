"""Flask app launcher"""
import os


from tracker.app import create_app


if __name__ == "__main__":
    environment = os.getenv("FLASK_ENV", "development")
    app = create_app(environment)
    app.run()
