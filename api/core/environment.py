import os

# [ENV]
PROD_ENV = "prod"
LOCAL_ENV = "local"

ENV = LOCAL_ENV
DEBUG = True
IS_DEV = True if ENV and ENV != PROD_ENV else False


# SQLALCHEMY_DATABASE_URI = os.environ.get('SQLALCHEMY_DATABASE_URI')
SQLALCHEMY_DATABASE_URI = os.environ.get('DATABASE')
print("SQLALCHEMY_DATABASE_URI: ", SQLALCHEMY_DATABASE_URI)
GRAPHQL_API = os.environ.get('GRAPHQL_API')

SNOWFLAKE_EXTERNAL_ACCOUNT = os.environ.get('SNOWFLAKE_EXTERNAL_ACCOUNT')
SNOWFLAKE_CURATOR_DATABASE = os.environ.get('SNOWFLAKE_CURATOR_DATABASE')
SNOWFLAKE_CURATOR_DEFAULT_WAREHOUSE = os.environ.get('SNOWFLAKE_CURATOR_DEFAULT_WAREHOUSE')

ADMIN_SNOWFLAKE_USER = os.environ.get('ADMIN_SNOWFLAKE_USER')
ADMIN_SNOWFLAKE_PASSWORD = os.environ.get('ADMIN_SNOWFLAKE_PASSWORD')
ADMIN_SNOWFLAKE_ACCOUNT = os.environ.get('ADMIN_SNOWFLAKE_ACCOUNT')
ADMIN_SNOWFLAKE_REGION = os.environ.get('ADMIN_SNOWFLAKE_REGION')
ADMIN_SNOWFLAKE_DATABASE = os.environ.get('ADMIN_SNOWFLAKE_DATABASE')
ADMIN_SNOWFLAKE_WAREHOUSE = os.environ.get('ADMIN_SNOWFLAKE_WAREHOUSE')

OWNER_SNOWFLAKE_USER = os.environ.get('OWNER_SNOWFLAKE_USER')
OWNER_SNOWFLAKE_PASSWORD = os.environ.get('OWNER_SNOWFLAKE_PASSWORD')
OWNER_SNOWFLAKE_ACCOUNT = os.environ.get('OWNER_SNOWFLAKE_ACCOUNT')
OWNER_SNOWFLAKE_REGION = os.environ.get('OWNER_SNOWFLAKE_REGION')
OWNER_SNOWFLAKE_DATABASE = os.environ.get('OWNER_SNOWFLAKE_DATABASE')
OWNER_SNOWFLAKE_WAREHOUSE = os.environ.get('OWNER_SNOWFLAKE_WAREHOUSE')
