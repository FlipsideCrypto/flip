import datetime
from typing import Optional
from pydantic import BaseModel, UUID4


# Shared properties
class User(BaseModel):
    id: UUID4
    user_name: str
    organization_id: UUID4
    organization_name: str

    created_at: Optional[datetime.datetime]
    updated_at: Optional[datetime.datetime]
    deleted_at: Optional[datetime.datetime]

    class Config:
        orm_mode = True

