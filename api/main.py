from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

from actions.api_v1.api import api_router

app = FastAPI(
    title="FLIP API",
    description="This API provides a REST interface to FLIP CLI.",
)

origins = [    
    "http://localhost",
    "http://localhost:3000",
    "http://127.0.0.1",
    "http://127.0.0.1:3000"
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

app.include_router(api_router)
