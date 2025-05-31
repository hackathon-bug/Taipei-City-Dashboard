<script setup>
import { onMounted, ref } from "vue";
import mapboxGl from "mapbox-gl";
import { MapboxOverlay } from "@deck.gl/mapbox";
import { PolygonLayer } from "@deck.gl/layers";

mapboxGl.accessToken = import.meta.env.VITE_MAPBOXTOKEN;

const mapContainer = ref(null);
let map = null;

onMounted(() => {
	if(!mapContainer.value) return;

	map = new mapboxGl.Map({
		container: mapContainer.value,
		style: "mapbox://styles/mapbox/dark-v10",
		center: [121.55, 25.05],
		zoom: 11,
		pitch: 60,
		bearing: -10,
	});

	map.on("load", () => {
		map.addControl(new mapboxGl.NavigationControl());

		// 加入 DEM 地形圖層（地勢圖）
		map.addSource('mapbox-dem', {
			type: 'raster-dem',
			url: 'mapbox://mapbox.terrain-rgb',
			tileSize: 512,
			maxzoom: 14,
		});
		map.setTerrain({ source: 'mapbox-dem', exaggeration: 1.5 });

		// 加入陰影層增加立體感
		map.addLayer({
			id: 'hillshading',
			source: 'mapbox-dem',
			type: 'hillshade'
		});

		// 加入天空圖層（可選，強化立體感）
		map.addLayer({
			id: 'sky',
			type: 'sky',
			paint: {
				'sky-type': 'atmosphere',
				'sky-atmosphere-sun': [0.0, 0.0],
				'sky-atmosphere-sun-intensity': 15
			}
		});

		// 加入模擬的淹水範圍 Deck.gl PolygonLayer
		const overlay = new MapboxOverlay({
			layers: [
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
						}
					],
					getPolygon: d => d.polygon,
					getFillColor: [0, 120, 255, 180],
					stroked: true,
					getLineColor: [0, 80, 200],
					lineWidthMinPixels: 2
				})
			]
		});

		map.addControl(overlay);
	});
});
</script>

<template>
	<div ref="mapContainer" class="map-container"></div>
</template>

<style scoped lang="scss">
@import 'mapbox-gl/dist/mapbox-gl.css';

.map-container {
	width: 100%;
	height: 100%;
	border-radius: 16px;
	box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
	overflow: hidden;
}
</style>
