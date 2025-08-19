# 🗓️ Go 語言 14 天速成學習日曆表

> 適合資深工程師，從零基礎快速進入 Production Ready 開發。

---

## Week 1：語法基礎 & Idiomatic Go

### Day 1｜Go 安裝 & 專案結構
- 安裝 Go、設定 VS Code
- 建立專案（`go mod init`）
- **練習**：CLI 計算器（加減乘除）

### Day 2｜基本語法（變數/型別/函式）
- 變數、常數、函式、多回傳值
- slice / map 操作
- **練習**：slice ↔ map ↔ JSON 轉換器

### Day 3｜struct / 方法 / interface
- struct 與方法
- interface（duck typing）
- **練習**：Shape interface + Circle & Rectangle 實作

### Day 4｜error / panic / defer
- 錯誤處理哲學
- panic / recover
- defer 用法（資源釋放）
- **練習**：檔案讀寫工具，錯誤需 wrap & defer 關閉

### Day 5｜併發基礎
- goroutines
- channels（buffered/unbuffered）
- select
- context（timeout / cancel）
- **練習**：worker pool（可取消任務）

### Day 6｜標準庫
- net/http
- encoding/json
- **練習**：HTTP JSON API（GET /hello）

### Day 7｜測試
- 單元測試：`testing`
- Benchmark 測試
- **練習**：為 API 寫單元測試 + benchmark 測試

---

## Week 2：實戰 & 高階應用

### Day 8｜Web 框架：Gin
- Gin router
- middleware（logging, recovery）
- **練習**：REST API skeleton（CRUD）

### Day 9｜資料庫：GORM + SQLite/Postgres
- ORM 基礎
- DB migration / CRUD
- **練習**：待辦清單 CRUD API

### Day 10｜身份驗證：JWT
- JWT 登入流程
- middleware 實作
- **練習**：API with JWT login

### Day 11｜高併發 & 快取
- Redis 操作
- 限流（token bucket）
- **練習**：短網址服務（含 Redis 快取 & 限流）

### Day 12｜效能調優
- pprof
- benchmark
- **練習**：壓測 API，定位瓶頸並優化

### Day 13｜Docker 化專案
- multi-stage build
- 單一 Go binary 容器
- **練習**：Docker 化 REST API

### Day 14｜雲端部署 & 綜合專案
- 部署到 Heroku / Fly.io / Kubernetes
- **最終專案**：待辦清單微服務  
  - Gin + GORM + Postgres  
  - JWT Auth  
  - Docker 部署  
  - pprof 壓測  

---

## ✅ 學習建議
1. 每天 commit → Push 到 GitHub 當作品集  
2. 用測試驅動 → 每個模組至少一個單元測試  
3. 從第 5 天開始 → 習慣用 `go test -bench` / `pprof` 找瓶頸  
4. 錯誤處理要 wrap，API 要帶 context，goroutines 要可控  




## 賽道紀錄
Go 學習賽道（基礎語法 → 併發 → Web API → 微服務 → 效能優化 → 雲原生）