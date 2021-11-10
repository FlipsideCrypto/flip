from typing import Any
import re
import datetime

from sqlalchemy.ext.declarative import as_declarative, declared_attr
from sqlalchemy import Column, ForeignKey, Integer, String, DateTime
from sqlalchemy.dialects.postgresql import UUID


@as_declarative()
class Base:
    id = Column(UUID(as_uuid=True), primary_key=True, index=True)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)
    updated_at = Column(DateTime, default=datetime.datetime.utcnow)
    deleted_at = Column(DateTime, default=None, nullable=True)

    __name__: str
    # Generate __tablename__ automatically
    @declared_attr
    def __tablename__(cls) -> str:
        # Convert CamelCase --> SnakeCase
        snake_case_table_name = re.sub(r'(?<!^)(?=[A-Z])', '_', cls.__name__).lower()
        # Pluralize
        return f"{snake_case_table_name}s"
