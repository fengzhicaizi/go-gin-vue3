/**
 * 初始化three
 */
import * as THREE from 'three';
import { OrbitControls } from 'three/examples/jsm/controls/OrbitControls';
import { EffectComposer } from 'three/examples/jsm/postprocessing/EffectComposer';
import { RenderPass } from 'three/examples/jsm/postprocessing/RenderPass';
import Stats from 'three/examples/jsm/libs/stats.module';
import { GUI } from 'dat.gui';

type animationCBType = (opt: { scene: THREE.Scene; camera: THREE.Camera }) => void;

class ThreeX {
	static len = 10000; // 默认长度
	// dom
	public element: HTMLElement;
	public clientWidth: number;
	public clientHeight: number;
	// scene
	public scene: THREE.Scene;
	// camera
	public camera: THREE.Camera;
	// render
	public renderer: THREE.WebGLRenderer;
	// clock
	public clock: THREE.Clock;
	// animationMixer
	public animationMixer: THREE.AnimationMixer;
	// control
	public orbitControl: OrbitControls | undefined;
	// light
	public ambientLight: THREE.AmbientLight | undefined;
	public directionalLight: THREE.DirectionalLight | undefined;
	public pointLight: THREE.PointLight | undefined;
	// helper
	public axeshelper: THREE.AxesHelper | undefined;
	public gridHelper: THREE.GridHelper | undefined;
	public directionalLightHelper: THREE.DirectionalLightHelper | undefined;
	public pointLightHelper: THREE.PointLightHelper | undefined;
	// GUI
	public GUI: GUI | undefined;
	// animate callback
	private animationCB: animationCBType | undefined;
	// effectComposer
	public effectComposer: EffectComposer;
	public stats: Stats | undefined;

