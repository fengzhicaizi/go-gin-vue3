export type openFnArgsType = {
	type?: 'edit' | 'add';
	data?: AuthBodyType & { id: number };
	ok?: () => void;
	cancel?: () => void;
};
