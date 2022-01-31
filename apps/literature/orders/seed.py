#!/usr/bin/env python3
from faker import Faker
from orders.app import db, Publisher, Literature


def initdb():
    """Initialize an empty database"""
    db.create_all()


def seed_publisher(count: int):
    """Add some publishers to the database"""
    faker = Faker()
    for index in range(count):
        pub = Publisher(
            name=faker.name(), phone=faker.phone_number(), email=faker.email()
        )
        db.session.add(pub)

    db.session.commit()
    Publisher.query.all()


def seed_literature(count):
    """Add literature to the database"""
    faker = Faker()
    for index in range(count):
        lit = Literature(name=faker.sentence(nb_words=4))
        db.session.add(lit)

    db.session.commit()
    Literature.query.all()


def main():
    """Main for seeding"""
    initdb()
    seed_publisher(10)
    seed_literature(25)


if __name__ == "__main__":
    main()