	/**
	 * 初始化构造函数
	 * @param element HTMLElement
	 */
	constructor(element: HTMLElement) {
		this.element = element;
		this.clientWidth = element.clientWidth;
		this.clientHeight = element.clientHeight;
		this.scene = this.createScene();
		this.clock = this.createClock();
		this.camera = this.createPerspectiveCamera();
		this.renderer = this.createRenderer();
		this.animationMixer = this.createAnimationMixer();
		this.effectComposer = this.createEffectComposer();
		this.startRender();
	}
	/**
	 * 创建场景
	 */
	private createScene() {
		return new THREE.Scene();
	}
	/**
	 * 创建时钟对象
	 */
	private createClock() {
		return new THREE.Clock();
	}
	/**
	 * 创建相机
	 */
	private createPerspectiveCamera() {
		const fov = 90;
		const aspect: number = this.clientWidth / this.clientHeight;
		const near = 0.1;
		const far: number = ThreeX.len * 10;
		const camera = new THREE.PerspectiveCamera(fov, aspect, near, far);
		camera.position.set(0, ThreeX.len, ThreeX.len);
		camera.lookAt(0, 0, 0);
		return camera;
	}
	/**
	 * 创建渲染器
	 */
	private createRenderer() {
		const renderer = new THREE.WebGLRenderer({ antialias: true, alpha: false });
		renderer.setSize(this.clientWidth, this.clientHeight);
		// this.renderer.setPixelRatio(window.devicePixelRatio);
		this.element.appendChild(renderer.domElement);
		return renderer;
	}
	/**
	 * 创建AnimationMixer
	 */
	private createAnimationMixer() {
		return new THREE.AnimationMixer(this.scene);
	}
	/**
	 * 添加后期处理对象
	 */
	private createEffectComposer() {
		const renderPass = new RenderPass(this.scene, this.camera);
		const effectComposer = new EffectComposer(this.renderer);
		effectComposer.addPass(renderPass);
		return effectComposer;
	}
	/**
	 * 渲染
	 */
	public startRender() {
		this.renderer.setAnimationLoop(() => {
			// 两种动画形式 集中式
			if (this.animationCB) {
				this.animationCB({
					scene: this.scene,
					camera: this.camera,
				});
			}
			// 分散式
			this.animationMixer.update(this.clock.getDelta());

			// stats更新
			if (this.stats) this.stats.update();
			this.effectComposer.render();
		});
	}
	/**
	 * 停止渲染
	 */
	public stopRender() {
		this.renderer.setAnimationLoop(null);
	}
	/**
	 * 绑定轨道控制器
	 * @returns this
	 */
	public bindOrbitControlsControl(): ThreeX {
		this.orbitControl = new OrbitControls(this.camera, this.renderer.domElement);
		return this;
	}
	/**
	 * 添加环境光
	 * @param color 默认 #ffffff
	 * @param intensity 默认 0.3
	 * @returns this
	 */
	public addAmbientLight(color: THREE.ColorRepresentation = new THREE.Color('#ffffff'), intensity = 0.3): ThreeX {
		this.ambientLight = new THREE.AmbientLight(color, intensity);
		this.scene.add(this.ambientLight);
		return this;
	}
	/**
	 * 添加平行光
	 * @param opt
	 * color 默认 #ffffff
	 * intensity 默认 1
	 * needHelper 默认 false
	 * @returns
	 */
	public addDirectionalLight(opt?: { color?: THREE.ColorRepresentation; intensity?: number; needHelper?: boolean }): ThreeX {
		const defaultOpt = {
			color: new THREE.Color('#ffffff'),
			intensity: 1,
			needHelper: false,
		};
		const options = Object.assign(defaultOpt, opt);
		this.directionalLight = new THREE.DirectionalLight(options.color, options.intensity);
		this.directionalLight.position.set(500, 500, 500);
		if (options.needHelper) {
			this.addDirectionalLightHelper(this.directionalLight);
		} else {
			this.scene.add(this.directionalLight);
		}
		return this;
	}
	/**
	 * 添加点光源
	 * @param opt
	 * color = #ffffff
	 * intensity = 1
	 * needHelper = false
	 * @returns
	 */
	public addPointLight(opt?: { color?: THREE.ColorRepresentation; intensity?: number; decay?: number; needHelper?: boolean }): ThreeX {
		const defaultOpt = {
			color: new THREE.Color('#ffffff'),
			intensity: 1,
			decay: 100,
			needHelper: false,
		};
		const options = Object.assign(defaultOpt, opt);
		this.pointLight = new THREE.PointLight(options.color, options.intensity);
		this.pointLight.position.set(100, 100, 100);
		if (options.needHelper) {
			this.addPointLightHelper(this.pointLight);
		} else {
			this.scene.add(this.pointLight);
		}
		return this;
	}
	/**
	 * 添加AxesHelper
	 * @param size
	 * @returns
	 */
	public addAxesHelper(size: number = ThreeX.len): ThreeX {
		this.axeshelper = new THREE.AxesHelper(size);
		this.scene.add(this.axeshelper);
		return this;
	}
	/**
	 * 添加gridHelper
	 * @param size
	 * @param divisions
	 * @returns
	 */
	public addGridHelper(size: number = ThreeX.len, divisions: number = ThreeX.len): ThreeX {
		this.gridHelper = new THREE.GridHelper(size, divisions);
		this.scene.add(this.gridHelper);
		return this;
	}
	/**
	 * 添加平行光helper
	 * @param light
	 * @returns
	 */
	public addDirectionalLightHelper(light: THREE.DirectionalLight): ThreeX {
		this.directionalLightHelper = new THREE.DirectionalLightHelper(light, 10);
		this.scene.add(this.directionalLightHelper);
		return this;
	}
	/**
	 * 添加点光源helper
	 * @param light
	 * @returns
	 */
	public addPointLightHelper(light: THREE.PointLight): ThreeX {
		this.pointLightHelper = new THREE.PointLightHelper(light, 10, new THREE.Color('#ffffff'));
		this.scene.add(this.pointLightHelper);
		return this;
	}
	/**
	 *
	 */
	public addStats(): ThreeX {
		this.stats = new Stats();
		this.stats.dom.style.position = 'absolute';
		this.stats.dom.style.left = '24px';
		this.stats.dom.style.top = '24px';
		this.element.appendChild(this.stats.dom);
		return this;
	}
	/**
	 * 添加gui
	 */
	public addGUI(): ThreeX {
		this.GUI = new GUI({ autoPlace: false });
		const style = this.GUI.domElement.style;
		style.position = 'absolute';
		style.top = '24px';
		style.right = '24px';
		this.element.appendChild(this.GUI.domElement);
		return this;
	}
	/**
	 * 添加动画
	 * @param callback 随requestAnimationFrame函数被调用
	 */
	public addAnimation(callback: animationCBType) {
		this.animationCB = callback;
	}
	/**
	 * 监听事件，重新渲染
	 */
	public reRender(dom: HTMLElement) {
		this.clientWidth = dom.clientWidth;
		this.clientHeight = dom.clientHeight;
		this.renderer.setSize(this.clientWidth, this.clientHeight);
	}
}

export { ThreeX };
