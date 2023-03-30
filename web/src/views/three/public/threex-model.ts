/**
 * 模型
 */
import * as THREE from 'three';
import { ThreeX } from './threex';
import { LineGeometry } from 'three/examples/jsm/lines/LineGeometry';
import { LineMaterial } from 'three/examples/jsm/lines/LineMaterial';
import { Line2 } from 'three/examples/jsm/lines/Line2';
import { geoMercator } from 'd3-geo';

class Model {
	private threeX: ThreeX;
	constructor(threeX: ThreeX) {
		this.threeX = threeX;
	}
	/**
	 * 光环底座
	 */
	public createRing() {
		const group = new THREE.Group();
		group.name = 'ring';
		const animationDuretion = 20;
		const launchPoints: THREE.Vector3[] = [];

		// innerRole
		{
			const innerRoleGroup = new THREE.Group();
			innerRoleGroup.name = 'innerRole';
			const radius = ThreeX.len;

			// 画点
			{
				const circleGroup = new THREE.Group();
				const geomerty = new THREE.CircleGeometry(radius / 200, 100);
				const material = new THREE.MeshBasicMaterial({
					color: new THREE.Color('#00d3c3'),
					side: THREE.DoubleSide,
				});
				const circle = new THREE.Mesh(geomerty, material);

				const curve = new THREE.EllipseCurve(0, 0, radius, radius, 0, 2 * Math.PI, false, 0);
				curve.getPoints(100).map((point: THREE.Vector2) => {
					const mesh = circle.clone();
					mesh.position.set(point['x'], 0, point['y']);
					mesh.rotateX(THREE.MathUtils.radToDeg(90));
					circleGroup.add(mesh);
				});

				innerRoleGroup.add(circleGroup);
			}

			// 画环
			{
				const ringGroup = new THREE.Group();
				const num = 3;
				const ringRaius = 200;

				const material = new THREE.MeshBasicMaterial({
					color: new THREE.Color('#00d3c3'),
					side: THREE.DoubleSide,
					transparent: true,
					opacity: 0.2,
				});

				new Array(num).fill({}).map((_, index: number) => {
					const group = new THREE.Group();
					const thetaStartDeg = THREE.MathUtils.degToRad((360 / num) * index);
					const thetaLengthDeg = THREE.MathUtils.degToRad(30);

					{
						const geomerty = new THREE.RingGeometry(radius - ringRaius, radius + ringRaius, 1000, 8, thetaStartDeg, thetaLengthDeg);
						const ring = new THREE.Mesh(geomerty, material);
						ring.rotateX(THREE.MathUtils.degToRad(90));
						group.add(ring);
					}

					new Array(2).fill({}).map((_, index1) => {
						const geomerty = new THREE.CircleGeometry(ringRaius, 32, THREE.MathUtils.degToRad(180 + index1 * 210 + (360 / num) * index), THREE.MathUtils.degToRad(180));
						const circle = new THREE.Mesh(geomerty, material);
						circle.rotateX(THREE.MathUtils.degToRad(90));
						circle.position.set(radius * Math.cos(thetaStartDeg + thetaLengthDeg * index1), 0, radius * Math.sin(thetaStartDeg + thetaLengthDeg * index1));
						group.add(circle);
					});
					ringGroup.add(group);
				});

				innerRoleGroup.add(ringGroup);
			}

			group.add(innerRoleGroup);

			// 动画
			{
				const rotateTime = [0, animationDuretion / 2, animationDuretion];
				const rotateValues = [0, -Math.PI, -2 * Math.PI];
				const rotateFrameTrack = new THREE.NumberKeyframeTrack('ring/innerRole.rotation[y]', rotateTime, rotateValues);
				const animationClip = new THREE.AnimationClip('rotateAnimation', animationDuretion, [rotateFrameTrack]);
				// 异步
				requestAnimationFrame(() => {
					const animationAction = this.threeX.animationMixer.clipAction(animationClip);
					animationAction.play();
				});
			}
		}

		// outRole
		{
			const outRoleGroup = new THREE.Group();
			outRoleGroup.name = 'outRole';
			const radius = 1.1 * ThreeX.len;

			// 画小圆点
			{
				const circleGroup = new THREE.Group();
				const geomerty = new THREE.CircleGeometry(ThreeX.len / 200, 100);
				const material = new THREE.MeshBasicMaterial({
					color: new THREE.Color('#00d3c3'),
					side: THREE.DoubleSide,
				});
				const circle = new THREE.Mesh(geomerty, material);

				const curve = new THREE.EllipseCurve(0, 0, radius, radius, 0, 2 * Math.PI, false, 0);
				curve.getPoints(50).map((point: THREE.Vector2) => {
					const mesh = circle.clone();
					mesh.position.set(point['x'], 0, point['y']);
					mesh.rotateX(THREE.MathUtils.radToDeg(90));
					circleGroup.add(mesh);
				});

				outRoleGroup.add(circleGroup);
			}

			// 画大圆点
			{
				const planeGroup = new THREE.Group();
				const geomerty = new THREE.PlaneGeometry(ThreeX.len / 50, ThreeX.len / 100);
				const material = new THREE.MeshBasicMaterial({
					color: new THREE.Color('#00d3c3'),
					side: THREE.DoubleSide,
				});
				const plane = new THREE.Mesh(geomerty, material);

				const curve = new THREE.EllipseCurve(0, 0, radius, radius, 0, 2 * Math.PI, false, 0);
				curve.getPoints(10).map((point: THREE.Vector2) => {
					const mesh = plane.clone();
					mesh.position.set(point['x'], 0, point['y']);
					mesh.lookAt(new THREE.Vector3(0, 0, 0));
					mesh.rotateX(THREE.MathUtils.degToRad(90));
					mesh.rotateZ(THREE.MathUtils.degToRad(90));
					planeGroup.add(mesh);
				});

				outRoleGroup.add(planeGroup);
			}

			// 画环
			{
				const geomerty = new THREE.RingGeometry(radius - 30, radius + 30, 1000);
				const material = new THREE.MeshBasicMaterial({
					color: new THREE.Color('#00d3c3'),
					side: THREE.DoubleSide,
					transparent: true,
					opacity: 0.1,
				});
				const ring = new THREE.Mesh(geomerty, material);
				ring.rotateX(THREE.MathUtils.degToRad(90));
				outRoleGroup.add(ring);
			}

			group.add(outRoleGroup);

			// 动画
			{
				const rotateTime = [0, animationDuretion / 2, animationDuretion];
				const rotateValues = [0, Math.PI, 2 * Math.PI];
				const rotateFrameTrack = new THREE.NumberKeyframeTrack('ring/outRole.rotation[y]', rotateTime, rotateValues);
				const animationClip = new THREE.AnimationClip('rotateAnimation', animationDuretion, [rotateFrameTrack]);
				requestAnimationFrame(() => {
					const animationAction = this.threeX.animationMixer.clipAction(animationClip);
					animationAction.play();
				});
			}
		}

		// background
		{
			const radius = ThreeX.len;
			const bgGroup = new THREE.Group();
			bgGroup.name = 'background';

			{
				const plane = new THREE.PlaneGeometry(2 * ThreeX.len, 2 * ThreeX.len, 20, 20);
				plane.rotateX(THREE.MathUtils.degToRad(90));
				{
					const geomerty = new THREE.CircleGeometry(radius / 200, 10);
					const material = new THREE.MeshBasicMaterial({
						color: new THREE.Color('#00d3c3'),
						transparent: true,
						opacity: 1,
						side: THREE.DoubleSide,
					});
					const mesh = new THREE.Mesh(geomerty, material);
					const position = plane.attributes.position as THREE.BufferAttribute;
					for (let i = 0; i < position.count; i++) {
						const m = mesh.clone();
						const copyMaterial = material.clone();
						m.material = copyMaterial;
						const x = position.getX(i);
						const y = position.getY(i);
						const z = position.getZ(i);
						m.position.set(x, y, z);
						m.rotateX(THREE.MathUtils.degToRad(90));
						// 点到中心点的距离
						const distence = new THREE.Vector3(x, y, z).distanceTo(new THREE.Vector3(0, 0, 0));
						// 距离中心点越近亮度越高
						if (distence > ThreeX.len) {
							copyMaterial.opacity = 0;
						} else {
							const opacity = ((ThreeX.len - distence) * 0.5) / ThreeX.len;
							copyMaterial.opacity = opacity;
							// 收集可发射的点, 供彩带发射用,条件（len>point>len/2）
							if (distence > ThreeX.len / 2) {
								launchPoints.push(new THREE.Vector3(x, -2000, z));
							}
						}
						bgGroup.add(m);
					}
				}
			}
			group.add(bgGroup);
		}

		// ribbon
		{
			const ribbonGroup = new THREE.Group();
			ribbonGroup.name = 'ribbon';

			{
				const mesh = this.createNumMesh();

				if (mesh) {
					mesh.rotateX(THREE.MathUtils.degToRad(90));
					mesh.rotateZ(THREE.MathUtils.degToRad(270));

					new Array(10).fill({}).map((_, index) => {
						setTimeout(() => {
							const randomIndex = Math.floor(Math.random() * launchPoints.length);
							const point = launchPoints[randomIndex];
							const cloneMesh = mesh.clone();
							const name = `ribbon-${index}`;
							cloneMesh.name = name;
							cloneMesh.position.set(point.x, point.y, point.z);
							ribbonGroup.add(cloneMesh);

							// 动画
							const times = [0, 20];
							const values = [point.x, -2000, point.z, point.x, 10000, point.z];

							const keyFrameTrack = new THREE.VectorKeyframeTrack(`ring/ribbon/${name}.position`, times, values);
							const animationClip = new THREE.AnimationClip('ribbonAnimation', 20, [keyFrameTrack]);

							requestAnimationFrame(() => {
								const animationAction = this.threeX.animationMixer.clipAction(animationClip);
								animationAction.play();
							});
						}, index * 2000);
					});
				}
			}

			group.add(ribbonGroup);
		}

		this.threeX.scene.add(group);
	}
	/**
	 * 创建湖南3D地图
	 */
	public createHNMap() {
		const loader = new THREE.FileLoader();
		const hnGroup = new THREE.Group();
		hnGroup.name = '湖南地图';
		this.threeX.scene.add(hnGroup);

		// 边界
		loader.load('/geojson/hunan/430000.json', json => {
			const data = JSON.parse(<string>json);
			const { features } = data;
			const group = new THREE.Group();
			group.name = '边界';

			const shapeGroup = new THREE.Group();
			shapeGroup.name = '区域块';
			shapeGroup.rotateX(THREE.MathUtils.degToRad(90));
			shapeGroup.position.setY(1100);

			const lineGroup = new THREE.Group();
			lineGroup.position.setY(1105);
			lineGroup.name = '上层线';

			const line2Group = new THREE.Group();
			line2Group.position.setY(95);
			line2Group.scale.set(1.001, 1.001, 1.001);
			line2Group.name = '下层线';

			const cityGroup = new THREE.Group();
			cityGroup.name = '城市';
			cityGroup.position.setY(1150);

			group.add(shapeGroup, lineGroup, line2Group, cityGroup);
			hnGroup.add(group);

			if (features && Array.isArray(features)) {
				const materialExtrude = new THREE.MeshPhongMaterial({
					color: new THREE.Color('#003f81'),
					transparent: true,
					opacity: 1,
				});

				const line2Material = new LineMaterial({
					color: 0x49bef1,
					linewidth: 2,
					transparent: true,
					resolution: new THREE.Vector2(this.threeX.clientWidth, this.threeX.clientHeight),
				});

				data.features.map((feature: any) => {
					const { geometry, properties } = feature;

					// 画模型
					if (geometry && geometry.coordinates && Array.isArray(geometry.coordinates) && geometry.coordinates.length > 0) {
						geometry.coordinates.map((f: any) => {
							if (f && Array.isArray(f)) {
								f.map(f1 => {
									const shape = new THREE.Shape();
									const lines: number[] = [];

									if (f1 && Array.isArray(f1)) {
										f1.map((point, index) => {
											// 经纬度转世界坐标系

											const p = this.geoToWorld(point);
											if (index === 0) {
												shape.moveTo(p[0], p[1]);
											}

											// 抽稀
											shape.lineTo(p[0], p[1]);
											lines.push(p[0], 0, p[1]);
										});
									}

									// 画extrude
									{
										const geomerty = new THREE.ExtrudeGeometry(shape, {
											depth: 1000,
											bevelEnabled: false,
										});

										const mesh = new THREE.Mesh(geomerty, materialExtrude);
										shapeGroup.add(mesh);
									}

									// 画line
									{
										const bufferGeometry = new THREE.BufferGeometry();
										bufferGeometry.setAttribute('position', new THREE.BufferAttribute(new Float32Array(lines), 3));
										const line = new THREE.Line(bufferGeometry, new THREE.LineBasicMaterial());
										const line2geomerty = new LineGeometry();

										line2geomerty.fromLine(line);
										const line2 = new Line2(line2geomerty, line2Material);
										line2.computeLineDistances();
										lineGroup.add(line2);

										const line3Material = line2Material.clone();
										line3Material.linewidth = 4;
										const line3 = new Line2(line2geomerty, line3Material);
										line2Group.add(line3);
									}
								});
							}
						});
					}

					// 画城市
					if (properties) {
						const { fullname, center } = properties;
						const point = this.geoToWorld(center);
						const cGroup = new THREE.Group();
						cGroup.name = fullname;
						cityGroup.add(cGroup);

						// 城市圆点
						{
							const geomerty = new THREE.CircleGeometry(100, 32);
							const material = new THREE.MeshBasicMaterial({
								color: new THREE.Color('#00ffbe'),
								transparent: true,
								opacity: 1,
								side: THREE.DoubleSide,
							});
							const circle = new THREE.Mesh(geomerty, material);
							circle.position.set(point[0], 0, point[1]);
							circle.rotateX(THREE.MathUtils.degToRad(270));
							cGroup.add(circle);
						}

						//  城市外圈
						{
							const geomerty = new THREE.RingGeometry(130, 140, 32);
							const material = new THREE.MeshBasicMaterial({
								color: new THREE.Color('#00ffbe'),
								transparent: true,
								opacity: 1,
								side: THREE.DoubleSide,
							});
							const ring = new THREE.Mesh(geomerty, material);
							ring.position.set(point[0], 0, point[1]);
							ring.rotateX(THREE.MathUtils.degToRad(270));
							cGroup.add(ring);
						}

						// 光柱
						{
							const geomerty = new THREE.ConeGeometry(20, 2000, 32);
							const material = new THREE.MeshBasicMaterial({
								color: new THREE.Color('#00ffbe'),
								transparent: true,
								opacity: 1,
								side: THREE.DoubleSide,
							});
							const cone = new THREE.Mesh(geomerty, material);
							cone.position.set(point[0], 1000, point[1]);
							cGroup.add(cone);
						}

						// 光环
						{
							const ringGroup = new THREE.Group();
							cityGroup.add(ringGroup);
							const geomerty = new THREE.RingGeometry(200, 220, 32);
							const material = new THREE.MeshBasicMaterial({
								color: new THREE.Color('#00ffbe'),
								transparent: true,
								opacity: 1,
								side: THREE.DoubleSide,
							});
							const ring = new THREE.Mesh(geomerty, material);
							ring.name = 'ring1';
							const ring2 = ring.clone();
							ring2.geometry = new THREE.RingGeometry(270, 290, 32);
							const ring3 = ring.clone();
							ring3.geometry = new THREE.RingGeometry(350, 370, 32);
							ringGroup.add(ring, ring2, ring3);

							// ringGroup.scale.set()
							ringGroup.rotateX(THREE.MathUtils.degToRad(90));
							ringGroup.position.set(point[0], 0, point[1]);

							// 动画
							{
								const scaleTimes = [0, 2];
								const scaleValues = [0, 2];
								const opacityValues = [1, 0];
								const frameTrack1 = new THREE.VectorKeyframeTrack('.scale[x]', scaleTimes, scaleValues);
								const frameTrack2 = new THREE.VectorKeyframeTrack('.scale[y]', scaleTimes, scaleValues);
								const frameTrack3 = new THREE.VectorKeyframeTrack('ring1.material.opacity', scaleTimes, opacityValues);
								const animationClip = new THREE.AnimationClip('ringScale', 2, [frameTrack1, frameTrack2, frameTrack3]);
								requestAnimationFrame(() => {
									const animationAction = this.threeX.animationMixer.clipAction(animationClip, ringGroup);
									animationAction.play();
								});
							}
						}

						// 文字
						{
							const canvas = this.createTextCanvas(fullname);
							if (!canvas) return;
							const map = new THREE.CanvasTexture(canvas);
							map.needsUpdate = true;
							const material = new THREE.SpriteMaterial({
								map,
							});
							const text = new THREE.Sprite(material);
							text.position.set(point[0], 2300, point[1]);
							text.scale.set(2000, 400, 1);
							cGroup.add(text);
						}
					}
				});
			}
		});
	}
	/**
	 * 数字上升效果
	 */
	public createNumMesh() {
		const canvas = document.createElement('canvas');
		canvas.width = 800;
		canvas.height = 100;
		const c = canvas.getContext('2d');
		if (!c) return;

		let step = 0;
		const gradient = c.createLinearGradient(0, 0, 800, 0);
		gradient.addColorStop(0, 'rgba(0,255,255,1)');
		gradient.addColorStop(1, 'rgba(0,255,255,0.2)');
		c.font = 'bold 100px 宋体';
		c.shadowColor = 'rgba(0,0,0,0)';
		c.textBaseline = 'middle';
		c.shadowOffsetX = 0;
		c.shadowOffsetY = 0;
		c.shadowBlur = 30;
		c.shadowColor = 'rgba(255, 255, 255, 1)';
		c.fillStyle = gradient;
		c.fillText(createRandomNum(12), 25, 50);

		const texture = new THREE.CanvasTexture(canvas);
		const geomerty = new THREE.PlaneGeometry(2000, 200);
		const material = new THREE.MeshBasicMaterial({
			transparent: true,
			opacity: 1,
			map: texture,
			side: THREE.DoubleSide,
		});
		const mesh = new THREE.Mesh(geomerty, material);
		mesh.rotateX(THREE.MathUtils.degToRad(270));

		animate();
		function animate() {
			step++;
			if (step % 5 === 0) {
				c?.clearRect(0, 0, 800, 100);
				c?.fillText(createRandomNum(12), 25, 50);
				mesh.material.map = new THREE.CanvasTexture(canvas);
			}
			requestAnimationFrame(() => {
				animate();
			});
		}

		function createRandomNum(size: number) {
			return new Array(size)
				.fill(0)
				.map(n => {
					n = Math.floor(Math.random() * 10);
					return n;
				})
				.join('');
		}

		return mesh;
	}
	/**
	 * 创建文字图片
	 */
	public createTextCanvas(text: string) {
		const canvas = document.createElement('canvas');
		canvas.width = 700;
		canvas.height = 100;
		const c = canvas.getContext('2d');
		if (!c) return;

		c.font = 'bold 60px 宋体';
		c.textBaseline = 'middle';
		c.textAlign = 'center';
		c.fillStyle = '#ffffff';
		c.fillText(text, 350, 50);

		return canvas;
	}
	/**
	 * geo转世界坐标
	 */
	public geoToWorld(point: [number, number]) {
		const center: [number, number] = [111.5970038, 27.6456252];
		const projection = geoMercator()
			.center(center)
			.scale(ThreeX.len * 15);
		return <[number, number]>projection(point);
	}
	/**
	 * 测试
	 * @returns
	 */
	public createTest() {
		const geomerty = new THREE.BoxGeometry(1000, 1000, 1000);
		const material = new THREE.MeshBasicMaterial({
			color: new THREE.Color('#ff0000'),
		});
		const mesh = new THREE.Mesh(geomerty, material);
		const mesh1 = mesh.clone();
		mesh1.position.set(3000, 0, 0);
		mesh1.material = new THREE.MeshBasicMaterial({
			color: new THREE.Color('#00ff00'),
		});
		this.threeX.scene.add(mesh, mesh1);
	}
}

export { Model };
