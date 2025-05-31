<template>
  <div class="example-container">
    <h1>FloodHotspotMap 使用範例</h1>

    <!-- 直接使用 FloodHotspotMap 組件 -->
    <div class="map-wrapper">
      <FloodHotspotMap />
    </div>
  </div>
</template>

<script setup>
import FloodHotspotMap from '../dashboardComponent/components/develop/FloodHotspotMap.vue';

// 這個範例展示了如何直接在 Vue 組件中使用 FloodHotspotMap
</script>

<style scoped lang="scss">
.example-container {
  padding: 20px;

  h1 {
    margin-bottom: 20px;
  }

  .map-wrapper {
    width: 100%;
    height: 500px;
    border-radius: 16px;
    overflow: hidden;
  }
}
</style>
