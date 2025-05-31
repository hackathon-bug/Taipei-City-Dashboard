# FloodHotspotMap 快速入門指南

## 組件簡介

FloodHotspotMap 是一個用於顯示淹水熱點區域的地圖組件，基於 Mapbox GL 和 Deck.gl 開發。

## 快速使用方法

### 方法 1: 直接使用

1. 導入組件:
   ```javascript
   import FloodHotspotMap from './dashboardComponent/components/develop/FloodHotspotMap.vue';
   ```

2. 在模板中使用:
   ```html
   <div style="width: 100%; height: 500px;">
     <FloodHotspotMap />
   </div>
   ```

### 方法 2: 整合到儀表板系統

1. 在 DashboardComponent.vue 中註冊組件
2. 創建組件配置
3. 在儀表板中使用

## 示例代碼

請參考以下文件:
- 直接使用示例: `src/examples/FloodHotspotMapExample.vue`
- 完整使用文檔: `src/docs/FloodHotspotMapUsage.md`

## 主要功能

- 3D 地形顯示
- 淹水區域可視化
- 支持地圖導航控制
- 可自定義淹水區域