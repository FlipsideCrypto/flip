from pydantic import BaseModel
from typing import Optional

from .user import User


class AuthRequest(BaseModel):
    username: str
    password: str

class AuthResponse(BaseModel):
    user: User
    jwt: str
