docker-compose up --detach
echo "starting container. Waiting for container to finish starting"
sleep 10
echo "finished waiting. loading db"
docker exec -it ece356-final-project_db_1 bash -c "cd var/lib/mysql-files/;mysql -uroot -proot test_db < ./DatabaseLoadingScripts/load-data.sql;"
docker exec -it ece356-final-project_db_1 bash -c "mysql -uroot -proot;"