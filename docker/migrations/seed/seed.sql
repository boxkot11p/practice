INSERT INTO User(
  UserID,
  Name,
  Coin
) 
VALUES
('b5f836df-acbe-422e-393d-31b570fd2403', 'ユーザー', 1800);

INSERT INTO Item(
  ItemID,
  Name
)
VALUES
('1f295f46-0cb9-fdc0-9124-853c0e4881a2', 'アイテムA');

INSERT INTO Monster(
  MonsterID,
  Name
)
VALUES
('b0ecc978-f30b-6757-0816-5a4badfe67ff', 'モンスターA'),
('dbae98d8-9cee-174d-b855-f2c89e14bb02', 'モンスターB'),
('69b09472-765d-7006-1847-f1571557cef0', '特定のモンスター');

INSERT INTO UserMonster(
  UserMonsterID,
  UserID,
  MonsterID,
  Level
)
VALUES
('22e8c4e4-5ed6-bcfa-0a90-8ac9ec721461', 'b5f836df-acbe-422e-393d-31b570fd2403', 'b0ecc978-f30b-6757-0816-5a4badfe67ff', 4),
('93ac47bf-a92e-27df-a045-46f386f2cda3', 'b5f836df-acbe-422e-393d-31b570fd2403', 'dbae98d8-9cee-174d-b855-f2c89e14bb02', 7);

INSERT INTO Mission(
  MissionID,
  OrderConditionMissionID,
  GiftItemID,
  Name,
  Category,
  ResetCycle,
  ResetWeek,
  ResetHour,
  ResetTime,
  ConditionLevelMonsterID,
  ConditionLevel,
  ConditionLevelHaveMonsterNumber,
  ConditionItemID,
  ConditionTargetMonsterID,
  ConditionSubjugationMonsterNumber,
  ConditionHaveCoin,
  GiftCoin
) 
VALUES 
('a75fd9e5-6cbc-2932-2b65-8249b241e952', "", "", '特定のモンスターを倒す', "BATTLE", 'NONE', "", 0, 0, "", 0, 0, "", '69b09472-765d-7006-1847-f1571557cef0', 0, 0, 100),
('d638c9ec-eacb-b38b-5d7e-8ed6dcfef67a', "", '1f295f46-0cb9-fdc0-9124-853c0e4881a2', '2000コイン貯める', "COIN", 'NONE', "", 0, 0, "", 0, 0, "", "", 0, 2000, 0),
('ca829cec-5a11-381a-e5d1-56c99c005790', "", "", 'モンスターAのレベルが5', 'LEVEL', "NONE", "", 0, 0, 'b0ecc978-f30b-6757-0816-5a4badfe67ff', 5, 0, "", "", 0, 0, 100),
('0225de9e-9589-b026-251b-daa523c8731b', "", "", 'レベル5以上のモンスターが2体', "LEVEL", 'NONE', "", 0, 0, "", 5, 2, "", "", 0, 0, 100),
('8469eaf3-d0ad-bee3-4891-3b8dafad6b02', "", "", '任意のモンスターを10回倒す', "BATTLE", 'WEEKLY', 'MONDAY', 0, 0, "", 0, 0, "", "", 10, 0, 100),
('90d49481-bfa4-296a-8841-1f58b101ddc5', "", "", 'ログイン', "LOGIN", 'DAILY', "", 4, 0, "", 0, 0, "", "", 0, 0, 100),
('ef63a85a-f120-f511-52e8-627c72aed692', '0225de9e-9589-b026-251b-daa523c8731b', "", 'アイテムAを所有する', "ITEM", 'NONE', "", 0, 0, "", 0, 0, '1f295f46-0cb9-fdc0-9124-853c0e4881a2', "", 0, 0, 100);