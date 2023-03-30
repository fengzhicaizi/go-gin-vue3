type ConfType = {
	baseUrl: string; // 后端api请求地址
	iconUrl: string; // iconfont地址
	fullScreenId: string; // full-screen 锚点id
	cesiumIon: string; // cesiumkey
};

const conf: ConfType = {
	baseUrl: 'http://localhost:9001',
	iconUrl: '//at.alicdn.com/t/c/font_3617313_l8x6wg0frts.js',
	fullScreenId: 'full-screen',
	cesiumIon: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJlN2M2MGJjNS0wMzZlLTQ1NjYtODQ0Mi1kZjQ2NGIwZDMwMmYiLCJpZCI6MTI0MDE3LCJpYXQiOjE2NzU5MzA3MTR9.lwJu5x5VjD2SX50Tlasg-w4qbhF000GdG06XD8fl5Tk`,
};

export default conf;
