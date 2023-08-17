# 2023 noneland backend interview

## 前情提要

你現在在處理交易所風控後台的新 api 開設，以下有 `本次測驗需要新增的兩隻 api 規格`

```
問題1
此專案目的是處理哪間公司的交易所風控? 與問題2有關
```

並且有提供第三方 `XX交易所` api 文件的情況下

```
問題2
此處定義的第三方是什麼?
1.自家公司的交易所, 只是架構以 microservices 的方式建立
2.真的是別人家的系統, 沒辦法要求對方配合我
```

以下是 PM 提出的規格

## 基本規格

1. 第一隻 api，同時回傳兩種餘額
   1. 需要顯示 `XX交易所` 的 `現貨` 帳戶 USDT 餘額
   1. 需要顯示 `XX交易所` 的 `合約` 帳戶 USDT 餘額
2. 需要顯示 `現貨` 帳戶轉出轉入紀錄
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
為什麼需要 保存 `6年` 內的所有交易紀錄
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

6年又是怎麼定義
衍伸性的資料 常常會進行時間聚合 by 年 月 日 週

保存 `6年` 內的所有交易紀錄
是希望我提出什麼解決方案? 
單純文字敘述, coding, 畫架構圖? 
我覺得很模糊, 不知道此問題想要什麼答案
```

## 備註

可使用任何第三方工具、服務、套件來設計

## 加分題

1. 目前前端 app 的幣種報價使用的是 `現貨` 相關的 api，並且 api 存在呼叫限制，後台的呼叫不應該影響報價邏輯
2. 請撰寫可被測試的程式碼，或是直接附上測試程式
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

#### api 使用限制

- REQUEST_WEIGHT 單位時間請求權重之和的上限
- RAW_REQUESTS 單位時間請求次數上限

```
問題6
REQUEST_WEIGHT, REQUEST_WEIGHT 有什麼差別?

告知我 ratelimter 目的是希望我回答什麼?
希望風控後台設計一個方式
不要去觸發額度限制?
```

GET `{{ url1 }}/exchangeInfo`

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

response:

```json
{
  "free": "10.12345"
}
```

#### 現貨帳戶轉入轉出紀錄

GET `{{ url1 }}/spot/transfer/records`

request 參數：

名稱 | 類型 | 是否必填 | 描述
----|------|--------|---------
startTime | LONG | NO |
endTime | LONG | NO |
endTime | LONG | NO |
current | LONG | NO | 當前回傳頁數，預設為 1
size | LONG | NO | 回傳筆數，預設 10，最大 100

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

response:

```json
{
  "free": "10.12345"
}
```



