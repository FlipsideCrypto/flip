FROM tiangolo/uvicorn-gunicorn-fastapi:python3.7

ARG ENVIRONMENT
ARG ENV
ARG VERSION

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1
ENV PYTHONPATH=/app
ENV ENVIRONMENT="${ENVIRONMENT}"
ENV ENV="${ENV}"
ENV DD_ENV="${ENV}"
ENV DD_VERSION="${VERSION}"

EXPOSE 80

COPY requirements*.txt ./
RUN pip install --upgrade --no-cache-dir pip && \
    pip install --upgrade --no-cache-dir -r requirements.txt


COPY . /api
