# FloodHotspotMap 整合到儀表板系統詳細指南

本指南提供了將 FloodHotspotMap 組件整合到儀表板系統的詳細步驟和示例代碼。

## 目錄

1. [前置準備](#前置準備)
2. [創建組件配置](#創建組件配置)
3. [在視圖中使用組件](#在視圖中使用組件)
4. [測試和故障排除](#測試和故障排除)
5. [進階配置](#進階配置)

## 前置準備

確保 FloodHotspotMap 組件已經在 DashboardComponent.vue 中正確註冊。這一步已經完成，您可以在 DashboardComponent.vue 中看到以下代碼：

```javascript
// 導入組件
import FloodHotspotMap from './components/develop/FloodHotspotMap.vue';

// 在 returnChartComponent 函數中註冊
case "FloodHotspotMap":
  return svg ? MapLegendSvg : FloodHotspotMap;
```

## 創建組件配置

### 步驟 1: 創建基本配置對象

首先，創建一個包含必要屬性的配置對象：

```javascript
const floodMapConfig = {
  id: "flood-hotspot-map",           // 唯一標識符
  name: "淹水熱點地圖",               // 顯示名稱
  index: "flood-hotspot-map",        // 索引名稱（用於API請求）
  city: "tpe",                       // 城市代碼
  source: "台北市政府",               // 數據來源
  short_desc: "顯示台北市淹水熱點區域", // 簡短描述
  time_from: "static",               // 數據時間類型（static表示靜態數據）
  update_freq: 1,                    // 更新頻率
  update_freq_unit: "day",           // 更新頻率單位
  chart_config: {                    // 圖表配置
    types: ["FloodHotspotMap"]       // 指定使用FloodHotspotMap組件
  },
  map_config: [                      // 地圖配置
    {
      index: "flood-hotspot",        // 地圖圖層索引
      type: "polygon",               // 圖層類型
      city: "tpe"                    // 城市代碼
    }
  ],
  chart_data: {}                     // 圖表數據（可以為空）
};
```

### 步驟 2: 在 Store 中使用配置

您可以通過以下方式將配置添加到 contentStore：

#### 方法 1: 添加到現有儀表板

如果您想將組件添加到現有儀表板，可以使用以下代碼：

```javascript
import { useContentStore } from "../store/contentStore";

const contentStore = useContentStore();

// 獲取當前儀表板
const currentDashboard = contentStore.currentDashboard;

// 將組件配置添加到儀表板
currentDashboard.components.push(floodMapConfig);

// 更新儀表板
contentStore.editCurrentDashboard();
```

#### 方法 2: 創建新儀表板

如果您想創建一個包含此組件的新儀表板，可以使用以下代碼：

```javascript
import { useContentStore } from "../store/contentStore";
import { useDialogStore } from "../store/dialogStore";

const contentStore = useContentStore();
const dialogStore = useDialogStore();

// 清除編輯儀表板
contentStore.clearEditDashboard();

// 設置儀表板屬性
contentStore.editDashboard.name = "淹水監測儀表板";
contentStore.editDashboard.icon = "water_drop";

// 添加組件
contentStore.editDashboard.components.push(floodMapConfig);

// 創建儀表板
contentStore.createDashboard();
```

## 在視圖中使用組件

### 在 MapView 中使用

在 MapView.vue 中，您可以直接使用 DashboardComponent 組件來顯示 FloodHotspotMap：

```vue
<template>
  <div class="map">
    <div class="map-charts">
      <DashboardComponent
        :config="floodMapConfig"
        mode="map"
        :info-btn="true"
        :toggle-on="true"
        @toggle="(value, map_config) => handleToggle(value, map_config)"
      />
    </div>
    <MapContainer />
  </div>
</template>

<script setup>
import { ref } from "vue";
import DashboardComponent from "../dashboardComponent/DashboardComponent.vue";
import MapContainer from "../components/map/MapContainer.vue";
import { useMapStore } from "../store/mapStore";

const mapStore = useMapStore();

// 創建配置對象
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

// 處理組件切換
function handleToggle(value, map_config) {
  if (!map_config[0]) {
    return;
  }
  if (value) {
    mapStore.addToMapLayerList(map_config);
  } else {
    mapStore.clearByParamFilter(map_config);
    mapStore.turnOffMapLayerVisibility(map_config);
  }
}
</script>
```

### 在自定義視圖中使用

您也可以在自定義視圖中使用 FloodHotspotMap 組件：

```vue
<template>
  <div class="custom-view">
    <h1>淹水監測系統</h1>
    <div class="dashboard-container">
      <DashboardComponent
        :config="floodMapConfig"
        mode="large"
        :info-btn="true"
      />
    </div>
  </div>
</template>

<script setup>
import DashboardComponent from "../dashboardComponent/DashboardComponent.vue";

// 創建配置對象
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
</script>

<style scoped>
.custom-view {
  padding: 20px;
}

.dashboard-container {
  margin-top: 20px;
  height: 500px;
}
</style>
```

## 測試和故障排除

### 測試整合

1. 確保 FloodHotspotMap 組件在 DashboardComponent.vue 中正確註冊
2. 創建配置對象並添加到儀表板
3. 在視圖中使用 DashboardComponent 顯示組件
4. 檢查地圖是否正確顯示
5. 測試地圖交互功能

### 常見問題和解決方案

1. **地圖不顯示**
   - 檢查 Mapbox Token 是否正確設置
   - 確保 chart_config.types 包含 "FloodHotspotMap"
   - 檢查控制台是否有錯誤

2. **地圖顯示但沒有淹水區域**
   - 檢查 PolygonLayer 的 data 屬性是否正確設置
   - 確保坐標值在正確的範圍內

3. **組件無法切換**
   - 確保 map_config 正確設置
   - 檢查 handleToggle 函數是否正確實現

## 進階配置

### 自定義淹水區域數據

您可以通過修改 FloodHotspotMap.vue 文件中的 PolygonLayer 數據來自定義淹水區域：

```javascript
new PolygonLayer({
  id: 'flood-polygon',
  data: [
    {
      polygon: [
        [121.54, 25.06],
        [121.56, 25.06],
        [121.56, 25.04],
        [121.54, 25.04]
      ]
    },
    {
      polygon: [
        [121.51, 25.03],
        [121.53, 25.03],
        [121.53, 25.01],
        [121.51, 25.01]
      ]
    }
  ],
  getPolygon: d => d.polygon,
  getFillColor: [0, 120, 255, 180],
  stroked: true,
  getLineColor: [0, 80, 200],
  lineWidthMinPixels: 2
})
```

### 連接實時數據

如果您有實時淹水數據，可以修改 FloodHotspotMap.vue 文件以從 API 獲取數據：

```javascript
import { onMounted, ref } from "vue";
import mapboxGl from "mapbox-gl";
import { MapboxOverlay } from "@deck.gl/mapbox";
import { PolygonLayer } from "@deck.gl/layers";
import axios from "axios";

const floodData = ref([]);

onMounted(async () => {
  try {
    // 從 API 獲取淹水數據
    const response = await axios.get('/api/flood-data');
    floodData.value = response.data.map(item => ({
      polygon: item.coordinates
    }));
    
    // 初始化地圖
    // ...
    
    // 使用獲取的數據創建 PolygonLayer
    const overlay = new MapboxOverlay({
      layers: [
        new PolygonLayer({
          id: 'flood-polygon',
          data: floodData.value,
          getPolygon: d => d.polygon,
          getFillColor: [0, 120, 255, 180],
          stroked: true,
          getLineColor: [0, 80, 200],
          lineWidthMinPixels: 2
        })
      ]
    });
    
    map.addControl(overlay);
  } catch (error) {
    console.error("Error fetching flood data:", error);
  }
});
```

這個詳細指南應該能幫助您完成 FloodHotspotMap 組件的整合過程。如果您有任何問題，請參考項目文檔或聯繫開發團隊。