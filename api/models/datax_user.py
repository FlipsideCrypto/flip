from sqlalchemy import Column, String, ForeignKey
from sqlalchemy.dialects.postgresql import UUID

from db.base_class import Base


class DataxUser(Base):

    user_id = Column(UUID(as_uuid=True), ForeignKey("users.id"), index=True)
    datax_name = Column("datax_name", String)
    account = Column("account", String)
    username = Column("username", String)
    password = Column("password", String)
    database = Column("database", String)
    warehouse = Column("warehouse", String)
