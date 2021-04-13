presSetup:
	sh ./ShScripts/preSetup.sh

startContainerAndLoad:
	sh ./ShScripts/setUpScript.sh

connectToContainer:
	docker exec -it ece356-final-project_db_1 bash -c "mysql -uroot -proot;"
	
installCLI:
	cd ./SQLClient && go install