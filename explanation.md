## プログラムの簡単な説明

- 起動プログラム
  - https://github.com/boxkot11p/practice/blob/main/app/cmd/main.go
- grpc apiのhandler
  - https://github.com/boxkot11p/practice/tree/main/app/infra/handler
- grpc apiのprotoファイル
  - https://github.com/boxkot11p/practice/blob/main/app/proto/api/service.proto
- usecaseについて
  - https://github.com/boxkot11p/practice/tree/main/app/usecase
  - serviceで単体のusecaseを定義
  - domainとentityはserviceから利用されるEntity
    - domainは対象ドメインのロジックを記載
    - entityは操作対象のデータの情報と値のロジックを記載
  - repositoryはinfra操作の抽象
- repositoryについて
  - https://github.com/boxkot11p/practice/tree/main/app/infra/repository/database
  - databaseとのやり取りをしている

## ミッションの拡張に関しての観点

- ミッションの種類の追加のしやすさをするために下記の観点で実装しました
  - Strategyパターンで実装することでミッションの処理を抽象化しました
  
    上記を行うことで新しくミッションを追加した際も抽象化したインターフェースを利用して実装ができます
    また抽象化することで処理は隠蔽されるので簡単に実装を切り替えれます
    - 対象のpackage
    - https://github.com/boxkot11p/practice/tree/main/app/usecase/domain
  - リセットサイクルやミッション種別の値を抽象化しました

    ミッションはリセットサイクルやミッション種別が複数あるので値を抽象化し値が持つ知識を凝集しています

    例えばリセットサイクルの種類が増えた場合は値のロジックを対象のpackageに足せばよく、かつ処理の凝集から単体テストが容易なので保守性をある程度担保できるかと考えています
    - 対象package
    - https://github.com/boxkot11p/practice/tree/main/app/usecase/entity/value
  - DBのEntityとUseCaseのEntityを分けています

    DBとUseCaseでは利用する値は似ていますが使われる目的が全く違うので分けています

    上記を行うことでDBの制約に引っ張られることなくEntityを定義でき、かつRepository層でDIが出来るので変換に関する知識が凝集されます

    仮にEntityの値を抽象化したいという場合はRepository層で変換を行えばいいということです
    上記が正解かはわかりませんが、知識はどこかで凝集することで結合度は低くなると考えています
    - 対象package
    - https://github.com/boxkot11p/practice/tree/main/app/usecase/entity
    - https://github.com/boxkot11p/practice/tree/main/app/infra/repository/database

- 時間が足りずにできなかったが理想的な実装について
  - test codeの導入
    各モジュールのテスタビリティが担保されることで継続的に開発可能な環境を構築するのが理想です
    なぜならテスタビリティが悪い = 保守性が低いという意味だと考えているのでテストを導入したかったです
  - DB設計について
    - DBからデータを取りアプリケション上で処理していますが、DBから無駄な情報を取っていますしDB設計についてもテーブルの正規化をやめてArray型を使いRelationを持てばパフォーマンスが良くなりますが時間が足りませんでした
