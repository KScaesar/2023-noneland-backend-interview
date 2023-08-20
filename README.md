# 2023 noneland backend interview

## project layout

```text
├── internal
│   ├── di          <-- Composition Root
│   ├ 
│   ├                   特定商業邏輯的程式碼, 類似 Clean Architecture
│   ├── api         <-- http restful.   Adapter Layer & Framework & Driver  
│   ├── app         <-- Use Case Layer
│   ├── entity      <-- Entity Layer
│   ├── external    <-- 3rd party api.  Adapter Layer & Framework & Driver  
│   └── database    <-- gorm.           Adapter Layer & Framework & Driver  
│
└── pkg             <-- tool package: 類似 AOP 精神, 用於公司所有專案的程式碼
     ├── errors
     ├── config.go
     ├── db.go
     ├── gin.go
     ├── general_dto.go
     ├── http_client.go
     └── validator.go
```

## 作答結果

**一般**  
- [ok] 回傳兩種餘額 `{url}/api/v1/exchange/summary/balance`
   - [test code](./internal/api/exchange_test.go#L33)
   - [impl code](./internal/external/exchange_service.go#L31)
- [▲] 需要顯示 `現貨` 帳戶轉出轉入紀錄 `{url}/api/v1/exchange/spot/transactions`
   - 完成一半, 不了解情境, 無法作答
   - [impl code](./internal/app/tx_backup.go#L27)

**加分**  
- [fail] 後台的呼叫不應該影響報價邏輯
   - 看完公司回覆還是覺得不可能做到, 可能我誤解使用情境
- [ok] 請撰寫可被測試的程式碼
- [ok] 架構調整

## 前情提要

假定我們自己的 `中心化` 交易所，已有錢包 app 完成初步現貨交易

現在正在處理『風險控制後台』管理介面開發，並提供 api 給『風險控制後台』前端開發人員串接

『風險控制後台』主要處理邏輯為，用戶下單現貨時，會在第三方開合約來做避險動作，因此需要有個後台讓風險控制人員觀看

我們會需要呼叫 `第三方XX交易所` 報價商後台所需要的參數

以下是 PM 提出的規格

## 基本規格

1. 第一隻 api，同時回傳兩種餘額
   1. 需要顯示 `XX交易所` 的 `現貨` 帳戶 USDT 餘額
   1. 需要顯示 `XX交易所` 的 `合約` 帳戶 USDT 餘額
2. 需要顯示 `現貨` 帳戶轉出轉入紀錄（`第三方XX交易所僅提供一個月內的交易紀錄查詢`）
   1. 根據法律遵循，我們應該保存 `6年` 內的所有交易紀錄

```
問題3
風控系統 api 的 url 要怎麼定義? 與問題2有關
日後需要串接不同的交易所嗎?
response 需要顯示是哪一家交易所嗎?
可以提供風控系統 response 範例?

1.不需要串接不同交易所
{{url}}/summary/balance

2.需要串接不同交易所
{{url}}/{{exchangeId}}/summary/balance
```

```
問題4
保存 `6年` 內的所有交易紀錄
這句話讓我感覺上面所謂的第三方, 不是真的第三方
而是自家公司就是交易所

前情提要說： 是處理交易所風控後台
既然是風控業務
那為什麼有 保存 `6年` 內的所有交易紀錄的需求
交易紀錄的保存
這應該是交易所伺服器要負責的吧
和風控後台的業務關聯是什麼?

我認知的風控應該是基於既成事實(已有的資料)
進行stream etl資料處理
獲得交易紀錄衍伸性的資料
意思是衍伸性的資料也要保留?

所以作業所說的後台是什麼?
1.風控後台
2.交易所後台


衍伸性的資料的6年又是怎麼定義
olap常常會進行時間聚合 by 年 月 日 週

保存 `6年` 內的所有交易紀錄
是希望我提出什麼解決方案? 
單純文字敘述, coding, 畫架構圖? 
我覺得很模糊, 不知道此問題想要什麼答案
```

## 備註

可使用任何第三方工具、服務、套件來設計

補充一下，簡單來說目前後台會需要呼叫 `第三方XX交易` 來`取得`資料並於後台`顯示`

而 `第三方XX交易` 有 `現貨` 與 `合約` 兩個 api 的 endpoint

有任何題目上的問題都可以寫信發問

## 加分題

1. 目前前端 app 的現貨買賣幣種報價使用的是 `現貨` 相關的 api，並且 api 存在呼叫限制，後台的呼叫不應該影響報價邏輯
   這邊可以畫出流程圖、架構圖或是直接程式碼都可以，是個開放式的問題
2. 請撰寫可被測試的程式碼，或是直接附上測試程式
   預期你可以使用 mock 的方式來處理第三方，並通過 di 的方式注入，但使用其他方式也行
3. 架構也能調整，假設你覺得有更好的改法

```
問題5
不應該影響報價邏輯 是什麼意思?
是否指後台的呼叫, 不應該佔據 交易所 api ratelimter 的額度用量?

如果交易所不是自家公司
然後要求給使用者的限制要維持原本 ratelimter 的額度用量 (因為題目說不應該影響報價邏輯, 所以我疑惑何謂報價邏輯?)
這樣的需求不可能作到
除非交易所另外開一個 api 給後台使用
```

## XX交易所 api 文件

### 現貨帳戶 api

#### 前端 app 現貨下單的 api

api 權重：5

這邊為加分題，前端 app 使用的 endpoint 與下方 `取得現貨帳戶餘額` 一致

因此會受限於下方 `{{ url1 }}/exchangeInfo` 的限制

#### api 使用限制

- REQUEST_WEIGHT 單位時間請求權重之和的上限
- RAW_REQUESTS 單位時間請求次數上限

```
問題6
REQUEST_WEIGHT, REQUEST_WEIGHT 有什麼差別?

告知我 ratelimter 目的是希望我回答什麼?
希望風控後台設計一個方式
不要去觸發額度限制?

這個 現貨 api 也會被 前端呼叫?
還是只專屬後台使用?
```

GET `{{ url1 }}/exchangeInfo`

api 權重：0

request:

Header: `Authorization: Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==`

response:

```json
{
  "timezone": "UTC",
  "serverTime": 1565246363776,
  "rateLimits": [
     {
        "rateLimitType": "REQUEST_WEIGHT",
        "interval": "MINUTE",
        "intervalNum": 1,
        "limit": 1200
     },
     {
        "rateLimitType": "RAW_REQUESTS",
        "interval": "MINUTE",
        "intervalNum": 5,
        "limit": 6100
     }
  ]
}
```

#### 取得現貨帳戶餘額

GET `{{ url1 }}/spot/balance`

```
問題7
userId 從哪邊帶入? 沒有 userId 如何指定抓到特定使用者的資料
HEADER, token or cookie?
風控後台會有第三方的系統 userId 的資料?
理論上風控系統只會有自己系統的 userId
回到問題2, 第三方是指什麼?
難道要做 oauth2.0, 先跟第三方拿user資料?
這個作業需要考慮 userId 是怎麼來的?
```

api 權重：5

request:

Header: `Authorization: Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==`

response:

```json
{
  "free": "10.12345"
}
```

#### 現貨帳戶轉入轉出紀錄

GET `{{ url1 }}/spot/transfer/records`

api 權重：5

request 參數：

名稱 | 類型 | 是否必填 | 描述
----|------|--------|---------
startTime | LONG | NO |
endTime | LONG | NO |
endTime | LONG | NO |
current | LONG | NO | 當前回傳頁數，預設為 1
size | LONG | NO | 回傳筆數，預設 10，最大 100

Header: `Authorization: Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==`

response:

- status: PENDING (等待), CONFIRMED (成功), FAILED (失敗);

```json
{
   "rows": [
      {
         "amount": "0.10000000",
         "asset": "BNB",
         "status": "CONFIRMED",
         "timestamp": 1566898617,
         "txId": 5240372201,
         "type": "IN"
      },
      {
         "amount": "5.00000000",
         "asset": "USDT",
         "status": "CONFIRMED",
         "timestamp": 1566888436,
         "txId": 5239810406,
         "type": "OUT"
      },
      {
         "amount": "1.00000000",
         "asset": "EOS",
         "status": "CONFIRMED",
         "timestamp": 1566888403,
         "txId": 5239808703,
         "type": "IN"
      }
   ],
   "total": 3
}
```

### 合約帳戶 api

#### api 使用限制

- REQUEST_WEIGHT 單位時間請求權重之和的上限
- RAW_REQUESTS 單位時間請求次數上限

GET `{{ url2 }}/exchangeInfo`

api 權重：0

request:

Header: `Authorization: Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==`

response:

```json
{
  "timezone": "UTC",
  "serverTime": 1565246363776,
  "rateLimits": [
     {
        "rateLimitType": "REQUEST_WEIGHT",
        "interval": "MINUTE",
        "intervalNum": 1,
        "limit": 1200
     },
     {
        "rateLimitType": "RAW_REQUESTS",
        "interval": "MINUTE",
        "intervalNum": 5,
        "limit": 6100
     }
  ]
}
```


#### 取得合約帳戶餘額

GET `{{ url2 }}/futures/balance`

api 權重：5

request:

Header: `Authorization: Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==`

response:

```json
{
  "free": "10.12345"
}
```



