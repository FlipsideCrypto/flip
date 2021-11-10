from fastapi import (
    APIRouter,
    Depends,
    status
)

from db.session import SessionLocal
from actions.deps import get_db, get_current_user
from schemas import data_exchange as data_exchange_schema
from schemas import user as user_schema
from services import datax_service


router = APIRouter(tags=["data_exchange"])


@router.get(
    "/creds",
    summary="Get Data-Exchange Creds",
    response_model=data_exchange_schema.DataExchangeCreds,
    status_code=status.HTTP_200_OK,
)
async def get_creds(
    db: SessionLocal = Depends(get_db),
    user: user_schema.User = Depends(get_current_user)
):
    datax_user = datax_service.getsert_datax_user(db, user.id, user.user_name)
    return data_exchange_schema.DataExchangeCreds(
        account=datax_user.account,
        exchange_name=datax_user.datax_name,
        url=f"https://{datax_user.account}.snowflakecomputing.com",
        username=datax_user.username,
        password=datax_user.password,
        database=datax_user.database,
        warehouse=datax_user.warehouse
    )
