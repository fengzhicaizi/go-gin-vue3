import * as THREE from 'three';
import { ThreeX } from '../public/threex';
import { Model } from '../public/threex-model';
import { Pass } from '../public/threex-pass';

function start() {
	const dom = document.getElementById('full-screen');

	if (dom) {
		const threeX = new ThreeX(dom).bindOrbitControlsControl().addAmbientLight(new THREE.Color('#ffffff'), 1).addDirectionalLight().addAxesHelper().addGUI().addStats();

		const { GUI, orbitControl } = threeX;

		// orbitControl
		if (orbitControl) {
			orbitControl.maxPolarAngle = Math.PI / 2;
			orbitControl.minDistance = 3000; // 相机最近距离
			orbitControl.maxDistance = 20000; // 相机最远距离
		}

		// model
		const model = new Model(threeX);
		// model.createTest();
		model.createRing();
		model.createHNMap();

		// 后期处理
		const pass = new Pass(threeX);
		// pass.glitchPass();
		// pass.luminosityShader();
		// pass.unrealBloomPass();
		// pass.flimPass();
		console.log(pass);

		if (GUI) {
			// const cameraFolder = GUI.addFolder('相机');
			// cameraFolder.add(camera.position, 'x', 0, 20000).name('x');
			// cameraFolder.add(camera.position, 'y', -20000, 20000).name('y');
			// cameraFolder.add(camera.position, 'z', 0, 20000).name('z');
			// const unrealBloomPassFolder = GUI.addFolder('辉光');
			// unrealBloomPassFolder.add(unrealBloomPass, 'radius', 0, 1000);
			// unrealBloomPassFolder.add(unrealBloomPass, 'strength', 0, 10);
			// unrealBloomPassFolder.add(unrealBloomPass, 'threshold', 0, 1);
		}

		dom.addEventListener('fullscreenchange', function () {
			const dom = document.getElementById('full-screen');
			if (dom) {
				threeX.reRender(dom);
			}
		});
	}
}

export { start };
