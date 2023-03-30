import * as Cesium from 'cesium';
import config from '@/conf';

export const startCesium = () => {
	(window as any).CESIUM_BASE_URL = '/';
	Cesium.Ion.defaultAccessToken = config.cesiumIon;
	const viewer = new Cesium.Viewer('full-screen', {
		terrainProvider: Cesium.createWorldTerrain(),
	});
	// Add Cesium OSM Buildings, a global 3D buildings layer.
	// const buildingTileset = viewer.scene.primitives.add(Cesium.createOsmBuildings());
	// Fly the camera to San Francisco at the given longitude, latitude, and height.
	viewer.camera.flyTo({
		destination: Cesium.Cartesian3.fromDegrees(-122.4175, 37.655, 400),
		orientation: {
			heading: Cesium.Math.toRadians(0.0),
			pitch: Cesium.Math.toRadians(-15.0),
		},
	});
};
