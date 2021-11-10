import random

from core.environment import (
    SNOWFLAKE_EXTERNAL_ACCOUNT, 
    SNOWFLAKE_CURATOR_DATABASE,
    SNOWFLAKE_CURATOR_DEFAULT_WAREHOUSE
)
from uuid import uuid4


def get_account():
    return SNOWFLAKE_EXTERNAL_ACCOUNT


def get_database():
    return SNOWFLAKE_CURATOR_DATABASE


def get_warehouse():
    return SNOWFLAKE_CURATOR_DEFAULT_WAREHOUSE


letters = 'abcdefghijklmnopqrstuvwxyz'

def gen_password():
    p1 = str(uuid4())[:13].replace('-', '').upper()
    pf = []
    inserted_upper = False
    inserted_lower = False
    for i, x in enumerate(p1):
        is_number = False
        try:
            int(x)
        except ValueError:
            is_number = True
        
        if is_number:
            pf.append(x)
            continue

        if inserted_upper is False:
            pf.append(letters[random.randint(0, len(letters)-1)].upper())
            inserted_upper = True
        elif inserted_lower is False:
            pf.append(letters[random.randint(0, len(letters)-1)].lower())
            inserted_lower = True
        else:
            pf.append(x)

    return "".join(pf)


def clean_username(user_name):
    return user_name.\
        replace("#", '').\
        replace("@", '').\
        replace("/", '').\
        replace("!", '').\
        replace("%", '').\
        replace("(", '').\
        replace(")", '').\
        replace("^", '').\
        replace("[", '').\
        replace("]", '').\
        replace("{", '').\
        replace("}", '').\
        replace("|", '').\
        replace("~", '').\
        replace("`", '').\
        replace("'", '').\
        replace('"', '').\
        replace(';', '').\
        replace(':', '').\
        replace('+', '').\
        replace('=', '').\
        replace('*', '').\
        replace('>', '').\
        replace('<', '').\
        replace(',', '').\
        replace("-", '_').\
        upper()