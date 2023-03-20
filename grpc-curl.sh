# 土曜の午後10時にログイン(2023/3/18)
grpcurl -plaintext -d \
'{
    "userID": "b5f836df-acbe-422e-393d-31b570fd2403",
    "createdAt": 1679176800
}' \
 localhost:9000 api.ActionService/Login

# 土曜の午後10時にモンスターAが任意のモンスターを倒す(2023/3/18)
grpcurl -plaintext -d \
'{
    "userID": "b5f836df-acbe-422e-393d-31b570fd2403",
    "opponentMonsterID": "b0ecc978-f30b-6757-0816-5a4badfe67ff",
    "createdAt": 1679176800
}' \
 localhost:9000 api.ActionService/Battle

# 土曜の午後10時にモンスターAのレベルが2上がる(2023/3/18)
grpcurl -plaintext -d \
'{
    "userID": "b5f836df-acbe-422e-393d-31b570fd2403",
    "monsterID": "b0ecc978-f30b-6757-0816-5a4badfe67ff",
    "level": 2,
    "createdAt": 1679176800
}' \
 localhost:9000 api.ActionService/LevelUp

# 日曜の午後10時にログイン(2023/3/19)
grpcurl -plaintext -d \
'{
    "userID": "b5f836df-acbe-422e-393d-31b570fd2403",
    "createdAt": 1679263200
}' \
 localhost:9000 api.ActionService/Login

 # 日曜の午後10時にモンスターBが特定のモンスターを倒す(2023/3/19)
grpcurl -plaintext -d \
'{
    "userID": "b5f836df-acbe-422e-393d-31b570fd2403",
    "opponentMonsterID": "69b09472-765d-7006-1847-f1571557cef0",
    "createdAt": 1679263200
}' \
 localhost:9000 api.ActionService/Battle

 # 日曜の午後10時にモンスターBがレベル１上がる(2023/3/19)
grpcurl -plaintext -d \
'{
    "userID": "b5f836df-acbe-422e-393d-31b570fd2403",
    "monsterID": "dbae98d8-9cee-174d-b855-f2c89e14bb02",
    "level": 1,
    "createdAt": 1679263200
}' \
 localhost:9000 api.ActionService/LevelUp