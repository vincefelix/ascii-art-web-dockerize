echo "----------------------------------------------------------------------------------"
echo "docker image build -f Dockerfile -t tryhere ."
docker image build -f Dockerfile -t tryhere .
echo "----------------------------------------------------------------------------------"
echo "                                                                                  "
echo "----------------------------------------------------------------------------------"
echo "docker container run -p 8080 --detach --name Dockerfile tryhere"
docker container run -p 8080 --detach --name Dockerfile tryhere
echo "----------------------------------------------------------------------------------"
echo "                                                                                  "
echo "----------------------------------------------------------------------------------"
echo "docker exec -it Dockerfile /bin/bash"
docker exec -it Dockerfile /bin/bash
echo "----------------------------------------------------------------------------------"


