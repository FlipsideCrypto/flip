from pydantic import BaseModel


class SnowflakeUser(BaseModel):
    account: str
    username: str
    password: str
    database: str
    warehouse: str


class DataExchangeCreds(BaseModel):
    url: str
    exchange_name: str
    account: str
    username: str
    password: str
    database: str
    warehouse: str
