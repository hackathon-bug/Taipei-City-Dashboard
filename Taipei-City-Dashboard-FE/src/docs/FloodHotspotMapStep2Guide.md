# FloodHotspotMap 整合到儀表板系統 - 步驟 2 之後的操作指南

本文檔針對 FloodHotspotMap 組件整合到儀表板系統的步驟 2（創建組件配置）之後的操作提供詳細說明。

## 步驟 2 之後該如何做

在完成步驟 2（創建組件配置）後，您需要執行以下操作來完成 FloodHotspotMap 組件的整合：

### 1. 將組件配置添加到儀表板

您有兩種方式可以將 FloodHotspotMap 組件添加到儀表板系統：

#### 方式 A: 添加到現有儀表板

```javascript
// 導入 contentStore
import { useContentStore } from "../store/contentStore";

// 獲取 store 實例
const contentStore = useContentStore();

// 獲取當前儀表板
const currentDashboard = contentStore.currentDashboard;

// 將組件配置添加到儀表板
currentDashboard.components.push(floodMapConfig);

// 更新儀表板
contentStore.editCurrentDashboard();
```

#### 方式 B: 創建新儀表板

```javascript
// 導入所需的 store
import { useContentStore } from "../store/contentStore";
import { useDialogStore } from "../store/dialogStore";

// 獲取 store 實例
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

### 2. 在視圖中使用組件

完成配置添加後，您可以在視圖中使用 DashboardComponent 來顯示 FloodHotspotMap 組件：

```vue
<template>
  <DashboardComponent
    :config="floodMapConfig"
    mode="map"
    :info-btn="true"
    :toggle-on="true"
    @toggle="(value, map_config) => handleToggle(value, map_config)"
  />
</template>

<script setup>
import DashboardComponent from "../dashboardComponent/DashboardComponent.vue";
import { useMapStore } from "../store/mapStore";

const mapStore = useMapStore();

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

### 3. 測試組件整合

完成上述步驟後，您應該測試組件是否正確整合：

1. 確認地圖是否正確顯示
2. 測試切換功能是否正常工作
3. 檢查淹水區域是否正確渲染

如果遇到問題，請檢查：
- 配置對象是否正確
- mapStore 是否正確處理地圖圖層
- 控制台是否有錯誤信息

### 4. 自定義組件（可選）

如果需要，您可以自定義 FloodHotspotMap 組件的行為：

#### 修改淹水區域數據

您可以修改 FloodHotspotMap.vue 文件中的 PolygonLayer 數據來自定義淹水區域：

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
    // 添加更多淹水區域
  ],
  getPolygon: d => d.polygon,
  getFillColor: [0, 120, 255, 180],
  stroked: true,
  getLineColor: [0, 80, 200],
  lineWidthMinPixels: 2
})
```

#### 連接實時數據

如果您有實時淹水數據，可以修改 FloodHotspotMap.vue 文件以從 API 獲取數據：

```javascript
import axios from "axios";

// 在 onMounted 中添加
try {
  // 從 API 獲取淹水數據
  const response = await axios.get('/api/flood-data');
  const floodData = response.data.map(item => ({
    polygon: item.coordinates
  }));
  
  // 使用獲取的數據創建 PolygonLayer
  const overlay = new MapboxOverlay({
    layers: [
      new PolygonLayer({
        id: 'flood-polygon',
        data: floodData,
        // 其他配置...
      })
    ]
  });
  
  map.addControl(overlay);
} catch (error) {
  console.error("Error fetching flood data:", error);
}
```

## 完整流程總結

1. 創建組件配置對象（步驟 2）
2. 將配置添加到現有儀表板或創建新儀表板
3. 在視圖中使用 DashboardComponent 顯示組件
4. 測試組件整合
5. 根據需要自定義組件

按照這些步驟，您應該能夠成功將 FloodHotspotMap 組件整合到儀表板系統中。