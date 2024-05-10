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
    uuid restaurant "店ID"
    uuid organizer "主催者"
    uuid[] perticipant "参加者"
    int perticipant_num "参加人数"
    timestamp conducted_at "開催日時"
    string picture "関連画像"
    string comment "備考"
  }

  Restaurant {
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

- Nomikuiテーブルの参加者,Userテーブルの募集/参加関連を`Entry`テーブルで解決
- Restaurantテーブルの`[]string tags`を解決→別テーブルを作る
- Userテーブルの`[]uuid favorite`を解決→別テーブルを作る

```mermaid
---
title: nomikui-first
---
erDiagram

  Nomikui {
    uuid id PK
    uuid restaurant "店ID"
    uuid organizer "主催者"
    int perticipant_num "参加人数"
    timestamp conducted_at "開催日時"
    string picture "関連画像"
    string comment "備考"
  }

  Entry {
    uuid userid "対象ユーザID"
    uuid nomikuiid "対象nomikuiID"
    bool represent "募集者か"
    timestamp created_at "応募日時"
  }

  Restaurant {
    uuid id PK
    string name "店名"
    string area "エリア"
    string genre "ジャンル"
    int perticipantnum "参加人数"
    timestamp conducted_at "開催日時"
    string picture "関連画像"
    string comment "備考"
  }

  Tags {
    uuid restaurant "店ID"
    string content "タグ名"
  }

  User {
    uuid id PK
    string name "ユーザ名"
    string traQid "認証済traQid"
    int userexp "ユーザexp"
  }

  Favorite {
    uuid userid "ユーザID"
    uuid restaurantid "店ID"
    timestamp conducted_at "開催日時"
  }

```

## second:第二正規化

- エリア,ジャンル番号の実装

```mermaid
---
title: nomikui-second
---
erDiagram

  Nomikui {
    uuid id PK
    uuid restaurant "店ID"
    uuid organizer "主催者"
    int perticipant_num "参加人数"
    timestamp conducted_at "開催日時"
    bool isopen "募集中か"
    string picture "関連画像"
    string comment "備考"
  }

  Entry {
    uuid userid "対象ユーザID"
    uuid nomikuiid "対象nomikuiID"
    bool represent "募集者か"
    timestamp created_at "応募日時"
  }

  Restaurant {
    uuid id PK
    string name "店名"
    uuid areaid "エリアid"
    uuid genreid "ジャンルid"
    int perticipantnum "参加人数"
    timestamp conducted_at "開催日時"
    string picture "関連画像"
    string comment "備考"
  }

  Area {
    uuid id "エリアid"
    strinf areaname "エリア名"
  }

  Genre {
    uuid id "ジャンルid"
    strinf areaname "ジャンル名"
  }

  Tags {
    uuid restaurant "店ID"
    string content "タグ名"
  }

  User {
    uuid id PK
    string name "ユーザ名"
    string traQid "認証済traQid"
    int userexp "ユーザexp"
  }

  Favorite {
    uuid userid "ユーザID"
    uuid restaurantid "店ID"
    timestamp conducted_at "開催日時"
  }

```

## third:第三正規化

<!-- ```mermaid
---
title:nomikui-second
---

``` -->
