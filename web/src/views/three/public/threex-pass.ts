/**
 * 后期处理通道
 */
import * as THREE from 'three';
import { ShaderPass } from 'three/examples/jsm/postprocessing/ShaderPass';
import { GlitchPass } from 'three/examples/jsm/postprocessing/GlitchPass';
import { OutlinePass } from 'three/examples/jsm/postprocessing/OutlinePass';
import { BloomPass } from 'three/examples/jsm/postprocessing/BloomPass';
import { FilmPass } from 'three/examples/jsm/postprocessing/FilmPass';
import { LuminosityShader } from 'three/examples/jsm/shaders/LuminosityShader';
import { TriangleBlurShader } from 'three/examples/jsm/shaders/TriangleBlurShader';
import { UnrealBloomPass } from 'three/examples/jsm/postprocessing/UnrealBloomPass';
import { ThreeX } from './threex';

class Pass {
	threeX: ThreeX;
	/**
	 * 构造函数
	 * @param  threeX ThreeX
	 */
	constructor(threeX: ThreeX) {
		this.threeX = threeX;
	}

	// pass
	/**
	 * 电脉冲故障风效果
	 */
	public glitchPass() {
		const pass = new GlitchPass();
		pass.renderToScreen = true;
		this.threeX.effectComposer.addPass(pass);
	}
	/**
	 * 描边效果
	 * 勾勒出场景中的物体轮廓
	 */
	public outlinePass(scene: THREE.Scene, camera: THREE.Camera, selectedObjects: THREE.Object3D[]) {
		const pass = new OutlinePass(new THREE.Vector2(10, 10), scene, camera);
		pass.renderToScreen = true; // 渲染至缓存
		pass.visibleEdgeColor = new THREE.Color(76, 148, 156); //可见边缘的颜色
		pass.hiddenEdgeColor = new THREE.Color(0, 0, 0); // 不可见边缘的颜色
		pass.edgeGlow = 1.0; // 发光强度
		pass.usePatternTexture = false; // 是否使用纹理
		pass.edgeThickness = 2.0; // 边缘浓度
		pass.edgeStrength = 2.0; // 边缘强度
		pass.pulsePeriod = 0; // 闪烁频率，值越大频率越低

		pass.selectedObjects = selectedObjects; // 一般生产环境中会配合鼠标事件，来改变selectedObjects，使选中的物体描边

		this.threeX.effectComposer.addPass(pass);
	}
	/**
	 * 辉光
	 * 该通道通过增强场景中明亮的区域来模拟真实世界中的摄像机
	 */
	public unrealBloomPass(): UnrealBloomPass {
		const pass = new UnrealBloomPass(new THREE.Vector2(this.threeX.clientWidth, this.threeX.clientHeight), 2, 0, 0.2);
		pass.renderToScreen = true;

		this.threeX.effectComposer.addPass(pass);

		return pass;
	}
	/**
	 * 泛光
	 * 该通道通过增强场景中明亮的区域来模拟真实世界中的摄像机
	 */
	public bloomPass() {
		const opt = {
			strength: 2.5,
			kernelSize: 25,
			sigma: 0.1,
		};

		const pass = new BloomPass(opt.strength, opt.kernelSize, opt.sigma);
		pass.renderToScreen = true;

		this.threeX.effectComposer.addPass(pass);
	}
	/**
	 * 通过扫描线和失真来模拟电视屏幕效果
	 */
	public flimPass() {
		const opt = {
			noiseIntensity: 0.35,
			scanlinesIntensity: 0.025,
			scanlinesCount: 648,
			grayscale: 1,
		};
		const pass = new FilmPass(opt.noiseIntensity, opt.scanlinesCount, opt.scanlinesCount, opt.grayscale);

		this.threeX.effectComposer.addPass(pass);
	}
	/**
	 * 着色器
	 */
	public luminosityShader() {
		const shader = new ShaderPass(LuminosityShader);
		this.threeX.effectComposer.addPass(shader);
	}
	/**
	 *
	 */
	public triangleBlurShader() {
		const shader = new ShaderPass(TriangleBlurShader);
		this.threeX.effectComposer.addPass(shader);
	}
}

export { Pass };
