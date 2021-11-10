from sqlalchemy import Column, String
from sqlalchemy.dialects.postgresql import UUID

from db.base_class import Base


class User(Base):

    org_id = Column("org_id", UUID(as_uuid=True))
    username = Column("username", String, index=True)
