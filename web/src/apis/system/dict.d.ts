export type DictType = {
	id?: number;
	name: string;
	type: string;
	remark: string;
};

export type DictData = {
	id?: number;
	label: string;
	value: string;
	sort: number;
};

export type queryAllDictDataType = {
	typeId?: number;
	type?: string;
};
