"""Publisher API"""
from flask import Blueprint, Response, json, request
from tracker.models import publisher
from tracker.models.publisher import Publisher, PublisherSchema

publisher_api = Blueprint("publishers", __name__)
publisher_schema = PublisherSchema()


def custom_response(response, status_code):
    """Custom response"""
    return Response(
        mimetype="application/json", response=json.dumps(response), status=status_code
    )


@publisher_api.route("/", methods=["POST"])
def create():
    """Create a Publisher"""
    data = request.get_json()
    """print(f"here: {publisher_schema.load(req_data)}")
    data, error = publisher_schema.load(req_data)
    if error:
        return custom_response(error, 400)
    """
    publisher_exists = Publisher.get_publisher_by_email(data.get("email"))
    if publisher_exists:
        message = {"error": "Publisher already exists with that email"}
        return custom_response(message, 400)

    publisher = Publisher(data)
    publisher.add()
    return custom_response(f"{publisher}", 201)


@publisher_api.route("/<int:publisher_id>", methods=["GET"])
def get_one_publisher(publisher_id):
    """Get a single publisher by id"""
    publisher = Publisher.get_publisher(publisher_id)
    if not publisher:
        return custom_response({"error": "user not found"}, 404)

    publisher_json = publisher_schema.dump(publisher)
    return custom_response(publisher_json, 200)
