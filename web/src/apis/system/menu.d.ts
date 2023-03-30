type QueryMenusParamsType = {
	//
};

export type AddMenuBodyType = {
	pid: number;
	name: string;
	code: string;
	path: string;
	sort: number;
	mark: string;
	icon: string;
};

export type UpdateMenuBodyType = AddMenuBodyType & {
	id: number;
};
