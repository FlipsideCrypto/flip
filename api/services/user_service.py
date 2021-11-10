import requests
import sqlalchemy

from schemas import user as user_schema
from core.environment import GRAPHQL_API
from sqlalchemy.orm import session
from models import User


def get_user(jwt: str) -> user_schema.User:
    payload="{\"query\":\"{\\n    me { \\n        id,\\n        username,\\n        email,\\n        organization {\\n            id,\\n            name\\n        }\\n    }\\n}\",\"variables\":{}}"
    headers = {
        'authorization': jwt,
        'Content-Type': 'application/json'
    }

    response = requests.request("POST", GRAPHQL_API, headers=headers, data=payload)
    if response.status_code != 200:
        return None

    user = response.json()
    if not user or (user and not user['data']['me']):
        return None

    return user_schema.User(
        id=user['data']['me']['id'],
        user_name=user['data']['me']['username'],
        organization_id=user['data']['me']['organization']['id'],
        organization_name=user['data']['me']['organization']['name']
    )


def get_jwt(username: str, password: str) -> str:
    payload='{"query":"mutation UserLoginMutation($input: UserLoginInput\u0021) {\\n  userLogin(input: $input) {\\n    token\\n  }\\n}\\n","variables":{"input":{"email":"'+username+'","password":"'+password+'"}},"operationName":"UserLoginMutation"}'
    headers = {
        'Content-Type': 'application/json'
    }

    response = requests.request("POST", GRAPHQL_API, headers=headers, data=payload)
    if response.status_code != 200:
        return None

    auth_data = response.json()
    if auth_data and auth_data.get('errors'):
        return None

    return auth_data['data']['userLogin']['token']


def save_db_user(db: session, user_id: str, username: str, org_id: str) -> User:
    query = db.query(User).filter(User.id == user_id)
    try:
        return query.one()
    except sqlalchemy.exc.NoResultFound:
        pass

    u = User(
        id=user_id,
        username=username,
        org_id=org_id
    )

    db.add(u)
    db.flush()
    db.commit()
    return u
