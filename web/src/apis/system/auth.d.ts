// 登录
export type LoginBodyType = {
	username: string;
	password: string;
};

// 获取所有用户信息
export type QueryParamsType = {
	pageSize: number;
	page: number;
	username?: string;
	realName?: string;
	phone?: string;
};

// 添加人员
export type AuthBodyType = {
	username: string;
	roleId: number;
	realName?: string;
	phone?: string;
};

// 更新人员
export type UpdateAuthBodyType = AuthBodyType & {
	id: number | null;
	password?: string;
};
