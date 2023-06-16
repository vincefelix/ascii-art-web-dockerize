echo "----------------------------------------------------------------------------------"
echo "1) Building the image: docker image build -f Dockerfile -t audit ."
docker image build -f Dockerfile -t audit .
echo "----------------------------------------------------------------------------------"
echo "                                                                                  "
echo "----------------------------------------------------------------------------------"
echo "2) Executing the container: docker container run -p 8080 --detach --name Dockerfile audit"
docker container run -p 8080:8080 --detach --name Dockerfile audit
echo "----------------------------------------------------------------------------------"
echo "                                                                                  "
echo "----------------------------------------------------------------------------------"
echo "3) Checking the running containers: docker ps"
docker ps
echo "----------------------------------------------------------------------------------"
echo "                                                                                  "
echo "----------------------------------------------------------------------------------"
echo "4) Entering the container: docker exec -it Dockerfile /bin/bash"
echo "4-a) Now enter the "ls -l" command"
docker exec -it Dockerfile /bin/bash
echo "----------------------------------------------------------------------------------"
echo "                                                                                  "
echo "----------------------------------------------------------------------------------"
echo "5) Inspecting the metadatas: docker inspect audit"
docker inspect audit
echo "----------------------------------------------------------------------------------"


