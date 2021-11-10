from fastapi import Depends, HTTPException, Security
from fastapi.security.api_key import APIKeyQuery, APIKeyHeader
import sqlalchemy
from sqlalchemy.orm import Session
from typing import Generator

from services import user_service
from schemas import user as user_schema
from db.session import SessionLocal


def get_db() -> Generator:
    db = None
    try:
        db = SessionLocal()
        yield db
    finally:
        if db:
            db.close()


api_key_header = APIKeyHeader(name="authorization", auto_error=False)


async def get_jwt(
    api_key_header: str = Security(api_key_header),  
) -> str:
    return api_key_header


def get_current_user(
    db: Session = Depends(get_db),
    jwt: str = Depends(get_jwt)
) -> user_schema.User:
    if not jwt:
        raise HTTPException(status_code=403, detail=f"A JWT is required to authenticate.")

    user = user_service.get_user(jwt)
    if not user:
        raise HTTPException(status_code=403, detail=f"Invalid JWT. User is not authenticated!")

    return user
