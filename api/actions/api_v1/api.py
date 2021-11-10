from fastapi import APIRouter

from actions.api_v1.endpoints import health_checks, users, data_exchange

api_router = APIRouter()
api_router.include_router(health_checks.router, prefix="/checks", tags=["health_checks"])
api_router.include_router(users.router, prefix="/users", tags=["users"])
api_router.include_router(data_exchange.router, prefix="/data-exchange", tags=["data_exchange"])
