docker run -d --restart=always \
    --name ipfs_host \
    -v `pwd`/ipfs-staging:/export \
    -v `pwd`/ipfs-data:/data/ipfs \
    -p 4001:4001 \
    -p 8080:8080 \
    -p 127.0.0.1:5001:5001 \
    ipfs/go-ipfs:latest
