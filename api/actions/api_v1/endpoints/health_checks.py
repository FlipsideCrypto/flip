from fastapi import APIRouter, status, Depends

from schemas.requests import PingResponse


router = APIRouter()


@router.get(
    "/ping",
    summary="Inspect health",
    response_model=PingResponse,
    status_code=status.HTTP_200_OK,
)
async def ping():
    """Inspects if the API instance is running."""
    return {"ping": "pong"}
