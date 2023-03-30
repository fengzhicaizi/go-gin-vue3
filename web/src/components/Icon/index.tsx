import { createFromIconfontCN } from '@ant-design/icons-vue';
// import { defineComponent } from 'vue';
import conf from '@/conf/index';

const MyIcon = createFromIconfontCN({
	extraCommonProps: {
		class: 'icon-style',
	},
	scriptUrl: conf.iconUrl, // 在 iconfont.cn 上生成
});

export default MyIcon;
