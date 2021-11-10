import sqlalchemy
from sqlalchemy.orm import session
from uuid import UUID, uuid4

from models import DataxUser, get_uid
from schemas.data_exchange import SnowflakeUser
from utils import snowflake_utils
from core.environment import (
    ADMIN_SNOWFLAKE_USER,
    ADMIN_SNOWFLAKE_PASSWORD,
    ADMIN_SNOWFLAKE_ACCOUNT,
    ADMIN_SNOWFLAKE_REGION,
    ADMIN_SNOWFLAKE_DATABASE,
    ADMIN_SNOWFLAKE_WAREHOUSE
)
from core import snowflake_db


def generate_snowflake_account(sf_db, user_name: str) -> SnowflakeUser:
    sf_u = SnowflakeUser(
        account=snowflake_utils.get_account(),
        username=snowflake_utils.clean_username(user_name),
        password=snowflake_utils.gen_password(),
        database=snowflake_utils.get_database(),
        warehouse=snowflake_utils.get_warehouse()
    )

    sf_db.create_curator(sf_u.username, sf_u.password, warehouse=sf_u.warehouse)
    sf_db.create_curator_schema(sf_u.database, sf_u.username, sf_u.username)
    return sf_u


def getsert_datax_user(db: session, user_id: UUID, user_name: str) -> DataxUser: 
    query = db.query(DataxUser).filter(DataxUser.user_id == user_id)
    try:
        return query.one()
    except sqlalchemy.exc.NoResultFound:
        pass

    sf_db = snowflake_db.SnowFlakeDB(
        user=ADMIN_SNOWFLAKE_USER,
        password=ADMIN_SNOWFLAKE_PASSWORD,
        account=ADMIN_SNOWFLAKE_ACCOUNT,
        database=ADMIN_SNOWFLAKE_DATABASE,
        warehouse=ADMIN_SNOWFLAKE_WAREHOUSE
    )
    
    snowflake_user = generate_snowflake_account(sf_db, user_name)
    datax_user = DataxUser(
        id=get_uid(),
        user_id=user_id,
        datax_name="FLIPSIDE",
        account=snowflake_user.account,
        username=snowflake_user.username,
        password=snowflake_user.password,
        database=snowflake_user.database,
        warehouse=snowflake_user.warehouse,
    )
    db.add(datax_user)
    db.flush()
    db.commit()
    sf_db.commit()
    return datax_user
