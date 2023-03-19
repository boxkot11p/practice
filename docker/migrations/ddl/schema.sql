CREATE TABLE User (
  UserID STRING(MAX) NOT NULL,
  Name STRING(MAX) NOT NULL,
  Coin INT64 NOT NULL
) PRIMARY KEY(UserID);

CREATE TABLE Item (
  ItemID STRING(MAX) NOT NULL,
  Name STRING(MAX) NOT NULL
) PRIMARY KEY (ItemID);

CREATE TABLE Monster (
  MonsterID STRING(MAX) NOT NULL,
  Name STRING(MAX) NOT NULL
) PRIMARY KEY(MonsterID);

CREATE TABLE UserItem (
  UserID STRING(MAX) NOT NULL,
  ItemID STRING(MAX) NOT NULL
) PRIMARY KEY(UserID, ItemID);

CREATE TABLE UserMonster (
  UserMonsterID STRING(MAX) NOT NULL,
  UserID STRING(MAX) NOT NULL,
  MonsterID STRING(MAX) NOT NULL,
  Level INT64 NOT NULL,
) PRIMARY KEY(UserMonsterID);

CREATE TABLE Mission (
  MissionID STRING(MAX) NOT NULL,
  OrderConditionMissionID STRING(MAX),
  GiftItemID STRING(MAX),
  Name STRING(MAX) NOT NULL,
  Category STRING(MAX) NOT NULL,
  ResetCycle STRING(MAX) NOT NULL,
  ResetWeek STRING(MAX),
  ResetHour INT64,
  ResetTime INT64,
  ConditionLevelMonsterID STRING(MAX),
  ConditionLevel INT64,
  ConditionLevelHaveMonsterNumber INT64,
  ConditionItemID STRING(MAX),
  ConditionTargetMonsterID STRING(MAX),
  ConditionSubjugationMonsterNumber INT64,
  ConditionHaveCoin INT64,
  GiftCoin INT64
) PRIMARY KEY (MissionID);

CREATE TABLE UserMissionAchievement (
  UserMissionAchievementID STRING(MAX) NOT NULL,
  UserID STRING(MAX) NOT NULL,
  MissionID STRING(MAX) NOT NULL,
  CreatedAt TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true)
) PRIMARY KEY (UserMissionAchievementID);

CREATE TABLE UserBattleHistory (
  UserActionHistoryID STRING(MAX) NOT NULL,
  UserID STRING(MAX) NOT NULL,
  MonsterID STRING(MAX),
  CreatedAt INT64 NOT NULL
) PRIMARY KEY (UserActionHistoryID);;