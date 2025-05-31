# FloodHotspotMap 組件使用說明

FloodHotspotMap 是一個基於 Mapbox GL 和 Deck.gl 的地圖組件，用於顯示淹水熱點區域。本文檔提供了如何在專案中使用此組件的詳細說明。

## 目錄

1. [組件功能](#組件功能)
2. [直接使用](#直接使用)
3. [整合到儀表板系統](#整合到儀表板系統)
4. [自定義配置](#自定義配置)
5. [依賴項](#依賴項)

## 組件功能

FloodHotspotMap 組件提供以下功能：

- 顯示 3D 地形圖層
- 使用 Deck.gl 的 PolygonLayer 顯示淹水區域
- 支持地圖導航控制
- 可自定義淹水區域的位置和樣式

## 直接使用

### 步驟 1: 導入組件

在您的 Vue 文件中導入 FloodHotspotMap 組件：

```vue
<script setup>
import FloodHotspotMap from '../dashboardComponent/components/develop/FloodHotspotMap.vue';
</script>
```

### 步驟 2: 在模板中使用

在模板中使用組件，並確保提供足夠的高度和寬度：

```vue
<template>
  <div class="map-container">
    <FloodHotspotMap />
  </div>
</template>

<style scoped>
.map-container {
  width: 100%;
  height: 500px;
  border-radius: 16px;
  overflow: hidden;
}
</style>
```

### 完整示例

請參考 `src/examples/FloodHotspotMapExample.vue` 文件，其中包含了直接使用 FloodHotspotMap 的完整示例。

## 整合到儀表板系統

要將 FloodHotspotMap 整合到儀表板系統中，需要進行以下步驟：

### 步驟 1: 在 DashboardComponent.vue 中註冊組件

打開 `src/dashboardComponent/DashboardComponent.vue` 文件，並進行以下修改：

1. 導入 FloodHotspotMap 組件：

```javascript
import FloodHotspotMap from './components/develop/FloodHotspotMap.vue';
```

2. 在 `returnChartComponent` 函數中添加 FloodHotspotMap 的處理：

```javascript
case "FloodHotspotMap":
  return svg ? MapLegendSvg : FloodHotspotMap;
```

### 步驟 2: 創建組件配置

在使用儀表板系統時，需要為 FloodHotspotMap 創建配置對象。以下是一個示例配置：

```javascript
const floodMapConfig = {
  id: "flood-hotspot-map",
  name: "淹水熱點地圖",
  index: "flood-hotspot-map",
  city: "tpe",
  source: "台北市政府",
  short_desc: "顯示台北市淹水熱點區域",
  time_from: "static",
  update_freq: 1,
  update_freq_unit: "day",
  chart_config: {
    types: ["FloodHotspotMap"]
  },
  map_config: [
    {
      index: "flood-hotspot",
      type: "polygon",
      city: "tpe"
    }
  ],
  chart_data: {}
};
```

### 步驟 3: 在儀表板中使用

在儀表板視圖中使用配置：

```vue
<DashboardComponent
  :config="floodMapConfig"
  mode="map"
  :info-btn="true"
  :toggle-on="true"
  @toggle="(value, map_config) => handleToggle(value, map_config)"
/>
```

## 自定義配置

您可以通過修改 FloodHotspotMap.vue 文件來自定義地圖的配置：

### 地圖樣式

修改 `style` 屬性來更改地圖的基本樣式：

```javascript
style: "mapbox://styles/mapbox/dark-v10", // 可以改為 light-v10, streets-v11 等
```

### 初始位置和縮放級別

修改 `center` 和 `zoom` 屬性來設置初始位置和縮放級別：

```javascript
center: [121.55, 25.05], // [經度, 緯度]
zoom: 11,                // 縮放級別
```

### 淹水區域

修改 PolygonLayer 的 `data` 屬性來自定義淹水區域：

```javascript
data: [
  {
    polygon: [
      [121.54, 25.06],
      [121.56, 25.06],
      [121.56, 25.04],
      [121.54, 25.04]
    ]
  }
  // 可以添加更多多邊形
]
```

### 淹水區域樣式

修改 PolygonLayer 的樣式屬性來自定義淹水區域的外觀：

```javascript
getFillColor: [0, 120, 255, 180], // RGBA 顏色
stroked: true,                    // 是否顯示邊框
getLineColor: [0, 80, 200],       // 邊框顏色
lineWidthMinPixels: 2             // 邊框寬度
```

## 依賴項

FloodHotspotMap 組件依賴以下庫：

- mapbox-gl: 用於基本地圖顯示
- @deck.gl/mapbox: 用於 Mapbox GL 與 Deck.gl 的集成
- @deck.gl/layers: 提供 PolygonLayer 等圖層

這些依賴項已經在項目的 package.json 中定義，無需額外安裝。