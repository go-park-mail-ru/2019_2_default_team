sudo docker network create default-net

cd ./database_docker_compose
sudo docker-compose up -d

cd ../

sudo docker pull maershov/default_team_comment
sudo docker pull maershov/default_team_session
sudo docker pull maershov/default_team

sudo docker-compose up