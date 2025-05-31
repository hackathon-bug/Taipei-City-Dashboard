<script setup>
import mapboxGl from "mapbox-gl";
import {onMounted, ref} from "vue";
import axios from "axios";

const mapContainer = ref(null);
let map;

onMounted(async () => {
	try{
		const response = await axios({
			method: "GET",
			url: ``
		});
		const {data} = response;

		mapboxGl.accessToken = import.meta.env.VITE_MAPBOXTOKEN;
		map = new mapboxGl.Map({
			container: mapContainer.value,
			style: "mapbox://styles/mapbox/light-v11",
			center: [121.5654, 25.053],
			zoom: 8,
			pitch: 45, // 添加傾斜視角以便看到地形效果
		});

		map.on("load", () => {
			// 加入地形效果
			map.addSource('mapbox-dem', {
				type: 'raster-dem',
				url: 'mapbox://mapbox.terrain-rgb',
				tileSize: 512,
				maxzoom: 14
			});

			// 設定地形效果（含放大倍率）
			map.setTerrain({ source: 'mapbox-dem', exaggeration: 1.5 });

			// 加入天空圖層（強化立體感）
			map.addLayer({
				id: 'sky',
				type: 'sky',
				paint: {
					'sky-type': 'atmosphere',
					'sky-atmosphere-sun': [0.0, 0.0],
					'sky-atmosphere-sun-intensity': 15
				}
			});
		});

		data.forEach(eq => {
			new mapboxGl.Marker({color: "red"})
				.setLngLat([eq.lng, eq.lat])
				.setPopup(
					new mapboxGl.Popup().setHTML(
						`<strong>震度：${eq.magnitude}</strong><br>深度：${eq.depth}km<br>時間：${eq.time}`
					)
				).addTo(map);
		})
	}catch(error){
		console.error("Error fetching data:", error);
	}
});
</script>

<template>
  <div
    ref="mapContainer"
    class="mapbox"
  />
</template>

<style scoped lang="scss">
.mapbox {
	width: 100%;
	height: 100%;
}
</style>
