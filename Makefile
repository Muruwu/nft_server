SERVER_NAME = mc_nft_server

server:
	#kitex -type protobuf -service ${SERVER_NAME} -I idls/ idls/mc_nft.proto
	kitex_external -type protobuf -service ${SERVER_NAME} -I idls/ idls/mc_nft.proto