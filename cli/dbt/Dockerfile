FROM fishtownanalytics/dbt:0.20.1
WORKDIR /support
COPY . .
RUN mkdir /root/.dbt
COPY profiles.yml /root/.dbt
WORKDIR /sql_models
ENTRYPOINT [ "bash"]