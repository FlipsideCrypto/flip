FROM fishtownanalytics/dbt:0.20.1
WORKDIR /support
COPY . .
RUN mkdir /root/.dbt
COPY profiles.yml /root/.dbt
# RUN apt-get update
# RUN apt-get -y install curl zsh
WORKDIR /sql_models
# RUN sh -c "$(wget https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh -O -)"
# RUN sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
ENTRYPOINT [ "bash"]