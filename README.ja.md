# Nihon's Kitchen

Nihon's Kitchen へようこそ！このプロジェクトは、日本に住んでいる英語圏の人たちが、日本の食材を使った素晴らしいレシピを英語で共有できるようにするために作られました。日本に来たばかりの人にとって、日本の食品ラベルのパッケージを読んで理解するのは難しいかもしれません。そこで、ユーザーがバーコードをスキャンして英訳を入力することで、他の人たちを助けるためのデータベースを作りました。

### フロントエンドのレポジトリーはこちら

https://github.com/nihonskitchen/frontend

## リンク

Nihon's Kitchen へのアクセスはこちら：

https://nihonskitchen-prod.web.app

## アプリの動作確認方法

アプリをインストールしてコードの動作を確認したい場合は、このリポジトリをフォークして以下を実行してください。

```
go get https://github.com/nihonskitchen/serverside

go run main.go
```

## 機能

1. レシピの投稿と閲覧
2. バーコードスキャナによる製品情報の検索と共有
3. レシピの材料をショッピングリストに追加して、オンラインで購入
4. ユーザープロフィールの作成と管理
5. レシピ検索

## 追加予定の機能

1. レシピを「お気に入り」リストに追加
2. レシピへのコメント
3. レシピの評価
4. レシピの価格表示
5. レシピの編集と削除
6. 冷蔵庫にある材料を元にレシピ提案

## 使用した技術

1. Go: https://golang.org/
2. Fiber: https://docs.gofiber.io/
3. Google App Engine: https://cloud.google.com/appengine
4. Firebase: https://firebase.google.com
   - Firebase Hosting
   - Firebase Authentication
   - Cloud Storage
   - Cloud Firestore
