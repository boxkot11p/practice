#!/bin/bash

echo 'ユーザーの情報'
echo 'select * from User u\G' | docker-compose exec -T spanner-cli spanner-cli -p mission-project -i mission-instance -d mission-database
echo ''
echo 'ユーザーが所持するアイテム情報'
echo 'select ui.UserID, i.Name from UserItem ui left join Item i on ui.ItemID = ui.ItemID\G' | docker-compose exec -T spanner-cli spanner-cli -p mission-project -i mission-instance -d mission-database
echo ''
echo 'ユーザーが所有するモンスターの情報'
echo 'select um.UserID, m.Name, um.Level from UserMonster um left join Monster m on um.MonsterID = m.MonsterID\G' | docker-compose exec -T spanner-cli spanner-cli -p mission-project -i mission-instance -d mission-database
echo ''
echo 'ミッション達成状況の情報'
echo 'select uma.UserID, m.Name, uma.CreatedAt from UserMissionAchievement uma left join Mission m on uma.MissionID = m.MissionID order by uma.CreatedAt asc\G' | docker-compose exec -T spanner-cli spanner-cli -p mission-project -i mission-instance -d mission-database