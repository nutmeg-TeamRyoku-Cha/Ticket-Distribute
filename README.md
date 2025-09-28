# コミット
- コミットメッセージはissue番号を載せる
- コミットメッセージは行った開発を端的にわかりやすく書く（長すぎないように注意する）
- コミットメッセージラベルを付ける
    - [add] file or directory の追加
    - [mod] file or directory の編集
    - [fix] file or directory のバグや軽微な修正
    - [del] file or directory の削除
    - [otr] その他
- ex)
    - `git commit -m "[add] model group (#1)"`
    - `git commit -m "[fix] login page (#2)"`
    - `git commit -m "[mod] mypage (#3)"`

# ブランチ
- mainブランチ
    安定ブランチ，本番用ブランチ
- developブランチ
    開発用ブランチ，開発段階での安定ブランチ，これを公開するときに安定ブランチにマージ

- feature ブランチ  
    開発するときはdevelopブランチからfeatureブランチを切る
    - feature/[NAME]/[ISUEE_NUM]-[TITLE]
        - 機能の追加や変更などを行うブランチ，developブランチから派生
        - ex) feature/dodo/1-create-view-env
    - fix/[NAME]/[ISUEE_NUM]-[TITLE]
        - バグの修正などを行うブランチ，developブランチから派生
        - ex) fix/dodo/2-fix-view-env