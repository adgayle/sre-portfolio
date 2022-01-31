"""Initialize orders"""
from orders.seed import initdb, seed_publisher, seed_literature

__version__ = "0.1.0"


initdb()
seed_publisher(10)
seed_literature(25)
