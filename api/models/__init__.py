from uuid import uuid4

from .user import User
from .datax_user import DataxUser


def get_uid() -> str:
    return str(uuid4())
