from fastapi import (
    APIRouter,
    Depends,
    HTTPException,
    status
)

from db.session import SessionLocal
from actions.deps import get_db, get_current_user
from schemas import user as user_schema
from schemas import auth as auth_schema
from services import user_service, datax_service


router = APIRouter(tags=["users"])


@router.get(
    "/me",
    summary="Get Current User",
    response_model=user_schema.User,
    status_code=status.HTTP_200_OK,
)
async def me(
    db: SessionLocal = Depends(get_db),
    user: user_schema.User = Depends(get_current_user)
):
    return user


@router.post(
    "/authenticate",
    summary="Authenticate",
    response_model=auth_schema.AuthResponse,
    status_code=status.HTTP_200_OK,
)
async def authenticate(
    auth_req: auth_schema.AuthRequest,
    db: SessionLocal = Depends(get_db),
):
    jwt = user_service.get_jwt(auth_req.username, auth_req.password)
    if not jwt:
        raise HTTPException(status_code=400, detail="Login failed. Check username/password!")

    user = user_service.get_user(jwt)
    if not user:
        raise HTTPException(status_code=403, detail=f"User does not exist for provided JWT!")

    # Save user to db if new user
    u = user_service.save_db_user(db, user.id, user.user_name, user.organization_id)
    datax_service.getsert_datax_user(db, user.id, user.user_name)

    return auth_schema.AuthResponse(
        user=user,
        jwt=jwt
    )
   