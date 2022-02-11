"""Tests for publisher"""
from faker import Faker
from tracker.models.publisher import Publisher


def test_new_publisher():
    """Create a new Publisher and test that it is defined correctly"""
    fake = Faker()

    data = {
        "name": fake.name(),
        "email": fake.ascii_email(),
        "phone": fake.phone_number(),
        "password": fake.password(),
    }
    publisher = Publisher(data)
    assert publisher.name is data.get("name"), "Incorrect name returned"
    assert publisher.email is data.get("email"), "Incorrect email returned"
    assert publisher.phone is data.get("phone"), "Incorrect phone returned"
    assert publisher.check_password(
        data.get("password")
    ), "Incorrect password hash returned"
