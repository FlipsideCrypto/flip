from uuid import UUID
from enum import Enum

from pydantic import BaseModel, Field, validator, UUID4
from typing import Any, Optional, List


class PingResponse(BaseModel):
    ping: str = Field("pong", description="The health status")
