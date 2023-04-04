import request from '@/utils/request';

export const uploadApi = (url: string, formData: FormData) => {
	return request.post(url, {
		file: formData,
	});
};
