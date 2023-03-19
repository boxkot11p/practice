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
('22e8c4e4-5ed6-bcfa-0a90-8ac9ec721461', '1f295f46-0cb9-fdc0-9124-853c0e4881a2', 'b0ecc978-f30b-6757-0816-5a4badfe67ff', 4),
('93ac47bf-a92e-27df-a045-46f386f2cda3', '1f295f46-0cb9-fdc0-9124-853c0e4881a2', '69b09472-765d-7006-1847-f1571557cef0', 7);

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
('a75fd9e5-6cbc-2932-2b65-8249b241e952', NULL, NULL, '特定のモンスターを倒す', "BATTLE", 'NONE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '69b09472-765d-7006-1847-f1571557cef0', NULL, NULL, 100),
('d638c9ec-eacb-b38b-5d7e-8ed6dcfef67a', NULL, '1f295f46-0cb9-fdc0-9124-853c0e4881a2', '2000コイン貯める', "COIN", 'NONE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 2000, NULL),
('ca829cec-5a11-381a-e5d1-56c99c005790', NULL, NULL, 'モンスターAのレベルが5', 'NONE', "LEVEL", NULL, NULL, NULL, 'b0ecc978-f30b-6757-0816-5a4badfe67ff', 5, NULL, NULL, NULL, NULL, NULL, 100),
('0225de9e-9589-b026-251b-daa523c8731b', NULL, NULL, 'レベル5以上のモンスターが2体', "LEVEL", 'NONE', NULL, NULL, NULL, NULL, 5, 2, NULL, NULL, NULL, NULL, 100),
('8469eaf3-d0ad-bee3-4891-3b8dafad6b02', NULL, NULL, '任意のモンスターを10回倒す', "BATTLE", 'WEEKLY', 'MONDAY', 0, 0, NULL, NULL, NULL, NULL, NULL, 10, NULL, 100),
('90d49481-bfa4-296a-8841-1f58b101ddc5', NULL, NULL, 'ログイン', "LOGIN", 'DAILY', NULL, 4, 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 100),
('ef63a85a-f120-f511-52e8-627c72aed692', '0225de9e-9589-b026-251b-daa523c8731b', NULL, 'アイテムAを所有する', "ITEM", 'NONE', NULL, NULL, NULL, NULL, NULL, NULL, '69b09472-765d-7006-1847-f1571557cef0', NULL, NULL, NULL, 100);