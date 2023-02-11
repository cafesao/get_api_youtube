# Build image
docker build -f build/package/Dockerfile -t getyoutube .

# Init container temporary
id=$(docker create getyoutube)

# Copy executable for ./build
docker cp $id:/app/getYoutube ./build/

# Delete container and image
docker rm -v $id
docker rmi getyoutube