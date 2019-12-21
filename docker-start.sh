cd ./database_docker_compose
docker-compose up -d

cd ../

docker pull maershov/default_team_comment
docker pull maershov/default_team_session
docker pull maershov/default_team

docker-compose up