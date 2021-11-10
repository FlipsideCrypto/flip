from terminaltables import AsciiTable
import snowflake.connector


class SnowFlakeDB(object):

    def __init__(self, user=None, password=None, account=None, region=None, database=None, warehouse=None):
        self.user = user
        self.password = password
        self.account = account
        self.region = region
        self.database = database
        self.warehouse = warehouse

    def set_role(self, role):
        self.role = role

    def print_results(self, cols, rows):
        if not cols or not rows:
            return
        table_data = [cols] + [row for row in rows]
        table = AsciiTable(table_data)
        print(table.table)

    def create_datashare(self, share_name, comment=''):
        try:
            self.execute(f"""
                CREATE SHARE "{share_name}" COMMENT='{comment}'
            """)
        except snowflake.connector.errors.ProgrammingError as e:
            if 'already exists' in e.msg:
                return
            raise Exception(e)
        self.commit()

    def attach_db_to_datashare(self, share_name, dbname):
        self.execute(f"""
            GRANT USAGE ON DATABASE "{dbname}" TO SHARE "{share_name}"
        """)
        self.commit()

    def attach_schema_to_datashare(self, share_name, dbname, schema):
        self.execute(f"""
            GRANT USAGE ON SCHEMA "{dbname}"."{schema}" TO SHARE "{share_name}"
        """)
        self.commit()

    def attach_table_to_datashare(self, share_name, dbname, schema, table):
        self.execute(f"""
            GRANT SELECT ON TABLE "{dbname}"."{schema}"."{table}" TO SHARE "{share_name}"
        """)
        self.commit()

    def attach_view_to_datashare(self, share_name, dbname, schema, view):
        self.execute(f"""
            GRANT SELECT ON VIEW "{dbname}"."{schema}"."{view}" TO SHARE "{share_name}"
        """)
        self.commit()

    def attach_sf_account_to_datashare(self, share_name, sf_account):
        self.execute(f"""
            ALTER SHARE "{share_name}" ADD ACCOUNTS = "{sf_account}";
        """)
        self.commit()

    def import_share(self, share_account, share_name, dbname, roles=[]):
        try:
            self.execute(f"""
                CREATE DATABASE "{dbname}" FROM SHARE {share_account}."{share_name}"
            """)
            self.commit()
        except snowflake.connector.errors.ProgrammingError as e:
            if 'already exists' in e.msg:
                pass
            else:
                raise e

        for role in roles:
            self.execute(f"""
                GRANT IMPORTED PRIVILEGES ON DATABASE "{dbname}" TO ROLE "{role}";
            """)
        
        self.commit()

    def create_curator(self, username, password, role="PUBLIC", warehouse="DEFAULT", namespace="community"):
        self.set_role("ACCOUNTADMIN")
        try:
            self.bulk_execute(
                [
                    f"""CREATE USER {username} PASSWORD = '{password}' DEFAULT_ROLE = "{role}" DEFAULT_WAREHOUSE = '{warehouse}' DEFAULT_NAMESPACE = '{namespace}' MUST_CHANGE_PASSWORD = FALSE""",
                    f"""GRANT ROLE "{role}" TO USER {username}"""
                ]
            )
        except snowflake.connector.errors.ProgrammingError as e:
            if 'already exists' in e.msg:
                self.execute(
                    f"""ALTER USER {username} SET PASSWORD = '{password}' MUST_CHANGE_PASSWORD = FALSE"""
                )
                return
            else:
                raise e

    
    def create_curator_schema(self, dbname, schema, username):
        self.set_role("ACCOUNTADMIN")
        try:
            self.execute(f"""
                    CREATE SCHEMA "{dbname}"."{schema}"
            """)
        except snowflake.connector.errors.ProgrammingError as e:
            if 'already exists' in e.msg:
                pass
            else:
                raise e

        self.commit()

        self.bulk_execute([
            f"""
            GRANT USAGE ON SCHEMA "{dbname}"."{schema}" TO ROLE public
            """,
            f"""
            GRANT ALL PRIVILEGES ON SCHEMA "{dbname}"."{schema}" TO ROLE public
            """
        ])
        self.commit()

    def get_tables_and_views(self, dbname, schema_name):
        cols, rows = self.execute(f"""
            SELECT * FROM "{dbname}"."INFORMATION_SCHEMA"."TABLES" where TABLE_SCHEMA = '{schema_name}'
        """)

        data = []
        for row in rows:
            d = {}
            for i, r in enumerate(row):
                d[cols[i]] = r
            data.append(d)
        
        return data

    @property
    def conn(self):
        if not hasattr(self, '_conn') or (hasattr(self, '_conn') and self._conn is None):
            self._conn = self._open_conn()
        return self._conn

    def _open_conn(self):
        creds = dict(
            user=self.user,
            password=self.password,
            account=self.account,
            database=self.database,
            warehouse=self.warehouse,
            autocommit=False
        )

        if self.region:
            creds['region'] = self.region

        return snowflake.connector.connect(**creds)

    def close_conn(self):
        self.conn.close()
        self._conn = None

    def commit(self):
        self.conn.commit()

    def execute(self, sql):
        cs = self.conn.cursor()

        exception = None
        try:
            if hasattr(self, 'role') and self.role:
                cs.execute("USE ROLE {};".format(self.role))

            cs.execute(sql)
            cols = [col[0] for col in cs.description]
            rows = cs.fetchall()
        except Exception as e:
            exception = e
            cols = None
            rows = None
        finally:
            cs.close()

        if exception:
            raise exception

        return cols, rows

    def bulk_execute(self, sql_statements):
        cs = self.conn.cursor()

        exception = None
        try:
            if hasattr(self, 'role') and self.role:
                cs.execute("USE ROLE {};".format(self.role))

            cs.execute("BEGIN")
            for sql in sql_statements:
                cs.execute(sql)
            self.conn.commit()
        except Exception as e:
            self.conn.rollback()
            exception = e
            cols = None
            rows = None
        finally:
            cs.close()

        if exception:
            raise exception

    @property
    def version(self):
        cs = self.conn.cursor()
        version = None
        try:
            cs.execute("SELECT current_version()")
            version = cs.fetchone()[0]
        finally:
            cs.close()
        return version
