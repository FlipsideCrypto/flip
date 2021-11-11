import json
import logging

from core.snowflake_db import SnowFlakeDB
from core.environment import (
    ADMIN_SNOWFLAKE_USER,
    ADMIN_SNOWFLAKE_ACCOUNT,
    ADMIN_SNOWFLAKE_PASSWORD,
    ADMIN_SNOWFLAKE_WAREHOUSE,
    ADMIN_SNOWFLAKE_DATABASE,
    OWNER_SNOWFLAKE_USER,
    OWNER_SNOWFLAKE_PASSWORD,
    OWNER_SNOWFLAKE_ACCOUNT,
    OWNER_SNOWFLAKE_REGION,
    OWNER_SNOWFLAKE_DATABASE,
    OWNER_SNOWFLAKE_WAREHOUSE
)


DATASHARE_NAME = "FLIPSIDE"

def run():
    with open("/app/data/schemas.json", "r") as f:
        schemas = json.loads(f.read())

    owner_sf = SnowFlakeDB(
        user=OWNER_SNOWFLAKE_USER,
        password=OWNER_SNOWFLAKE_PASSWORD,
        account=OWNER_SNOWFLAKE_ACCOUNT,
        region=OWNER_SNOWFLAKE_REGION,
        database=OWNER_SNOWFLAKE_DATABASE,
        warehouse=OWNER_SNOWFLAKE_WAREHOUSE
    )
    owner_sf.set_role("ACCOUNTADMIN")
    owner_sf.create_datashare(DATASHARE_NAME, comment="Datasets curated by the Flipside Analytics team.")
    for s in schemas:
        dbname = s['database'].upper()
        schema = s['schema'].upper()
        print(f"attaching '{dbname}'.'{schema}'...")
        owner_sf.attach_db_to_datashare(DATASHARE_NAME, dbname)
        owner_sf.attach_schema_to_datashare(DATASHARE_NAME, dbname, schema)
        objects = owner_sf.get_tables_and_views(dbname, schema)

        for o in objects:
            print(f"attaching {o['TABLE_TYPE'].lower()}'{dbname}'.'{o['TABLE_SCHEMA']}'.'{o['TABLE_NAME']}'...")
            
            if o['TABLE_TYPE'] == 'VIEW':
                if s['exclude_views'] is True:
                    continue
                try:
                    owner_sf.attach_view_to_datashare(DATASHARE_NAME, dbname, o['TABLE_SCHEMA'], o['TABLE_NAME'])
                except Exception as e:
                    logging.warning(e.msg)
            else:
                try:
                    owner_sf.attach_table_to_datashare(DATASHARE_NAME, dbname, o['TABLE_SCHEMA'], o['TABLE_NAME'])
                except Exception as e:
                    logging.warning(e.msg)

    owner_sf.attach_sf_account_to_datashare(DATASHARE_NAME, "AVA40942")

    e_db = SnowFlakeDB(
        user=ADMIN_SNOWFLAKE_USER,
        password=ADMIN_SNOWFLAKE_PASSWORD,
        account=ADMIN_SNOWFLAKE_ACCOUNT,
        database=ADMIN_SNOWFLAKE_DATABASE,
        warehouse=ADMIN_SNOWFLAKE_WAREHOUSE
    )
    e_db.set_role("ACCOUNTADMIN")
    e_db.import_share(OWNER_SNOWFLAKE_ACCOUNT, DATASHARE_NAME, DATASHARE_NAME, roles=["PUBLIC", "ACCOUNTADMIN"])


if __name__ == '__main__':
    run()