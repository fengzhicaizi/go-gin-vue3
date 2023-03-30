export type AuthStateBaseType = {
	id: number;
	username: string;
	realName: string;
	phone: string;
	roleId: number;
	isRoot: boolean;
};

export type AuthStateRoleType = {
	id: number;
	menuIds: string;
	name: string;
};

export type AuthStateMenuType = {
	id: number;
	pid: number;
	name: string;
	code: string;
	path: string;
	icon: string;
	sort: number;
	children: AuthStateMenuType[];
};

export type AuthStateType = {
	auth: AuthStateBaseType | null;
	role: AuthStateRoleType | null;
	menu: AuthStateMenuType | null;
};
