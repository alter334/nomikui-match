## draft:正規化とか考えずとりあえず欲しい情報書き出す

- nomikui
    - ID
    - 店
    - 主催者
    - 参加者
    - 希望人数
    - 開催日時
    - 締切日時
    - 備考

- 店データ
    - ID
    - 名前
    - 場所
    - タグ
    - ジャンル
    - 価格帯
    - 収容人数
    - レビュー
    - 関連画像

- ユーザに紐付くデータ
    - ID(traQUUID)
    - 過去のnomikui募集
    - 過去のnomikui参加
    - お気に入り店
    - nomikuiEXP

<!-- draft状態 -->

```mermaid
---
title: nomikui-draft
---
erDiagram

  Nomikui {
    uuid id PK
    string name "店名"
    string area "エリア"
    string[] tags"タグ"
    string genre "ジャンル"
    int perticipantnum "参加人数"
    timestamp conducted_at "開催日時"
    string picture "関連画像"
    string comment "備考"
  }

  Restaurant {
    uuid id PK
    string name "店名"
    uuid organizer "主催者"
    uuid[] perticipant "参加者"
    int perticipant_num "参加人数"
    timestamp conducted_at "開催日時"
    string picture "関連画像"
    string comment "備考"
  }

  User {
    uuid id PK
    string name "ユーザ名"
    string traQid "認証済traQid"
    uuid[] open_nomikui "nomikui募集"
    uuid[] taken_nomikui "nomikui参加"
    uuid[] favorite "お気に入り店"
    int userexp "ユーザexp"
  
  }
```
## first:第一正規化


